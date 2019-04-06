package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

const (
	Count        = 100
	TotalWorkers = 5
)

func Work(id int, wg *sync.WaitGroup, reqChan <-chan []byte, respChan chan<- string) {
	log.Printf("worker #%d: started\n", id)

	for data := range reqChan {
		s := fmt.Sprintf("%x", md5.Sum(data))

		log.Printf("worker #%d: sending: %s\n", id, s)

		respChan <- s
	}

	log.Printf("worker #%d: done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	reqChan := make(chan []byte, TotalWorkers)
	respChan := make(chan string, TotalWorkers)

	requests := r()

	go func() {
		for _, b := range requests {
			reqChan <- b
		}
		close(reqChan)
	}()

	wg.Add(TotalWorkers)
	for id := 1; id <= TotalWorkers; id++ {
		// starting workers
		go Work(id, &wg, reqChan, respChan)
	}

	go resChanCloser(&wg, respChan)

	log.Println("results:", getResults(respChan))
}

func resChanCloser(wg *sync.WaitGroup, respChan chan<- string) {
	wg.Wait() // goroutine is blocked

	log.Println("all workers are done, closing respChan")

	close(respChan) // sending a "signal" to func getResults(), that there will be no more messages.
}

func getResults(respChan <-chan string) []string {
	res := make([]string, 0, Count)

	for {
		s, ok := <-respChan
		if !ok {
			break
		}

		log.Println("getResults: got from workers:", s)

		res = append(res, s) // or write to DB...
	}

	return res
}

func r() [][]byte {
	res := make([][]byte, 0, Count)

	for i := 1; i <= Count; i++ {
		b := make([]byte, 256)
		rand.Read(b)
		res = append(res, b)
	}

	log.Println("generated requests")

	return res
}
