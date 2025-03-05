package worker

import (
	"sync"
	"threadpool/task"
	worker "threadpool/workerManager"
)

type Pool struct {
	tasks   chan *task.Task
	results chan *task.Task
	workers []worker.Worker
	wg      sync.WaitGroup
}

func NewPool(tasksNum int, workers []worker.Worker) *Pool {
	pool := &Pool{
		tasks:   make(chan *task.Task, tasksNum),
		results: make(chan *task.Task, tasksNum),
		workers: workers,
	}

	for i := 0; i < len(workers); i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			pool.ProcessTask()
		}()
	}

	return pool
}

func (p *Pool) ProcessTask() {
	for task := range p.tasks {
		for _, worker := range p.workers {
			if worker.CanProcess(task) {
				worker.Process(task)
				p.results <- task
				break
			}
		}
	}
}

func (p *Pool) Close() {
	close(p.tasks)
	p.wg.Wait()
	close(p.results)
}

func (p *Pool) AddTask(task *task.Task) {
	p.tasks <- task
}

func (p *Pool) GetResults() chan *task.Task {
	return p.results
}
