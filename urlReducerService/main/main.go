package main

import (
	"awesomeProject/controller"
	"io"
	"net/http"
	"sync"
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
		_, err := writer.Write([]byte("localhost:8081/" + longurl))
		if err != nil {
			return
		}
	} else {

		value := request.URL.Query().Get("url")
		shorturl := urlMapper.LongShortMap[value]
		writer.Header().Set("Location", "localhost:8081/"+shorturl)

		writer.WriteHeader(300)

	}
}
