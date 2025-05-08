package command

import (
	"task_1/internal/line"
	"time"
)

const (
	CommandType_ML int = iota
	CommandType_MR
	CommandType_IF_FLAG
	CommandType_GOTO
)

type Command struct {
	Type int
	Args int
}

func ML(robotID int, l *line.Line) {
	time.Sleep(1 * time.Second)

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if robotID == 1 {
		l.Robot1.Index--
	} else {
		l.Robot2.Index--
	}
}

func MR(robotID int, l *line.Line) {
	time.Sleep(1 * time.Second)

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if robotID == 1 {
		l.Robot1.Index++
	} else {
		l.Robot2.Index++
	}
}

func IF_FLAG(robotID int, l *line.Line) bool {
	time.Sleep(1 * time.Second)

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if robotID == 1 {
		return l.Robot1.Index == l.BlackCellIndex
	} else {
		return l.Robot2.Index == l.BlackCellIndex
	}
}

func GOTO(n int) int {
	return n
}
