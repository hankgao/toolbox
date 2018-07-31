package skytb

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// getWalletBalance returns balance (coins and hours) of a given wallet, coin type should be provided
func getWalletBalance(wn, ct string) (uint64, uint64) {

	ctd := CoinTypeDetails(ct)

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%4d/wallet/balance?id=%s", ctd.WebInterfacePort, wn))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	r := bufio.NewReader(resp.Body)

	r.ReadString('\n')
	r.ReadString('\n')
	c, _ := r.ReadString('\n')
	h, _ := r.ReadString('\n')

	cs := strings.Split(c, ":")[1]
	hs := strings.Split(h, ":")[1]

	// rempve space, '\n', ','
	cs = RemoveGarbage(cs)
	hs = RemoveGarbage(hs)

	coins, err := strconv.ParseUint(cs, 10, 64)
	if err != nil {
		panic(err)
	}
	hours, err := strconv.ParseUint(hs, 10, 64)
	if err != nil {
		panic(err)
	}

	return coins, hours
}

// CreateWallet creates a wallet from a seed. You can give a name for the wallet, which will be used to replace the
// system generated name
func CreateWallet(cn, label, seed string) string {

	v := url.Values{}
	v.Set("seed", seed)
	v.Set("label", label)
	v.Set("scan", "1")

	body := v.Encode()

	ctd := CoinTypeDetails(cn)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%d/wallet/create", ctd.WebInterfacePort), strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-CSRF-Token", getCsrfToken(cn))

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

// CreateDistributionWallets creates 100 wallets from the initial distribution seeds
// walllet files have to be renamed into format like this: distribiton001.wlt
func CreateDistributionWallets(cn string, seeds []string) []string {

	user, _ := user.Current()
	wallets := make([]string, 100)

	for i := 0; i < 100; i++ {
		wlt := CreateWallet(cn, fmt.Sprintf("distribution%03d", i), seeds[i])
		f := filepath.Join(user.HomeDir, fmt.Sprintf("/.%s/wallets/%s", cn, wlt))
		t := filepath.Join(user.HomeDir, fmt.Sprintf("/.%s/wallets/distribution%03d.wlt", cn, i))
		err := os.Rename(f, t)
		if err != nil {
			panic(err)
		}
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

	fn = fn[0 : len(fn)-1]

	return fn
}
