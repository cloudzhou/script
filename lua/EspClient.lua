--------------------------------------
-- EspClient module for NODEMCU
-- LICENCE: http://opensource.org/licenses/MIT
-- cloudzhou<wuyunzhou@espressif.com>
-- inspired by yangbo
--------------------------------------

--[[
get EspClient.lua from remote:
function receive(conn, pl)
    print("1\n")
    print(pl)
end
conn = net.createConnection(net.TCP, false) 
conn:on("receive", receive)
conn:connect(80, "115.29.202.58")
conn:send("GET /static/script/EspClient.lua HTTP/1.0\r\nHost: iot.espressif.cn\r\n"
    .."Connection: close\r\nAccept: */*\r\n\r\n")


demo.lua:

require("EspClient")
EspClient.init("513d09340e29eb61f91f5cb4e717682c48444d6f")

-- datapoint
function light(datastreamName, datapoint)
   print("test function!"..datapoint.x)
end
EspClient.datapoint("light", light)

-- rpc
function rpc(action, parameters)
   print("test function!"..action)
end
EspClient.rpc("action", rpc)

-- run
EspClient.run()

]]--

local moduleName = ...
local M = {}
_G[moduleName] = M

local conn = nil
local devicekey = nil
local server = '115.29.202.58' -- iot.espressif.cn
local port = 8000

local rpcMapFunc = {}
local datapointMapFunc = {}

local keepAliveTime = 50000
local isConnected = false
local isDebug = false
local buffer = nil

local function getStr(str, key)
    for k in string.gmatch(str, '.*"'..key..'" *: *"(%w+)".*') do
        if k ~= nil then
            return k
        end
    end
    return nil
end

local function getNumber(str, key)
    for k in string.gmatch(str, '.*"'..key..'" *: *(%d+).*') do
        if k ~= nil then
            return k
        end
    end
    return nil
end

local function connect()
    conn = net.createConnection(net.TCP, false)
    conn:on('connection', function(sck, response)
        print('connected at '..tmr.now())
        isConnected = true
    end)
    conn:on('disconnection', function(sck, response)
        print('disconnect at '..tmr.now())
        isConnected = false
        connect()
    end)
    conn:on('receive', function(sck, response)
        route(response)
    end)
    conn:on('sent', function(sck, response)
        print('sent at '..tmr.now())
    end)
    print('connecting at '..tmr.now())
    conn:connect(port, server)
end

local function route(response)
    buffer = buffer..response
    local i, j = string.find(buffer, '\n')
    if i == nil then
        return nil
    end
    local line = string.sub(buffer, 1, i-1)
    buffer = string.sub(buffer, i+1, -1)
    local path = getStr(response, 'path')
    local nonce = getNumber(response, 'nonce')
    local rpc = string.gmatch(path, '/v1/device/rpc/?')()
    if rpc then
        action = getStr(response, 'action')
        func = rpcMapFunc[action]
        if func ~= nil then
            local result = func(action, {})
            if result and nonce then
                conn:send('{"status": 200, "nonce": '..nonce..'}')
            end
        end
    end
    local datastreamName = string.gmatch(path, '/v1/datastreams/([a-z-_.]+)/datapoint/?')()
    if datastreamName then
        func = datapointMapFunc(datastreamName)
        if func ~= nil then
            local datapoint = {}
            datapoint['x'] = getNumber(response, 'x')
            datapoint['y'] = getNumber(response, 'y')
            datapoint['z'] = getNumber(response, 'z')
            datapoint['k'] = getNumber(response, 'k')
            datapoint['l'] = getNumber(response, 'l')
            local result = func(datastreamName, datapoint)
            if result and nonce then
                conn:send('{"status": 200, "nonce": '..nonce..'}')
            end
        end
    end

    print('unsupport command')
    return nil
end

local function keepAlive()
    local pingstr = '{"path": "/v1/ping/", "method": "GET", "meta": {"Authorization": "token '..devicekey..'"}}\n'
    if isConnected == true then
        conn:send(pingstr)
    else
        connectServer()         
    end
end
----
function M.init(devicekey)
    if devicekey == nil then
        assert(false, 'devicekey must be valid')
    end
    devicekey = devicekey
end

function M.run()
    connect()
    tmr.alarm(1, keepAliveTime, 1, function() 
        keepAlive()
    end)
end

----
function M.onDatapoint(datastreamName, datapointFunc)
    datapointMapFunc[datastreamName] = datapointFunc
end

function M.onRpc(action, rpcFunc)
    rpcMapFunc[action] = rpcFunc
end

----
function M.setKeepAliveTime(keepAliveTime)
    keepAliveTime = keepAliveTime
end

function M.setIsDebug(isDebug)
    isDebug = isDebug
end


