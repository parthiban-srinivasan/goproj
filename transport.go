package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

type countRequest struct {
	Str string `json:"s"`
}

type countResponse struct {
	Vi int `json:"v"`
}

func makeCountEndpoint(srv StringService) endpoint.Endpoint {
         return func(ctx context.Context, request interface{}) (interface{}, error) {
                  req := request.(countRequest)
                  v   := srv.Count(req.Str)
                  return countResponse{v}, nil
         }
}

func decodeCountRequest(r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}