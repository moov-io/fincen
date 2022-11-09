// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package batch

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/cash_payments"
	"github.com/moov-io/fincen/pkg/currency_transaction"
	"github.com/moov-io/fincen/pkg/exempt_designation"
	"github.com/moov-io/fincen/pkg/financial_accounts"
	"github.com/moov-io/fincen/pkg/suspicious_activity"
)

func NewReport(args ...string) *EFilingBatchXML {

	reportXml := EFilingBatchXML{}

	rType := "UNKNOWN"
	if len(args) > 0 {
		rType = args[0]
	}

	if rType == fincen.ReportSubmission {
		reportXml.StatusCode = "A"
	} else {
		if fincen.CheckInvolved(rType, fincen.Report112, fincen.Report111, fincen.Report110, fincen.Report114, fincen.Form8300) {
			reportXml.FormTypeCode = rType
		}
	}

	return &reportXml
}

func CreateReportWithBuffer(buf []byte) (*EFilingBatchXML, error) {

	reportXml := EFilingBatchXML{}

	err := xml.Unmarshal(buf, &reportXml)
	if err == nil {
		return &reportXml, nil
	}

	err = json.Unmarshal(buf, &reportXml)
	if err == nil {
		return &reportXml, nil
	}

	return nil, errors.New("unable to create batch, invalid input data")
}

func CreateReportWithFile(path string) (*EFilingBatchXML, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", path, err)
	}

	r, err := CreateReportWithBuffer(raw)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file: %w", err)
	}

	return r, nil
}

type EFilingBatchXML struct {
	XMLName                 xml.Name                 `xml:"EFilingBatchXML"`
	SeqNum                  fincen.SeqNumber         `xml:"SeqNum,attr,omitempty" json:",omitempty"`
	StatusCode              string                   `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	TotalAmount             float64                  `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
	PartyCount              int64                    `xml:"PartyCount,attr,omitempty" json:",omitempty"`
	ActivityCount           int64                    `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
	AccountCount            int64                    `xml:"AccountCount,attr,omitempty" json:",omitempty"`
	ActivityAttachmentCount int64                    `xml:"ActivityAttachmentCount,attr,omitempty" json:",omitempty"`
	AttachmentCount         int64                    `xml:"AttachmentCount,attr,omitempty" json:",omitempty"`
	JointlyOwnedOwnerCount  int64                    `xml:"JointlyOwnedOwnerCount,attr,omitempty" json:",omitempty"`
	NoFIOwnerCount          int64                    `xml:"NoFIOwnerCount,attr,omitempty" json:",omitempty"`
	ConsolidatedOwnerCount  int64                    `xml:"ConsolidatedOwnerCount,attr,omitempty" json:",omitempty"`
	Attrs                   []xml.Attr               `xml:",any,attr"`
	FormTypeCode            string                   `xml:"FormTypeCode,omitempty" json:",omitempty"`
	Activity                []fincen.ElementActivity `xml:"Activity,omitempty" json:",omitempty"`
	EFilingSubmissionXML    *EFilingSubmissionXML    `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
}

