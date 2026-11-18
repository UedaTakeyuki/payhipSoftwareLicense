package license

import (
	"log"
)

func ShowDisable(licenseKey string) {
	payhipResult, err := Disable(licenseKey)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Disable response", payhipResult)
	}

}

func Disable(licenseKey string) (payhipResult PayhipResponse, err error) {
	const url = "https://payhip.com/api/v2/license/disable"

	resp, err := client.R().
		SetHeader("product-secret-key", productSecretKey).
		SetFormData(map[string]string{"license_key": licenseKey}).
		SetResult(&payhipResult).
		Put(url)
	err = confirmResponse(resp, err, payhipResult)
	return
}
