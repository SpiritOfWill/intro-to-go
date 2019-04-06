package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"sync"
)

const (
	Count        = 50
	TotalWorkers = 5
)

func Work(i int, wg *sync.WaitGroup, reqChan <-chan []byte, respChan chan<- string) {
	fmt.Printf("spainning worker #%d\n", i)

	for data := range reqChan {
		s := fmt.Sprintf("%x", md5.Sum(data))

		fmt.Printf("worker #%d: sending: %s\n", i, s)

		respChan <- s
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	reqChan := make(chan []byte, TotalWorkers)
	respChan := make(chan string, Count)

	wg.Add(TotalWorkers)
	for i := 1; i <= TotalWorkers; i++ {
		// starting workers
		go Work(i, &wg, reqChan, respChan)
	}

	requests := r()

	go func() {
		for _, b := range requests {
			reqChan <- b
		}
		close(reqChan)
	}()

	go resChanCloser(&wg, respChan)

	fmt.Println(getResults(respChan))
}

func resChanCloser(wg *sync.WaitGroup, respChan chan<- string) {
	wg.Wait() // goroutine is blocked

	// all workers are done

	close(respChan) // sending a "signal" to func getResults(), that there will be no more messages.
}

func getResults(respChan <-chan string) []string {
	res := make([]string, 0, Count)

	for {
		s, ok := <-respChan
		if !ok {
			break
		}

		fmt.Println("got from workers:", s)

		res = append(res, s) // or write to DB...
	}

	return res
}

func r() [][]byte {
	res := make([][]byte, Count)

	for i := 1; i <= TotalWorkers; i++ {
		b := make([]byte, 256)
		rand.Read(b)
		res = append(res, b)
	}

	return res
}
