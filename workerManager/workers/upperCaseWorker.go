package workers

import (
	"strconv"
	"strings"
	"threadpool/task"
	"unicode"
)

type UpperCaseWorker struct{}

func (ucw UpperCaseWorker) Process(task *task.Task) {
	task.Result = strings.ToUpper(task.Data)
}

func (ucw UpperCaseWorker) CanProcess(task *task.Task) bool {
	if _, err := strconv.Atoi(task.Data); err != nil {
		return !unicode.IsUpper([]rune(task.Data)[0])
	}
	return false
}