type dummyXML struct {
	XMLName    xml.Name
	Attrs      []xml.Attr       `xml:",any,attr"`
	SeqNum     fincen.SeqNumber `xml:"SeqNum,attr"`
	StatusCode string           `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	Content    []byte           `xml:",innerxml"`
}

type EFilingBatchUnmarshal struct {
	XMLName                 xml.Name              `xml:"EFilingBatchXML"`
	SeqNum                  fincen.SeqNumber      `xml:"SeqNum,attr,omitempty" json:",omitempty"`
	StatusCode              string                `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	TotalAmount             float64               `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
	PartyCount              int64                 `xml:"PartyCount,attr,omitempty" json:",omitempty"`
	ActivityCount           int64                 `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
	AccountCount            int64                 `xml:"AccountCount,attr,omitempty" json:",omitempty"`
	ActivityAttachmentCount int64                 `xml:"ActivityAttachmentCount,attr,omitempty" json:",omitempty"`
	AttachmentCount         int64                 `xml:"AttachmentCount,attr,omitempty" json:",omitempty"`
	JointlyOwnedOwnerCount  int64                 `xml:"JointlyOwnedOwnerCount,attr,omitempty" json:",omitempty"`
	NoFIOwnerCount          int64                 `xml:"NoFIOwnerCount,attr,omitempty" json:",omitempty"`
	ConsolidatedOwnerCount  int64                 `xml:"ConsolidatedOwnerCount,attr,omitempty" json:",omitempty"`
	Attrs                   []xml.Attr            `xml:",any,attr"`
	FormTypeCode            string                `xml:"FormTypeCode,omitempty" json:",omitempty"`
	Activity                []dummyXML            `xml:"Activity,omitempty" json:",omitempty"`
	EFilingSubmissionXML    *EFilingSubmissionXML `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
}

type batchAttr struct {
	TotalAmount             float64
	PartyCount              int64
	ActivityCount           int64
	AccountCount            int64
	ActivityAttachmentCount int64
	AttachmentCount         int64
	JointlyOwnedOwnerCount  int64
	NoFIOwnerCount          int64
	ConsolidatedOwnerCount  int64
}

func (r *EFilingBatchXML) copy(org EFilingBatchUnmarshal) {
	// copy object
	r.XMLName = org.XMLName
	r.SeqNum = org.SeqNum
	r.Attrs = org.Attrs
	r.StatusCode = org.StatusCode
	r.TotalAmount = org.TotalAmount
	r.PartyCount = org.PartyCount
	r.ActivityCount = org.ActivityCount
	r.AccountCount = org.AccountCount
	r.ActivityAttachmentCount = org.ActivityAttachmentCount
	r.AttachmentCount = org.AttachmentCount
	r.JointlyOwnedOwnerCount = org.JointlyOwnedOwnerCount
	r.NoFIOwnerCount = org.NoFIOwnerCount
	r.ConsolidatedOwnerCount = org.ConsolidatedOwnerCount
	r.FormTypeCode = org.FormTypeCode
	r.EFilingSubmissionXML = org.EFilingSubmissionXML
}

func (r *EFilingBatchXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var dummy EFilingBatchUnmarshal
	if err := d.DecodeElement(&dummy, &start); err != nil {
		return err
	}

	r.copy(dummy)

	for i := range dummy.Activity {
		act := dummy.Activity[i]

		buf, err := xml.Marshal(&act)
		if err != nil {
			return fincen.NewErrValueInvalid("Activity")
		}

		constructor := activityConstructor[r.FormTypeCode]
		if constructor == nil {
			return fincen.NewErrValueInvalid("FormTypeCode")
		}

		elm := constructor()
		if err = xml.Unmarshal(buf, elm); err != nil {
			return fincen.NewErrValueInvalid("Activity")
		}

		r.Activity = append(r.Activity, elm)
	}

	return nil
}

type EFilingBatchMarshal struct {
	XMLName                 xml.Name              `xml:"EFilingBatchXML"`
	SeqNum                  fincen.SeqNumber      `xml:"SeqNum,attr,omitempty" json:",omitempty"`
	StatusCode              string                `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	TotalAmount             float64               `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
	PartyCount              int64                 `xml:"PartyCount,attr,omitempty" json:",omitempty"`
	ActivityCount           int64                 `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
	AccountCount            int64                 `xml:"AccountCount,attr,omitempty" json:",omitempty"`
	ActivityAttachmentCount int64                 `xml:"ActivityAttachmentCount,attr,omitempty" json:",omitempty"`
	AttachmentCount         int64                 `xml:"AttachmentCount,attr,omitempty" json:",omitempty"`
	JointlyOwnedOwnerCount  int64                 `xml:"JointlyOwnedOwnerCount,attr,omitempty" json:",omitempty"`
	NoFIOwnerCount          int64                 `xml:"NoFIOwnerCount,attr,omitempty" json:",omitempty"`
	ConsolidatedOwnerCount  int64                 `xml:"ConsolidatedOwnerCount,attr,omitempty" json:",omitempty"`
	Attrs                   []xml.Attr            `xml:",any,attr"`
	FormTypeCode            string                `xml:"fc2:FormTypeCode,omitempty" json:",omitempty"`
	ActivitiesContent       []byte                `xml:",innerxml"`
	EFilingSubmissionXML    *EFilingSubmissionXML `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
}

func (r EFilingBatchXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	dummy := EFilingBatchMarshal{
		XMLName:                 r.XMLName,
		SeqNum:                  r.SeqNum,
		StatusCode:              r.StatusCode,
		TotalAmount:             r.TotalAmount,
		PartyCount:              r.PartyCount,
		ActivityCount:           r.ActivityCount,
		AccountCount:            r.AccountCount,
		ActivityAttachmentCount: r.ActivityAttachmentCount,
		AttachmentCount:         r.AttachmentCount,
		JointlyOwnedOwnerCount:  r.JointlyOwnedOwnerCount,
		NoFIOwnerCount:          r.NoFIOwnerCount,
		ConsolidatedOwnerCount:  r.ConsolidatedOwnerCount,
		Attrs:                   r.Attrs,
		FormTypeCode:            r.FormTypeCode,
		EFilingSubmissionXML:    r.EFilingSubmissionXML,
	}

	for index := 0; index < len(dummy.Attrs); index++ {
		switch dummy.Attrs[index].Name.Local {
		case "schemaLocation", "xsi", "fc2":
			dummy.Attrs = append(dummy.Attrs[:index], dummy.Attrs[index+1:]...)
			index--
		}
	}

	switch dummy.FormTypeCode {
	case fincen.Form8300:
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xsi:schemaLocation",
			},
			Value: "www.server.gov/base https://www.fincen.gov/base/EFL_8300XBatchSchema.xsd",
		})
	case fincen.Report110:
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xsi:schemaLocation",
			},
			Value: "www.fincen.gov/base https://www.fincen.gov/base/EFL_DOEPXBatchSchema.xsd",
		})
	case fincen.Report112:
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xsi:schemaLocation",
			},
			Value: "www.fincen.gov/base https://www.fincen.gov/sites/default/files/schema/base/EFL_CTRXBatchSchema.xsd",
		})
	case fincen.Report111:
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xsi:schemaLocation",
			},
			Value: "www.fincen.gov/base https://www.fincen.gov/base/EFL_FinCENSARXBatchSchema.xsd",
		})
	case fincen.Report114:
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xsi:schemaLocation",
			},
			Value: "www.fincen.gov/base/EFL_FBARXBatchSchema.xsd",
		})
	}

	if fincen.CheckInvolved(r.FormTypeCode, fincen.Report112, fincen.Report111, fincen.Report110, fincen.Report114, fincen.Form8300) {
		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xmlns:xsi",
			},
			Value: "http://www.w3.org/2001/XMLSchema-instance",
		})

		dummy.Attrs = append(dummy.Attrs, xml.Attr{
			Name: xml.Name{
				Local: "xmlns:fc2",
			},
			Value: "www.fincen.gov/base",
		})
		// Batch report don't support sequence number
		dummy.SeqNum = 0
		// force namespace prefix
		start.Name.Local = "fc2:EFilingBatchXML"
	}

	for _, act := range r.Activity {
		content := []byte{'\n'}
		if converted, err := fincen.MarshalIndent(act, fincen.DefaultXMLIntent, fincen.DefaultXMLIntent); err == nil {
			content = append(content, converted...)
		}
		if len(content) > 1 {
			dummy.ActivitiesContent = content
		}
	}

	return e.EncodeElement(&dummy, start)
}

func (r *EFilingBatchXML) AppendActivity(act fincen.ElementActivity) error {
	if act == nil {
		return errors.New("invalid activity")
	}

	if !fincen.CheckInvolved(r.FormTypeCode, fincen.Report112, fincen.Report111, fincen.Report110, fincen.Report114, fincen.Form8300) ||
		r.FormTypeCode != act.FormTypeCode() {
		return errors.New("invalid form type")
	}

	r.Activity = append(r.Activity, act)

	return nil
}

func (r *EFilingBatchXML) fieldInclusionReport() error {

	if len(r.Activity) < 1 {
		return fincen.NewErrValueInvalid("Activity")
	}

	return nil
}

func (r *EFilingBatchXML) fieldInclusionSubmission() error {
	if r.EFilingSubmissionXML == nil {
		return fincen.NewErrValueInvalid("EFilingSubmissionXML")
	}

	return nil
}

func (r EFilingBatchXML) generateAttrs() batchAttr {
	s := batchAttr{}
	// The count of all <Activity> elements in the batch
	s.ActivityCount = int64(len(r.Activity))

	for _, activity := range r.Activity {
		s.TotalAmount += activity.TotalAmount()

		switch r.FormTypeCode {
		case fincen.Form8300:
			// The count of all <Party> elements in the batch where the
			// <ActivityPartyTypeCode> is equal to 16, 23, 4, 3, and 8 (combined)
			s.PartyCount += activity.PartyCount("16", "23", "4", "3", "8")
		case fincen.Report110:
			// The count of all <Party> elements in the batch where the
			//<ActivityPartyTypeCode> is equal to 3, 11, 12, and 45 (combined)
			s.PartyCount += activity.PartyCount("3", "11", "12", "45")
		case fincen.Report112:
			// The total count of all <Party> elements recorded in the batch file.
			s.PartyCount += activity.PartyCount()
		case fincen.Report111:
			// The count of all <Party> elements in the batch where the
			// <ActivityPartyTypeCode> is equal to “33” (Subject)
			s.PartyCount += activity.PartyCount("33")

		case fincen.Report114:
			// The total count of <Party> elements where the <ActivityPartyTypeCode> element is equal to “41”
			s.PartyCount += activity.PartyCount("41")

			// AccountCount. The total count of all <Account> elements recorded in the batch file.
			if elm, ok := activity.(*financial_accounts.ActivityType); ok {
				s.AccountCount = int64(len(elm.Account))
			}

			// The total count of <Party> elements where the <ActivityPartyTypeCode>
			// element is equal to “42”
			s.JointlyOwnedOwnerCount += activity.PartyCount("42")

			// The total count of <Party> elements where the <ActivityPartyTypeCode> element is
			// equal to “43”
			s.NoFIOwnerCount += activity.PartyCount("43")

			// The total count of <Party> elements where the <ActivityPartyTypeCode>
			// equal to “44”
			s.NoFIOwnerCount += activity.PartyCount("44")
		}

	}

	return s
}

func (r *EFilingBatchXML) GenerateAttrs() error {

	s := r.generateAttrs()

	r.ActivityCount = s.ActivityCount
	r.TotalAmount = s.TotalAmount
	r.PartyCount = s.PartyCount
	r.ActivityAttachmentCount = s.ActivityAttachmentCount
	r.AttachmentCount = s.AttachmentCount
	r.JointlyOwnedOwnerCount = s.JointlyOwnedOwnerCount
	r.NoFIOwnerCount = s.NoFIOwnerCount
	r.ConsolidatedOwnerCount = s.ConsolidatedOwnerCount

	return nil
}

func (r *EFilingBatchXML) validateAttrs() error {

	s := r.generateAttrs()

	if r.ActivityCount != s.ActivityCount {
		return fincen.NewErrValueInvalid("ActivityCount")
	}

	if r.AccountCount != s.AccountCount {
		return fincen.NewErrValueInvalid("AccountCount")
	}

	if r.TotalAmount != s.TotalAmount {
		return fincen.NewErrValueInvalid("TotalAmount")
	}

	if r.PartyCount != s.PartyCount {
		return fincen.NewErrValueInvalid("PartyCount")
	}

	if r.ActivityAttachmentCount != s.ActivityAttachmentCount {
		return fincen.NewErrValueInvalid("ActivityAttachmentCount")
	}

	if r.AttachmentCount != s.AttachmentCount {
		return fincen.NewErrValueInvalid("AttachmentCount")
	}

	if r.JointlyOwnedOwnerCount != s.JointlyOwnedOwnerCount {
		return fincen.NewErrValueInvalid("JointlyOwnedOwnerCount")
	}

	if r.NoFIOwnerCount != s.NoFIOwnerCount {
		return fincen.NewErrValueInvalid("NoFIOwnerCount")
	}

	if r.ConsolidatedOwnerCount != s.ConsolidatedOwnerCount {
		return fincen.NewErrValueInvalid("ConsolidatedOwnerCount")
	}

	return nil
}

// Validate args:
//
//	1: disableValidateAttrs
func (r EFilingBatchXML) Validate(args ...string) error {

	if r.StatusCode == "A" {
		// FinCEN XML Acknowledgement
		if err := r.fieldInclusionSubmission(); err != nil {
			return err
		}
	} else {

		// FinCEN XML Batch Reporting
		if !fincen.CheckInvolved(r.FormTypeCode, fincen.Report112, fincen.Report111, fincen.Report110, fincen.Report114, fincen.Form8300) {
			return fincen.NewErrValueInvalid("FormTypeCode")
		}

		if err := r.fieldInclusionReport(); err != nil {
			return err
		}
	}

	if _, err := fincen.ValidateSeqNumbers(&r); err != nil {
		return err
	}

	if len(args) == 0 {
		if err := r.validateAttrs(); err != nil {
			return err
		}
	}

	return fincen.Validate(&r, args...)
}

func (r EFilingBatchXML) GenerateSeqNumbers(args ...int) error {
	seqNum := fincen.SeqNumber(1)
	if len(args) > 0 {
		seqNum = fincen.SeqNumber(args[0])
	}
	return fincen.GenerateSeqNumbers(&r, &seqNum)
}

func (r *EFilingBatchXML) SetSeqNumber(number fincen.SeqNumber) error {
	return errors.New("unsupported sequence number")
}

type EFilingSubmissionXML struct {
	XMLName            xml.Name             `xml:"EFilingSubmissionXML"`
	SeqNum             fincen.SeqNumber     `xml:"SeqNum,attr"`
	StatusCode         string               `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	EFilingActivityXML []EFilingActivityXML `xml:"EFilingActivityXML"`
}

