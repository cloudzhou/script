curl http://nation.ppsj.com.cn/jiangshu/index18_list.html|iconv -f gbk -t utf8 > 0.html
for x in {1..188}; do
    curl http://nation.ppsj.com.cn/jiangshu/index18_list115.html|iconv -f gbk -t utf8 > $x.html
    sleep 0.1
done

# curl http://mp.ppsj.com.cn/brand1/yilezhiye.html | iconv -f gbk -t utf-8 | replace gb2312 utf8 > 0_9.html; sleep 0.1
# w3m -dump -I utf-8 -O utf-8 0_9.html 0_9.txt
