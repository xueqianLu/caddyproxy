package types

type CreateWebsite struct {
	Domain   string `json:"domain"`
	Resource string `json:"resource"` // Resource url, it is a zip file, eg. http://example.com/website.zip
}

type ForwardWebsite struct {
	Domain string `json:"domain"`
	Target string `json:"target"`
}
