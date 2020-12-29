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
	"io/ioutil"
	"net/http"
	"strings"
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
			err = fmt.Errorf("workspace name required")
		}
		ws = workspace
		name = fullname

	}
	return
}

// HandleRequest to execute a request.
func HandleRequest(req *http.Request) (int, []byte, error) {
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

// HandleBody to raw unmarshal body.
func HandleBody(req *http.Request, to interface{}) error {

	_, body, err := HandleRequest(req)
	if err != nil {
		return fmt.Errorf("can't not get the response body")
	}
	err = json.Unmarshal(body, &to)
	if err != nil {
		return fmt.Errorf("can't not unmarshal the body")
	}

	return nil
}

// HandleText to get response which in pure text format.
func HandleText(req *http.Request, pattern string) error {

	code, body, err := HandleRequest(req)
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

	code, _, err := HandleRequest(req)
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

// Create to do a request for create.
func Create(req *http.Request) error {
	code, _, err := HandleRequest(req)
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

// Delete handle response for delete action.
func Delete(req *http.Request, status map[int]string) error {

	code, _, err := HandleRequest(req)
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
