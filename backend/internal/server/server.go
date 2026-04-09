package server

import (
	"fmt"
	"net/http"
)

func HeartBeatHandler(h http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("No implementation error")
}
