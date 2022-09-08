package httpservers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	<-time.After(time.Second * 5)
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Message string `json:"message"`
	}{"Hello from golang"}

	json.NewEncoder(w).Encode(data)
}

func helloWithContext(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("helloWithContext started")
	defer fmt.Println("helloWithContext ended")

	fakeJobDuration := time.Second * 5

	select {
	case <-time.After(fakeJobDuration):
		{
			hello(w, req)

		}
	case <-ctx.Done():
		{
			err := ctx.Err()
			fmt.Printf("err: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.Header)
}

func Run() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", handleHeaders)
	http.HandleFunc("/hello-with-context", helloWithContext)

	http.ListenAndServe(":8080", nil)
}
