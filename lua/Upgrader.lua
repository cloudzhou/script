--------------------------------------
-- Upgrader module for NODEMCU
-- LICENCE: http://opensource.org/licenses/MIT
-- cloudzhou<wuyunzhou@espressif.com>
--------------------------------------

--[[
require('Upgrader')
Upgrader.update('EspClient.lua', 'http://115.29.202.58/static/lua/EspClient.lua')
Upgrader.updateEspClient()
]]--

local moduleName = ...
local M = {}
_G[moduleName] = M

local header = ''
local isTruncated = false
local function save(filename, response)
    if isTruncated then
        file.writeline(response)
        return
    end
    header = header..response
    local i, j = string.find(header, '\r\n\r\n')
    if i == nil or j == nil then
        return
    end
    prefixBody = string.sub(header, j+1, -1)
    file.writeline(prefixBody)
    header = ''
    isTruncated = true
    return
end

----
function M.update(filename, url)
    file.open(filename, 'w')
    local ip, port, path = string.gmatch(url, 'http://([0-9.]+):?([0-9]*)(/.*)')()
    if ip == nil then
        return false
    end
    if port == nil then
        prot = 80
    end
    if path == '' then
        path = '/'
    end
    socket = net.createConnection(net.TCP, false)
    socket:on('receive', function(sck, response)
        save(filename, response)
    end)
    socket:on('disconnection', function(sck, response)
        print('disconnect at '..tmr.now())
        header = ''
        isTruncated = false
        file.close()
    end)
    socket:connect(port, ip)
    conn:send('GET '..path..' HTTP/1.0\r\nHost: '..ip..'\r\n'
    ..'Connection: close\r\nAccept: */*\r\n\r\n')
end

function M.updateEspClient()
    return M.update('EspClient.lua', 'http://115.29.202.58/static/lua/EspClient.lua')
end

