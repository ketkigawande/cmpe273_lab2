package main

import (  
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)
type Request struct {
	   Name   string `json:"name"`
}

type Response struct {
	Greeting string `json:"greeting"`	
}

func postt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   
    request:= Request{}
    response := Response{}
    json.NewDecoder(r.Body).Decode(&request)
    response.Greeting = "Hello, "+request.Name
	// Marshal provided interface into JSON structure
    uj, _ := json.Marshal(response)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)
}



func main() {  
    r := httprouter.New()
r.POST("/hello", postt)

	 server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: r,
    }
    server.ListenAndServe()

 }
// curl -H "Content-Type: application/json" -X POST -d '{"name":"ketki"}' http://localhost:8080 