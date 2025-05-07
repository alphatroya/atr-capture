package env

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

type Config struct {
	Path string `json:"path"`
}

func (e Config) journalsFolder() string {
	return path.Join(e.Path, "journals")
}

func (e Config) pagesFolder() string {
	return path.Join(e.Path, "pages")
}

func (e Config) TodayJournalPath() string {
	currentDate := time.Now()
	return path.Join(e.journalsFolder(), currentDate.Format("2006_01_02")+".md")
}

func (e Config) PagePath(noteTitle string) string {
	return path.Join(e.pagesFolder(), noteTitle+".md")
}

func LoadConfig() (Config, string, error) {
	var e Config
	configDir, err := os.UserConfigDir()
	if err != nil {
		return e, "", fmt.Errorf("CheckEnvs: failed to find user config dir location, err=%w", err)
	}
	configPath := path.Join(configDir, "atr-capture", "config.json")
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return e, configPath, fmt.Errorf("CheckEnvs: failed to find config json file, err=%w", err)
	}

	if err := json.Unmarshal(jsonData, &e); err != nil {
		return e, configPath, fmt.Errorf("CheckEnvs: failed to unmarshal config file, err=%w", err)
	}
	return e, configPath, nil
}
