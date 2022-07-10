package domain

import "fmt"

type Car struct {
	Brand string
	Model string
}

func (c *Car) String() string {
	return fmt.Sprintf("%s - %s", c.Brand, c.Model)
}

type CarService interface {
	GetCar() (*Car, error)
}
