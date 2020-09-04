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
	"os"

	"github.com/yuhangch/geoserver-cli/config"
)

const (
	datastoresPattern         string = "%s/rest/workspaces/%s/datastores"
	datastorePattern          string = "%s/rest/workspaces/%s/datastores/%s"
	datastoreCreateShpPattern string = "%s/rest/workspaces/%s/datastores/%s/%s.%s?configure=%s"
)

// StoreMethod represent upload method for create datastore.
// type StoreMethod string

// const (
// 	FileMethod     StoreMethod = "file"
// 	URLMethod      StoreMethod = "url"
// 	ExternalMethod StoreMethod = "external"
// )

var (
	datastoreDeleteStatus map[int]string = map[int]string{
		200: "Deleted",
		403: "Datastore or related is not empty (and recurse not true)",
		404: "Datastore not exist",
	}
)

// DataStoresResponse represent datastore list.
type DataStoresResponse struct {
	DataStores map[string][]Entry `json:"datastores"`
}

// Fmt to fmt print workspace list.
func (d *DataStoresResponse) Fmt() string {
	s := "DataStores:\n"
	for i, v := range d.DataStores["dataStore"] {
		s += fmt.Sprintf("  - %d %s \n", i, v.Name)
	}
	return s
}

// DataStoresGet to get datastores list for workspaces.
func DataStoresGet(cfg *config.Config, ws string) error {
	url := fmt.Sprintf(datastoresPattern, cfg.ServerURL(), ws)
	method := "GET"

	req := NewRequest(cfg, method, url, nil)
	var r DataStoresResponse
	Get(req, &r)

	fmt.Println(r.Fmt())
	return nil
}

// DataStoresCreate to get datastores list for workspaces.
func DataStoresCreate(cfg *config.Config, ws, name, path, method, format, configure string) error {
	url := fmt.Sprintf(datastoreCreateShpPattern, cfg.ServerURL(), ws, name, method, format, configure)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can't access such file: %w", err)
	}

	req := NewReqContain(cfg, "PUT", url, "application/zip", f)
	DoCreate(req)
	return nil
}

// DataStoreDelete to delete a datastore.
func DataStoreDelete(cfg *config.Config, ws, name string, r bool) error {
	url := fmt.Sprintf(datastorePattern, cfg.ServerURL(), ws, name)
	url += fmt.Sprintf("?recurse=%t", r)
	method := "DELETE"
	req := NewRequest(cfg, method, url, nil)
	Delete(req, datastoreDeleteStatus)
	return nil
}

// TODO: multi-type body
// DataStorePut to put an exist datastore.
// func DataStorePut() {

// }
