package check_client

import (
	entity "url-checker/internal/domain"
)

func convertStatus(status int) entity.Status {
	if status >= 200 && status < 300 {
		return entity.Available
	}

	res := entity.NotCheck
	if status > 0 && status < 200 {
		res = entity.NotAvailable
	}

	if status >= 300 && status < 400 {
		res = entity.Moved
	}

	if status >= 400 {
		res = entity.NotAvailable
	}

	return res
}
