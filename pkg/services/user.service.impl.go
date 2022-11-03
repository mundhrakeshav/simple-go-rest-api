package services

import (
	"context"
	"errors"
	"mundhrakeshav/go-http/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) (int64, error)
	DeleteUser(*string) (int64, error)
}

const userNameKey = "user_name"
const userAgeKey = "user_age"
const userAddrKey = "user_address"

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	err := u.usercollection.FindOne(u.ctx, bson.D{bson.E{
		Key: userNameKey, Value: *name,
	}}).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) UpdateUser(user *models.User) (int64, error) {
	filter := bson.D{primitive.E{Key: userNameKey, Value: user.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: userNameKey, Value: user.Name}, primitive.E{Key: userAgeKey, Value: user.Age}, primitive.E{Key: userAddrKey, Value: user.Address}}}}
	result, err := u.usercollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if result.MatchedCount != 1 {
		return 0, errors.New("no matched document found for update")
	}
	return result.ModifiedCount, nil

}

func (u *UserServiceImpl) DeleteUser(name *string) (int64, error) {
	filter := bson.D{primitive.E{Key: userNameKey, Value: name}}
	result, err := u.usercollection.DeleteOne(u.ctx, filter)
	return result.DeletedCount, err
}
