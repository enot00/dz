package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dz_server/internal/domain/event"
	"github.com/go-chi/chi/v5"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}
		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

//роутер вызывает эту функцию
func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}
		event, err := (*c.service).Create(body) //вызываем в сервисе
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}
		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		event, err := (*c.service).Update(id, body) //вызываем в сервисе
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
}

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}
		event, err := (*c.service).Delete(id)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}
		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}
}
