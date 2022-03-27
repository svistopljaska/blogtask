package api

import (
	"blogtask/data"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts", CreatePost).Methods("POST")

	return router
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := data.GetBlogPosts()
	if err != nil {
		w.WriteHeader(500)
	}

	j, err := json.Marshal(&posts)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	//прочитаем тело запроса
	var err error
	j := make([]byte, r.ContentLength)
	j, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	defer r.Body.Close()

	//Преобразуем json в структуру data.BlogPost
	post := data.BlogPost{}

	err = json.Unmarshal(j, &post)
	if err != nil {
		w.WriteHeader(500)
	}

	//сохраним пост в БД
	err = post.Save()
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)

}
