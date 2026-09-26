package deco

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

func TestRequestCancellationDoesNotPoisonCachedClient(t *testing.T) {
	oldURL := baseURL
	t.Cleanup(func() { baseURL = oldURL })
	c := New("test-deco")
	started := make(chan struct{})
	var calls int32
	c.c.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if atomic.AddInt32(&calls, 1) == 1 {
			close(started)
			<-r.Context().Done()
			return nil, r.Context().Err()
		}
		if r.Context().Err() != nil {
			t.Error("original client inherited canceled context")
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"ok":true}`)), Header: make(http.Header)}, nil
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		var result map[string]interface{}
		done <- c.WithContext(ctx).doPost("/test", EndpointArgs{}, nil, &result)
	}()
	<-started
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v, want canceled request", err)
	}
	var result map[string]interface{}
	if err := c.doPost("/test", EndpointArgs{}, nil, &result); err != nil || result["ok"] != true {
		t.Fatalf("next request through original client failed: %v %v", result, err)
	}
}
