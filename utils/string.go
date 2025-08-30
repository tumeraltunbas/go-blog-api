package utils

import (
	"strconv"

	"github.com/tumeraltunbas/go-blog/config"
)

func GetPort() string {
	config := config.Get()
	return ":" + strconv.Itoa(config.Port)
}
