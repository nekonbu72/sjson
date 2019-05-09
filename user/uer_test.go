package user

import (
	"log"
	"testing"
)

const (
	id   string = "1"
	name string = "Bob"
)

func TestNewJSON(t *testing.T) {
	j, err := newJOSN()
	if err != nil {
		t.Errorf("newUser: %v\n", err)
	}

	if j.User.Id != id {
		t.Errorf("id: %v\n", j.User.Id)
	}

	if j.User.Name != name {
		t.Errorf("name: %v\n", j.User.Name)
	}
	log.Printf("%+v\n", j)
}
