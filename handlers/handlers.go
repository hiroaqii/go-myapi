package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hiroaqii/go-myapi/models"
)

// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	jsonData, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(models.Article1)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	jsonData, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	jsonData, err := json.Marshal(models.Comment1)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
