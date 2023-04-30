package usecases

func (s *service) SearchByFlowAndAuthor(flow, author string) ([]string, error) {

	return s.searchEngine.FindByFlowAndAuthor(flow, author)

}
