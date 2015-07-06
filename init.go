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

	flag.StringVar(&redisHost, "redis_host", "127.0.0.1:6379", "redis host")
	flag.StringVar(&redisPasswd, "redis_passwd", "", "redis password")

	once.Do(initializing)
}

func initializing() {
	Instance = NewRedisCache(redisHost, redisPasswd, defaultExpiration)
}
