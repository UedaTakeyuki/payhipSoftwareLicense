package license

import (
	"log"
)

func ShowVerify(licenseKey string) {
	payhipResult, err := Verify(licenseKey)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Verify response", payhipResult)
	}
}

func Verify(licenseKey string) (payhipResult PayhipResponse, err error) {
	const url = "https://payhip.com/api/v2/license/verify"

	resp, err := client.R().
		SetHeader("product-secret-key", productSecretKey).
		SetQueryParam("license_key", licenseKey).
		SetResult(&payhipResult).
		Get(url)
	err = confirmResponse(resp, err, payhipResult)
	return
}
