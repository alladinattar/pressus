package usecases

import (
	"github.com/pressus/config"
	"github.com/pressus/models/presenters"
	"github.com/pressus/repository/queue"
	"github.com/pressus/repository/search-engine"
)

type Service interface {
	GetFlows() ([]presenters.FlowObj, error)
	GetArticlesByFlow(flow string) ([]presenters.ArticleObj, error)
	SaveArticlesByFlow(flow string) (bool, error)
	ProcessLinks()
	ProcessLinksFromResultQueue()
	SearchByFlowAndAuthor(flow, author string) ([]string, error)
	SearchByFlowAndDate(flow, from, until string) ([]string, error)
	GetEnv() *config.Env
}

type service struct {
	env          *config.Env
	repoTasks    queue.QueueRepo
	repoResult   queue.QueueRepo
	searchEngine search.EngineRepo
}

func NewService(env *config.Env, repoTasks queue.QueueRepo, repoResult queue.QueueRepo, searchEngine search.EngineRepo) Service {
	return &service{env, repoTasks, repoResult, searchEngine}
}

func (s *service) GetEnv() *config.Env {
	return s.env
}
