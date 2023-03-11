package config

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

type Env struct {
	SessionStore *session.Store
	Config       *Config
	DB           *sql.DB
}

func NewEnv() *Env {

	return &Env{
		SessionStore: session.New(session.Config{
			CookieHTTPOnly: true,
			Expiration:     time.Hour * 24 * 30,
		}),
		Config: &Config{},
		DB:     &sql.DB{},
	}
}
