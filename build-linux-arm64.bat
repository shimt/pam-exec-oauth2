set GOOS=linux
set GOARCH=arm64

go build -o "pam-exec-oauth2.%GOOS%-%GOARCH%"