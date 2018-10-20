package uadmin

import (
	"encoding/base32"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// getOTP is a function that generates TOTP using github.com/pquerna/otp
// Parametrers:
//   - seed: OTP seed is base32
//   - digits: the number of degits for the OTP
//   - algoriths: "sha1", "sha256", "sha512"
//   - skew: the number of minutes to search around the OTP
//   - period: the number of seconds for the OTP to change
func getOTP(seed string, digits int, algorithm string, skew uint, period uint) string {
	seed = strings.Replace(base32.StdEncoding.EncodeToString([]byte(seed)), "=", "", -1)
	var algo otp.Algorithm
	algorithm = strings.ToLower(algorithm)
	switch algorithm {
	case "sha1":
		algo = otp.AlgorithmSHA1
	case "sha256":
		algo = otp.AlgorithmSHA256
	case "sha512":
		algo = otp.AlgorithmSHA512
	default:
		Trail(ERROR, "getOTP: Unable to generate otp, unknown hash algorithms (%s)", algorithm)
		return ""
	}
	opts := totp.ValidateOpts{
		Algorithm: algo,
		Digits:    otp.Digits(digits),
		Skew:      skew,
		Period:    period,
	}

	pass, err := totp.GenerateCodeCustom(seed, time.Now().UTC(), opts)
	if err != nil {
		Trail(ERROR, "Unable to generate OTP. %s", err)
		return ""
	}
	return pass
}

func verifyOTP(pass, seed string, digits int, algorithm string, skew uint, period uint) bool {
	seed = strings.Replace(base32.StdEncoding.EncodeToString([]byte(seed)), "=", "", -1)
	var algo otp.Algorithm
	algorithm = strings.ToLower(algorithm)
	switch algorithm {
	case "sha1":
		algo = otp.AlgorithmSHA1
	case "sha256":
		algo = otp.AlgorithmSHA256
	case "sha512":
		algo = otp.AlgorithmSHA512
	default:
		Trail(ERROR, "getOTP: Unable to generate otp, unknown hash algorithms (%s)", algorithm)
		return false
	}
	opts := totp.ValidateOpts{
		Algorithm: algo,
		Digits:    otp.Digits(digits),
		Skew:      skew,
		Period:    period,
	}

	valid, err := totp.ValidateCustom(pass, seed, time.Now().UTC(), opts)
	if err != nil {
		Trail(ERROR, "Unable to verify OTP. %s", err)
		return false
	}
	return valid
}
