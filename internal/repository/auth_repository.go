package repository

import (
	"mkpticket/internal/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Repository[entity.User]
	FindByEmail(db *gorm.DB, email string) (*entity.User, error)
}

type authRepositoryImpl struct {
	RepositoryImpl[entity.User]
}

func NewAuthRepositoryImpl() AuthRepository {
	return &authRepositoryImpl{}
}

func (a *authRepositoryImpl) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	return &user, db.Where("email = ?", email).Take(&user).Error
}
