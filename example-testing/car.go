package demo

import "errors"

//Car struct
type Car struct {
	Name  string
	Price float32
}

//SetName set car anme
func (c *Car) SetName(name string) string {
	if name != "" {
		c.Name = name
	}
	return c.Name
}

// new Object
func New(name string, price float32) (*Car, error) {
	if name == "" {
		return nil, errors.New("missing the name")
	}
	return &Car{
		Name:  name,
		Price: price,
	}, nil
}
