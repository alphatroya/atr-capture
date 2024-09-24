package env

import (
	"fmt"
	"os"
)

const knowledgeBaseEnv = "KNOWLEDGE_BASE"

func CheckEnvs() error {
	_, ok := os.LookupEnv(knowledgeBaseEnv)
	if !ok {
		return fmt.Errorf("Failed to find \"%s\" env variable. It is REQUIRED for app", knowledgeBaseEnv)
	}
	return nil
}
