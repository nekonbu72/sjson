package user

import (
	"github.com/nekonbu72/sjson/sjson"
)

type JSON struct {
	User struct {
		Id   string `JSON:"id"`
		Name string `JSON:"name"`
	} `JSON:"user"`
}

const path string = "test.JSON"

func newJOSN() (*JSON, error) {
	j := new(JSON)
	// var u user
	err := sjson.OpenDecode(path, j)
	return j, err
}
