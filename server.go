package main 

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

// Start with a server that has one endpoint returning `"hello"`. Just `net/http`, nothing else, maybe 15 lines. 
// Get it running and hit it from your browser. 
// Then add a second endpoint that returns JSON. 
// Then add an endpoint that accepts a POST with a JSON body and stores it in memory. 
// Then swap in-memory storage for SQLite. 
// Then put it behind your Astro frontend. 
// Each step is small, each one teaches one thing, and at the end you have a real backend you understand top to bottom — and a reusable pattern for every future project.


func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", jsonResponse)
	log.Fatal(http.ListenAndServe(":8010", nil))
	
}

func rootHandler(w http.ResponseWriter, r *http.Request)   {
	fmt.Println(r)
	str := "what do you want?"
	fmt.Fprint(w, str)
}

func jsonResponse(w http.ResponseWriter, r *http.Request)   {
	mp1 := map[string]string {
		"hello": "world",
	}
	fmt.Println(r)
	w.Header().Set("Content-Type", "application/json")
	// mpjson := json.Marshal(mp1)
	json.NewEncoder(w).Encode(mp1)
	// fmt.Fprintf(json.NewEncoder(w).Encode(v), mp1)
}


