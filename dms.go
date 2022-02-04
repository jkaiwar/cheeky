package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

var tmpl *template.Template

var kstate int = 0

func init() {
	tmpl = template.Must(template.ParseFiles("g.html"))
}

func main() {
	PORT := ":8081"
	http.HandleFunc("/", index)
	http.HandleFunc("/flip", flop)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	d := struct {
		Title        string
		Color        string
		Instructions string
	}{}
	if kstate == 0 {
		d.Title = "The Kitchen is Unoccupied"
		d.Color = "Green"
		d.Instructions = "Click to reserve the kitchen for five minutes"
	}
	if kstate == 1 {
		d.Title = "The Kitchen is Occupied by Covid Believers"
		d.Color = "Red"
		d.Instructions = "You must keep clicking every five minutes"
	}
	tmpl.ExecuteTemplate(w, "g.html", d)
}

func flop(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	kstate = 1
	dur := time.Duration(5) * time.Minute
	f := func() {
		kstate = 0
	}
	time.AfterFunc(dur, f)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
