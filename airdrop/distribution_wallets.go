package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func read100Seeds(fn string) []string {
	seeds := make([]string, 100)

	f, err := os.Open(fn)
	defer f.Close()

	reader := bufio.NewReader(f)

	if err != nil {
		panic(err)
	}

	seed := ""
	for index := 0; ; index++ {
		seed, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		seeds[index] = seed[0 : len(seed)-1]
	}

	return seeds
}

func createWallet(label, seed string) string {

	v := url.Values{}
	v.Set("seed", seed)
	v.Set("label", label)
	v.Set("scan", "1")

	body := v.Encode()

	req, err := http.NewRequest("POST", urlWalletCreate, strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-CSRF-Token", getCsrfToken())

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return fileName(string(data))
}

func createDistributionWallets(seeds []string) []string {
	wallets := make([]string, 100)

	for i := 0; i < 100; i++ {
		wallets[i] = createWallet(fmt.Sprintf("distribution%03d", i), seeds[i])
	}

	return wallets
}

func fileName(js string) string {
	fn := ""
	r := bufio.NewReader(strings.NewReader(js))

	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(l, "filename") {
			l = strings.Replace(l, "\"", "", -1)
			l = strings.Replace(l, "filename", "", 1)
			l = strings.Replace(l, ": ", "", 1)
			l = strings.Replace(l, " ", "", -1)
			fn = strings.Replace(l, ",", "", 1)
			fmt.Printf("%s\n", fn)
			break
		}
	}

	return fn
}
