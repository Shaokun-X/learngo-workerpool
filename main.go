package main

import (
	"bufio"
	"flag"
	"os"
	worker "shaokun-x/exercise/worker"
)

func main() {
	urls, concurrency := parse()
	pool := worker.NewWorkerPool(concurrency)

	tasks := []worker.Task{}
	for _, url := range urls {
		tasks = append(tasks, &worker.DownloadTask{Name: url, URL: url})
	}
	pool.Submit(tasks)
}

func parse() ([]string, uint) {
	filePath := flag.String("f", "", "The file to read URLs from.")
	concurrency := flag.Uint("c", 4, "The number of concurrency.")
	flag.Parse()

	urls := []string{}

	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	return urls, *concurrency
}
