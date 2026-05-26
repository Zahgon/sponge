package initial

import (
	"github.com/go-dev-frame/sponge/pkg/app"
	//"github.com/go-dev-frame/sponge/internal/rpcclient"
)

// Close releasing resources after service exit
func Close(servers []app.IServer) []app.Close { _ = "STUB: not implemented"; return nil }

// close server

// close the rpc client connection
// example:
//closes = append(closes, func() error {
//	return rpcclient.CloseServerNameExampleRPCConn()
//})

// close tracing

//nolint

// close logger
