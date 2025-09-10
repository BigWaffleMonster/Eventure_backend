package user

import (
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_definitions"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService interface {
	GetByID(id uuid.UUID) (*UserView, results.Result)
	Update(id uuid.UUID, data *UserUpdateInput) results.Result
	Delete(id uuid.UUID) results.Result
}

type userService struct {
	Uof UnitOfWork
}

func NewUserService(uof UnitOfWork) UserService {
	return &userService{
		Uof: uof,
	}
}

func (s *userService) GetByID(id uuid.UUID) (*UserView, results.Result) {
	var userView UserView

	data, result := s.Uof.Repository().GetByID(id)
	if result.IsFailed {
		return nil, result
	}

	copier.Copy(&userView, &data)

	return &userView, results.NewResultOk()
}

func (s *userService) Update(id uuid.UUID, data *UserUpdateInput) results.Result {
	user, result := s.Uof.Repository().GetByID(id)
	if result.IsFailed {
		return result
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
			return results.NewUnauthorizedError("Login or password is invalid")
		}

		user.Password = password
	}

	return s.Uof.Repository().Update(user)
}

func (s *userService) Delete(id uuid.UUID) results.Result {
	domainEventData, result := domain_events_definitions.NewUserDeleted(id)

	if result.IsFailed {
		return result
	}

	return s.Uof.DomainEventStore().AddToStore(domainEventData)
}
