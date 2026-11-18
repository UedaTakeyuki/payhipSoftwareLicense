package license

import (
	"time"

	"resty.dev/v3"
)

// refer https://payhip.com/product/edit/tRXlM
var productSecretKey string

// Struct for mapping the response from Payhip API.
type PayhipResponse struct {
	Data struct {
		Enabled     bool      `json:"enabled"`
		ProductLink string    `json:"product_link"`
		LicenseKey  string    `json:"license_key"`
		BuyerEmail  string    `json:"buyer_email"`
		Uses        int       `json:"uses"`
		Date        time.Time `json:"date"`
	} `json:"data"`
}

var client = resty.New().SetTimeout(10 * time.Second)

func SetProcuctSecret(s string) {
	productSecretKey = s
}
