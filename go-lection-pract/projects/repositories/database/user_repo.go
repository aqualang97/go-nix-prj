package database

import "my-uuid/repositories/models"

type UserDBRepository struct {
}

func (udbr UserDBRepository) GetByEmail(email string) (user models.User) {
	panic("implemented me")
}
