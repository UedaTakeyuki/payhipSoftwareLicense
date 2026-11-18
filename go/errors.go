//go:build !no_license_check

package license

import (
	"errors"
	"log"

	"resty.dev/v3"
)

// Define custom errors so callers can easily determine the cause of an error.
var (
	ErrhostUnreachable = errors.New("host unreachable")
	ErrAuthFailed      = errors.New("authentication failed")
	ErrRateLimited     = errors.New("rate limited")
	ErrServerError     = errors.New("server error")
	ErrBadLicenseKey   = errors.New("bad license key")
	ErrInvalidLicense  = errors.New("invalid license key")
	ErrDisabledLicense = errors.New("license key is disabled")
)

func ConfirmResponse(resp *resty.Response, err error, payhipResult PayhipResponse) (result error) {
	if err != nil {
		log.Printf("Request Error: %v", err)
		return
	}

	// 1. Check for HTTP error status codes (400+) using IsStatusFailure().
	if resp.IsStatusFailure() {
		statusCode := resp.StatusCode()

		switch statusCode {
		case 401, 403:
			result = ErrAuthFailed
			log.Printf("❌ Authentication error（HTTP %d）: Incorrect 'productSecretKey'\n", statusCode)
		case 429:
			result = ErrRateLimited
			log.Println("⚠️ Rate limit exceeded: Too many requests. Please wait a moment.")
		default:
			if statusCode >= 500 {
				result = ErrServerError
				log.Printf("⚠️ Server error（HTTP %d）:Payhip's system may be temporarily down.\n", statusCode)
			} else {
				result = ErrBadLicenseKey
				log.Printf("❌ Bad license key（HTTP %d）: Invalid license key.\n", statusCode)
			}
		}
		return
	}

	// 2. Verify Payhip response (returns {} when the license key is invalid)
	if payhipResult.Data.LicenseKey == "" {
		result = ErrInvalidLicense
		log.Println("❌ Invalid license key.（Payhip returns {}}）")
		return
	}

	// 3. Check if the existing key is disabled.
	if !payhipResult.Data.Enabled {
		result = ErrDisabledLicense
		log.Println("❌ This license key has been disabled.")
		return
	}

	return
}
