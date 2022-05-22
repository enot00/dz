package event

import (
	"fmt"
	// "log"
	// "github.com/upper/db/v4/adapter/postgresql"
)

// var settings = postgresql.ConnectionURL{
// 	Database: `test_base`,
// 	Host:     `127.0.0.1:54322`,
// 	User:     `enot00`,
// 	Password: `320201`,
// }

//
type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	// sess, err := postgresql.Open(settings)
	// if err != nil {
	// 	log.Fatal("Open: ", err)
	// }
	// defer sess.Close()
	// fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)
	events := make([]Event, EventsCount)
	for i := int64(0); i < EventsCount; i++ {
		events[i] = Event{
			Id:   i + 1,
			Name: fmt.Sprintf("Event #%d", i+1),
		}
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	if id <= EventsCount {
		return &Event{
			Id:   id,
			Name: fmt.Sprintf("Event #%d", id),
		}, nil
	} else {
		return nil, nil
	}
}
