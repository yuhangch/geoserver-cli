package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/yuhangch/geoserver-cli/config"
)

// NewRequest to build a request from a server config.
func NewRequest(cfg *config.Config, method, url string, payload io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	auth, err := cfg.BasicAuth()
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+auth)
	return req
}

// Get to unmarshal body.
func Get(req *http.Request, to interface{}) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &to)
}

// GetTxt to get response in pure text.
func GetText(req *http.Request, pattern string) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {

	}

	defer res.Body.Close()
	if res.StatusCode == 401 {
		fmt.Println("execute faild")
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(fmt.Sprintf(pattern, string(body)))
}

// Do just do a request.
func Do(req *http.Request, s, f string) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {

	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		fmt.Println(s)
	} else {
		fmt.Println(f)
	}
}

// Del to get response in del response.
func Del(req *http.Request) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {

	}

	defer res.Body.Close()
	status := res.StatusCode
	switch status {
	case 200:
		fmt.Println("Success workspace deleted")
	case 403:
		fmt.Println("Workspace or related Namespace is not empty (and recurse not true)")
	case 404:
		fmt.Println("Workspace doesn’t exist")
	case 405:
		fmt.Println("Can’t delete default workspace")
	default:
		fmt.Println("Unknown error")
	}
}
