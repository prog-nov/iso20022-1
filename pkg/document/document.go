// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package document

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/moov-io/iso20022/pkg/acmt_v01"
	"github.com/moov-io/iso20022/pkg/acmt_v02"
	"github.com/moov-io/iso20022/pkg/acmt_v03"
	"github.com/moov-io/iso20022/pkg/admi_v01"
	"github.com/moov-io/iso20022/pkg/admi_v02"
	"github.com/moov-io/iso20022/pkg/auth_v01"
	"github.com/moov-io/iso20022/pkg/auth_v02"
	"github.com/moov-io/iso20022/pkg/camt_v01"
	"github.com/moov-io/iso20022/pkg/camt_v03"
	"github.com/moov-io/iso20022/pkg/camt_v04"
	"github.com/moov-io/iso20022/pkg/camt_v05"
	"github.com/moov-io/iso20022/pkg/camt_v06"
	"github.com/moov-io/iso20022/pkg/camt_v07"
	"github.com/moov-io/iso20022/pkg/camt_v08"
	"github.com/moov-io/iso20022/pkg/camt_v09"
	"github.com/moov-io/iso20022/pkg/camt_v10"
	"github.com/moov-io/iso20022/pkg/head_v01"
	"github.com/moov-io/iso20022/pkg/head_v02"
	"github.com/moov-io/iso20022/pkg/pacs_v04"
	"github.com/moov-io/iso20022/pkg/pacs_v06"
	"github.com/moov-io/iso20022/pkg/pacs_v07"
	"github.com/moov-io/iso20022/pkg/pacs_v08"
	"github.com/moov-io/iso20022/pkg/pacs_v09"
	"github.com/moov-io/iso20022/pkg/pacs_v10"
	"github.com/moov-io/iso20022/pkg/pacs_v11"
	"github.com/moov-io/iso20022/pkg/pain_v01"
	"github.com/moov-io/iso20022/pkg/pain_v05"
	"github.com/moov-io/iso20022/pkg/pain_v08"
	"github.com/moov-io/iso20022/pkg/pain_v09"
	"github.com/moov-io/iso20022/pkg/pain_v10"
	"github.com/moov-io/iso20022/pkg/pain_v11"
	"github.com/moov-io/iso20022/pkg/reda_v01"
	"github.com/moov-io/iso20022/pkg/remt_v02"
	"github.com/moov-io/iso20022/pkg/remt_v04"
	"github.com/moov-io/iso20022/pkg/utils"
)

// Document interface for ISO 20022
type Iso20022Document interface {
	// Validate will be process validation check of document
	Validate() error

	// NameSpace check will return xmlns of document
	NameSpace() string
}

func NewDocumentBySpace(space string) Iso20022Document {
	if doc := newDocumentPart1(space); doc != nil {
		return doc
	}
	if doc := newDocumentPart2(space); doc != nil {
		return doc
	}
	if doc := newDocumentPart3(space); doc != nil {
		return doc
	}
	if doc := newDocumentPart4(space); doc != nil {
		return doc
	}

	return nil
}

