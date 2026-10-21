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

type UrlRepository struct {
	collection *mongo.Collection
}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{
		collection: config.DB.Collection("url"),
	}
}

func (r *UrlRepository) Create(url *models.Url) error {
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)

	defer cancel()

	_, err := r.collection.InsertOne(ctx, url)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.ErrUrlExists
		}
		return err
	}

	return nil
}

func (r *UrlRepository) FindByShortCode(code string) (*models.Url, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	defer cancel()

	var url models.Url

	err := r.collection.FindOne(
		ctx,
		bson.D{{
			Key:   "short_code",
			Value: code,
		}},
	).Decode(&url)

	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *UrlRepository) FindByUserID(userId bson.ObjectID) ([]models.Url, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)

	defer cancel()

	cursor, err := r.collection.Find(
		ctx,
		bson.D{{
			Key:   "user_id",
			Value: userId,
		}},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	urls := make([]models.Url, 0)

	if err := cursor.All(ctx, &urls); err != nil {
		return nil, err
	}

	return urls, nil
}

func (r *UrlRepository) DeleteByIDAndUserID(id, userId bson.ObjectID) error {
	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	defer cancel()

	result, err := r.collection.DeleteOne(
		ctx,
		bson.D{
			{
				Key:   "_id",
				Value: id,
			},
			{
				Key:   "user_id",
				Value: userId,
			},
		},
	)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *UrlRepository) IncrementClickCount(shortCode string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	result, err := r.collection.UpdateOne(
		ctx,
		bson.D{
			{Key: "short_code", Value: shortCode},
		},
		bson.D{{
			Key: "$inc", Value: bson.D{
				{Key: "clicks", Value: 1},
			},
		}},
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
