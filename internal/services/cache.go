package services

type ICache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
}

type Cache struct {
}

func NewCache() *ICache {
	return nil
}
