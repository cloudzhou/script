local moduleName = ...
local M = {}
_G[moduleName] = M

local handlers = {}
local connections = {}

local function response(status, body)
    local length = string.len(body)
    return 'HTTP/1.0 '..status..' OK\r\nServer: esphttpd/0.9\r\nContent-Type: text/html\r\nContent-Length: '..length..'\r\nConnection: close\r\n\r\n'..body
end

local function urlDecode(str)
  str = string.gsub(str, '+', ' ')
  str = string.gsub(str, '%%(%x%x)', function(h) return string.char(tonumber(h, 16)) end)
  str = string.gsub(str, '\r\n', '\n')
  return str
end

local function setmode(mode, ssid, pwd)
    if mode == 'softap' then
        wifi.setmode(wifi.SOFTAP)
        wifi.ap.config({ssid=ssid, pwd=pwd})
    end
    if mode == 'station' then
        wifi.setmode(wifi.STATION)
        wifi.sta.config(ssid, pwd)
    end
    return wifi.sta.getip()
end

local function config(conn, path, method, data)
    if method == 'POST' then
        data = urlDecode(data)
        ssid = string.gmatch(data, 'ssid=([^&]+)')()
        password = string.gmatch(data, 'password=([^&]+)')()
        mode = string.gmatch(data, 'mode=([^&]+)')()
        setmode(mode, ssid, password)
        tmr.dealy(1000)
    end
    local modestr = 'nil'
    local mode = wifi.getmode()
    local ip = wifi.sta.getip()
    if mode == 1 then
        modestr = 'station'
    elseif mode == 2 then
        modestr = 'softap'
    end
    file.open('config.html', 'r')
    body = ''
    while true do
        local line = file.readline()
        if line == nil then
            break
        end
        line = string.gsub(line, '{{mode}}', mode)
        line = string.gsub(line, '{{ip}}', ip)
        body = body .. line
    end
    conn:send(response(200, body))
end

local function getConnPair(conn)
    local i = 0
    while i < 10 do
        local key = 'k'..i
        local value = connections[key]
        if conn == nil and value == nil then
            return key, value
        end
        if conn ~= nil and value ~= nil and value.c == conn then
            return key, value
        end
        i = i + 1
    end
    return nil, nil
end

local function connection(conn, data)
    key, value = getConnPair(nil)
    if key == nil then
        conn:send(response(500, 'Too Many Connections'))
        conn:close()
        return
    end
    value = {c=conn, t=tmr.now(), p=nil, l=0, m=nil, h='', b=nil} -- path, length, method, header, body
    connections[key] = value
end

local function receive(conn, data)
    key, value = getConnPair(conn)
    if key == nil or value == nil then
        conn:send(response(500, 'Too Many Connections'))
        conn:close()
        return
    end
    value.t = tmr.now()
    if value.b ~= nil then
        value.b = value.b .. data
        if string.len(value.b) > 256 then
            conn:send(response(400, 'Too Looong'))
            disconnection(conn, data)
        end
    else
        value.h = value.h .. data
        if string.len(value.h) > 256 then
            conn:send(response(400, 'Too Looong'))
            disconnection(conn, data)
        end
        local i, j = string.find(value.h, '\r\n\r\n')
        if i == nil then
            return false
        end
        value.b = string.sub(value.h, i+4+1, -1)
        value.h = '\r\n' .. string.sub(value.h, 1, i-1)
        value.m, value.p = 'GET', string.gmatch(value.h, '\r\nGET ([0-9a-zA-Z.-_/]+) HTTP/1.+')()
        if value.p == nil then
            value.m, value.p = 'POST', string.gmatch(value.h, '\r\nPOST ([0-9a-zA-Z.-_/]+) HTTP/1.+')()
        end
        if value.p == nil then
            disconnection(conn, '')
        end
        if value.m == 'POST' then
            value.l = string.gmatch(value.h, '\r\nContent-Length: ([0-9]+)')()
            if value.l == nil then
                disconnection(conn, '')
            end
            value.l = tonumber(value.l)
        end
    end
    if value.m == nil then
        return
    end
    if (value.m == 'GET') or (value.m == 'POST' and value.b ~= nil and string.len(value.b) == value.l) then
        func = handlers[value.p]
        if func == nil then
            conn:send(response(404, '404 Not Found'))
        end
        func(conn, value.p, value.m, value.b)
        disconnection(conn, data)
    end
end

local function disconnection(conn, data)
    key, value = getConnPair(conn)
    if key == nil then
        return
    end
    connections[key] = nil
end

local function connCheck()
    local expired = tmr.now() - 10000
    local i = 0
    while i < 10 do
        local key = 'k'..i
        local value = connections[key]
        if value and value.t < expired then
            disconnection(value.c, '')
        end
    end
end

local function route(conn)
    conn:on('connection', connection)
    conn:on('receive', receive)
    conn:on('disconnection', disconnection)
    tmr.alarm(1, 5000, 1, function() 
        connCheck()
    end)
end
----
function M.handle(path, func)
    handlers[path] = func
end

function M.run()
    srv = net.createServer(net.TCP) 
    srv:listen(80, route)
    M.handle('/config/', config)
end
----
