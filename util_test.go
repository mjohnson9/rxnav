package rxnav_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/nightexcessive/rxnav"
)

func TestCheckConnection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := rxnav.CheckConnection(ctx)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMaxConns(t *testing.T) {
	transport := http.DefaultTransport.(*http.Transport)
	t.Logf("Max idle connections: %d", transport.MaxIdleConns)
	t.Logf("Max idle connections per host: %d", transport.MaxIdleConnsPerHost)
	t.Logf("Max total connections per host: %d", transport.MaxConnsPerHost)
}
