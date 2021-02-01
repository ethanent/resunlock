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
	fmt.Println("===== ResUnlock =====")
	
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

	fmt.Println("Checking status...")

	blocked, err := requests.CheckBlocked(c)

	if err != nil {
		panic(err)
	}

	if !blocked {
		fmt.Println("You are already authorized.")
		return
	}

	fmt.Println("Not authorized. Submitting requests...")

	if err := requests.SubmitAuthRest(c); err != nil {
		panic(err)
	}

	fmt.Println("OK. Authorization requested. Waiting for confirmation...")

	for {
		b, err := requests.CheckBlocked(c)

		if err != nil {
			panic(err)
		}

		if b == false {
			fmt.Println("\nAuthorization confirmed.")
			return
		}

		fmt.Print(".")

		time.Sleep(time.Millisecond * 700)
	}
}
