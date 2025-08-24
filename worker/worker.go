package worker

import (
	"log"
	"strconv"
)

type Worker struct {
	*log.Logger
	id     string
	signal chan int
}

func (w *Worker) Execute(task Task) {
	w.Printf("Starting task: %s\n", task.GetName())
	_, err := task.Run()
	if err != nil {
		w.Printf(
			"Failed to run task: %s\n",
			task.GetName(),
		)
	}
	i, err := strconv.Atoi(w.id)
	if err != nil {
		w.Panicln(err)
	}
	w.Printf("Finished task: %s\n", task.GetName())
	w.signal <- i
}