func (r EFilingSubmissionXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

func (r *EFilingSubmissionXML) SetSeqNumber(number fincen.SeqNumber) error {
	r.SeqNum = number
	return nil
}

type EFilingActivityXML struct {
	XMLName                 xml.Name                  `xml:"EFilingActivityXML"`
	SeqNum                  fincen.SeqNumber          `xml:"SeqNum,attr"`
	BSAID                   fincen.RestrictNumeric14  `xml:"BSAID"`
	EFilingActivityErrorXML []EFilingActivityErrorXML `xml:"EFilingActivityErrorXML"`
}

func (r EFilingActivityXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

func (r *EFilingActivityXML) SetSeqNumber(number fincen.SeqNumber) error {
	r.SeqNum = number
	return nil
}

type EFilingActivityErrorXML struct {
	XMLName              xml.Name                   `xml:"EFilingActivityErrorXML"`
	SeqNum               fincen.SeqNumber           `xml:"SeqNum,attr"`
	ErrorContextText     *fincen.RestrictString4000 `xml:"ErrorContextText,omitempty" json:",omitempty"`
	ErrorElementNameText *fincen.RestrictString512  `xml:"ErrorElementNameText,omitempty" json:",omitempty"`
	ErrorLevelText       *fincen.RestrictString50   `xml:"ErrorLevelText,omitempty" json:",omitempty"`
	ErrorText            *fincen.RestrictString525  `xml:"ErrorText,omitempty" json:",omitempty"`
	ErrorTypeCode        *fincen.RestrictString50   `xml:"ErrorTypeCode,omitempty" json:",omitempty"`
}

func (r EFilingActivityErrorXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

func (r *EFilingActivityErrorXML) SetSeqNumber(number fincen.SeqNumber) error {
	r.SeqNum = number
	return nil
}

type constructorFunc func() fincen.ElementActivity

var (
	activityConstructor = map[string]constructorFunc{
		fincen.Report112: func() fincen.ElementActivity { return &currency_transaction.ActivityType{} },
		fincen.Report111: func() fincen.ElementActivity { return &suspicious_activity.ActivityType{} },
		fincen.Report110: func() fincen.ElementActivity { return &exempt_designation.ActivityType{} },
		fincen.Report114: func() fincen.ElementActivity { return &financial_accounts.ActivityType{} },
		fincen.Form8300:  func() fincen.ElementActivity { return &cash_payments.ActivityType{} },
	}
)
