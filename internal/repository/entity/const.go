package entity

import (
	"errors"
)

type Status int

const (
	NotCheck Status = iota
	Available
	NotAvailable
	Moved
	ServerError
)

var NoDataErr = errors.New("No Data")
