package mapping

import (
	"time"

	"go.openly.dev/pointy"

	entity "url-checker/internal/domain"
	inMemoryBd "url-checker/internal/repository/in_memory_bd/entity"
)

func URLBdToURLInfoMapping(in inMemoryBd.UrlInBd) entity.UrlInfo {

	var lastCheck *time.Time
	if in.LastCheck != 0 {
		lastCheck = pointy.Pointer(time.Unix(in.LastCheck, 0))
	}

	return entity.UrlInfo{
		URL:       in.URL,
		Duration:  time.Duration(in.Duration) * time.Millisecond,
		Headers:   in.Headers,
		LastCheck: lastCheck,
		Status:    entity.Status(in.Status),
	}
}
