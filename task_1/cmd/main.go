package main

import (
	"context"
	"fmt"
	"sync"
	cmd "task_1/internal/commands"
	line "task_1/internal/line"
	"time"
)

const (
	maxSizeOfField = 5
	MR             = cmd.CommandType_MR
	ML             = cmd.CommandType_ML
	IF_FLAG        = cmd.CommandType_IF_FLAG
	GOTO           = cmd.CommandType_GOTO
)

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	ctx, cancel := context.WithCancel(context.Background())

	line := line.NewLine(mu)
	line.SetRandomIndexes(maxSizeOfField)

	program := []cmd.Command{
		{Type: ML},            // 0
		{Type: IF_FLAG},       // 1
		{Type: GOTO, Args: 4}, // 2
		{Type: GOTO, Args: 0}, // 3
		{Type: ML},            // 4
		{Type: GOTO, Args: 4}, // 5
	}

	wg.Add(4)

	go startRobot(ctx, 1, program, line, wg)

	go startRobot(ctx, 2, program, line, wg)

	go line.Print(ctx, wg)

	go func() {
		defer wg.Done()
		for {
			time.Sleep(100 * time.Millisecond)

			mu.Lock()
			met := line.Robot1.Index == line.Robot2.Index
			mu.Unlock()

			if met {
				fmt.Println("Роботы встретились")
				cancel()
				break
			}
		}
	}()

	wg.Wait()
}

func startRobot(ctx context.Context, robotID int, program []cmd.Command, line *line.Line, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(program); i++ {
		select {
		case <-ctx.Done():
			return
		default:
			command := program[i]

			switch command.Type {
			case ML:
				cmd.ML(robotID, line)
			case MR:
				cmd.MR(robotID, line)
			case IF_FLAG:
				if !cmd.IF_FLAG(robotID, line) {
					i++
				}
			case GOTO:
				i = cmd.GOTO(command.Args - 1)
			}
		}
	}
}
