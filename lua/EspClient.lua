--------------------------------------
-- EspClient module for NODEMCU
-- LICENCE: http://opensource.org/licenses/MIT
-- cloudzhou<wuyunzhou@espressif.com>
--------------------------------------

--[[
get EspClient.lua from remote server via http:
script = ''
conn = net.createConnection(net.TCP, false) 
conn:on("receive", function(conn, pl) script = script..pl end)
conn:connect(80, "115.29.202.58")
conn:send("GET /static/script/Upgrader.lua HTTP/1.0\r\nHost: iot.espressif.cn\r\n"
    .."Connection: close\r\nAccept: */*\r\n\r\n")
i, j = string.find(script, '\r\n\r\n')
script = string.sub(script, j+1, -1)
file.open("EspClient.lua", "w")
file.write(script)
file.close()
node.restart()

demo.lua:
require("EspClient")
EspClient.init("513d09340e29eb61f91f5cb4e717682c48444d6f")
EspClient.run()

-- send datapoint
datapoint = {}
datapoint['x'] = 1
EspClient.datapoint("light", datapoint)

-- on datapoint
function light(datastreamName, datapoint)
   print("test function!"..datapoint.x)
end
EspClient.onDatapoint("light", light)

-- on rpc
function rpc(action, parameters)
   print("test function!"..action)
end
EspClient.onRpc("action", rpc)

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

local keepAliveTime = 60000
local isConnected = false
local isDebug = false
local buffer = nil

local function getStr(str, key)
    for k in string.gmatch(str, '.*"'..key..'" *: *"([^"]+)".*') do
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
        identify()
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
        return false
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
                conn:send('{"status": 200, "nonce": '..nonce..'}\n')
            end
        end
        return true
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
                conn:send('{"status": 200, "nonce": '..nonce..'}\n')
            end
        end
        return true
    end

    print('unsupport command')
    return false
end

local function identify()
    local identifystr = '{"path": "/v1/device/identify/", "method": "POST", "meta": {"Authorization": "token '..devicekey..'"}}\n'
    if isConnected == true then
        conn:send(identifystr)
    end
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
    if devicekey == nil or devicekey == '' then
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
function M.datapoint(datastreamName, datapoint)
    datapointStr = ''
    if datapoint.at ~= nil then
        datapointStr = datapointStr..'"at": "'..datapoint.at..'", '
    end
    if datapoint.x ~= nil then
        datapointStr = datapointStr..'"x": '..datapoint.x..', '
    end
    if datapoint.y ~= nil then
        datapointStr = datapointStr..'"y": '..datapoint.y..', '
    end
    if datapoint.z ~= nil then
        datapointStr = datapointStr..'"z": '..datapoint.z..', '
    end
    if datapoint.k ~= nil then
        datapointStr = datapointStr..'"k": '..datapoint.k..', '
    end
    if datapoint.l ~= nil then
        datapointStr = datapointStr..'"l": '..datapoint.l..', '
    end
    if datapointStr == '' then
        return
    end
    datapointStr = string.sub(datapointStr, 0, -3)
    conn:send('{"path": "/v1/datastreams/'..datastreamName..'/datapoint/", "method": "POST", "meta": {"Authorization": "token '..devicekey..'"}, "body": {"datapoint":{'..datapointStr..'}}}\n')
end

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

