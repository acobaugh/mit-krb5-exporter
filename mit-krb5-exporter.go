package main

import (
	"github.com/Tkanos/gonfig"
	//"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

type Cfg struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
}

func main() {
	// read config file/env
	cfg := Cfg{}
	err := gonfig.GetConf("mit-krb5-exporter.json", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// oauth2 client
	oauthConf := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		Scopes:       []string{},
	}

	ctx := context.Background()
	tok, err := oauthConf.Token(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%s\n", tok)
}
