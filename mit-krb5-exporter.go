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
	TokenURL     string
}

func main() {
	// read config file/env
	cfg := Cfg{}
	err := gonfig.GetConf("mit-krb5-exporter.json", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Read config file")

	// oauth2 client
	oauthConf := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		Scopes:       []string{},
	}

	client := oauthConf.Client(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(cfg.TokenURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.Status)
}
