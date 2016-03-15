package main

import (
	"fmt"
	"github.com/joonazan/tykki"
)

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

type Field int

func (f Field) Contains(pos tykki.Pos) bool {
	return abs(pos.X) <= int(f) && abs(pos.Y) <= int(f) && abs(pos.X+pos.Y) <= int(f)
}

func getRing(radius int) []tykki.Pos {
	if radius == 0 {
		return []tykki.Pos{{0, 0}}
	}
	results := []tykki.Pos{}
	dx := -radius
	dy := radius

	addRes := func() {
		results = append(results,
			tykki.Pos{dx, dy},
			tykki.Pos{-dx, -dy},
		)
	}
	addRes()
	for j := 0; j < 3; j++ {
		hops := radius
		if j == 2 {
			hops -= 1
		}
		for i := 0; i < hops; i++ {
			if j <= 1 {
				dx += 1
			}
			if j >= 1 {
				dy -= 1
			}
			addRes()
		}
	}
	return results
}

func main() {

	var field Field

	tykki.Run("NTN",
		func(config tykki.Config) {
			field = Field(config.FieldRadius)
		},
		func(bots []tykki.Bot, events []tykki.Event) []tykki.Action {
			for _, e := range events {
				fmt.Println(e)
			}
			fmt.Println()

			actions := make([]tykki.Action, len(bots))
			for i, b := range bots {
				actions[i] = b.Move(tykki.Pos{1, 0})
			}
			return actions
		},
	)
}
