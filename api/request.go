/*
 * Copyright © 2020 Yuhang Chen <i@yuhang.ch>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"fmt"
	"github.com/yuhangch/geoserver-cli/config"
	"io"
	"net/http"
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





