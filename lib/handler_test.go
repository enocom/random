package random_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	random "github.com/enocom/random/lib"
)

func BenchmarkRootHandler(b *testing.B) {
	store := random.NewStore(0)
	h := random.NewRootHandler(store)
	recoder := httptest.NewRecorder()
	req := &http.Request{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(recoder, req)
	}
}
