package repositories

import (
	"context"
	"go-shortener/config"
	"go-shortener/models"
	"go-shortener/utils"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: config.DB.Collection("users"),
	}
}

func (r *UserRepository) Create(
	user *models.User,
) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if user.ID.IsZero() {
		user.ID = bson.NewObjectID()
	}

	_, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.ErrEmailExists
		}

		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	var user models.User

	err := r.collection.FindOne(
		ctx,
		bson.D{{Key: "email", Value: email}},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
