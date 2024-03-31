package main

import (
	"io"
	"net/http"
	"sync"
	"urlReduceService/controller"
)

var (
	urlMapper = controller.UrlMapper{
		LongShortMap: make(map[string]string),
		ShortLongMap: make(map[string]string),
	}
	once sync.Once
)

func main() {

	http.HandleFunc("/url", handle)
	http.ListenAndServe(":8081", nil)

}

func handle(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		url, _ := io.ReadAll(request.Body)
		longurl := urlMapper.Post(string(url))
		_, err := writer.Write([]byte(longurl))
		if err != nil {
			return
		}
	} else {

		value := request.URL.Query().Get("url")
		shorturl := urlMapper.LongShortMap[value]
		writer.Header().Set("Location", shorturl)
		writer.WriteHeader(301)

	}
}
