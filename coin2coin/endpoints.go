package main

import (
	"net/http"
	"net/url"
)

// APIEndpoint represents an endpoint
type APIEndpoint struct {
	HTTPMethod     string   `json:"http_method"` // GET or PUT
	Path           string   `json:"explorer_path"`
	SkycoinPath    string   `json:"skycoin_path"` // Not useed for now, just in case in the future
	QueryArgs      []string `json:"query_args,omitempty"`
	Description    string   `json:"description"`
	ExampleRequest string   `json:"example_request"`
	// This string will be parsed into a map[string]interface{} in order to render newlines
	ExampleResponse string `json:"-"`
}

// OrderRequest represents an order sent from the web page
type OrderRequest struct {
	CoinFrom    string   `json:"coin_from"` // mzcoin, system will generate an address of this type
	CoinTo      string   `json:"coin_to"`   // shellcoin
	MailingInfo struct { // optional, if buyer doenst' give coin address, then mailing address must be provided
		BuyerName    string `json:"buyer_name"` //
		BuyerPhone   string `json:"buyer_phone"`
		BuyerAddress string `json:"buyer_address"`
	}
	BuyerCoinAddress string `json:"buyer_coin_address"` // optional, if not provided, a mailing address must be provided
}

// BindingAddress represents the response of /api/order endpoint
// address will displayed on the webpage both in literal format and QR code
type BindingAddress struct {
	CoinType    string `json:"coin_type"`
	CoinAddress string `json:"coin_address"`
}

var apiEndpoints = []APIEndpoint{
	{
		HTTPMethod:     "GET",
		Path:           "/api/coinpair",
		SkycoinPath:    "",
		Description:    "Returns all effecive coin pairs",
		ExampleRequest: "/api/coinpair",
		ExampleResponse: `{
			"coin_pairs": [
					{
						"from_coin": "shellcoin",
						"to_coin": "mzcoin",
						"exchange_rate" : 1.01
					}
				]
		}`,
	},

	{
		HTTPMethod:     "PUT",
		Path:           "/api/order",
		SkycoinPath:    "",
		Description:    "Check validity of an order and returns a binding address or error code",
		ExampleRequest: "/api/order",
		ExampleResponse: `{
				"coin_type": "shellcoin",
				"coin_address": "1dweirsdf213423423aekfkandfqewr",
		}`,
	},
}

func (s APIEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query url.Values
	if s.QueryArgs != nil {
		query = url.Values{}
		for _, s := range s.QueryArgs {
			query.Add(s, r.URL.Query().Get(s))
		}
	}

}
