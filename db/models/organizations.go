package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organizations typ definition
type Organizations struct {
	Name           string               `bson:"name,omitempty"`
	Email          string               `bson:"email,omitempty"`
	Contact        string               `bson:"contact,omitempty"`
	LogoLink       string               `bson:"logoLink,omitempty"`
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedOn      time.Time            `bson:"createdOn,omitempty"`
	Active         bool                 `bson:"active,omitempty"`
	AllowedSources []primitive.ObjectID `bson:"allowedSources,omitempty"`
	SourceTag      string               `bson:"sourceTag,omitempty"`
}
