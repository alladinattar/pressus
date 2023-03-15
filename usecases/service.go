package usecases

import (
	"github.com/pressus/config"
	"github.com/pressus/models/presenters"
	"github.com/pressus/repository"
)

type Service interface {
	GetFlows() ([]presenters.FlowObj, error)
	GetArticlesByFlow(flow string) ([]presenters.ArticleObj, error)
	SaveArticlesByFlow(flow string) (bool, error)
	GetEnv() *config.Env
}

type service struct {
	env  *config.Env
	repo repository.QueueRepo
}

func NewService(env *config.Env, repo repository.QueueRepo) Service {
	return &service{env, repo}
}

func (s *service) GetEnv() *config.Env {
	return s.env
}