func newDocumentPart1(space string) Iso20022Document {
	switch space {
	case utils.DocumentAcmt03600101NameSpace:
		return &acmt_v01.DocumentAcmt03600101{}
	case utils.DocumentAcmt02200102NameSpace:
		return &acmt_v02.DocumentAcmt02200102{}
	case utils.DocumentAcmt02300102NameSpace:
		return &acmt_v02.DocumentAcmt02300102{}
	case utils.DocumentAcmt02400102NameSpace:
		return &acmt_v02.DocumentAcmt02400102{}
	case utils.DocumentAcmt03000102NameSpace:
		return &acmt_v02.DocumentAcmt03000102{}
	case utils.DocumentAcmt03300102NameSpace:
		return &acmt_v02.DocumentAcmt03300102{}
	case utils.DocumentAcmt03500102NameSpace:
		return &acmt_v02.DocumentAcmt03500102{}
	case utils.DocumentAcmt03700102NameSpace:
		return &acmt_v02.DocumentAcmt03700102{}
	case utils.DocumentAcmt00700103NameSpace:
		return &acmt_v03.DocumentAcmt00700103{}
	case utils.DocumentAcmt00800103NameSpace:
		return &acmt_v03.DocumentAcmt00800103{}
	case utils.DocumentAcmt00900103NameSpace:
		return &acmt_v03.DocumentAcmt00900103{}
	case utils.DocumentAcmt01000103NameSpace:
		return &acmt_v03.DocumentAcmt01000103{}
	case utils.DocumentAcmt01100103NameSpace:
		return &acmt_v03.DocumentAcmt01100103{}
	case utils.DocumentAcmt01200103NameSpace:
		return &acmt_v03.DocumentAcmt01200103{}
	case utils.DocumentAcmt01300103NameSpace:
		return &acmt_v03.DocumentAcmt01300103{}
	case utils.DocumentAcmt01400103NameSpace:
		return &acmt_v03.DocumentAcmt01400103{}
	case utils.DocumentAcmt01500103NameSpace:
		return &acmt_v03.DocumentAcmt01500103{}
	case utils.DocumentAcmt01600103NameSpace:
		return &acmt_v03.DocumentAcmt01600103{}
	case utils.DocumentAcmt01700103NameSpace:
		return &acmt_v03.DocumentAcmt01700103{}
	case utils.DocumentAcmt01800103NameSpace:
		return &acmt_v03.DocumentAcmt01800103{}
	case utils.DocumentAcmt01900103NameSpace:
		return &acmt_v03.DocumentAcmt01900103{}
	case utils.DocumentAcmt02000103NameSpace:
		return &acmt_v03.DocumentAcmt02000103{}
	case utils.DocumentAcmt02100103NameSpace:
		return &acmt_v03.DocumentAcmt02100103{}
	case utils.DocumentAcmt02700103NameSpace:
		return &acmt_v03.DocumentAcmt02700103{}
	case utils.DocumentAcmt02800103NameSpace:
		return &acmt_v03.DocumentAcmt02800103{}
	case utils.DocumentAcmt02900103NameSpace:
		return &acmt_v03.DocumentAcmt02900103{}
	case utils.DocumentAcmt03100103NameSpace:
		return &acmt_v03.DocumentAcmt03100103{}
	case utils.DocumentAcmt03200103NameSpace:
		return &acmt_v03.DocumentAcmt03200103{}
	case utils.DocumentAcmt03400103NameSpace:
		return &acmt_v03.DocumentAcmt03400103{}
	case utils.DocumentAdmi00200101NameSpace:
		return &admi_v01.DocumentAdmi00200101{}
	case utils.DocumentAdmi00400101NameSpace:
		return &admi_v01.DocumentAdmi00400101{}
	case utils.DocumentAdmi00500101NameSpace:
		return &admi_v01.DocumentAdmi00500101{}
	}
	return nil
}

