package usecases

import (
	"github.com/pressus/config"
	"github.com/pressus/models/presenters"
	"github.com/pressus/repository/queue"
	"github.com/pressus/repository/search-engine"
)

type Service interface {
	GetFlows() ([]presenters.FlowObj, error)
	GetArticlesByFlow(flow string) ([]presenters.ArticleLink, error)
	SaveArticlesByFlow(flow string) (bool, error)
	ProcessLinks()
	ProcessLinksFromResultQueue()
	GetEnv() *config.Env
}

type service struct {
	env          *config.Env
	repo         queue.QueueRepo
	searchEngine search.EngineRepo
}

func NewService(env *config.Env, repo queue.QueueRepo, searchEngine search.EngineRepo) Service {
	return &service{env, repo, searchEngine}
}

func (s *service) GetEnv() *config.Env {
	return s.env
}
