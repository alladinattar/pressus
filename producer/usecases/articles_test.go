package usecases

import (
	"github.com/pressus/config"
	"testing"
)

func BenchmarkService_GetArticlesByFlow(b *testing.B) {

	env := config.NewEnv()
	env.Config.Api.Port = ":3000"
	env.Config.Parser.DefaultRoute = "https://journal.tinkoff.ru"
	service_obj := NewService(env)
	for i := 0; i < 3; i++ {
		service_obj.GetArticlesByFlow("emigration-all")
	}
}
