package api

import (
	"fmt"

	"github.com/yuhangch/geoserver-cli/config"
)

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

const (
	layersRootPattern string = "%s/rest/layers"
	layersWsPattern   string = "%s/rest/workspaces/%s/layers"
	layerPattern      string = "%s/rest/workspaces/%s/layers/%s"
)

var (
	layersDeleteStatus map[int]string = map[int]string{
		200: "Deleted",
		403: "Layer may be related is not empty (and recurse not true)",
		500: "Layer may be related is not empty (and recurse not true)",
		404: "Layer doesn’t exist",
	}
)

// LayersResponse represent layers list.
type LayersResponse struct {
	Layers map[string][]Entry `json:"layers"`
}

// Fmt to fmt print layers list.
func (d *LayersResponse) Fmt() string {
	s := "layers:\n"
	for i, v := range d.Layers["layer"] {
		s += fmt.Sprintf("  - %d %s \n", i, v.Name)
	}
	return s
}

// LayersGet to get layers list.
func LayersGet(cfg *config.Config, ws string) error {
	var url string
	if len(ws) > 0 {
		fmt.Print(fmt.Sprintf("[%s] ", ws))
		url = fmt.Sprintf(layersWsPattern, cfg.ServerURL(), ws)
	} else {
		url = fmt.Sprintf(layersRootPattern, cfg.ServerURL())

	}
	method := "GET"

	req := NewRequest(cfg, method, url, nil)
	var r LayersResponse
	Get(req, &r)

	fmt.Println(r.Fmt())
	return nil
}

// LayerDelete to delete a layer.
func LayerDelete(cfg *config.Config, ws, name string) error {

	url := fmt.Sprintf(layerPattern, cfg.ServerURL(), ws, name)
	// fmt.Println(url)
	method := "DELETE"
	req := NewRequest(cfg, method, url, nil)

	Delete(req, layersDeleteStatus)
	return nil
}
