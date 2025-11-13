package reader

import (
	"encoding/json"
	"fmt"
)

type Variant struct {
	Text string `json:"text"`
	Goto string `json:"gotoname"`
}

type Stage struct {
	Name     string    `json:"name"`
	Text     string    `json:"text"`
	Variants []Variant `json:"variants"`
}
type Novel struct {
	Name   string  `json:"name"`
	Stages []Stage `json:"stages"`
}

func R(jsonfile string) (Novel, error) {
	var novel Novel
	err := json.Unmarshal([]byte(jsonfile), &novel)
	if err != nil {
		return novel, fmt.Errorf("error with reading json file: %w", err)
	}

	return novel, nil
}
