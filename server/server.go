package server

import (
	"github.com/negaihoshi/daigou/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
