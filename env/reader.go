package env

import (
	"fmt"
	"os"
	"path"
	"time"
)

type Envs struct {
	path string
}

func (e Envs) journalsFolder() string {
	return path.Join(e.path, "journals")
}

func (e Envs) pagesFolder() string {
	return path.Join(e.path, "pages")
}

func (e Envs) TodayJournalPath() string {
	currentDate := time.Now()
	return path.Join(e.journalsFolder(), currentDate.Format("2006_01_02")+".md")
}

func (e Envs) PagePath(noteTitle string) string {
	return path.Join(e.pagesFolder(), noteTitle+".md")
}

func CheckEnvs() (Envs, error) {
	var e Envs
	var ok bool
	const knowledgeBaseEnv = "KNOWLEDGE_BASE"
	e.path, ok = os.LookupEnv(knowledgeBaseEnv)
	if !ok {
		return e, fmt.Errorf("failed to find \"%s\" env variable. It is REQUIRED for app", knowledgeBaseEnv)
	}
	return e, nil
}
