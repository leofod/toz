package service

import (
	repository "toz/pkg/url/repository"
)

// var list SingleList

type Url interface {
	Create(full string) (string, error)
	GetFull(short string) (full string, err error)
}

type ServiceUrl struct {
	Url
}

func NewServiceUrl(rep *repository.UrlRepository) *ServiceUrl {
	return &ServiceUrl{
		Url: NewUrlService(rep.Url, InitList(), "zzzzzzzzzzz"),
	}
}
