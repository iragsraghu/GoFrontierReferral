package urlGenerate

import (
	"fmt"
	"net/url"
)

func Link(referral_code string) {
	base, err := url.Parse("http://www.example.com")
	if err != nil {
		return
	}

	// Path params
	base.Path += "referral"

	// Query params
	params := url.Values{}
	params.Add("q", referral_code)
	base.RawQuery = params.Encode()

	fmt.Printf("%q\n", base.String())
}
