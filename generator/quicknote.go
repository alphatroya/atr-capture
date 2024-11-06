package generator

import "time"

func GenerateQuickNoteTitle(t time.Time) string {
	return t.Format("20060102150405")
}
