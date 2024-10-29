package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"auth-server/internal/auth/model"
	"auth-server/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService() *UserService {
	collection := database.GetMongoCollection("hi", "hi")
	return &UserService{collection: collection}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Retrieved all users successfully.")
	return users, nil
}
