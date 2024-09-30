package run

import (
	"caddyproxy/openapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func GetCommand() *cobra.Command {
	runtimeCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the server to provide api.",
		Run:   runCommand,
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	setFlags(runtimeCmd)

	return runtimeCmd
}

func setlog(path string) func() {
	if path == "" {
		return nil
	}
	// logrus log to file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return func() {
		file.Close()
	}
}

func setFlags(cmd *cobra.Command) {

	cmd.Flags().StringVar(
		&params.logPath,
		logFlag,
		"",
		"the log file path",
	)

	cmd.Flags().IntVar(
		&params.port,
		portFlag,
		9000,
		"the api service used port",
	)

	cmd.Flags().StringVar(
		&params.host,
		hostFlag,
		"0.0.0.0",
		"the api service used host",
	)

	cmd.Flags().StringVar(
		&params.downloadDir,
		downloadDirFlag,
		"/root/data/download",
		"the download directory",
	)

	cmd.Flags().StringVar(
		&params.caddyUrl,
		caddyUrlFlag,
		"",
		"caddy api url",
	)
	cmd.Flags().StringVar(
		&params.caddyRoot,
		caddyRootFlag,
		"",
		"caddy configure root path",
	)
}

func runCommand(cmd *cobra.Command, _ []string) {
	closeFunc := setlog(params.logPath)
	defer func() {
		if closeFunc != nil {
			closeFunc()
		}
	}()

	api := openapi.NewOpenAPI(&openapi.Config{
		Host:    params.host,
		Port:    params.port,
		TempDir: params.downloadDir,
	})
	if err := api.Run(); err != nil {
		log.WithError(err).Error("api service exit")
	}
}
