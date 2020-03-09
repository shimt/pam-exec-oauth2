module github.com/shimt/pam-exec-oauth2

replace golang.org/x/oauth2 => ./vendor/golang.org/x/oauth2

require (
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/shimt/go-simplecli v0.0.4
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/appengine v1.6.5 // indirect
)

go 1.13
