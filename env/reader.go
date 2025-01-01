package env

import (
	"fmt"
	"os"
	"time"
)

type Envs struct {
	path string
}

func (e Envs) journalsFolder() string {
	return e.path + "journals/"
}

func (e Envs) PagesFolder() string {
	return e.path + "pages/"
}

func (e Envs) TodayJournalPath() string {
	currentDate := time.Now()
	return e.journalsFolder() + currentDate.Format("2006_01_02") + ".md"
}

func CheckEnvs() (Envs, error) {
	var e Envs
	var ok bool
	const knowledgeBaseEnv = "KNOWLEDGE_BASE"
	e.path, ok = os.LookupEnv(knowledgeBaseEnv)
	if !ok {
		return e, fmt.Errorf("Failed to find \"%s\" env variable. It is REQUIRED for app", knowledgeBaseEnv)
	}
	return e, nil
}
