package handlers

import (
	"fmt"
	"net/http"
	_ "oss.navercorp.com/taeun-ju/go-swag-test/docs"
)

// @Summary hello handler
// @Description hello test
// @Tags auth
// @ID hello-world
// @Accept  json
// @Produce  json
// @Success 200
// @Router /hello [get]
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
