package entity

import (
	"errors"
)

type Status int

const (
	Available Status = iota
	NotAvailable
	Moved
	ServerError
)

var NoDataErr = errors.New("No Data")
