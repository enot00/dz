package event

import (
	"fmt"
	"net/url"
	"strconv"
)

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(params url.Values) (*Event, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Create(params url.Values) (*Event, error) {
	parse_price, err := strconv.Atoi(params["price"][0])
	if err != nil {
		fmt.Printf("ParseFail.Create(): %s", err)
	}
	event := Event{
		Master:  params["master"][0],
		Title:   params["title"][0],
		Descrip: params["descrip"][0],
		Price:   parse_price,
	}
	return (*s.repo).Create(event) //вызываем в репозитории
}
