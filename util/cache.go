package util

import (
	"github.com/muesli/cache2go"
)
func NewCache()*cache2go.CacheTable  {
	 return  cache2go.Cache("myCache")
}
