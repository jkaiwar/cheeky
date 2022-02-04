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

var timestr string = "0"

var timestrcounter int = 0

func init() {
	tmpl = template.Must(template.ParseFiles("g.html"))
}

func main() {
	PORT := ":8083"
	http.HandleFunc("/", index)
	http.HandleFunc("/flip", flop)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	d := struct {
		Title        string
		Color        string
		Instructions string
		Url          string
		ImgMargin    string
		Font         string
		Time         string
	}{}
	if kstate == 0 {
		d.Title = "The Kitchen will be occupied by Joe Rogan's chosen people"
		d.Color = "Gold"
		d.Instructions = "Click to reserve the kitchen for five minutes"
		d.Url = "https://i.ytimg.com/vi/u5kfo7MgAtQ/maxresdefault.jpg"
		d.ImgMargin = "0%"
		d.Font = "sans-serif"
		d.Time = "0"

	}
	if kstate == 1 {
		d.Title = "The Kitchen is Occupied by Covid Believers"
		d.Color = "Red"
		d.Instructions = "You must keep clicking every five minutes"
		d.Url = "https://i.pinimg.com/originals/13/3d/b4/133db4f9d60cfb7f52c00f8bec546149.png"
		d.ImgMargin = "-14%"
		d.Font = "Comic Neue"
		d.Time = timestr
	}
	tmpl.ExecuteTemplate(w, "g.html", d)
}

func flop(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	kstate = 1
	timestr = time.Now().Format("01-02-2006 15:04:05")
	timestrcounter = timestrcounter + 1
	dur := time.Duration(5) * time.Minute
	f := func() {
		timestrcounter = timestrcounter - 1
		if timestrcounter == 0 {
			kstate = 0
		}
	}
	time.AfterFunc(dur, f)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
