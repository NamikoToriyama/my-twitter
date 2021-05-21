package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/mercari/Week6/Week6/NamikoToriyama/app/domain"
	"github.com/mercari/Week6/Week6/NamikoToriyama/app/repository"
	"github.com/mercari/Week6/Week6/NamikoToriyama/app/service"
	"google.golang.org/appengine"
)

var dirPath = "./front/"
var tmpl = template.Must(template.ParseGlob(dirPath + "*.html"))

func init() {

	ur := repository.NewTweetRepository()
	us := service.NewTweetService(ur)
	th := domain.NewTweetDomain(us)
	reg := NewTweetHandler(th)

	// Handler
	http.HandleFunc("/tweet", reg.tweetPage)
	http.HandleFunc("/detail", reg.tweetDetailPage)
	http.HandleFunc("/delete", reg.tweetDeletePage)
	http.HandleFunc("/edit", reg.editPage)
	http.HandleFunc("/update", reg.updateTweetPage)

	http.HandleFunc("/my", myHandler)
	http.HandleFunc("/signup", indexHandler)
	http.HandleFunc("/", reg.tweetListPage)
}

func main() {
	appengine.Main()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	fmt.Println("listening on port: ", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// indexHandler ... Sign up page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("front/signup.html"))

	if err := t.ExecuteTemplate(w, "signup.html", getEnvVars()); err != nil {
		log.Fatal(err)
	}

}

// myHandler ... My page.
func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my" {
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("front/myPage.html"))
	if err := t.ExecuteTemplate(w, "myPage.html", getEnvVars()); err != nil {
		log.Fatal(err)
	}
}

// getEnvVars ... Get environment variables as a map.
func getEnvVars() map[string]string {
	return map[string]string{
		"apiKey":            os.Getenv("API_KEY"),
		"authDomain":        os.Getenv("AUTH_DOMAIN"),
		"databaseURL":       os.Getenv("DATABASE_URL"),
		"projectId":         os.Getenv("PROJECT_ID"),
		"storageBucket":     os.Getenv("STORAGE_BUCKET"),
		"messagingSenderId": os.Getenv("MESSAGING_SENDER_ID"),
		"appId":             os.Getenv("APP_ID"),
	}
}
