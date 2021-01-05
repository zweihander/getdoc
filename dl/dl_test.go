package dl

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"
)

type unusableHTTPClient struct{}

func (u unusableHTTPClient) Do(req *http.Request) (*http.Response, error) {
	panic("should not be called")
}

func TestClient_Download(t *testing.T) {
	c, err := NewClient(Options{
		Client:   unusableHTTPClient{},
		Path:     filepath.Join("_testdata", "121.db"),
		Readonly: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := c.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	data, err := c.Get(context.Background(), 0, "schema")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Error("blank")
	}
}
