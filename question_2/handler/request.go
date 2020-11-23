package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetDataMovie - get movie from api
func GetDataMovie(searchWord, pagination string) ([]byte, error) {
	resp, err := http.Get(url(searchWord, pagination))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func url(searchWord, pagination string) string {
	return fmt.Sprintf(
		"%s/?apikey=%s=%s&page=%s",
		os.Getenv("BASE_URL"),
		os.Getenv("API_KEY"),
		searchWord,
		pagination,
	)
}
