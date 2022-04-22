package main

import (
	"GoFrontierReferral/entity"
	"GoFrontierReferral/randomNumber"
	"GoFrontierReferral/repository"
	"GoFrontierReferral/serialNumber"
	"GoFrontierReferral/urlGenerate"
	"GoFrontierReferral/utility"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	repo repository.ReferralRepository = repository.NewRepository()
)

func GenerateLink(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var referral entity.Referral
	// err := json.NewDecoder(request.Body).Decode(&referral)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{"error": "Error Generating code"}`))
	// 	return
	// }

	rand.Seed(time.Now().UnixNano())
	referral_code := randomNumber.RandomString(6) // Referral Code
	number := serialNumber.SerialNumber("dummy")  // Generate serial number
	serial_number := utility.TrimString(number)   // Trim the string

	referral.ReferralCode = referral_code // Assiging referral code
	referral.SerialNumber = serial_number // Assiging serial number
	repo.Save(&referral)                  // Save all data to firestore
	response.WriteHeader(http.StatusOK)   // Send response
	json.NewEncoder(response).Encode(referral)

	// Referral link generation
	if referral_code != "" {
		urlGenerate.Link(referral_code)
	} else {
		fmt.Println("Referral code is empty")
	}
}
