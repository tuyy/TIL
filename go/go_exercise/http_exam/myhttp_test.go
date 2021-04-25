package exam

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	resp, err := http.Get("http://www.naver.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(b))

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/tuyy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	http.ListenAndServe(":8080", nil)
}

func TestContext(t *testing.T) {
	http.HandleFunc("/context", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println("server: hello handler started")
		defer fmt.Println("server: hello handler ended")

		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintf(w, "hello\n")
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println("server:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
