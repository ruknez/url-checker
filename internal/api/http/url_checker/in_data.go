package http

import (
	entity "url-checker/internal/domain"
)

type inResource struct {
	Url string `json:"url"`
}

type outStatus struct {
	Status entity.Status `json:"status"`
}
