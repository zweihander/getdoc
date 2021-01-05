package getdoc

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/getdoc/dl"
)

type unusableHTTPClient struct{}

func (u unusableHTTPClient) Do(req *http.Request) (*http.Response, error) {
	panic("should not be called")
}

func TestExtract(t *testing.T) {
	c, err := dl.NewClient(dl.Options{
		Client:   unusableHTTPClient{},
		Path:     filepath.Join("dl", "_testdata", "121.db"),
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

	doc, err := Extract(context.Background(), c)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 121, doc.Index.Layer)
}
