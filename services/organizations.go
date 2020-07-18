package services

import (
	"context"
	"strings"
	"time"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/wolf00/user_lms/constants"
	"github.com/wolf00/user_lms/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wolf00/user_lms/db"
	"github.com/wolf00/user_lms/db/models"
	organization "github.com/wolf00/user_lms/proto/organization"
)

// OrganizationService blah
type OrganizationService struct {
}

// Create blah
func (o *OrganizationService) Create(ctx context.Context, req *organization.NewOrganizationRequest, rsp *organization.Response) error {
	validationErrors := o.organizationValidationErrors(req)
	if len(validationErrors) > 0 {
		rsp.Message = strings.Join(validationErrors, ", ")
		rsp.Status = false
		return nil
	}
	newOrganization := models.Organizations{}
	newOrganization.Email = req.Email
	newOrganization.Name = req.Name
	newOrganization.SourceTag = req.SourceTag
	newOrganization.Contact = req.Contact
	newOrganization.LogoLink = req.LogoLink
	err := createOrganization(ctx, newOrganization, req.Password)
	if err != nil {
		return err
	}
	rsp.Message = "organization created successfully"
	rsp.Status = true
	return nil
}

// GetByID blah
func (o *OrganizationService) GetByID(ctx context.Context, req *organization.OrganizationByIdRequest, rsp *organization.OrganizationDetails) error {
	organizationID, err := primitive.ObjectIDFromHex(req.OrganizationId)
	if err != nil {
		log.Error(err)
		return err
	}
	filter := bson.M{"_id": organizationID}
	dbOrganization, err := organizationByFilter(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.Id = dbOrganization.ID.Hex()
	rsp.Contact = dbOrganization.Contact
	rsp.Email = dbOrganization.Email
	rsp.LogoLink = dbOrganization.LogoLink
	rsp.Name = dbOrganization.Name
	rsp.SourceTag = dbOrganization.SourceTag
	allowedSources := []string{}
	for i := 0; i < len(dbOrganization.AllowedSources); i++ {
		sourceDetails, err := sourceByFilter(ctx, bson.M{"_id": dbOrganization.AllowedSources[i]})
		if err != nil {
			log.Error(err)
			return err
		}
		allowedSources = append(allowedSources, sourceDetails.SourceTag)
	}
	rsp.AllowedSources = allowedSources
	return nil
}

// Update blah
func (o *OrganizationService) Update(ctx context.Context, req *organization.UpdateOrganizationRequest, rsp *organization.Response) error {
	existingOrganization, err := organizationByID(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if utilities.ValidateString(req.Name) {
		existingOrganization.Name = req.Name
	}
	if utilities.ValidateString(req.Contact) {
		existingOrganization.Contact = req.Contact
	}
	if utilities.ValidateString(req.LogoLink) {
		existingOrganization.LogoLink = req.LogoLink
	}
	if len(req.AllowedSources) > 0 {
		sFilter := bson.M{
			"sourceTag": bson.M{
				"$in": req.AllowedSources,
			},
		}
		sources, err := sourcesByFilter(ctx, sFilter)
		if err != nil {
			log.Error(err)
			return err
		}
		sourceIds := []primitive.ObjectID{}
		for i := 0; i < len(sources); i++ {
			sourceIds = append(sourceIds, sources[i].ID)
		}
		existingOrganization.AllowedSources = sourceIds
	}
	updateData := bson.D{
		{"$set", bson.D{
			{"name", existingOrganization.Name},
			{"contact", existingOrganization.Contact},
			{"logoLink", existingOrganization.LogoLink},
			{"allowedSources", existingOrganization.AllowedSources},
		}},
	}
	_, err = updateOrganization(ctx, existingOrganization.ID, updateData)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.Message = "Organization updated successfully"
	rsp.Status = true
	return nil
}

func (o *OrganizationService) organizationValidationErrors(req *organization.NewOrganizationRequest) []string {
	validationErrors := []string{}
	if !utilities.ValidateString(req.Contact) {
		validationErrors = append(validationErrors, constants.EmptyContact)
	}
	if !utilities.ValidateString(req.Name) {
		validationErrors = append(validationErrors, "organization name cannot be empty")
	}
	if !utilities.ValidateString(req.SourceTag) {
		validationErrors = append(validationErrors, "source tag cannot be empty")
	}
	if !utilities.ValidateString(req.Email) {
		validationErrors = append(validationErrors, constants.EmptyEmail)
	} else {
		if !utilities.ValidateEmail(req.Email) {
			validationErrors = append(validationErrors, constants.InvalidEmailFormat)
		}
	}
	return validationErrors
}

func createOrganization(ctx context.Context, newOrganization models.Organizations, userPassword string) error {
	newOrganization.CreatedOn = time.Now()
	newOrganization.Active = true
	newOrganization.AllowedSources = []primitive.ObjectID{}
	dbClient := db.ConnectMongo(ctx)
	organizationCollection := dbClient.Collection(constants.Organizations)
	sourceCollection := dbClient.Collection(constants.Sources)
	userCollection := dbClient.Collection(constants.Users)

	return dbClient.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			log.Error(err)
			return err
		}
		organizationID, err := organizationCollection.InsertOne(sessionContext, newOrganization)
		if err != nil {
			log.Error(err)
			return err
		}
		newSource := models.Sources{
			SourceTag:    newOrganization.SourceTag,
			SystemSource: false,
		}
		_, err = sourceCollection.InsertOne(sessionContext, newSource)
		if err != nil {
			log.Error(err)
			sessionContext.AbortTransaction(sessionContext)
			return err
		}
		newUser := models.Users{
			Email:          newOrganization.Email,
			Contact:        newOrganization.Contact,
			FirstName:      "Jhon",
			LastName:       "Doe",
			Active:         true,
			CreatedOn:      time.Now(),
			OrganizationID: organizationID.InsertedID.(primitive.ObjectID),
			Password:       userPassword,
		}
		_, err = userCollection.InsertOne(sessionContext, newUser)
		if err != nil {
			log.Error(err)
			sessionContext.AbortTransaction(sessionContext)
			return err
		}
		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
}

func organizationByID(ctx context.Context, organizationID string) (models.Organizations, error) {
	var organizationDetails models.Organizations
	orgID, err := primitive.ObjectIDFromHex(organizationID)
	if err != nil {
		log.Error(err)
		return organizationDetails, err
	}
	return organizationByFilter(ctx, bson.M{"_id": orgID})
}

func organizationByFilter(ctx context.Context, filter interface{}) (models.Organizations, error) {
	helper := db.Organizations(ctx)
	var organization models.Organizations
	err := helper.FindOne(ctx, filter).Decode(&organization)

	if err != nil {
		return organization, err
	}

	return organization, nil
}

func updateOrganization(ctx context.Context, organizationID primitive.ObjectID, updateFieldAndValues interface{}) (*mongo.UpdateResult, error) {
	helper := db.Organizations(ctx)
	return helper.UpdateOne(ctx, bson.M{"_id": organizationID}, updateFieldAndValues)
}
