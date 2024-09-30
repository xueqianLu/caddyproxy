package caddy

type NewWebsiteCaddyFile struct {
	DomainName string
	RootDir    string
}

type StaticWebsiteCreatePayload struct {
	DomainName       string
	WebsiteCaddyFile string
	RootDir          string
}
