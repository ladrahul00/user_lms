package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sources typ definition
type Sources struct {
	SourceTag    string             `bson:"sourceTag,omitempty"`
	SystemSource bool               `bson:"systemSource,omitempty"`
	ID           primitive.ObjectID `bson:"_id,omitempty"`
}
