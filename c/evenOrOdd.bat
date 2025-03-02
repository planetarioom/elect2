@echo off 
setlocal EnableDelayedExpansion
set /p num=type number: 
for /L %%i in (0,1,%num%) do (
    set /a remainder=%%i %% 2
    if !remainder! == 0 (
	echo %%i is even
    ) else (
	echo %%i is odd
    )
)
endlocal 
pause