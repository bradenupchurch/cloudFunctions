curl https://www.google.com
if %errorlevel% == 0 (
	echo(Ok. HTTP 500
) else (
	echo(Error. HTTP 404
)
