package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

// ParseID parses ID from the request parameters
func ParseID(c *gin.Context) (bson.ObjectID, error) {
	ID := c.Param("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return bson.NilObjectID, errors.New("ID is required")
	}

	objID, _ := bson.ObjectIDFromHex(ID)

	return objID, nil
}
