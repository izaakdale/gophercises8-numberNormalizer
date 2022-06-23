package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type PostNumberRequest struct {
	Number string `json:"number"`
}
type NumberResponse struct {
	ID     int64  `json:"id"`
	Number string `json:"number"`
}

var reg *regexp.Regexp

func init() {
	var err error

	reg, err = regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) PostNumber(w http.ResponseWriter, r *http.Request) {

	var req PostNumberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Fprintln(w, "Invalid json body")
	}

	dbNumber, err := server.store.AddNumber(context.Background(), req.Number)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbNumber)
}

func (server *Server) GetNormalizedNumbers(w http.ResponseWriter, r *http.Request) {

	dbNumbers, err := server.store.GetNumbers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var res []NumberResponse
	for _, n := range dbNumbers {
		resNum := new(NumberResponse)
		resNum.ID = n.ID

		normNum := reg.ReplaceAllString(n.Number, "")
		resNum.Number = normNum
		res = append(res, *resNum)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
