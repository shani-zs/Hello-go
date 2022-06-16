package Questions

import "fmt"

type MathHelper interface {
	Count() int
}
type Calculator struct {
	model string
}

func (c Calculator) Count() string {
	return "Counting is fun"
}
func CountStuff(m MathHelper) {
	fmt.Println(m.Count())
}
