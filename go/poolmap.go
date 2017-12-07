package util

import (
	"container/list"
	"sync"
	"time"
)

// poolmap 是一个通用的 pool 的 map
// 用于具备 key 对象的缓存，使用业务 id 作为 key
// 提供最大长度和生存时间的控制，lazy 回收
// 同时对于具体的 pool，也提供最大长度和生存时间
type Poolmap struct {
	// 数据结构
	elementMap map[int]*list.Element
	list       *list.List

	// 生存周期和最大长度，以及触发回收的长度
	poolDuration   time.Duration
	poolMaxLen     int
	poolReclaimLen int

	// list 移到最前面的求余参数，减少频繁的移动，以及更新时间戳
	moveToFrontMod int

	// 对应 pool 缓存对象的生存周期，最大长度
	resourceDuration time.Duration
	resourceMaxLen   int

	// 初始得到 object 函数
	new func(key int) (interface{}, error)

	sync.Mutex
}

func (this *Poolmap) GetResource(key int) *Resource {
	this.Lock()
	element, ok := this.elementMap[key]
	if !ok {
		this.Unlock()
		return this.newResource(key)
	}
	pool := element.Value.(*Pool)
	resource := pool.Get()
	if resource == nil {
		this.Unlock()
		return this.newResource(key)
	}
	this.Unlock()
	return resource
}

func (this *Poolmap) newResource(key int) *Resource {
	object, err := this.new(key)
	if object == nil || err != nil {
		return nil
	}
	resource := NewResource(object, time.Now().Add(this.resourceDuration))
	return resource
}

func (this *Poolmap) PutResource(key int, r *Resource) {
	if r == nil || r.object == nil {
		return
	}
	if r.HasErr() || r.IsExpired() {
		return
	}
	this.Lock()
	defer this.Unlock()
	element, ok := this.elementMap[key]
	if !ok {
		pool := NewPool(key, time.Now().Add(this.poolDuration), this.resourceMaxLen)
		pool.Put(r)
		element = this.list.PushFront(pool)
		this.elementMap[key] = element
		return
	}
	pool := element.Value.(*Pool)
	pool.Put(r)
	pool.touch = pool.touch + 1
	if pool.touch%this.moveToFrontMod == 0 {
		pool.expired = time.Now().Add(this.poolDuration)
		this.list.MoveToFront(element)
	}
	this.reclaim()
}

func (this *Poolmap) Reclaim() {
	this.Lock()
	defer this.Unlock()
	this.reclaim()
}

func (this *Poolmap) reclaim() {
	if this.list.Len() < this.poolMaxLen {
		return
	}
	var i int
	for {
		element := this.list.Back()
		if element == nil {
			break
		}
		pool := element.Value.(*Pool)
		if !pool.IsExpired() && i >= this.poolReclaimLen {
			break
		}
		i++
		this.list.Remove(element)
		delete(this.elementMap, pool.key)
	}
}

// 具体 pool 对象
// 提供最大长度和生存时间的控制，lazy 回收
type Pool struct {
	key     int
	expired time.Time
	maxLen  int
	list    *list.List
	touch   int
}

func NewPool(key int, expired time.Time, maxLen int) *Pool {
	pool := &Pool{key: key, expired: expired, maxLen: maxLen, list: list.New()}
	return pool
}

func (this *Pool) Get() *Resource {
	for {
		element := this.list.Front()
		if element == nil {
			return nil
		}
		this.list.Remove(element)
		resource := element.Value.(*Resource)
		if !resource.IsExpired() {
			return resource
		}
	}
}

func (this *Pool) Put(r *Resource) {
	if this.list.Len() > this.maxLen {
		return
	}
	this.list.PushBack(r)
}

func (this *Pool) IsExpired() bool {
	return time.Now().After(this.expired)
}

// 抽象的 pool 资源，可以是 conn，rsa.PrivateKey
type Resource struct {
	object  interface{}
	expired time.Time
	err     error
}

func NewResource(object interface{}, expired time.Time) *Resource {
	this := &Resource{object: object, expired: expired}
	return this
}

func (this *Resource) HasErr() bool {
	return this.err != nil
}

func (this *Resource) IsExpired() bool {
	return time.Now().After(this.expired)
}
