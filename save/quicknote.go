package save

import "time"

func generateQuickNoteTitle(t time.Time) string {
	return t.Format("20060102150405")
}
