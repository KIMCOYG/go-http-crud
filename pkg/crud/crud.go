package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func CreateData() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getArticles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
	fmt.Println("Endpoint Hit: createArticle")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article

	json.Unmarshal(reqBody, &article)

	id := article.Id

	var err error
	var num int

	num, err = strconv.Atoi(id)
	num = num - 1

	if err == nil {
		fmt.Println("Error updateArticle")
	}

	json.Unmarshal(reqBody, &article)

	for _, art := range Articles {
		if art.Id == id {
			Articles[num].Title = article.Title
			Articles[num].Desc = article.Desc
			Articles[num].Content = article.Content
		}
	}

	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint Hit: updateArticle")
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/create", createArticle).Methods("POST")
	router.HandleFunc("/read", getArticles).Methods("GET")
	router.HandleFunc("/read/{id}", getArticle).Methods("GET")
	router.HandleFunc("/update", updateArticle).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", router))
}
