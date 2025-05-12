package model

import (
	"time"
)

type User struct {
	ID        string    `firestore:"id"`
	FirstName string    `firestore:"first_name"`
	LastName  string    `firestore:"last_name"`
	Email     string    `firestore:"email"`
	CreatedAt time.Time `firestore:"created_at"`
}