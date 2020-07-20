package api

import (
	"fmt"
	"os"

	"github.com/yuhangch/geoserver-cli/config"
)

const (
	datastoresPattern         string = "%s/rest/workspaces/%s/datastores"
	datastorePattern          string = "%s/rest/workspaces/%s/datastores/%s"
	datastoreCreateShpPattern string = "%s/rest/workspaces/%s/datastores/%s/file.shp"
)

var (
	datastoreDeleteStatus map[int]string = map[int]string{
		200: "Deleted",
		403: "Datastore or related is not empty (and recurse not true)",
		404: "Datastore not exist",
	}
)

// DataStoreResponse represent datastore list.
type DataStoreResponse struct {
	DataStores map[string][]Entry `json:"datastores"`
}

// Fmt to fmt print workspace list.
func (d *DataStoreResponse) Fmt() string {
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
	var r DataStoreResponse
	Get(req, &r)

	fmt.Println(r.Fmt())
	return nil
}

// DataStoresCreate to get datastores list for workspaces.
func DataStoresCreate(cfg *config.Config, ws, name, path string) error {
	url := fmt.Sprintf(datastoreCreateShpPattern, cfg.ServerURL(), ws, name)
	method := "PUT"
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("can't access such file: %w", err)
	}

	req := NewZipRequest(cfg, method, url, f)
	DoCreate(req)
	return nil
}

// DataStoreDelete to delete a datastore.
func DataStoreDelete(cfg *config.Config, ws, name string, r bool) error {
	url := fmt.Sprintf(datastorePattern, cfg.ServerURL(), ws, name)
	url += fmt.Sprintf("?recurse=%t", r)
	// fmt.Println(url)
	method := "DELETE"
	req := NewRequest(cfg, method, url, nil)
	Del(req, datastoreDeleteStatus)
	return nil
}

// TODO: multi-type body
// DataStorePut to put an exist datastore.
// func DataStorePut() {

// }
