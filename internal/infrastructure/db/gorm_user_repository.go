package db

import (
	"realworld-go-gin/internal/domain/user"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) FindByEmail(email string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *GormUserRepository) Save(u *user.User) error {
	return r.db.Create(u).Error
}