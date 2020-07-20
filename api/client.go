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

// NewZipRequest to build a request from a server config.
func NewZipRequest(cfg *config.Config, method, url string, payload io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/zip")
	req.Header.Add("Accept", "application/json")
	auth, err := cfg.BasicAuth()
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+auth)
	return req
}

// Do to execute a request.
func Do(req *http.Request) (int, []byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, nil, err
	}
	return res.StatusCode, body, nil
}

// Get to unmarshal body.
func Get(req *http.Request, to interface{}) error {

	_, body, err := Do(req)
	if err != nil {
		return fmt.Errorf("can't not get the response body")
	}
	err = json.Unmarshal(body, &to)
	if err != nil {
		return fmt.Errorf("can't not unmarshal the body")
	}

	return nil
}

// GetText to get response in pure text.
func GetText(req *http.Request, pattern string) error {

	code, body, err := Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}

	if code == 401 {
		return fmt.Errorf("execute failed: %w", err)
	}
	fmt.Println(fmt.Sprintf(pattern, string(body)))
	return nil
}

// SimplyDo just do a request.
func SimplyDo(req *http.Request, s, f string) error {

	code, _, err := Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}

	if code == 200 {
		fmt.Println(s)
	} else {
		fmt.Println(f)
	}
	return nil
}

// DoCreate to do a request for create.
func DoCreate(req *http.Request) error {
	code, _, err := Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}
	if code == 201 {
		fmt.Println("Created")
	} else {
		fmt.Println("Failed to create")
	}
	return nil
}

// Del to get response in del response.
func Del(req *http.Request, status map[int]string) error {

	code, _, err := Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}

	if msg, ok := status[code]; ok {
		fmt.Println(msg)
		return nil
	}
	return fmt.Errorf("unknown error")

}
