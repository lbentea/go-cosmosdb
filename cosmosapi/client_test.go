package cosmosapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)

		// check default headers
		assert.NotNil(t, r.Header[HEADER_AUTH])
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/dbs/ToDoList", r.URL.Path)

	}))
	defer ts.Close()

	cfg := Config{
		MasterKey: TestKey,
	}
	c := New(ts.URL, cfg, nil, nil)

	_, err := c.GetDatabase(context.Background(), "ToDoList", nil)
	assert.NotNil(t, err)
}

func TestMaskHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "Bearer some-secret-token")
	header.Add("Content-Type", "application/json")

	maskedHeader := maskHeader(header)

	assert.Equal(t, "*", maskedHeader.Get(HEADER_AUTH))
	assert.Equal(t, "application/json", maskedHeader.Get("Content-Type"))
}
