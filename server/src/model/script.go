package model

import "encoding/json"

func makeResult(old any, new any) error {
	oldJson, err := json.Marshal(old)
	if err != nil {
		return err
	}
	err = json.Unmarshal(oldJson, new)
	if err != nil {
		return err
	}
	return err
}
