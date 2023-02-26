package usecases

import (
	"github.com/pressus/config"
)

type Service interface {
	GetData() []string
	GetEnv() *config.Env
}

type service struct {
	env *config.Env
}

func NewService(env *config.Env) Service {
	return &service{env}
}

func (s *service) GetEnv() *config.Env {
	return s.env
}
