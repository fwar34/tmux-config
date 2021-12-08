#!/usr/bin/lua

-- 1447/7919MB [          ]  2.0% 0.32 0.73 0.60
-- %d 数字 %D 非数字 %:转义特殊字符 %s 空白字符
io.input(io.stdin)
local out = io.read()
-- local out = "2065/7919MB [||        ] 28.5% 0.46 0.60 0.59"
local mem, cpu = out:match("(%d*MB)[%s%[%]|]*(%d+%.%d%%)")
print(mem, " ", cpu)
