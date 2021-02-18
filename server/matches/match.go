package matches

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Match representation
type Match struct {
	ID     string `json:"id"` // needs change to UUIDv4
	Sport  string `json:"sport"`
	League string `json:"league"`
	Teams  *Teams `json:"teams"`
}

// Teams in match
type Teams struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

// hardcoded data - need to implement database
var allMatches []Match = []Match{
	{ID: "1", Sport: "Football", League: "Serie A", Teams: &Teams{Home: "Juventus", Away: "Napoli"}},
	{ID: "2", Sport: "Football", League: "Premier League", Teams: &Teams{Home: "Liverpool", Away: "Arsenal"}},
}

// GetMatches - get all matches
func GetMatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allMatches)
}

// GetMatch - get single match by ID
func GetMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range allMatches {
		if (item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Match{})
}

// CreateMatch - create single match
func CreateMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var match Match
	_ = json.NewDecoder(r.Body).Decode(&match)
	allMatches = append(allMatches, match)
	json.NewEncoder(w).Encode(match)
}

// UpdateMatch - upadte match by ID
func UpdateMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var match Match
	params := mux.Vars(r)
	for index, item := range allMatches {
		if (item.ID) == params["id"] {
			allMatches = append(allMatches[:index], allMatches[index+1:]...)
			json.NewDecoder(r.Body).Decode(&match)
			match.ID = params["id"]
			allMatches = append(allMatches, match)
			json.NewEncoder(w).Encode(match)
			return
		}
	}
	json.NewEncoder(w).Encode(allMatches)
}

// DeleteMatch - delete match by ID
func DeleteMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range allMatches {
		if (item.ID) == params["id"] {
			allMatches = append(allMatches[:index], allMatches[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(allMatches)
}
