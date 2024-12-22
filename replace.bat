cd C:\Windows\System32\zh-CN
takeown /f user32.dll.mui
icacls user32.dll.mui /grant Administrators:F
ren user32.dll.mui user32.dll.mui.bak