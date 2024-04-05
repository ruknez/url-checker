package check_client

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

//go:generate ./../../../../bin/moq -stub -skip-ensure -pkg mocks -out ./mocks/get_url_statuser_mock.go . GetUrlStatuser:GetUrlStatuserMock
type GetUrlStatuser interface {
	GetUrlStatus(ctx context.Context, url string) (int, error)
}

type checkClient struct {
	r io.Reader
}

func (c *checkClient) GetUrlStatus(ctx context.Context, url string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, url, c.r)
	if err != nil {
		return 0, errors.Wrap(err, "http.NewRequestWithContext")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "DefaultClient.Do")
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
