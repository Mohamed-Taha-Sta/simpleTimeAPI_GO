package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Location    string `json:"location"`
	FullTime    string `json:"full_time"`
	Day         string `json:"day"`
	TimeInHours string `json:"time_in_hours"`
	DayInMonth  string `json:"dayInMonth"`
	Month       string `json:"month"`
	Year        string `json:"year"`
}

var timeDifferences = map[string]int{
	"USA":            -5,
	"Canada":         -4,
	"Brazil":         -3,
	"Argentina":      -3,
	"UK":             0,
	"Ireland":        0,
	"Portugal":       0,
	"Spain":          1,
	"France":         1,
	"Germany":        1,
	"Italy":          1,
	"Poland":         1,
	"SouthAfrica":    2,
	"Greece":         2,
	"Turkey":         3,
	"SaudiArabia":    3,
	"Iran":           3,
	"Pakistan":       5,
	"India":          5,
	"Bangladesh":     6,
	"Thailand":       7,
	"China":          8,
	"Japan":          9,
	"Australia":      10,
	"NewZealand":     12,
	"Mexico":         -6,
	"Cuba":           -5,
	"Colombia":       -5,
	"Peru":           -5,
	"Venezuela":      -4,
	"Chile":          -4,
	"Greenland":      -3,
	"Iceland":        0,
	"Norway":         1,
	"Sweden":         1,
	"Finland":        2,
	"Egypt":          2,
	"Russia":         3,
	"UAE":            4,
	"Afghanistan":    4,
	"SriLanka":       5,
	"Myanmar":        6,
	"Indonesia":      7,
	"Vietnam":        7,
	"Philippines":    8,
	"SouthKorea":     9,
	"PapuaNewGuinea": 10,
	"Fiji":           12,
	"Samoa":          13,
	"Tunisia":        1,
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	requestTime := time.Now().UTC()
	location := r.URL.Query().Get("location")
	var timeDiff int
	timeDiff, ok := timeDifferences[location]

	if !ok {
		http.Error(w, "Invalid location", http.StatusBadRequest)
		return
	}

	t := requestTime.Add(time.Duration(timeDiff) * time.Hour)
	timeResponse := TimeResponse{
		Location:    location,
		FullTime:    t.Format(time.RFC1123),
		Day:         t.Format("Monday"),
		TimeInHours: t.Format("15:04"),
		DayInMonth:  t.Format("02"),
		Month:       t.Format("January"),
		Year:        t.Format("2006"),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(timeResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}
