//Пакет представляет собой набор хэндлеров входящих запросов и регистрирует роуты
package api

import (
	"blogtask/data"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//InitRouter производит первоначальное сопостовление роутов и хэндлеров
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts", CreatePost).Methods("POST")

	return router
}

//GetPosts обрабатывает входящий запрос на получение списка постов
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := data.GetBlogPosts()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	j, err := json.Marshal(&posts)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//CreatePost обрабатывает запрос на сохранение нового поста
func CreatePost(w http.ResponseWriter, r *http.Request) {
	//прочитаем тело запроса
	var err error
	j := make([]byte, r.ContentLength)
	j, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	defer r.Body.Close()

	//Преобразуем json в структуру data.BlogPost
	post := data.BlogPost{}

	err = json.Unmarshal(j, &post)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	//сохраним пост в БД
	err = post.Save()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.WriteHeader(200)

}
