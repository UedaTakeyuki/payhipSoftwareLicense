package license

import (
	"log"
)

func ShowEnable(licenseKey string) {
	payhipResult, err := Enable(licenseKey)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Enable response", payhipResult)
	}

}

func Enable(licenseKey string) (payhipResult PayhipResponse, err error) {
	const url = "https://payhip.com/api/v2/license/enable"

	resp, err := client.R().
		SetHeader("product-secret-key", productSecretKey).
		SetFormData(map[string]string{"license_key": licenseKey}).
		SetResult(&payhipResult).
		Put(url)
	err = confirmResponse(resp, err, payhipResult)
	return
}
