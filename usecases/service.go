package usecases

import (
	"github.com/pressus/config"
)

type Service interface {
	GetFlows() ([]string, error)
	GetArticlesByFlow(flow string) ([]string, error)
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
