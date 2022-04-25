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

	rand.Seed(time.Now().UnixNano())
	referral_code := randomNumber.RandomString(6) // Referral Code
	number := serialNumber.SerialNumber("dummy")  // Generate serial number
	serial_number := utility.TrimString(number)   // Trim the string

	// Check if the serial number is already used
	serialNumbers, err := repo.FindAllSerialNumber()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error"}`))
		return
	}

	if contains(serialNumbers, serial_number) {
		response.WriteHeader(409) // data already exists
	} else {
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
}

// Getting all records from firestore
func ListReferrals(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var referrals []entity.Referral
	referrals, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error Generating code"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(referrals)
}

// Record exists or not
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
