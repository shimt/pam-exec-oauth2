// Copyright © 2017 Shinichi MOTOKI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/shimt/go-simplecli"
	"golang.org/x/oauth2"
)

var cli = simplecli.NewCLI()

func initCLI() {
	cli.CommandLine.String("client-id", "", "OAuth2 Client ID")
	cli.CommandLine.String("client-secret", "", "OAuth2 Client Secret")
	cli.CommandLine.StringArray("scopes", []string{}, "OAuth2 Scopes")
	cli.CommandLine.String("redirect-url", "", "OAuth2 Redirect URL")
	cli.CommandLine.String("endpoint-auth-url", "", "OAuth2 End Point Auth URL")
	cli.CommandLine.String("endpoint-token-url", "", "OAuth2 End Point Token URL")
	cli.CommandLine.String("username-format", "%s", "username format")

	cli.BindSameName("client-id")
	cli.BindSameName("client-secret")
	cli.BindSameName("scopes")
	cli.BindSameName("redirect-url")
	cli.BindSameName("endpoint-auth-url")
	cli.BindSameName("endpoint-token-url")
	cli.BindSameName("username-format")

	cli.Config.SetConfigFile(
		filepath.Join(cli.Application.Directory, cli.Application.Name+".yaml"),
	)
}

func init() {
	initCLI()
}

func main() {
	err := cli.Setup()
	cli.Exit1IfError(err)

	if cli.ConfigFile != "" {
		fmt.Println("Using config file:", cli.ConfigFile)
	}

	username := os.Getenv("PAM_USER")
	password := ""

	stdinScanner := bufio.NewScanner(os.Stdin)
	if stdinScanner.Scan() {
		password = stdinScanner.Text()
	}

	cli.Log.Debug("create oauth2Config")
	cli.Log.Debugf("ClientID: %s", cli.Config.GetString("client-id"))
	cli.Log.Debugf("ClientSecret: %s", cli.Config.GetString("client-secret"))
	cli.Log.Debugf("Scopes: %s", cli.Config.GetStringSlice("scopes"))
	cli.Log.Debugf("EndPoint.AuthURL: %s", cli.Config.GetString("endpoint-auth-url"))
	cli.Log.Debugf("EndPoint.TokenURL: %s", cli.Config.GetString("endpoint-token-url"))

	oauth2Config := oauth2.Config{
		ClientID:     cli.Config.GetString("client-id"),
		ClientSecret: cli.Config.GetString("client-secret"),
		Scopes:       cli.Config.GetStringSlice("scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  cli.Config.GetString("endpoint-auth-url"),
			TokenURL: cli.Config.GetString("endpoint-token-url"),
		},
	}

	extraParameters := url.Values{}

	for k, v := range cli.Config.GetStringMapString("extra-parameters") {
		extraParameters[k] = []string{v}
	}

	cli.Log.Debug("create oauth2Context")

	oauth2Context := context.Background()

	cli.Log.Debug("call PasswordCredentialsToken")

	oauth2Token, err := oauth2Config.PasswordCredentialsTokenEx(
		oauth2Context,
		fmt.Sprintf(cli.Config.GetString("username-format"), username),
		password,
		extraParameters,
	)

	cli.Exit1IfError(err)

	if !oauth2Token.Valid() {
		cli.Exit(1)
		cli.Log.Debug("OAuth2 authentication failed")
	}

	cli.Log.Debug("OAuth2 authentication success")
	cli.Exit(0)
}
