package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id int, user User) (User, error)
	DeleteUserByID(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id int, user User) (User, error) {
	var existingUser User
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	r.db.Save(&existingUser)

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id int) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
