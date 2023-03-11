package repository

import (
	"context"
	"fmt"

	domain "github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/SethukumarJ/go-gin-clean-arch/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDatabaseMongo struct {
	DB *mongo.Client
}

// Delete implements interfaces.UserRepository
func (*userDatabaseMongo) Delete(ctx context.Context, user domain.Users) error {
	panic("unimplemented")
}

// FindByID implements interfaces.UserRepository
func (*userDatabaseMongo) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

// Save implements interfaces.UserRepository
func (db *userDatabaseMongo) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Get the "users" collection.
	collection := db.DB.Database("mongo_demo").Collection("users")

	// Insert the user document.
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domain.Users{}, fmt.Errorf("error inserting user: %v", err)
	}

	// Get the ID of the inserted document and set it on the user.
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return domain.Users{}, fmt.Errorf("error getting inserted ID: %v", err)
	}
	fmt.Println("id", id)

	return user, nil
}

// FindAll implements interfaces.UserRepository
func (db *userDatabaseMongo) FindAll(ctx context.Context) ([]domain.UserResponse, error) {
	// Create a slice to hold the users.
	var users []domain.UserResponse

	collection := db.DB.Database("mongo_demo").Collection("users")
	// Retrieve all the documents from the "users" collection.
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding users: %v", err)
	}
	defer cursor.Close(ctx)

	// Loop through the documents and decode them into user structs.
// Loop through the documents and decode them into user structs.
for cursor.Next(ctx) {
    var user domain.UserResponse
    err := cursor.Decode(&user)
    if err != nil {
        return nil, fmt.Errorf("error decoding user: %v", err)
    }

    // Check if the ID field is present in the MongoDB document.
    if user.ID != nil {
       fmt.Println(user.ID,"/////////////")
    }

    users = append(users, user)
}

	// Check for any errors during the iteration.
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users cursor: %v", err)
	}

	return users, nil
}

func NewUserMongoRepository(DB *mongo.Client) interfaces.UserRepository {

	return &userDatabaseMongo{DB: DB}
}
