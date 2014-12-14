package homie

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func init() {
	http.HandleFunc("/", base)
	http.HandleFunc("/about", about)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/KF5JBH", arrl)
}

type Person struct {
	Alias       string `json:"alias"`
	Age         int    `json:"age"`
	Gh          string `json:"gh"`
	Slideshare  string `json:"slideshare"`
	Speakerdeck string `json:"speakerdeck"`
	Other       string `json:"other"`
	FreeDNS     string `json:"getfreedns"`
	Meta        Meta   `json:"meta"`
}

type Meta struct {
	Geo      string    `json:"geo"`
	Time     time.Time `json:"resp"`
	UnixTime int64     `json:"req_unix_nano"`
}

func base(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, maxage=30, s-maxage=60")

	mm := Person{Alias: "mmay",
		Age:         time.Now().Year() - 1990,
		Gh:          "https://github.com/mmay",
		Speakerdeck: "https://speakerdeck.com/mmay",
		Slideshare:  "http://www.slideshare.net/iammichaelmay",
		FreeDNS:     "https://dnsimple.com/r/4b4375226280a9",
		Meta: Meta{
			Time:     time.Now(),
			UnixTime: time.Now().Unix()}}

	result, _ := json.MarshalIndent(mm, "", "  ")
	io.WriteString(w, string(result))
}

func about(w http.ResponseWriter, r *http.Request) {

}
func blog(w http.ResponseWriter, r *http.Request) {

}

func arrl(w http.ResponseWriter, r *http.Request) {

}

func link(ref string, name string) string {
	a := []string{"<a href=http://", ref, ">", name, "<\a>"}
	return strings.Join(a, "")
}
