@echo off 
chcp 65001
if not "%OS%"=="Windows_NT" exit
title WindosActive

cd /D %~dp0

echo.
echo 如果不能上网请更换IP地址
echo 默认不会更改ip地址,检测是否连接
echo.
set /P change_ip=是否更换IP(y/[n]):
if "%change_ip%"=="y" (
python ./main.py ip
) else (
python ./main.py
)
pause