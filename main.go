package main

import (
	"crypto/tls"
	"fmt"
	"github.com/ethanent/resunlock/requests"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	jar, err := cookiejar.New(&cookiejar.Options{})

	if err != nil {
		panic(err)
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Jar:     jar,
		Timeout: time.Second * 5,
	}

	fmt.Println("Submitting requests...")

	if err := requests.SubmitAuthRest(c); err != nil {
		panic(err)
	}

	if err := requests.FetchSafeConnect(c); err != nil {
		panic(err)
	}

	fmt.Println("OK. Unlocked.")
}
