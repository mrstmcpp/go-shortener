package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserIDFromContext(c *gin.Context) bson.ObjectID {
	userID, exists := c.Get("userID")
	if !exists {
		return bson.NilObjectID
	}

	id, ok := userID.(bson.ObjectID)

	if !ok {
		return bson.NilObjectID
	}

	return id
}
