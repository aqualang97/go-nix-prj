package models

import "my-uuid/repositories/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) models.User
	Insert(user *models.User) error
}

type SupplierRepositoryInterface interface {
	GetAll() (suppliers []*models.Supplier)
	//...
}

//...
