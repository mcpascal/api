package cache

import lru "github.com/hashicorp/golang-lru"

var (
	Cache *lru.Cache
)

func Setup() {
	c, err := lru.New(100)
	if err != nil {
		panic(err)
	}
	Cache = c
}
