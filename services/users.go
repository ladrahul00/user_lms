package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/wolf00/user_lms/constants"
	"github.com/wolf00/user_lms/db"
	"github.com/wolf00/user_lms/db/models"
	user "github.com/wolf00/user_lms/proto/user"
	"github.com/wolf00/user_lms/utilities"
)

// UserService blah
type UserService struct {
}

// Create a new user
func (u *UserService) Create(ctx context.Context, req *user.NewUserRequest, rsp *user.UserResponse) error {
	validationErrors := u.userValidationErrors(req)
	if len(validationErrors) > 0 {
		return u.failure(strings.Join(validationErrors, ", "), rsp)
	}
	organization, err := organizationByFilter(ctx, bson.M{"name": req.OrganizationName})

	if err != nil {
		return u.failure("organization doesn't exist", rsp)
	}

	newUser := models.Users{
		Email:          req.Email,
		Contact:        req.Contact,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		OrganizationID: organization.ID,
		Password:       req.Password,
	}
	_, err = createUser(ctx, newUser)
	if err != nil {
		rsp.Message = "failed to create user"
		rsp.Status = false
	} else {
		rsp.Message = "user created successfully"
		rsp.Status = true
	}

	return err
}

// GetByID a new user
func (u *UserService) GetByID(ctx context.Context, req *user.UserByIdRequest, rsp *user.UserDetails) error {
	user, err := userByID(ctx, req.UserId)
	if err != nil {
		return err
	}
	u.userToUserDetail(user, rsp)
	return nil
}

// GetByEmail a new user
func (u *UserService) GetByEmail(ctx context.Context, req *user.UserByEmailRequest, rsp *user.UserDetails) error {
	if !utilities.ValidateEmail(req.Email) {
		return fmt.Errorf(constants.InvalidEmailFormat)
	}
	filter := bson.M{"email": req.Email}
	user, err := userByFilter(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}
	u.userToUserDetail(user, rsp)
	return nil
}

// Update a user
func (u *UserService) Update(ctx context.Context, req *user.UpdateUserRequest, rsp *user.UserResponse) error {
	existingUser, err := userByID(ctx, req.Id)
	if err != nil {
		return err
	}
	if utilities.ValidateString(req.FirstName) {
		existingUser.FirstName = req.FirstName
	}
	if utilities.ValidateString(req.LastName) {
		existingUser.LastName = req.LastName
	}
	if utilities.ValidateString(req.Contact) {
		existingUser.Contact = req.Contact
	}
	if utilities.ValidateString(req.Password) {
		existingUser.Password = req.Password
	}
	existingUser.Active = req.Active
	updateData := bson.D{
		{"$set", bson.D{
			{"firstName", existingUser.FirstName},
			{"lastName", existingUser.LastName},
			{"contact", existingUser.Contact},
			{"active", existingUser.Active},
			{"password", existingUser.Password},
		}},
	}
	_, err = updateUser(ctx, existingUser.ID, updateData)
	if err != nil {
		rsp.Message = "failed to update the user"
		rsp.Status = false
		return err
	}
	rsp.Message = "user updated successfully"
	rsp.Status = true
	return nil
}

func (u *UserService) userToUserDetail(dbUser models.Users, userDetails *user.UserDetails) {
	userDetails.Contact = dbUser.Contact
	userDetails.Email = dbUser.Email
	userDetails.FirstName = dbUser.FirstName
	userDetails.LastName = dbUser.LastName
	userDetails.OrganizationId = dbUser.OrganizationID.Hex()
	userDetails.Id = dbUser.ID.Hex()
	userDetails.Password = dbUser.Password
}

func (u *UserService) failure(message string, rsp *user.UserResponse) error {
	rsp.Message = message
	rsp.Status = false
	return nil
}

func (u *UserService) userValidationErrors(newUser *user.NewUserRequest) []string {
	validationErrors := []string{}
	if !utilities.ValidateString(newUser.Contact) {
		validationErrors = append(validationErrors, constants.EmptyContact)
	}
	if !utilities.ValidateString(newUser.FirstName) {
		validationErrors = append(validationErrors, constants.EmptyFirstName)
	}
	if !utilities.ValidateString(newUser.OrganizationName) {
		validationErrors = append(validationErrors, "organization name cannot be empty")
	}
	if !utilities.ValidateString(newUser.Email) {
		validationErrors = append(validationErrors, constants.EmptyEmail)
	} else {
		if !utilities.ValidateEmail(newUser.Email) {
			validationErrors = append(validationErrors, constants.InvalidEmailFormat)
		}
	}
	return validationErrors
}

func createUser(ctx context.Context, newUser models.Users) (*mongo.InsertOneResult, error) {
	newUser.Active = true
	newUser.CreatedOn = time.Now()
	helper := db.Users(ctx)
	return helper.InsertOne(ctx, newUser)
}

func userByID(ctx context.Context, id string) (models.Users, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	var user models.Users
	if err != nil {
		log.Error(err)
		return user, err
	}
	filter := bson.M{"_id": userID}
	user, err = userByFilter(ctx, filter)
	if err != nil {
		log.Error(err)
		return user, err
	}
	return user, nil
}

func userByFilter(ctx context.Context, filter interface{}) (models.Users, error) {
	helper := db.Users(ctx)
	var user models.Users
	err := helper.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func updateUser(ctx context.Context, userID primitive.ObjectID, updateFieldAndValues interface{}) (*mongo.UpdateResult, error) {
	helper := db.Users(ctx)
	return helper.UpdateOne(ctx, bson.M{"_id": userID}, updateFieldAndValues)
}
