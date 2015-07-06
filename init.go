/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          init.go
 * Description:   fork revel's cache
**/

package cache

import (
	"flag"
	"sync"
	"time"
)

var (
	defaultExpiration = time.Hour
	once              sync.Once
	redisHost         string
	redisPasswd       string
)

func init() {

	flag.StringVar(&redisHost, "host", "127.0.0.1:6379", "redis host")
	flag.StringVar(&redisPasswd, "password", "", "redis password")

	once.Do(initializing)
}

func initializing() {
	Instance = NewRedisCache(redisHost, redisPasswd, defaultExpiration)
}
