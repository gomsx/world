package world

import "fmt"

type World struct{}

func (World) Print() string {
	s := "world v2.0.3\n"
	fmt.Println(s)
	return s
}
