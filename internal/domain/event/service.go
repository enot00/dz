package event

import (
	"encoding/json"
	"fmt"
)

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(body []byte) (*Event, error)
	Update(id int64, body []byte) (*Event, error)
	Delete(id int64) (*Event, error)
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

func (s *service) Create(body []byte) (*Event, error) {
	var event Event
	err := json.Unmarshal(body, &event)
	if err != nil {
		fmt.Printf("JSON create: %s", err)
	}
	return (*s.repo).Create(event) //вызываем в репозитории
}

func (s *service) Update(id int64, body []byte) (*Event, error) {
	var event Event
	err := json.Unmarshal(body, &event)
	if err != nil {
		fmt.Printf("JSON update: %s", err)
	}
	return (*s.repo).Update(id, event)
}

func (s *service) Delete(id int64) (*Event, error) {
	return (*s.repo).Delete(id)
}
