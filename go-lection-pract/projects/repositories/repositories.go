package repositories

import "my-uuid/repositories/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) models.User
}
