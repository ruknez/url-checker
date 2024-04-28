package in_memory_bd

// UrlInBd структура которая хранится в самой бд (внутренния).
type UrlInBd struct {
	URL       string
	Duration  int64
	Headers   []string
	LastCheck int64
	Status    int
}
