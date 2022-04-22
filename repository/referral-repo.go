package repository

import (
	"GoFrontierReferral/entity"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type ReferralRepository interface {
	Save(post *entity.Referral) (*entity.Referral, error)
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
