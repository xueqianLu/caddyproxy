package run

const (
	hostFlag        = "host"
	portFlag        = "port"
	caddyUrlFlag    = "caddy-url"
	caddyRootFlag   = "caddy-root"
	downloadDirFlag = "download-dir"
	logFlag         = "log"
)

type serviceParam struct {
	host        string
	port        int
	caddyUrl    string
	caddyRoot   string
	logPath     string
	downloadDir string
}

var (
	params = &serviceParam{
		host:        "0.0.0.0",
		port:        9000,
		caddyUrl:    "",
		caddyRoot:   "",
		logPath:     "/root/data/service.log",
		downloadDir: "/root/data/download",
	}
)
