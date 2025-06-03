package user

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService interface {
	GetByID(id uuid.UUID) (*UserView, error)
	Update(id uuid.UUID, data *UserUpdateInput) error
	Remove(id uuid.UUID) error
}

type userService struct {
	Repo UserRepository
	DomainEventBus domain_events_abstractions.DomainEventBus
}

func NewUserService(repo UserRepository, eventBus domain_events_abstractions.DomainEventBus) UserService {
	return &userService{
		Repo: repo,
		DomainEventBus: eventBus,
	}
}

func (s *userService) GetByID(id uuid.UUID) (*UserView, error) {
	var userView UserView

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
	domainEventData, err := domain_events.NewUserDeletedDomainEvent(id)

	if err != nil {
		return err
	}

	return s.DomainEventBus.AddToStore(domainEventData)
}
