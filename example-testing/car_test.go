package demo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNew(t *testing.T) {
// 	c, err := New("tesla", 100)
// 	if err != nil {
// 		t.Error("got errors", err)
// 	}

// 	if c == nil {
// 		t.Error("car should be nil")
// 	}
// }

func TestNewWithAssert(t *testing.T) {
	c, err := New("tesla", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	assert.Equal(t, "tesl", c.Name)
}

type User struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *User
		sv    reflect.Value
	)

	model = &User{}
	sv = reflect.ValueOf(model)
	sv = sv.Elem()

	t.Log("reflect.ValueOf", sv.Kind())

	sv.FieldByName("UserId").SetString("123")
	sv.FieldByName("Name").SetString("leighj")
}
