package commands

import "example2/domain/aggregate"

type areaCommand struct {
	shape aggregate.Shape
}

func (c areaCommand) Execute() error {
	err := c.shape.Area()
	if err != nil {
		return err
	}
	return nil
}
