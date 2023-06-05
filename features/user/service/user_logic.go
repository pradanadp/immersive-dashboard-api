package service

import (
	userRepo "alta-immersive-dashboard/features/user/repository"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository userRepo.UserRepository
	validate       *validator.Validate
}

// CreateUser implements UserService.
func (us *userService) CreateUser(user userRepo.UserEntity) (uint, error) {
	if user.FullName == "" {
		return 0, errors.New("error, name is required")
	}
	if user.Email == "" {
		return 0, errors.New("error, email is required")
	}
	if user.Password == "" {
		return 0, errors.New("error, password is required")
	}

	userID, err := us.userRepository.Insert(user)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return userID, nil
}

// DeleteUser implements UserService.
func (us *userService) DeleteUser(userID uint) error {
	panic("unimplemented")
}

// GetAllUser implements UserService.
func (us *userService) GetAllUser() ([]userRepo.UserEntity, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (us *userService) GetUser(userID uint) (userRepo.UserEntity, error) {
	userEntity, err := us.userRepository.Select(userID)
	if err != nil {
		return userRepo.UserEntity{}, fmt.Errorf("error: %v", err)
	}
	return userEntity, nil
}

// Login implements UserService.
func (us *userService) Login(email string, password string) (userRepo.UserEntity, string, error) {
	if email == "" {
		return userRepo.UserEntity{}, "", errors.New("email is required")
	} else if password == "" {
		return userRepo.UserEntity{}, "", errors.New("password is required")
	}

	loggedInUser, accessToken, err := us.userRepository.Login(email, password)
	if err != nil {
		return userRepo.UserEntity{}, "", err
	}

	return loggedInUser, accessToken, nil
}

// UpdateUser implements UserService.
func (us *userService) UpdateUser(userID uint, updatedUser userRepo.UserEntity) error {
	panic("unimplemented")
}

func New(repo userRepo.UserRepository) UserService {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}