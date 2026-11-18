package license

import (
	"log"
)

func ShowUsage(licenseKey string) {
	payhipResult, err := Usage(licenseKey)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Usage response", payhipResult)
	}

}

func Usage(licenseKey string) (payhipResult PayhipResponse, err error) {
	const url = "https://payhip.com/api/v2/license/usage"

	resp, err := client.R().
		SetHeader("product-secret-key", productSecretKey).
		SetFormData(map[string]string{"license_key": licenseKey}).
		SetResult(&payhipResult).
		Put(url)
	err = confirmResponse(resp, err, payhipResult)
	return
}
