package repository

import (
	"context"

	domain "github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/SethukumarJ/go-gin-clean-arch/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

type userDatabaseMongo struct {
	DB *mongo.Client
}

// Delete implements interfaces.UserRepository
func (*userDatabaseMongo) Delete(ctx context.Context, user domain.Users) error {
	panic("unimplemented")
}

// FindAll implements interfaces.UserRepository
func (*userDatabaseMongo) FindAll(ctx context.Context) ([]domain.Users, error) {
	panic("unimplemented")
}

// FindByID implements interfaces.UserRepository
func (*userDatabaseMongo) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

// Save implements interfaces.UserRepository
func (*userDatabaseMongo) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	panic("unimplemented")
}

func NewUserMongoRepository(DB *mongo.Client) interfaces.UserRepository {
	return &userDatabaseMongo{DB}
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
