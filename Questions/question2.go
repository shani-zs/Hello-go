package Questions

type Cat struct {
	Color string
}

func (c Cat) ChangeColor() {
	c.Color = "pink"
}
