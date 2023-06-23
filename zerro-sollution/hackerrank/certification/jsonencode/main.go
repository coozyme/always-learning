package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type Manager struct {
	FullName       string `json:"full_name,omitemtpy"`
	Position       string `json:"position,omitemtpy"`
	Age            int32  `json:"age,omitemtpy"`
	YearsInCompany int32  `json:"years_in_company,omitemtpy"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	data, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	return reader, nil
}

func main() {
	m := &Manager{
		FullName:       "Ary",
		Position:       "BE",
		Age:            0,
		YearsInCompany: 0,
	}

	i, err := EncodeManager(m)

	if err != nil {
		log.Println("ERR", err)
	}
	b, _ := json.Marshal(m)
	i.Read(b)
}
