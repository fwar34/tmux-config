#!/usr/bin/lua

-- %d 数字 %D 非数字 %:转义特殊字符 %s 空白字符
io.input(io.stdin)
local out = io.read()
-- local out = "2065/7919MB [||        ] 28.5% 0.46 0.60 0.59"
local mem, cpu = out:match("(.*MB)[%s%[%]|]*(%d+%.%d%%)")
-- print(mem, " ", cpu)

local temp = function ()
    if os.execute('sensors > /dev/null 2>&1') then
        local t = io.popen('sensors')
        local output = t:read('*all')
        local temp = output:match("Package id 0:%s*+(%d*%.*%d*).*")
        if not temp then
            return nil
        end
        return temp .. ' °C'
    end
end

local core_temp = temp()
if core_temp then
    print(core_temp, ' | ', mem, ' | ', cpu)
else
    print(mem, ' | ', cpu)
end
