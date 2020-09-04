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
	"fmt"
	"strings"

	"github.com/yuhangch/geoserver-cli/config"
)

const (
	workspacesPattern string = "%s/rest/workspaces"
	workspacePattern  string = "%s/rest/workspaces/%s"
)

var (
	workspaceDeleteStatus map[int]string = map[int]string{
		200: "Success workspace deleted",
		403: "Workspace or related Namespace is not empty (and recurse not true)",
		404: "Workspace doesn’t exist",
		405: "Can’t delete default workspace",
	}
)

// WorkSpaceResponse the model of response of workspaces list.
type WorkSpaceResponse struct {
	Workspaces map[string][]Entry
}

// WorkSpace to represent workspace detail.
type WorkSpace struct {
	Name           string `json:"name"`
	Isolated       bool   `json:"isolated"`
	DataStores     string `json:"dataStores"`
	CoverageStores string `json:"coverageStores"`
	WmsStores      string `json:"wmsStores"`
	WmtsStores     string `json:"wmtsStores"`
}

// Fmt to fmt print workspace list.
func (w *WorkSpaceResponse) Fmt() string {
	s := "Workspaces:\n"
	for i, v := range w.Workspaces["workspace"] {
		s += fmt.Sprintf("  - %d %s \n", i, v.Name)
	}
	return s
}

// WorkSpacesGet handle workspaces query.
func WorkSpacesGet(cfg *config.Config) {
	url := fmt.Sprintf(workspacesPattern, cfg.ServerURL())
	method := "GET"

	req := NewRequest(cfg, method, url, nil)
	var w WorkSpaceResponse
	Get(req, &w)

	fmt.Println(w.Fmt())
}

// WorkSpacesPost handle create new workspace.
func WorkSpacesPost(cfg *config.Config, name string) {
	url := fmt.Sprintf(workspacesPattern, cfg.ServerURL())
	// fmt.Println(url)
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf("{\"workspace\":{\"name\":\"%s\"}}", name))

	req := NewRequest(cfg, method, url, payload)
	GetText(req, "workspace %s created")
}

// WorkSpacePut handle delete workspace.
func WorkSpacePut(cfg *config.Config, oldname, newname string) {

	url := fmt.Sprintf(workspacePattern, cfg.ServerURL(), oldname)
	// fmt.Println(url)
	method := "PUT"
	payload := strings.NewReader(fmt.Sprintf("{\"workspace\":{\"name\":\"%s\"}}", newname))

	req := NewRequest(cfg, method, url, payload)
	DoWithMsg(req, fmt.Sprintf("%s renamed to %s", oldname, newname), "Rename failed")
}

// WorkSpaceDelCode del response code .
type WorkSpaceDelCode struct {
	one int
}

// WorkSpaceDelete handle delete workspace.
func WorkSpaceDelete(cfg *config.Config, name string) {

	url := fmt.Sprintf(workspacePattern, cfg.ServerURL(), name)
	// fmt.Println(url)
	method := "DELETE"
	payload := strings.NewReader(fmt.Sprintf("{\"workspace\":{\"name\":\"%s\"}}", name))

	req := NewRequest(cfg, method, url, payload)
	Delete(req, workspaceDeleteStatus)
}
