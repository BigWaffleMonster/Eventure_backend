package utils

import (
	"errors"

	t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserDataFromCtx(c *gin.Context) (*t.UserDataCtx, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return nil, errors.New("missing userID in ctx")
	}

	email, exists := c.Get("email")
	if !exists {
		return nil, errors.New("missing email in ctx")
	}

	login, exists := c.Get("login")
	if !exists {
		return nil, errors.New("missing login in ctx")
	}

	userDataCtx := t.UserDataCtx{
		UserID: userID.(uuid.UUID),
		Email:  email.(string),
		Login:  login.(string),
	}

	return &userDataCtx, nil
}
