package main

import (
	"fmt"
	"totp"
)

func main() {

	secret := string("put OTP Secret here")
	otp := totp.GetTOTPToken(secret)

	fmt.Println(otp)

}
