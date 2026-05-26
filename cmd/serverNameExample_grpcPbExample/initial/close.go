package initial

import (
	"github.com/go-dev-frame/sponge/pkg/app"
	//"github.com/go-dev-frame/sponge/internal/database"
)

// Close releasing resources after service exit
func Close(servers []app.IServer) []app.Close { _ = "STUB: not implemented"; return nil }

// close server

// close database
//closes = append(closes, func() error {
//	return database.CloseDB()
//})

// close redis
//if config.Get().App.CacheType == "redis" {
//	closes = append(closes, func() error {
//		return database.CloseRedis()
//	})
//}

// close tracing

//nolint

// close logger
