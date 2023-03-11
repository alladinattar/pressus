package repository

import (
	"github.com/pressus/config"
)

type DBRepo interface {
	GetUser() error
}

type dbRepo struct {
	env *config.Env
}

func (d *dbRepo) GetUser() error {
	//TODO implement me
	panic("implement me")
}

func NewDBRepo(env *config.Env) DBRepo {
	return &dbRepo{env}
}

func (d *dbRepo) GetEnv() *config.Env {
	return d.env
}
