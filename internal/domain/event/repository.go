package event

import (
	"fmt"

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
	Update(id int64, event Event) (*Event, error)
	Delete(id int64) (*Event, error)
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
		fmt.Printf("Open: %s", err)
		return nil, err
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	var events []Event
	err = sess.Collection("events").Find().All(&events)
	if err != nil {
		fmt.Printf("Find All: %s", err)
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Printf("Open: %s", err)
		return nil, err
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)
	var event Event
	err = sess.Collection("events").Find(id).One(&event)
	if err != nil {
		fmt.Printf("Find one: %s", err)
	}
	return &event, nil
}
func (r *repository) Create(event Event) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Printf("Open: %s", err)
		return nil, err
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	col := sess.Collection("events")

	res_id, err := col.Insert(&event)
	//справедливости ради в lesson7 мы не использовали Insert только InsertInto
	if err != nil {
		fmt.Printf("Create: %s", err)
	}
	var res Event
	err = col.Find(res_id).One(&res) //возращаем записанный евент с базы
	if err != nil {
		fmt.Printf("Create return: %s", err)
	}
	//возмоно правильно возращать db.InsertResult??
	return &res, nil
}

func (r *repository) Update(id int64, event Event) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Printf("Open: %s", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	col := sess.Collection("events")
	err = col.Find(id).Update(event)
	if err != nil {
		fmt.Printf("Update: %s", err)
	}
	var res Event
	err = col.Find(id).One(&res) //возращаем обновленный евент с базы
	if err != nil {
		fmt.Printf("Update return: %s", err)
	}
	return &res, nil
}

func (r *repository) Delete(id int64) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Printf("Open: %s", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	col := sess.Collection("events")
	var res Event
	err = col.Find(id).One(&res) //возращаем евент который мы удалили для вывода отчета или отмене(повторного Create())
	if err != nil {
		fmt.Printf("Delete return: %s", err)
	}
	err = col.Find(id).Delete()
	if err != nil {
		fmt.Printf("Delete: %s", err)
	}
	return &res, nil
}
