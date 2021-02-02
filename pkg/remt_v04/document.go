// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentRemt00100104 struct {
	Xmlns   string              `xml:"xmlns,attr"`
	RmtAdvc RemittanceAdviceV04 `xml:"RmtAdvc"`
}

func (doc DocumentRemt00100104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00100104) NameSpace() string {
	return utils.DocumentRemt00100104NameSpace
}

func (doc DocumentRemt00100104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RmtAdvc RemittanceAdviceV04 `xml:"RmtAdvc"`
	}
	output.RmtAdvc = doc.RmtAdvc
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
