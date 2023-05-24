package service

import (
	"go-article/model"
	"go-article/repository"
)

type Service interface {
	GetAll() ([]model.Post, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]model.Post, error) {
	return s.repository.GetAll()
}
