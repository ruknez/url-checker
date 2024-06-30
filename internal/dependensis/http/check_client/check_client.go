package check_client

import (
	"context"
	"net/http"

	entity "url-checker/internal/domain"

	"github.com/pkg/errors"
)

type CheckClient struct{}

func NewCheckClient() *CheckClient {
	return &CheckClient{}
}

// GetUrlStatus делает запрос по урлу и возвращает его статус.
func (c *CheckClient) GetUrlStatus(ctx context.Context, url string) (entity.Status, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, url, http.NoBody)
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
