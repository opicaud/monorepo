package commands

import "example2/domain/provider/aggregate"

type AreaCommand struct {
	shape aggregate.Shape
}

func (c AreaCommand) Execute() error {
	err := c.shape.Area()
	if err != nil {
		return err
	}
	return nil
}
