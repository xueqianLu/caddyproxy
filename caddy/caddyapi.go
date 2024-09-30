package caddy

import (
	"bytes"
	"caddyproxy/types"
	"caddyproxy/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"text/template"
)

type CaddyAPI struct {
	Url     string // URL of the Caddy API
	RootDir string // Root directory of the Caddy to store website files and .caddy files.
	client  *http.Client
}

func NewCaddyAPI(url string, root string) *CaddyAPI {
	return &CaddyAPI{Url: url, RootDir: root, client: &http.Client{}}
}

func (c *CaddyAPI) buildPayloadForCreateWebsite(domain string, root string) (string, error) {
	// build the payload for creating a new website.
	buffer := bytes.NewBufferString("")
	cfg := StaticWebsiteCreatePayload{
		DomainName:       domain,
		RootDir:          root,
		WebsiteCaddyFile: filepath.Join(c.RootDir, fmt.Sprintf("%s.caddy", domain)),
	}
	tmpl, err := template.New("test").Parse(staticWebsiteCaddyFileTempl)
	if err != nil {
		log.WithError(err).Error("failed to parse website caddyfile template")
		return "", err
	}
	err = tmpl.Execute(buffer, cfg)
	if err != nil {
		log.WithError(err).Error("failed to build website caddyfile")
		return "", err
	}
	return buffer.String(), nil
}

func (c *CaddyAPI) newCaddyFile(domain string, root string) error {
	buffer := bytes.NewBufferString("")
	caddypath := filepath.Join(c.RootDir, fmt.Sprintf("%s.caddy", domain))
	// build CaddyFileTemplate.
	cfg := NewWebsiteCaddyFile{
		DomainName: domain,
		RootDir:    root,
	}
	tmpl, err := template.New("test").Parse(staticWebsiteCaddyFileTempl)
	if err != nil {
		log.WithError(err).Error("failed to parse website caddyfile template")
		return err
	}
	err = tmpl.Execute(buffer, cfg)
	if err != nil {
		log.WithError(err).Error("failed to build website caddyfile")
		return err
	}

	if err = utils.CreateFile(caddypath, buffer.String()); err != nil {
		log.WithError(err).Error("failed to create website caddyfile")
		return err
	}
	return nil
}

func (c *CaddyAPI) CreateWebsite(domain string, resource string) error {
	// unzip resource to the website root directory.
	websiteRoot := filepath.Join(c.RootDir, "websites", domain)
	if err := utils.Unzip(resource, websiteRoot); err != nil {
		return err
	}
	// create caddyfile.
	if err := c.newCaddyFile(domain, websiteRoot); err != nil {
		return err
	}

	// build api payload.
	payload, err := c.buildPayloadForCreateWebsite(domain, websiteRoot)
	if err != nil {
		return err
	}

	// send put request to the caddy api.
	path := fmt.Sprintf("%s/config/apps/http/servers/srv0/routes/%d", c.Url, 0)
	if err := c.put(path, []byte(payload)); err != nil {
		return err
	}

	return nil
}

func (c *CaddyAPI) ForwardWebsite(param types.ForwardWebsite) error {
	// todo: create a new site in Caddy and set the target to the domain.
	return nil
}

func (c *CaddyAPI) put(path string, data []byte) error {
	// send a get request to the Caddy API.
	req, err := http.NewRequest("PUT", path, bytes.NewBuffer(data))
	if err != nil {
		log.WithError(err).Error("failed to create request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.WithError(err).Error("failed to send request")
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("failed to read response body")
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status":  resp.StatusCode,
			"message": string(body),
		}).Error("failed to put data")
		return fmt.Errorf("failed to put data")
	}
	return nil
}
