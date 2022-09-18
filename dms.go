package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template
var timestr string = ""
var t *time.Timer

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
	t = time.NewTimer((0) * time.Second)
}

func main() {
	PORT := ":8081"
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/flip", flipHandler)
	http.HandleFunc("/api", apiHandler)
	go func() {
		for {
			<-t.C
			timestr = ""
		}
	}()
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	d := struct {
		Title        string
		Color        string
		Instructions string
		Url          string
		ImgMargin    string
		Font         string
		Time         string
	}{}
	if timestr == "" {
		d.Title = "The Kitchen is Free"
		d.Color = "Gold"
		d.Instructions = "Click to reserve the kitchen for five minutes"
		d.Font = "sans-serif"
		d.Time = "0"

	} else {
		d.Title = "The Kitchen is Occupied"
		d.Color = "Red"
		d.Instructions = "Reservation expiring soon"
		d.Font = "Comic Neue"
		d.Time = timestr
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", d)
}

func flipHandler(w http.ResponseWriter, r *http.Request) {
	timestr = time.Now().Format("01-02-2006 15:04:05")
	dur := time.Duration(5) * time.Minute
	t.Reset(dur)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if timestr == "" {
		w.Write([]byte("novid: false"))
	} else {
		w.Write([]byte("novid: true"))
	}
}
