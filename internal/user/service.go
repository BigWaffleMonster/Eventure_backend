package user

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/helpers"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService interface {
	GetByID(id uuid.UUID) (*UserGetView, error)
	Update(id uuid.UUID, data *UserUpdateInput) error
	Remove(id uuid.UUID) error
}

type userService struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) GetByID(id uuid.UUID) (*UserGetView, error) {
	var userView UserGetView

	data, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	copier.Copy(&userView, &data)

	return &userView, nil
}

func (s *userService) Update(id uuid.UUID, data *UserUpdateInput) error {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}

	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.UserName != nil {
		user.UserName = *data.UserName
	}
	if data.IsEmailConfirmed != nil {
		user.IsEmailConfirmed = *data.IsEmailConfirmed
	}

	if data.Password != nil {
		password, err := helpers.HashPassword(*data.Password)
		if err != nil {
			return errors.New("error with saving password")
		}

		user.Password = password
	}

	err = s.Repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Remove(id uuid.UUID) error {
	err := s.Repo.Remove(id)
	if err != nil {
		return err
	}

	return nil
}
