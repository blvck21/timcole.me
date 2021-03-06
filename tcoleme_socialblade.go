package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// TwitchStats Returns socialblade twitch stats for modesttim
func TwitchStats(w http.ResponseWriter, r *http.Request) {
	url := "https://api.socialblade.com/v2/twitch/statistics?query=statistics&username=modesttim&api=" + settings.Get("SOCIALBLADE_API")

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}
