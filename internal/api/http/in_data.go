package http

import (
	entity "url-checker/internal/domain"
)

type InResource struct {
	Url string `json:"url"`
}

type OutStatus struct {
	Status entity.Status `json:"status"`
}
