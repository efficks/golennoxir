package main

import (
	"encoding/json"
	"fmt"
	"github.com/efficks/lennoxHeatPump/lennox"
	"net/http"
)

func GetState(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET")
	fmt.Fprintln(w, "Welcome!")
}

func PutCoolState(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s CoolState
	err := decoder.Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lennoxState, e := s.Convert()
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", lennoxState)

	lennox.Apply(lennoxState)

	w.WriteHeader(http.StatusNoContent)
}
