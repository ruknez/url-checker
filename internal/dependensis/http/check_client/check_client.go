package check_client

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	entity "url-checker/internal/domain"
)

type checkClient struct {
	r io.Reader
}

func NewCheckClient() *checkClient {
	// TODO тут нужен ридер но не ясно зачем?
	return &checkClient{r: r}
}

func (c *checkClient) GetUrlStatus(ctx context.Context, url string) (entity.Status, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, url, c.r)
	if err != nil {
		return 0, errors.Wrap(err, "http.NewRequestWithContext")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "DefaultClient.Do")
	}

	defer resp.Body.Close()

	return convertStatus(resp.StatusCode), nil
}
