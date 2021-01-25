package user

import (
	"context"
	"demo/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	resource   *database.Resource
	collection *mongo.Collection
}

type Repository interface {
	Create(UserRequest) (User, error)
	Update(UpdateUserRequest, string) error
	GetAll() (Users, error)
	GetByID(string2 string) (*User, error)
	Login(LoginRequest) (*User, error)
}

func newRepoInstance(resource *database.Resource) Repository {
	collection := resource.DB.Collection("user")
	repository := &UserRepository{resource: resource, collection: collection}
	return repository
}

func (ur *UserRepository) GetByID(id string) (*User, error) {
	var user User

	ctx, cancel := initContext()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	err := ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Login(loginRequest LoginRequest) (*User, error) {
	var user User

	ctx, cancel := initContext()
	defer cancel()

	err := ur.collection.FindOne(ctx, bson.M{"email": loginRequest.Email, "password": loginRequest.Password}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetAll() (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Create(userRequest UserRequest) (User, error) {
	user := User{
		Id:   primitive.NewObjectID(),
		Name: userRequest.Name,
		Age:  userRequest.Age,
	}
	ctx, cancel := initContext()
	defer cancel()
	_, err := ur.collection.InsertOne(ctx, user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Update(request UpdateUserRequest, id string) error {
	newUser := bson.M{
		"name": request.Name,
		"age":  request.Age,
	}
	ctx, cancel := initContext()
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ur.collection.UpdateOne(ctx, bson.M{"_id": objID},newUser)
	if err != nil {
		return err
	}
	return nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
