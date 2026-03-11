package user

import "github.com/google/uuid"

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uuid.UUID) (*UserResponse, error) {
	user_raw, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	user := UserResponse{
		ID:    user_raw.ID,
		Login: user_raw.Login,
		Email: user_raw.Email,
	}

	return &user, nil
}
