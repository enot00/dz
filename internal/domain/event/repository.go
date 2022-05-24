package event

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `test_base`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: `320201`,
}

//
type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(events Event) (*Event, error)
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	eventsTable := sess.Collection("events")
	var events []Event
	err = eventsTable.Find().All(&events)
	if err != nil {
		log.Fatal("eventsTable.Find: ", err)
	}
	for i := range events {
		fmt.Printf("record #%d: %#v\n", i, events[i])
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	eventsTable := sess.Collection("events")
	var event Event
	res := eventsTable.Find(id)
	err = res.One(&event)
	if err != nil {
		log.Fatal("eventsTable.Find: ", err)
	}
	return &event, nil
}
func (r *repository) Create(events Event) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	_, err = sess.SQL().InsertInto("events").Values(&events).Exec()
	if err != nil {
		log.Fatal("eventsTable.Create: ", err)
	}
	return &events, nil
}
