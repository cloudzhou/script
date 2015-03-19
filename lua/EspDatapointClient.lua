local moduleName = ...
local M = {}
_G[moduleName] = M

local conn = nil
local devicekey = nil
local server = '115.29.202.58' -- iot.espressif.cn
local port = 8000

local keepAliveTime = 60000
local isConnected = false

local function getStr(str, key)
    for k in string.gmatch(str, '.*"'..key..'" *: *"([^"]+)".*') do
        if k ~= nil then
            return k
        end
    end
    return nil
end

local function getNumber(str, key)
    for k in string.gmatch(str, '.*"'..key..'" *: *([0-9.]+).*') do
        if k ~= nil then
            return k
        end
    end
    return nil
end

local function identify()
    local identifystr = '{"path": "/v1/device/identify/", "method": "POST", "meta": {"Authorization": "token '..devicekey..'"}}\n'
    if isConnected == true then
        conn:send(identifystr)
    end
end

local function route(response)
end

local function connect()
    conn = net.createConnection(net.TCP, false)
    conn:on('connection', function(sck, response)
        isConnected = true
        identify()
    end)
    conn:on('disconnection', function(sck, response)
        isConnected = false
        connect()
    end)
    conn:on('receive', function(sck, response)
        route(response)
    end)
    conn:on('sent', function(sck, response)
    end)
    conn:connect(port, server)
end

local function keepAlive()
    local pingstr = '{"path": "/v1/ping/", "method": "GET", "meta": {"Authorization": "token '..devicekey..'"}}\n'
    if isConnected == true then
        conn:send(pingstr)
    else
        connect()
    end
end
----
function M.init(key)
    if key == nil or key == '' then
        assert(false, 'need key')
    end
    devicekey = key
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

----
function M.setKeepAliveTime(t)
    keepAliveTime = t
end