func newDocumentPart2(space string) Iso20022Document {
	switch space {
	case utils.DocumentAdmi00600101NameSpace:
		return &admi_v01.DocumentAdmi00600101{}
	case utils.DocumentAdmi00700101NameSpace:
		return &admi_v01.DocumentAdmi00700101{}
	case utils.DocumentAdmi01100101NameSpace:
		return &admi_v01.DocumentAdmi01100101{}
	case utils.DocumentAdmi01700101NameSpace:
		return &admi_v01.DocumentAdmi01700101{}
	case utils.DocumentAdmi00400102NameSpace:
		return &admi_v02.DocumentAdmi00400102{}
	case utils.DocumentAdmi00900102NameSpace:
		return &admi_v02.DocumentAdmi00900102{}
	case utils.DocumentAdmi01000102NameSpace:
		return &admi_v02.DocumentAdmi01000102{}
	case utils.DocumentAuth00100101NameSpace:
		return &auth_v01.DocumentAuth00100101{}
	case utils.DocumentAuth00200101NameSpace:
		return &auth_v01.DocumentAuth00200101{}
	case utils.DocumentAuth00300101NameSpace:
		return &auth_v01.DocumentAuth00300101{}
	case utils.DocumentAuth01800102NameSpace:
		return &auth_v02.DocumentAuth01800102{}
	case utils.DocumentAuth01900102NameSpace:
		return &auth_v02.DocumentAuth01900102{}
	case utils.DocumentAuth02000102NameSpace:
		return &auth_v02.DocumentAuth02000102{}
	case utils.DocumentAuth02100102NameSpace:
		return &auth_v02.DocumentAuth02100102{}
	case utils.DocumentAuth02200102NameSpace:
		return &auth_v02.DocumentAuth02200102{}
	case utils.DocumentAuth02300102NameSpace:
		return &auth_v02.DocumentAuth02300102{}
	case utils.DocumentAuth02400102NameSpace:
		return &auth_v02.DocumentAuth02400102{}
	case utils.DocumentAuth02500102NameSpace:
		return &auth_v02.DocumentAuth02500102{}
	case utils.DocumentAuth02600102NameSpace:
		return &auth_v02.DocumentAuth02600102{}
	case utils.DocumentAuth02700102NameSpace:
		return &auth_v02.DocumentAuth02700102{}
	case utils.DocumentCamt10100101NameSpace:
		return &camt_v01.DocumentCamt10100101{}
	case utils.DocumentCamt10200101NameSpace:
		return &camt_v01.DocumentCamt10200101{}
	case utils.DocumentCamt10300101NameSpace:
		return &camt_v01.DocumentCamt10300101{}
	case utils.DocumentCamt10400101NameSpace:
		return &camt_v01.DocumentCamt10400101{}
	case utils.DocumentCamt03500103NameSpace:
		return &camt_v03.DocumentCamt03500103{}
	case utils.DocumentCamt06900103NameSpace:
		return &camt_v03.DocumentCamt06900103{}
	case utils.DocumentCamt07100103NameSpace:
		return &camt_v03.DocumentCamt07100103{}
	case utils.DocumentCamt08600103NameSpace:
		return &camt_v03.DocumentCamt08600103{}
	case utils.DocumentCamt01300104NameSpace:
		return &camt_v04.DocumentCamt01300104{}
	case utils.DocumentCamt01400104NameSpace:
		return &camt_v04.DocumentCamt01400104{}
	case utils.DocumentCamt01500104NameSpace:
		return &camt_v04.DocumentCamt01500104{}
	case utils.DocumentCamt01600104NameSpace:
		return &camt_v04.DocumentCamt01600104{}
	case utils.DocumentCamt01700104NameSpace:
		return &camt_v04.DocumentCamt01700104{}
	case utils.DocumentCamt02000104NameSpace:
		return &camt_v04.DocumentCamt02000104{}
	case utils.DocumentCamt03200104NameSpace:
		return &camt_v04.DocumentCamt03200104{}
	case utils.DocumentCamt03800104NameSpace:
		return &camt_v04.DocumentCamt03800104{}
	case utils.DocumentCamt07000104NameSpace:
		return &camt_v04.DocumentCamt07000104{}
	case utils.DocumentCamt01800105NameSpace:
		return &camt_v05.DocumentCamt01800105{}
	}
	return nil
}

