package project

import "gorm.io/gorm"

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}
