// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00300108 struct {
	Xmlns              string                       `xml:"xmlns,attr"`
	FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"FIToFICstmrDrctDbt"`
}

func (doc DocumentPacs00300108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00300108) NameSpace() string {
	return utils.DocumentPacs00300108NameSpace
}

func (doc DocumentPacs00300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"FIToFICstmrDrctDbt"`
	}
	output.FIToFICstmrDrctDbt = doc.FIToFICstmrDrctDbt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
