set GOOS=linux
set GOARCH=arm

go build -o "pam-exec-oauth2.%GOOS%-%GOARCH%"