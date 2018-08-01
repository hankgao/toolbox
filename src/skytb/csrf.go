package skytb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CsrfValue is something
type CsrfValue struct {
	CsrfToken string `json:"csrf_token"`
}

func getCsrfToken(coinName string) string {
	csrf := CsrfValue{}
	ctd := CoinTypeDetails(coinName)
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%4d/api/v1/csrf", ctd.WebInterfacePort))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(data, &csrf)

	return csrf.CsrfToken

}
