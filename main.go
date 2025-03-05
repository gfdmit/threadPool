package main

import (
	"fmt"
	"sync"
	"threadpool/task"
	worker "threadpool/workerManager"
	workerPool "threadpool/workerManager/workerPool"
	"threadpool/workerManager/workers"
)

func main() {
	tasksNum := 10
	workers := []worker.Worker{
		workers.UpperCaseWorker{},
		workers.ReverseStringWorker{},
		workers.MathWorker{Multiplier: 50},
	}
	workerPool := workerPool.NewPool(tasksNum, workers)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < tasksNum; i++ {
			task := &task.Task{ID: i, Data: fmt.Sprintf("%d", i)}
			workerPool.AddTask(task)
		}
	}()
	wg.Wait()
	workerPool.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range workerPool.GetResults() {
			fmt.Println(result.Result)
		}
	}()

	wg.Wait()
}
