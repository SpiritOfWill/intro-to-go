package main

import (
	"io/ioutil"
	"log"
	"testing"
)

var requests = r(Count)

func init() {
	log.SetOutput(ioutil.Discard)
}

func Benchmark_doAsync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doAsync(requests)
	}
}

func Benchmark_doSync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doSync(requests)
	}
}
