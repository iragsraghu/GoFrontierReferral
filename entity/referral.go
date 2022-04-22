package entity

// Post data structure
type Referral struct {
	ReferralCode string `json:"referal_code"`
	SerialNumber string `json:"serial_number.omitempty"`
}
