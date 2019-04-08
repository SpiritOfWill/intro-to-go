package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var requests = r(Count)

func init() {
	log.SetOutput(ioutil.Discard)
}

func BenchmarkDoAsync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoAsync(requests)
	}
}

func BenchmarkDoSync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoSync(requests)
	}
}

func Test_md5sum(t *testing.T) {
	tests := []struct {
		req  string
		want string
	}{
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"abcd", "e2fc714c4727ee9395f324cd2e7f331f"},
	}

	for _, tt := range tests {
		t.Run(tt.req, func(t *testing.T) {
			got := md5sum([]byte(tt.req))

			require.Equal(t, tt.want, got)
		})
	}
}
