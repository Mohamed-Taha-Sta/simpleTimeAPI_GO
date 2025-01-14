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
	"USA":                   -5,
	"Canada":                -4,
	"Brazil":                -3,
	"Argentina":             -3,
	"UK":                    0,
	"Ireland":               0,
	"Portugal":              0,
	"Spain":                 1,
	"France":                1,
	"Germany":               1,
	"Italy":                 1,
	"Poland":                1,
	"SouthAfrica":           2,
	"Greece":                2,
	"Turkey":                3,
	"SaudiArabia":           3,
	"Iran":                  3,
	"Pakistan":              5,
	"India":                 5,
	"Bangladesh":            6,
	"Thailand":              7,
	"China":                 8,
	"Japan":                 9,
	"Australia":             10,
	"NewZealand":            12,
	"Mexico":                -6,
	"Cuba":                  -5,
	"Colombia":              -5,
	"Peru":                  -5,
	"Venezuela":             -4,
	"Chile":                 -4,
	"Greenland":             -3,
	"Iceland":               0,
	"Norway":                1,
	"Sweden":                1,
	"Finland":               2,
	"Egypt":                 2,
	"Russia":                3,
	"UAE":                   4,
	"Afghanistan":           4,
	"SriLanka":              5,
	"Myanmar":               6,
	"Indonesia":             7,
	"Vietnam":               7,
	"Philippines":           8,
	"SouthKorea":            9,
	"PapuaNewGuinea":        10,
	"Fiji":                  12,
	"Samoa":                 13,
	"Tunisia":               1,
	"Tunis":                 1,
	"Sousse":                1,
	"Bizerte":               1,
	"Sfax":                  1,
	"Belgium":               1,
	"Netherlands":           1,
	"Switzerland":           1,
	"Austria":               1,
	"Denmark":               1,
	"Palestine":             2,
	"Lebanon":               2,
	"Jordan":                2,
	"Syria":                 2,
	"Iraq":                  3,
	"Kuwait":                3,
	"Qatar":                 3,
	"Bahrain":               3,
	"Oman":                  4,
	"Kazakhstan":            6,
	"Mongolia":              8,
	"Malaysia":              8,
	"Taiwan":                8,
	"NorthKorea":            9,
	"Guam":                  10,
	"SolomonIslands":        11,
	"Vanuatu":               11,
	"MarshallIslands":       12,
	"Nauru":                 12,
	"Kiribati":              14,
	"USA_EST":               -5,
	"USA_CST":               -6,
	"USA_MST":               -7,
	"USA_PST":               -8,
	"Canada_Pacific":        -8,
	"Canada_Mountain":       -7,
	"Canada_Central":        -6,
	"Canada_Eastern":        -5,
	"Canada_Atlantic":       -4,
	"Winnipeg":              -6,
	"QuebecCity":            -5,
	"Hamilton":              -5,
	"BuenosAires":           -3,
	"Guadalajara":           -6,
	"Monterrey":             -6,
	"Valencia":              1,
	"Seville":               1,
	"Zaragoza":              1,
	"Munich":                1,
	"Hamburg":               1,
	"Cologne":               1,
	"Stuttgart":             1,
	"Dusseldorf":            1,
	"Naples":                1,
	"Turin":                 1,
	"Palermo":               1,
	"Genoa":                 1,
	"Bologna":               1,
	"Florence":              1,
	"Catania":               1,
	"Bari":                  1,
	"Venice":                1,
	"Verona":                1,
	"Messina":               1,
	"Padua":                 1,
	"Trieste":               1,
	"Brescia":               1,
	"Prato":                 1,
	"Modena":                1,
	"ReggioCalabria":        1,
	"ReggioEmilia":          1,
	"Perugia":               1,
	"Livorno":               1,
	"Ravenna":               1,
	"Cagliari":              1,
	"Rimini":                1,
	"Salerno":               1,
	"Ferrara":               1,
	"Sassari":               1,
	"Latina":                1,
	"GiuglianoInCampania":   1,
	"Monza":                 1,
	"Syracuse":              1,
	"Bergamo":               1,
	"Pescara":               1,
	"Trento":                1,
	"Forli":                 1,
	"Vicenza":               1,
	"Terni":                 1,
	"Bolzano":               1,
	"Novara":                1,
	"Piacenza":              1,
	"Ancona":                1,
	"Arezzo":                1,
	"Udine":                 1,
	"LaSpezia":              1,
	"Pesaro":                1,
	"Alessandria":           1,
	"Pistoia":               1,
	"Pisa":                  1,
	"Catanzaro":             1,
	"GuidoniaMontecelio":    1,
	"Lucca":                 1,
	"TorreDelGreco":         1,
	"Treviso":               1,
	"BustoArsizio":          1,
	"Como":                  1,
	"Marsala":               1,
	"Grosseto":              1,
	"Pozzuoli":              1,
	"Varese":                1,
	"Fiumicino":             1,
	"Casoria":               1,
	"Asti":                  1,
	"CiniselloBalsamo":      1,
	"Caserta":               1,
	"Gela":                  1,
	"Aprilia":               1,
	"Ragusa":                1,
	"Pavia":                 1,
	"Cremona":               1,
	"Carpi":                 1,
	"QuartuSantElena":       1,
	"LameziaTerme":          1,
	"Imola":                 1,
	"Aquila":                1,
	"Massa":                 1,
	"Viterbo":               1,
	"Cosenza":               1,
	"Potenza":               1,
	"CastellammareDiStabia": 1,
	"Afragola":              1,
	"Vittoria":              1,
	"Crotone":               1,
	"Pomezia":               1,
	"Vigevano":              1,
	"Carrara":               1,
	"Caltanissetta":         1,
	"Viareggio":             1,
	"Fano":                  1,
	"Savona":                1,
	"Matera":                1,
	"Olbia":                 1,
	"Legnano":               1,
	"Agrigento":             1,
	"Sassuolo":              1,
	"Bagheria":              1,
	"Anzio":                 1,
	"Portici":               1,
	"Modica":                1,
	"Sanremo":               1,
	"Avellino":              1,
	"Teramo":                1,
	"Montesilvano":          1,
	"Siena":                 1,
	"Gallarate":             1,
	"Velletri":              1,
	"CavaDe’Tirreni":        1,
	"SanSevero":             1,
	"Aosta":                 1,
	"Cesena":                1,
	"NoviLigure":            1,
	"Tivoli":                1,
	"Cuneo":                 1,
	"Foligno":               1,
	"Torremaggiore":         1,
	"NoceraInferiore":       1,
	"PalazzoloAcreide":      1,
	"Lugo":                  1,
	"Rovigo":                1,
	"Pordenone":             1,
	"Acireale":              1,
	"MazaraDelVallo":        1,
	"Trapani":               1,
	"Caltagirone":           1,
	"Floridia":              1,
	"Ercolano":              1,
	"PortoEmpedocle":        1,
	"MolaDiBari":            1,
	"SanGiovanniRotondo":    1,
	"Manfredonia":           1,
	"SanteramoInColle":      1,
	"Monopoli":              1,
	"RuvoDiPuglia":          1,
	"Foggia":                1,
	"Barletta":              1,
	"Andria":                1,
	"Trani":                 1,
	"Bisceglie":             1,
	"CanosaDiPuglia":        1,
	"Corato":                1,
	"Molfetta":              1,
	"Bitonto":               1,
	"PaloDelColle":          1,
	"Altamura":              1,
	"GravinaInPuglia":       1,
	"Grottaglie":            1,
	"MartinaFranca":         1,
	"Brindisi":              1,
	"Taranto":               1,
	"Gallipoli":             1,
	"Manduria":              1,
	"Oria":                  1,
	"SanVitoDeiNormanni":    1,
	"Mesagne":               1,
	"FrancavillaFontana":    1,
	"Fasano":                1,
	"Ostuni":                1,
	"Cisternino":            1,
	"Carovigno":             1,
	"CeglieMessapica":       1,
	"Latiano":               1,
	"SanPietroVernotico":    1,
	"Torchiarolo":           1,
	"Squinzano":             1,
	"Trepuzzi":              1,
	"Surbo":                 1,
	"Lecce":                 1,
	"Cavallino":             1,
	"Vernole":               1,
	"Melendugno":            1,
	"Martano":               1,
	"CarpignanoSalentino":   1,
	"Maglie":                1,
	"Poggiardo":             1,
	"SantaCesareaTerme":     1,
	"Otranto":               1,
	"Castro":                1,
	"Tricase":               1,
	"Alessano":              1,
	"Presicce":              1,
	"Salve":                 1,
	"MorcianoDiLeuca":       1,
	"Patù":                  1,
	"CastrignanoDelCapo":    1,
	"Taviano":               1,
	"Racale":                1,
	"Melissano":             1,
	"Casarano":              1,
	"Matino":                1,
	"Parabita":              1,
	"Collepasso":            1,
	"Cutrofiano":            1,
	"Aradeo":                1,
	"Galatina":              1,
	"SoglianoCavour":        1,
	"CoriglianoD’Otranto":   1,
	"Melpignano":            1,
	"CastrignanoDe’Greci":   1,
	"Zollino":               1,
	"Sternatia":             1,
	"Martignano":            1,
	"Calimera":              1,
	"CapraricaDiLecce":      1,
	"Veglie":                1,
	"CampiSalentina":        1,
	"Carmiano":              1,
	"Arnesano":              1,
	"MonteroniDiLecce":      1,
	"Leverano":              1,
	"Copertino":             1,
	"Nardò":                 1,
	"Galatone":              1,
	"SantaMariaAlBagno":     1,
	"SantaCaterina":         1,
	"PortoCesareo":          1,
	"TorreLapillo":          1,
	"PuntaProsciutto":       1,
	"TorreColimena":         1,
	"SanPietroInBevagna":    1,
	"Maruggio":              1,
	"Campomarino":           1,
	"TorreOvo":              1,
	"MonacoDiBaviera":       1,
	"Berlino":               1,
	"Amburgo":               1,
	"Colonia":               1,
	"FrancoforteSulMeno":    1,
	"Stoccarda":             1,
	"Düsseldorf":            1,
	"Dortmund":              1,
	"Essen":                 1,
	"Brema":                 1,
	"Dresda":                1,
	"Lipsia":                1,
	"Hannover":              1,
	"Chongqing":             8,
	"Algiers":               1,
	"AbuDhabi":              4,
	"Amsterdam":             1,
	"Athens":                2,
	"Auckland":              12,
	"Baghdad":               3,
	"Bangkok":               7,
	"Barcelona":             1,
	"Beijing":               8,
	"Berlin":                1,
	"Bogota":                -5,
	"Brasilia":              -3,
	"Brisbane":              10,
	"Brussels":              1,
	"Bucharest":             2,
	"Budapest":              1,
	"Cairo":                 2,
	"Calgary":               -7,
	"CapeTown":              2,
	"Caracas":               -4,
	"Casablanca":            0,
	"Chicago":               -6,
	"Copenhagen":            1,
	"Dallas":                -6,
	"DarEsSalaam":           3,
	"Dhaka":                 6,
	"Doha":                  3,
	"Dubai":                 4,
	"Dublin":                0,
	"Edmonton":              -7,
	"Frankfurt":             1,
	"Geneva":                1,
	"Hanoi":                 7,
	"Harare":                2,
	"Havana":                -5,
	"Helsinki":              2,
	"HoChiMinhCity":         7,
	"HongKong":              8,
	"Honolulu":              -10,
	"Houston":               -6,
	"Istanbul":              3,
	"Jakarta":               7,
	"Jerusalem":             2,
	"Johannesburg":          2,
	"Karachi":               5,
	"Khartoum":              2,
	"Kingston":              -5,
	"KualaLumpur":           8,
	"KuwaitCity":            3,
	"Kyiv":                  2,
	"Lagos":                 1,
	"Lahore":                5,
	"LasVegas":              -8,
	"Lima":                  -5,
	"Lisbon":                0,
	"London":                0,
	"LosAngeles":            -8,
	"Madrid":                1,
	"Manila":                8,
	"Melbourne":             10,
	"MexicoCity":            -6,
	"Miami":                 -5,
	"Milan":                 1,
	"Montreal":              -5,
	"Moscow":                3,
	"Nairobi":               3,
	"Nassau":                -5,
	"NewOrleans":            -6,
	"NewYork":               -5,
	"Osaka":                 9,
	"Oslo":                  1,
	"Ottawa":                -5,
	"Paris":                 1,
	"Perth":                 8,
	"Philadelphia":          -5,
	"Phoenix":               -7,
	"Prague":                1,
	"Pretoria":              2,
	"Quito":                 -5,
	"RioDeJaneiro":          -3,
	"Riyadh":                3,
	"Rome":                  1,
	"SaintPetersburg":       3,
	"SanFrancisco":          -8,
	"Santiago":              -4,
	"SaoPaulo":              -3,
	"Seattle":               -8,
	"Seoul":                 9,
	"Shanghai":              8,
	"Singapore":             8,
	"Sofia":                 2,
	"Stockholm":             1,
	"Sydney":                10,
	"Taipei":                8,
	"Tallinn":               2,
	"Tashkent":              5,
	"Tokyo":                 9,
	"Toronto":               -5,
	"Vancouver":             -8,
	"Vienna":                1,
	"Warsaw":                1,
	"WashingtonDC":          -5,
	"Wellington":            12,
	"Zagreb":                1,
	"Zurich":                1,
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