func newDocumentPart3(space string) Iso20022Document {
	switch space {
	case utils.DocumentCamt02500105NameSpace:
		return &camt_v05.DocumentCamt02500105{}
	case utils.DocumentCamt02600105NameSpace:
		return &camt_v05.DocumentCamt02600105{}
	case utils.DocumentCamt02800105NameSpace:
		return &camt_v05.DocumentCamt02800105{}
	case utils.DocumentCamt03000105NameSpace:
		return &camt_v05.DocumentCamt03000105{}
	case utils.DocumentCamt03500105NameSpace:
		return &camt_v05.DocumentCamt03500105{}
	case utils.DocumentCamt03600105NameSpace:
		return &camt_v05.DocumentCamt03600105{}
	case utils.DocumentCamt03900105NameSpace:
		return &camt_v05.DocumentCamt03900105{}
	case utils.DocumentCamt04600105NameSpace:
		return &camt_v05.DocumentCamt04600105{}
	case utils.DocumentCamt04800105NameSpace:
		return &camt_v05.DocumentCamt04800105{}
	case utils.DocumentCamt04900105NameSpace:
		return &camt_v05.DocumentCamt04900105{}
	case utils.DocumentCamt05000105NameSpace:
		return &camt_v05.DocumentCamt05000105{}
	case utils.DocumentCamt05100105NameSpace:
		return &camt_v05.DocumentCamt05100105{}
	case utils.DocumentCamt05600105NameSpace:
		return &camt_v05.DocumentCamt05600105{}
	case utils.DocumentCamt06000105NameSpace:
		return &camt_v05.DocumentCamt06000105{}
	case utils.DocumentCamt02100106NameSpace:
		return &camt_v06.DocumentCamt02100106{}
	case utils.DocumentCamt02400106NameSpace:
		return &camt_v06.DocumentCamt02400106{}
	case utils.DocumentCamt02900106NameSpace:
		return &camt_v06.DocumentCamt02900106{}
	case utils.DocumentCamt03100106NameSpace:
		return &camt_v06.DocumentCamt03100106{}
	case utils.DocumentCamt03300106NameSpace:
		return &camt_v06.DocumentCamt03300106{}
	case utils.DocumentCamt03400106NameSpace:
		return &camt_v06.DocumentCamt03400106{}
	case utils.DocumentCamt04700106NameSpace:
		return &camt_v06.DocumentCamt04700106{}
	case utils.DocumentCamt05700106NameSpace:
		return &camt_v06.DocumentCamt05700106{}
	case utils.DocumentCamt05800106NameSpace:
		return &camt_v06.DocumentCamt05800106{}
	case utils.DocumentCamt05900106NameSpace:
		return &camt_v06.DocumentCamt05900106{}
	case utils.DocumentCamt00300107NameSpace:
		return &camt_v07.DocumentCamt00300107{}
	case utils.DocumentCamt00900107NameSpace:
		return &camt_v07.DocumentCamt00900107{}
	case utils.DocumentCamt01100107NameSpace:
		return &camt_v07.DocumentCamt01100107{}
	case utils.DocumentCamt01200107NameSpace:
		return &camt_v07.DocumentCamt01200107{}
	case utils.DocumentCamt01900107NameSpace:
		return &camt_v07.DocumentCamt01900107{}
	case utils.DocumentCamt02300107NameSpace:
		return &camt_v07.DocumentCamt02300107{}
	case utils.DocumentCamt08700107NameSpace:
		return &camt_v07.DocumentCamt08700107{}
	case utils.DocumentCamt00400108NameSpace:
		return &camt_v08.DocumentCamt00400108{}
	case utils.DocumentCamt00500108NameSpace:
		return &camt_v08.DocumentCamt00500108{}
	case utils.DocumentCamt00600108NameSpace:
		return &camt_v08.DocumentCamt00600108{}
	case utils.DocumentCamt00700108NameSpace:
		return &camt_v08.DocumentCamt00700108{}
	case utils.DocumentCamt00800108NameSpace:
		return &camt_v08.DocumentCamt00800108{}
	case utils.DocumentCamt01000108NameSpace:
		return &camt_v08.DocumentCamt01000108{}
	case utils.DocumentCamt02600108NameSpace:
		return &camt_v08.DocumentCamt02600108{}
	case utils.DocumentCamt02700108NameSpace:
		return &camt_v08.DocumentCamt02700108{}
	case utils.DocumentCamt03700108NameSpace:
		return &camt_v08.DocumentCamt03700108{}
	case utils.DocumentCamt05200108NameSpace:
		return &camt_v08.DocumentCamt05200108{}
	case utils.DocumentCamt05300108NameSpace:
		return &camt_v08.DocumentCamt05300108{}
	case utils.DocumentCamt05400108NameSpace:
		return &camt_v08.DocumentCamt05400108{}
	case utils.DocumentCamt05500109NameSpace:
		return &camt_v09.DocumentCamt05500109{}
	case utils.DocumentCamt05600109NameSpace:
		return &camt_v09.DocumentCamt05600109{}
	case utils.DocumentCamt02800110NameSpace:
		return &camt_v10.DocumentCamt02800110{}
	case utils.DocumentCamt02900110NameSpace:
		return &camt_v10.DocumentCamt02900110{}
	case utils.DocumentHead00100101NameSpace:
		return &head_v01.BusinessApplicationHeaderV01{}
	case utils.DocumentHead00100102NameSpace:
		return &head_v02.BusinessApplicationHeaderV02{}
	}
	return nil
}

