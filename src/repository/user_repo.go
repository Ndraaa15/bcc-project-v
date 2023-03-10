package repository

import (
	"bcc-project-v/src/entities"
	"bcc-project-v/src/model"
)

func (r *Repository) CreateUser(user *entities.User) error {
	if err := r.db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindUser(model *model.LoginUser) (*entities.User, error) {
	user := entities.User{}
	err := r.db.Debug().Where("username = ?", model.Username).Or("email = ?", model.Email).First(&user).Error
	return &user, err
}

func (r *Repository) FindUserByID(ID uint) (*entities.User, error) {
	user := entities.User{}
	err := r.db.Debug().First(&user, ID).Error
	return &user, err
}

func (r *Repository) UpdateUser(user *entities.User) error {
	err := r.db.Debug().Save(&user).Error
	return err
}
