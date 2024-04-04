package entity

import (
	"time"
)

type UrlInfo struct {
	URL       string
	Duration  time.Duration
	Headers   []string
	LastCheck *time.Time
	Status    *Status
}

type UrlInBd struct {
	URL       string
	Duration  time.Duration
	Headers   []string
	LastCheck int64
	Status    int
}
