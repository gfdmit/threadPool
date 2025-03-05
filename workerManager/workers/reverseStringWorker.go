package workers

import (
	"strconv"
	"threadpool/task"
)

type ReverseStringWorker struct{}

func (rsw ReverseStringWorker) Process(task *task.Task) {
	runes := []rune(task.Data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	task.Result = string(runes)
}

func (rsw ReverseStringWorker) CanProcess(task *task.Task) bool {
	if _, err := strconv.Atoi(task.Data); err != nil {
		return true
	}
	return false
}
