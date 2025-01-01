package draft

import (
	"encoding/json"
	"fmt"
	"os"
)

func restoreDraft() (Draft, bool) {
	if _, err := os.Stat(draftLocation); os.IsNotExist(err) {
		return Draft{}, false
	}

	data, err := os.ReadFile(draftLocation)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return Draft{}, false
	}

	var draft Draft
	if err := json.Unmarshal(data, &draft); err != nil {
		fmt.Printf("error unmarshaling JSON to Draft: %v\n", err)
		return Draft{}, false
	}

	return draft, true
}

func RestoreOrNewDraft(requestConfirm func() bool) Draft {
	d, draftExist := restoreDraft()
	if !draftExist {
		return d
	}

	if requestConfirm() {
		return d
	}
	DropDraft()
	return Draft{}
}

func DropDraft() {
	if _, err := os.Stat(draftLocation); err == nil {
		err = os.Remove(draftLocation)
		if err != nil {
			fmt.Printf("error removing draft file at %s: %v\n", draftLocation, err)
		} else {
			fmt.Println("Draft file successfully removed.")
		}
	}
}
