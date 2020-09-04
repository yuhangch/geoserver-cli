package api

/*
Copyright © 2020 Yuhang Chen <i@yuhang.ch>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yuhangch/geoserver-cli/config"
)

// ParseName parse name with workspace or not.
func ParseName(fullname, workspace string) (ws string, name string, err error) {
	if strings.Contains(fullname, ":") {
		strs := strings.Split(fullname, ":")
		if len(strs) < 2 {
			err = fmt.Errorf("illegal name")
		}
		ws = strs[0]
		name = strs[1]
		err = nil
	} else {
		if len(workspace) < 1 {
			err = fmt.Errorf("require a workspace name")
		}
		ws = workspace
		name = fullname

	}
	return
}

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

// NewReqAccept to build a request with custom accept.
func NewReqAccept(cfg *config.Config, method, url, accept string, payload io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", accept)
	auth, err := cfg.BasicAuth()
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+auth)
	return req
}

// NewReqContain to build a request with custom content type.
func NewReqContain(cfg *config.Config, method, url, content string, payload io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", content)
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

// DoWithMsg just do a request.
func DoWithMsg(req *http.Request, s, f string) error {

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

// Delete to get response in del response.
func Delete(req *http.Request, status map[int]string) error {

	code, _, err := Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}
	// fmt.Println(code)

	if msg, ok := status[code]; ok {
		fmt.Println(msg)
		return nil
	}
	return fmt.Errorf("unknown error")

}
