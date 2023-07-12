package container_test

import (
	"testing"

	"github.com/golobby/container/v3"
	"github.com/stretchr/testify/assert"
)

type Shape interface {
	SetArea(int)
	GetArea() int
}

type Circle struct {
	a int
}

func (c *Circle) SetArea(a int) {
	c.a = a
}

func (c Circle) GetArea() int {
	return c.a
}

type Database interface {
	Connect() bool
}

type MySQL struct{}

func (m MySQL) Connect() bool {
	return true
}

var instance = container.New()

func TestContainer_NamedSingleton(t *testing.T) {
	err := instance.NamedSingleton("theCircle", func() Shape {
		return &Circle{a: 13}
	})
	assert.NoError(t, err)

	var sh Shape
	err = instance.NamedResolve(&sh, "theCircle")
	assert.NoError(t, err)
	assert.Equal(t, sh.GetArea(), 13)
}

func TestContainer_Fill_Unexported_With_Struct_Pointer(t *testing.T) {
	err := instance.Singleton(func() Shape {
		return &Circle{a: 5}
	})
	assert.NoError(t, err)

	err = instance.Singleton(func() Database {
		return &MySQL{}
	})
	assert.NoError(t, err)

	myApp := struct {
		s Shape    `container:"type"`
		d Database `container:"type"`
		y int
	}{}

	err = instance.Fill(&myApp)
	assert.NoError(t, err)

	assert.IsType(t, &Circle{}, myApp.s)
	assert.IsType(t, &MySQL{}, myApp.d)
}