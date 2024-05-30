package setup

import (
	"flag"
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/config/logger_config"
	"gateway/pkg/config/path_config"
	"gateway/pkg/transport/http_handlers"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt"
	"os"
)

func AppSetup() {
	flag.Parse()
	setupLoggers()
	setPublicKey()
	http_handlers.Init()
}

func setupLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.InfoFileLogger.SetOutput(file)
	logger_config.ErrorFileLogger.SetOutput(file)
	logger_config.DebugFileLogger.SetOutput(file)
	setLoggersStyle()
}

func setLoggersStyle() {
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("[ERROR] [GATEWAY]").
		Foreground(lipgloss.Color("#ED1E26")).
		Bold(true)
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("[INFO] [GATEWAY]").
		Foreground(lipgloss.Color("#18B894")).
		Bold(true)
	log.DebugLevelStyle = lipgloss.NewStyle().
		SetString("[DEBUG] [GATEWAY]").
		Foreground(lipgloss.Color("#FFC60A")).
		Bold(true)
}

func setPublicKey() {
	bytes, err := os.ReadFile(*path_config.KeyPath)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	key, err := jwt.ParseECPublicKeyFromPEM(bytes)
	global_vars_config.PublicKey = key
}
