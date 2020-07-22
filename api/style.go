package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/yuhangch/geoserver-cli/config"
)

const (
	//application/vnd.ogc.sld+xml for SLD 1.0.0 SLDs
	fmtVndOgcSldXML string = "application/vnd.ogc.sld+xml"
	//application/vnd.ogc.se+xml for SLD 1.1.0 SLDs
	fmtOgcSeXML string = "application/vnd.ogc.se+xml"
	//application/vnd.geoserver.geocss+css for css styles
	fmtVndGeoServerGeocssCSS string = "application/vnd.geoserver.geocss+css"
	//application/vnd.geoserver.ysld+yaml for ysld styles
	fmtVndGeoServerYsldYaml string = "pplication/vnd.geoserver.ysld+yaml"
	//application/vnd.geoserver.mbstyle+json for mb styles
	fmtVndGeoServerMbStyleJSON string = "application/vnd.geoserver.mbstyle+json"

	stylesRootPattern string = "%s/rest/styles"
	stylesWsPattern   string = "%s/rest/workspaces/%s/styles"
	stylePattern      string = "%s/rest/workspaces/%s/styles/%s"
)

var (
	styleFmtDict map[string]string = map[string]string{
		"sld": fmtVndOgcSldXML,
	}
)

// Style to represent style.
type Style struct {
	Name      string `json:"name"`
	Workspace struct {
		Name string `json:"name"`
	} `json:"workspace"`
	Format          string `json:"format"`
	LanguageVersion struct {
		Version string `json:"version"`
	} `json:"languageVersion"`
	Filename string `json:"filename"`
}

// StyleGet to get a style.
func StyleGet(cfg *config.Config, ws, name string) (Style, error) {
	url := fmt.Sprintf(stylePattern, cfg.ServerURL(), ws, name)
	method := "GET"

	req := NewRequest(cfg, method, url, nil)
	var s map[string]Style
	Get(req, &s)

	// fmt.Printf("%+v\n", s["style"])
	return s["style"], nil
}

// StyleGetBody to get a style body.
func StyleGetBody(cfg *config.Config, ws, name string) ([]byte, error) {
	st, err := StyleGet(cfg, ws, name)
	if err != nil {
		return nil, err
	}
	if accept, ok := styleFmtDict[st.Format]; ok {

		url := fmt.Sprintf(stylePattern, cfg.ServerURL(), ws, name)
		method := "GET"

		req := NewReqAccept(cfg, method, url, accept, nil)
		_, body, err := Do(req)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		return nil, fmt.Errorf("unknown style format")
	}

}

// StyleEdit edit a style.
func StyleEdit(cfg *config.Config, ws, name string) error {
	style, err := StyleGet(cfg, ws, name)
	if err != nil {
		return err
	}
	body, err := StyleGetBody(cfg, ws, name)
	if err != nil {
		return err
	}
	tfpath := os.TempDir() + "/GCTL_STYLE_EDIT"
	err = ioutil.WriteFile(tfpath, body, 0644)
	if err != nil {
		return err
	}
	cmdv := exec.Command("vi", tfpath)
	cmdv.Stdin = os.Stdin
	cmdv.Stdout = os.Stdout
	cmdv.Run()
	cmdv.Wait()

	url := fmt.Sprintf(stylePattern, cfg.ServerURL(), ws, name)
	method := "PUT"
	done, err := os.Open(tfpath)
	if err != nil {
		return fmt.Errorf("read style from vim error")
	}
	req := NewReqContain(cfg, method, url, styleFmtDict[style.Format], done)
	DoWithMsg(req, "Updated", "Failed")
	return nil
}
