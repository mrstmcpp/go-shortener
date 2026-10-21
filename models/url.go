package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Url struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      bson.ObjectID `bson:"user_id" json:"-"`
	OriginalUrl string        `bson:"original_url" json:"originalUrl"`
	ShortCode   string        `bson:"short_code" json:"shortCode"`
	CreatedAt   time.Time     `bson:"created_at" json:"createdAt"`
	Clicks      int           `bson:"clicks" json:"clicks"`
}
