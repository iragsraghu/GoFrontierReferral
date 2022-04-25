package repository

import (
	"GoFrontierReferral/entity"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ReferralRepository interface {
	Save(post *entity.Referral) (*entity.Referral, error)
	FindAll() ([]entity.Referral, error)
	FindAllSerialNumber() ([]string, error)
}

type repo struct{}

// NewRepository
func NewRepository() ReferralRepository {
	return &repo{}
}

const (
	projectId      string = "gofrontier2"
	collectionName string = "referrals"
)

func (*repo) Save(referral *entity.Referral) (*entity.Referral, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ReferralCode": referral.ReferralCode,
		"SerialNumber": referral.SerialNumber,
	})

	if err != nil {
		log.Fatalf("Failed to adding a new post: %v", err)
		return nil, err
	}

	return referral, nil
}

func (*repo) FindAll() ([]entity.Referral, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	var referrals []entity.Referral
	var serialNumbers []string
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		referral := entity.Referral{
			ReferralCode: doc.Data()["ReferralCode"].(string),
			SerialNumber: doc.Data()["SerialNumber"].(string),
		}
		referrals = append(referrals, referral)
		serialNumbers = append(serialNumbers, referral.SerialNumber)
	}
	return referrals, nil
}

func (*repo) FindAllSerialNumber() ([]string, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	var serialNumbers []string
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		referral := entity.Referral{
			ReferralCode: doc.Data()["ReferralCode"].(string),
			SerialNumber: doc.Data()["SerialNumber"].(string),
		}
		serialNumbers = append(serialNumbers, referral.SerialNumber)
	}
	return serialNumbers, nil
}
