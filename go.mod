module github.com/shimt/pam-exec-oauth2

replace golang.org/x/oauth2 => ./vendor/golang.org/x/oauth2

require (
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/shimt/go-simplecli v0.0.3
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20191029155521-f43be2a4598c // indirect
	google.golang.org/appengine v1.6.5 // indirect
)

go 1.13
