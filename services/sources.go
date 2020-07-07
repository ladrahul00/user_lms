package services

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/wolf00/golang_lms/user/db"
	"github.com/wolf00/golang_lms/user/db/models"
	source "github.com/wolf00/golang_lms/user/proto/source"
	"go.mongodb.org/mongo-driver/bson"
)

// SourceServce blah
type SourceServce struct{}

// GetBySourceTag is a single request handler called via client.Call or the generated client code
func (e *SourceServce) GetBySourceTag(ctx context.Context, req *source.SourceRequest, rsp *source.SourceResponse) error {
	filter := bson.M{"sourceTag": req.SourceTag}
	sourceDetails, err := sourceByFilter(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.SourceTag = sourceDetails.SourceTag
	rsp.SystemSource = sourceDetails.SystemSource
	return nil
}

func sourceByFilter(ctx context.Context, filter interface{}) (models.Sources, error) {
	helper := db.Sources(ctx)
	var source models.Sources
	err := helper.FindOne(ctx, filter).Decode(&source)
	if err != nil {
		log.Error(err)
		return source, err
	}
	return source, nil
}

func sourcesByFilter(ctx context.Context, filter interface{}) ([]models.Sources, error) {
	sources := []models.Sources{}

	helper := db.Sources(ctx)
	curr, err := helper.Find(ctx, filter)

	if err != nil {
		return sources, err
	}

	for curr.Next(ctx) {
		var source models.Sources
		err = curr.Decode(&source)
		if err != nil {
			return sources, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}
