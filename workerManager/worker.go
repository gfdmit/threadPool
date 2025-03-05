package worker

import (
	"threadpool/task"
)

type Worker interface {
	Process(task *task.Task)
	CanProcess(task *task.Task) bool
}
