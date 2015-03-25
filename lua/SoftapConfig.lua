local moduleName = ...
local M = {}
_G[moduleName] = M

local router = {}

local function urlDecode(str)
  str = string.gsub(str, '+', ' ')
  str = string.gsub(str, '%%(%x%x)', function(h) return string.char(tonumber(h, 16)) end)
  str = string.gsub(str, '\r\n', '\n')
  return str
end

local function config(conn, path, method, data)
    if method == 'POST' then
    end
    local modestr = 'nil'
    local mode = wifi.getmode()
    local ip = wifi.sta.getip()
    if mode == 1 then
        modestr = 'station'
    else if mode == 2 then
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
    length = string.len(body)
    conn:send('HTTP/1.0 200 OK\r\nServer: esphttpd/0.9\r\nContent-Type: text/html\r\nContent-Length: '..length..'\r\nConnection: close\r\n\r\n')
    conn:send(body)
    conn:close()
end

----
function M.setmode(mode, ssid, pwd)
    if mode == 'SOFTAP' then
        wifi.setmode(wifi.SOFTAP)
        wifi.ap.config({ssid=ssid, pwd=pwd})
    end
    if mode == 'STATION' then
        wifi.setmode(wifi.STATION)
        wifi.sta.config(ssid, pwd)
    end
    return wifi.sta.getip()
end

function M.handle(path, func)
    router[path] = func
end

function M.run()
    M.handle('/config/', config)
end
----
