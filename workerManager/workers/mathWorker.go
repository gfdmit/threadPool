package workers

import (
	"log"
	"strconv"
	"threadpool/task"
)

type MathWorker struct {
	Multiplier int
}

func (mw MathWorker) Process(task *task.Task) {
	num, err := strconv.Atoi(task.Data)
	if err != nil {
		log.Printf("error while converting data to int: %v\n", err)
		task.Result = err.Error()
	} else {
		temp := num * mw.Multiplier
		task.Result = strconv.Itoa(temp)
	}
}

func (mw MathWorker) CanProcess(task *task.Task) bool {
	if _, err := strconv.Atoi(task.Data); err != nil {
		return false
	}
	return true
}
