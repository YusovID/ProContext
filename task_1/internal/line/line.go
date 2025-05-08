package line

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"task_1/internal/robot"
	"time"
)

type Line struct {
	Mutex          *sync.Mutex
	Robot1, Robot2 robot.Robot
	BlackCellIndex int
}

func NewLine(mu *sync.Mutex) *Line {
	return &Line{
		Mutex:  mu,
		Robot1: *robot.NewRobot(1),
		Robot2: *robot.NewRobot(2),
	}
}

func (l *Line) SetRandomIndexes(maxSizeOfField int) {
	l.Robot1.Index = rand.IntN(maxSizeOfField)
	l.Robot2.Index = rand.IntN(maxSizeOfField)

	for l.Robot1.Index == l.Robot2.Index || math.Abs(float64(l.Robot1.Index-l.Robot2.Index)) < 2 {
		l.Robot1.Index = rand.IntN(maxSizeOfField)
		l.Robot2.Index = rand.IntN(maxSizeOfField)
	}

	var max, min int

	if l.Robot1.Index > l.Robot2.Index {
		max = l.Robot1.Index
		min = l.Robot2.Index
	} else {
		max = l.Robot2.Index
		min = l.Robot1.Index
	}

	l.BlackCellIndex = rand.IntN(max-min-1) + (min + 1)
}

func (l *Line) Print(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			return
		default:
			var min, max int

			min = l.Robot1.Index
			if l.Robot2.Index < min {
				min = l.Robot2.Index
			}
			if l.BlackCellIndex < min {
				min = l.BlackCellIndex
			}

			max = l.Robot1.Index
			if l.Robot2.Index > max {
				max = l.Robot2.Index
			}
			if l.BlackCellIndex > max {
				max = l.BlackCellIndex
			}

			for i := min; i <= max; i++ {
				if i == l.BlackCellIndex {
					fmt.Print("*")
				} else if i == l.Robot1.Index {
					fmt.Print("1")
				} else if i == l.Robot2.Index {
					fmt.Print("2")
				} else {
					fmt.Print(" ")
				}
			}

			fmt.Print("\n")
		}
	}
}
