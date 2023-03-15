package usecases

import log "github.com/sirupsen/logrus"

func (s *service) SaveArticlesByFlow(flow string) (bool, error) {
	articles, err := s.GetArticlesByFlow(flow)
	if err != nil {
		return false, err
	}
	log.Info("Len: ", len(articles))
	for i, article := range articles {
		log.Infof("%d: %s", i, article.Link)
	}

	s.repo.SaveLinksToArticles(articles)
	if err != nil {
		return false, err
	}
	return true, nil
}
