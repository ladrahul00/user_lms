package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users typ definition
type Users struct {
	FirstName      string             `bson:"firstName,omitempty"`
	LastName       string             `bson:"lastName,omitempty"`
	Email          string             `bson:"email,omitempty"`
	Contact        string             `bson:"contact,omitempty"`
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedOn      time.Time          `bson:"createdOn,omitempty"`
	Active         bool               `bson:"active,omitempty"`
	OrganizationID primitive.ObjectID `bson:"organizationId,omitempty"`
}
