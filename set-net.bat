@echo off 
chcp 65001
if not "%OS%"=="Windows_NT" exit
title WindosActive

cd /D %~dp0

echo.
./simulate/bin/client.exe
pause