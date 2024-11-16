package env

import (
	"fmt"
	"os"
	"time"
)

const knowledgeBaseEnv = "KNOWLEDGE_BASE"

type Envs struct {
	path string
}

const journalsFolder = "journals/"
const pagesFolder = "pages/"

func (e Envs) TodayJournalPath() string {
	currentDate := time.Now()
	return e.path + journalsFolder + currentDate.Format("2006_01_02") + ".md"
}

func (e Envs) PagesPath() string {
	return e.path + pagesFolder
}

func CheckEnvs() (Envs, error) {
	var e Envs
	var ok bool
	e.path, ok = os.LookupEnv(knowledgeBaseEnv)
	if !ok {
		return e, fmt.Errorf("Failed to find \"%s\" env variable. It is REQUIRED for app", knowledgeBaseEnv)
	}
	return e, nil
}
