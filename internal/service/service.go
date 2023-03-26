package service

import (
	puregrade "github.com/ZaiPeeKann/auth-service_pg/internal/models"
	"github.com/ZaiPeeKann/auth-service_pg/internal/repository"
)

type Authorization interface {
	CreateUser(puregrade.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// type User interface {
// 	GetProfile(id int) (puregrade.User, error)
// 	FollowUser(id int) error
// 	MuteUser(id int) error
// }

type Review interface {
	GetAll() ([]puregrade.Review, error)
	GetOneByID(id int) (puregrade.Review, error)
	Create(review puregrade.Review) (int, error)
	Update(id int, title, body string) error
	Delete(id, userId int) error
}

type Product interface {
	GetAll() ([]puregrade.Product, error)
	GetOneByID(id int) (puregrade.Product, error)
	Create(product puregrade.Product) (int, error)
	Delete(id int) error
}

type Service struct {
	Authorization
	// User
	Review
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		// User:          NewUserService(repos),
		Review:  NewReviewService(repos),
		Product: NewProductService(repos),
	}
}
