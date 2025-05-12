package service

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UserService struct {
    client *firestore.Client
}

func NewUserService(client *firestore.Client) *UserService {
    return &UserService{client: client}
}

// getUserDetails retrieves user details from Firestore
func (s *UserService) GetUserDetails(ctx context.Context, userID string) (map[string]interface{}, error) {
	doc, err := s.client.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		log.Printf("Failed to get user %s: %v", userID, err)
		return nil, err
	}
	return doc.Data(), nil
}

// UpdateUserDetails updates user details in Firestore
func (s *UserService) UpdateUserDetails(ctx context.Context, userID string, data map[string]interface{}) error {
    _, err := s.client.Collection("users").Doc(userID).Set(ctx, data, firestore.MergeAll)
    if err != nil {
        log.Printf("Failed to update user %s: %v", userID, err)
        return err
    }
    return nil
}

// GetAllUsers retrieves all users from Firestore
func (s *UserService) GetAllUsers(ctx context.Context) ([]map[string]interface{}, error) {
	iter := s.client.Collection("users").Documents(ctx)
	defer iter.Stop()

	var users []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Printf("Failed to iterate users: %v", err)
			return nil, err
		}
		users = append(users, doc.Data())
	}

	return users, nil
}