func newDocumentPart4(space string) Iso20022Document {
	switch space {
	case utils.DocumentPacs01000104NameSpace:
		return &pacs_v04.DocumentPacs01000104{}
	case utils.DocumentPacs02800104NameSpace:
		return &pacs_v04.DocumentPacs02800104{}
	case utils.DocumentPacs00800106NameSpace:
		return &pacs_v06.DocumentPacs00800106{}
	case utils.DocumentPacs00200107NameSpace:
		return &pacs_v07.DocumentPacs00200107{}
	case utils.DocumentPacs00300108NameSpace:
		return &pacs_v08.DocumentPacs00300108{}
	case utils.DocumentPacs00800109NameSpace:
		return &pacs_v09.DocumentPacs00800109{}
	case utils.DocumentPacs00900109NameSpace:
		return &pacs_v09.DocumentPacs00900109{}
	case utils.DocumentPacs00200110NameSpace:
		return &pacs_v10.DocumentPacs00200110{}
	case utils.DocumentPacs00400110NameSpace:
		return &pacs_v10.DocumentPacs00400110{}
	case utils.DocumentPacs00700110NameSpace:
		return &pacs_v10.DocumentPacs00700110{}
	case utils.DocumentPacs00200111NameSpace:
		return &pacs_v11.DocumentPacs00200111{}
	case utils.DocumentPain00700101NameSpace:
		return &pain_v01.DocumentPain00700101{}
	case utils.DocumentPain01800101NameSpace:
		return &pain_v01.DocumentPain01800101{}
	case utils.DocumentPain00900105NameSpace:
		return &pain_v05.DocumentPain00900105{}
	case utils.DocumentPain01000105NameSpace:
		return &pain_v05.DocumentPain01000105{}
	case utils.DocumentPain01100105NameSpace:
		return &pain_v05.DocumentPain01100105{}
	case utils.DocumentPain01200105NameSpace:
		return &pain_v05.DocumentPain01200105{}
	case utils.DocumentPain01300105NameSpace:
		return &pain_v05.DocumentPain01300105{}
	case utils.DocumentPain01400105NameSpace:
		return &pain_v05.DocumentPain01400105{}
	case utils.DocumentPain01400108NameSpace:
		return &pain_v08.DocumentPain01400108{}
	case utils.DocumentPain01300108NameSpace:
		return &pain_v08.DocumentPain01300108{}
	case utils.DocumentPain00800109NameSpace:
		return &pain_v09.DocumentPain00800109{}
	case utils.DocumentPain00100110NameSpace:
		return &pain_v10.DocumentPain00100110{}
	case utils.DocumentPain00700110NameSpace:
		return &pain_v10.DocumentPain00700110{}
	case utils.DocumentPain00200111NameSpace:
		return &pain_v11.DocumentPain00200111{}
	case utils.DocumentReda06600101NameSpace:
		return &reda_v01.DocumentReda06600101{}
	case utils.DocumentReda06700101NameSpace:
		return &reda_v01.DocumentReda06700101{}
	case utils.DocumentReda06800101NameSpace:
		return &reda_v01.DocumentReda06800101{}
	case utils.DocumentReda06900101NameSpace:
		return &reda_v01.DocumentReda06900101{}
	case utils.DocumentReda07000101NameSpace:
		return &reda_v01.DocumentReda07000101{}
	case utils.DocumentReda07100101NameSpace:
		return &reda_v01.DocumentReda07100101{}
	case utils.DocumentReda07200101NameSpace:
		return &reda_v01.DocumentReda07200101{}
	case utils.DocumentReda07300101NameSpace:
		return &reda_v01.DocumentReda07300101{}
	case utils.DocumentRemt00100102NameSpace:
		return &remt_v02.DocumentRemt00100102{}
	case utils.DocumentRemt00200102NameSpace:
		return &remt_v02.DocumentRemt00200102{}
	case utils.DocumentRemt00100104NameSpace:
		return &remt_v04.DocumentRemt00100104{}
	}
	return nil
}

type documentDummy struct {
	XMLName xml.Name
	Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

// ParseIso20022Document will return a interface of ISO 20022 document after pass buffer
func ParseIso20022Document(buf []byte) (doc Iso20022Document, err error) {
	bType := utils.GetBufferFormat(buf)
	if bType == utils.DocumentTypeUnknown {
		return nil, utils.NewErrInvalidFileType()
	}

	var dummy documentDummy
	if bType == utils.DocumentTypeJson {
		err = json.Unmarshal(buf, &dummy)
	} else if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, &dummy)
	}
	if err != nil {
		return nil, err
	}

	// getting space
	// 1. getting from xml name
	var namespace string
	if len(dummy.XMLName.Space) > 0 {
		namespace = dummy.XMLName.Space
	}

	// 2. getting from xml attr
	if len(namespace) == 0 {
		for _, attr := range dummy.Attrs {
			if attr.Name.Local == utils.XmlDefaultNamespace {
				namespace = attr.Value
			}
		}
	}

	// 3. getting from xml attr with prefix
	if len(namespace) == 0 {
		for _, attr := range dummy.Attrs {
			if strings.Contains(attr.Name.Local, utils.XmlDefaultNamespace+":") {
				namespace = attr.Value
			}
		}
	}

	if namespace == "" {
		return nil, utils.NewErrOmittedNameSpace()
	}

	doc = NewDocumentBySpace(namespace)
	if doc == nil {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	if bType == utils.DocumentTypeJson {
		err = json.Unmarshal(buf, doc)
	} else if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, doc)
	}

	if err != nil {
		return nil, err
	}

	return
}
