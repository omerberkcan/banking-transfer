package api

import (
	"github.com/omerberkcan/banking-transfer/internal/service"
)

type Handler struct {
	Auth IAuthHandler
}

func NewHandler(s *service.Services) *Handler {
	return &Handler{
		Auth: authHandler{s.Auth},
	}
}
