package main

import (
	"bytes"
	"github.com/Tkanos/gonfig"
	"github.com/alexflint/go-arg"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Cfg struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	ServiceURL   string
	Timeout      int
}

type Args struct {
	Config string `arg:"required"`
	Key    string `arg:"required"`
	File   string `arg:"required"`
	Stdout bool
}

func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := parseConf(args.Config)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	// create an http client and get our oauth token
	client := oauthClient(ctx, cfg)
	resp, err := client.Get(cfg.TokenURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.Status)

	file, err := os.Open(args.File)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = uploadFile(client, cfg.ServiceURL, args.Key, file)
	if err != nil {
		log.Fatal(err)
	}
}

func parseConf(cfgFile string) (Cfg, error) {
	// read config file/env
	cfg := Cfg{Timeout: 60}
	err := gonfig.GetConf(cfgFile, &cfg)

	return cfg, err
}

func parseArgs() (Args, error) {
	var args Args
	err := arg.Parse(&args)

	return args, err
}

func oauthClient(ctx context.Context, c Cfg) *http.Client {
	// oauth2 client
	oauthConf := clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     c.TokenURL,
		Scopes:       []string{},
	}

	return oauthConf.Client(ctx)
}

func uploadFile(client *http.Client, url string, key string, file *os.File) (res *http.Response, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	var fw io.Writer
	if x, ok := file.(io.Closer); ok {
		defer x.Close()
	}
	// Add a file
	if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
		return
	}
	if _, err = io.Copy(fw, r); err != nil {
		return _, err
	}

	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)

	return
}
