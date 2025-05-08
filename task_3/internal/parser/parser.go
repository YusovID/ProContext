package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"task_3/internal/model"

	"golang.org/x/net/html/charset"
)

func ParseXML(xmlData []byte) (*model.ValCurs, error) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs model.ValCurs
	err := decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML: %v", err)
	}

	return &valCurs, nil
}
