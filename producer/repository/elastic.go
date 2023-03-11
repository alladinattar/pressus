package repository

import (
	"github.com/pressus/config"
)

type DBRepo interface {
	PushLinkForParsing() error
}

type dbRepo struct {
	env *config.Env
}

func (d *dbRepo) PushLinkForParsing() error {
	panic("implement me")
}

func NewDBRepo(env *config.Env) DBRepo {
	return &dbRepo{env}
}

func (d *dbRepo) GetEnv() *config.Env {
	return d.env
}
