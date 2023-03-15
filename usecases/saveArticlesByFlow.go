package usecases

func (s *service) SaveArticlesByFlow(flow string) (bool, error) {
	articles, err := s.GetArticlesByFlow(flow)
	if err != nil {
		return false, err
	}

	s.repo.SaveLinksToArticles(articles)
	if err != nil {
		return false, err
	}
	return true, nil
}
