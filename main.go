package main

import (
	"net/http"
	"os"
//	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)

	var srv StringService
	srv = stringService{}

	countHandler := httptransport.NewServer(
		ctx,
		makeCountEndpoint(srv),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
