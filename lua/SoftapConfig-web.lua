local moduleName = ...
local M = {}
_G[moduleName] = M

local handlers = {}

local function response(status, body)
    local length = string.len(body)
    return 'HTTP/1.0 '..status..' OK\r\nServer: esphttpd/0.9\r\nContent-Type: text/html\r\nContent-Length: '..length..'\r\nConnection: close\r\n\r\n'..body
end

local function urlDecode(str)
    if str == nil then
        return nil
    end
    str = string.gsub(str, '+', ' ')
    str = string.gsub(str, '%%(%x%x)', function(h) return string.char(tonumber(h, 16)) end)
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
        ssid = urlDecode(string.gmatch(data, 'ssid=([^&]+)')())
        password = urlDecode(string.gmatch(data, 'password=([^&]+)')())
        mode = urlDecode(string.gmatch(data, 'mode=([^&]+)')())
        setmode(mode, ssid, password)
        tmr.delay(3000000)
    end
    local modestr = 'nil'
    local mode = wifi.getmode()
    local ip = wifi.sta.getip()
    if mode == 1 then
        modestr = 'station'
    elseif mode == 2 then
        modestr = 'softap'
    end
    if ip == nil then
        ip = '...'
    end
    file.open('config.html', 'r')
    body = ''
    while true do
        local line = file.readline()
        if line == nil then
            break
        end
        line = string.gsub(line, '{{mode}}', modestr)
        line = string.gsub(line, '{{ip}}', ip)
        body = body .. line
    end
    conn:send(response(200, body))
end

local function receive(conn, data)
    local i, j = string.find(data, '\r\n\r\n')
    if i == nil then
        return false
    end
    local header = string.sub(data, 1, i-1)
    local body = string.sub(data, i+4, -1)
    local data = nil
    local method, path = 'GET', string.gmatch(header, 'GET ([0-9a-zA-Z.-_/]+) HTTP/1.+')()
    if path == nil then
        method, path = 'POST', string.gmatch(header, 'POST ([0-9a-zA-Z.-_/]+) HTTP/1.+')()
    end
    if path == nil then
        return
    end
    func = handlers[path]
    if func == nil then
        conn:send(response(404, '404 Not Found'))
        conn:close()
        return
    end
    func(conn, path, method, body)
    conn:close()
    conn = nil
    path = nil
    method = nil
    body = nil
end

local function route(conn)
    conn:on('receive', receive)
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

