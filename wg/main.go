package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"sync"
)

const (
	Count        = 10
	TotalWorkers = 5
)

func Work(i int, wg *sync.WaitGroup, inChannel <-chan []byte, resChannel chan<- string) {
	fmt.Printf("spainning worker #%d\n", i)

	for data := range inChannel {
		s := fmt.Sprintf("%x", md5.Sum(data))

		fmt.Printf("worker #%d: sending: %s\n", i, s)

		resChannel <- s
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	inChannel := make(chan []byte, TotalWorkers)
	resChannel := make(chan string, Count)

	wg.Add(TotalWorkers)
	for i := 1; i <= TotalWorkers; i++ {
		// starting workers
		go Work(i, &wg, inChannel, resChannel)
	}

	for _, b := range r() {
		inChannel <- b
	}
	close(inChannel)

	go resChanCloser(&wg, resChannel) // non blocking

	fmt.Println(getResults(resChannel))
}

func resChanCloser(wg *sync.WaitGroup, resChannel chan<- string) {
	wg.Wait() // goroutine is blocked

	// all workers are done

	close(resChannel) // sending a "signal" to func getResults(), that there will be no more messages.
}

func getResults(resChannel <-chan string) []string {
	res := make([]string, 0, Count)

	for {
		s, ok := <-resChannel
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
