package main

import (
	"encoding/json"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "oss.navercorp.com/taeun-ju/go-swag-test/docs"
	"oss.navercorp.com/taeun-ju/go-swag-test/pkg/handlers"
)

type errorResp struct {
	Error string `json:"error"`
}

type authParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResp struct {
	Token string `json:"token"`
}

// writeJSON provides function to format output response in JSON
func writeJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error Parsing JSON")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func decodePost(r *http.Request, structure interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Println("Error parsing post data")
	}
}

// authLogin godoc
// @Summary Auth Login
// @Description Auth Login
// @Tags auth
// @ID auth-login
// @Accept  json
// @Produce  json
// @Param authLogin body main.authParam true "Auth Login Input"
// @Success 200 {object} main.authResp
// @Router /auth/login [post]
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var param authParam
	decodePost(r, &param)

	failResp := errorResp{
		Error: "Wrong username/password",
	}
	writeJSON(w, 401, failResp)
}

// @title Go Restful API with Swagger
// @version 1.0
// @description Simple swagger implementation in Go HTTP

// @contact.name Linggar Primahastoko
// @contact.url http://linggar.asia
// @contact.email x@linggar.asia

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8082
// @BasePath /
func main() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/hello", handlers.HelloWorld)
	http.HandleFunc("/auth/login", AuthLogin)

	http.ListenAndServe(":8082", nil)
}