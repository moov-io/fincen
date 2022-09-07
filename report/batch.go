// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package batch

import (
	"encoding/xml"
	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/cash_payments"
	"github.com/moov-io/fincen/currency_transaction"
	"github.com/moov-io/fincen/exempt_designation"
	"github.com/moov-io/fincen/financial_accounts"
	"github.com/moov-io/fincen/suspicious_activity"
)

// validating elements
var (
	_ fincen.Element = (*EFilingBatchXML)(nil)
	_ fincen.Element = (*EFilingSubmissionXML)(nil)
	_ fincen.Element = (*EFilingActivityXML)(nil)
	_ fincen.Element = (*EFilingActivityErrorXML)(nil)
)

type EFilingBatchXML struct {
	XMLName              xml.Name                 `xml:"EFilingBatchXML"`
	SeqNum               fincen.SeqNumber         `xml:"SeqNum,attr"`
	StatusCode           string                   `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	TotalAmount          float64                  `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
	PartyCount           int64                    `xml:"PartyCount,attr,omitempty" json:",omitempty"`
	ActivityCount        int64                    `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
	Attrs                []xml.Attr               `xml:",any,attr"`
	FormTypeCode         string                   `xml:"FormTypeCode,omitempty" json:",omitempty"`
	Activity             []fincen.ElementActivity `xml:"Activity,omitempty" json:",omitempty"`
	EFilingSubmissionXML *EFilingSubmissionXML    `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
}

type dummyXML struct {
	XMLName    xml.Name
	Attrs      []xml.Attr       `xml:",any,attr"`
	SeqNum     fincen.SeqNumber `xml:"SeqNum,attr"`
	StatusCode string           `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	Content    []byte           `xml:",innerxml"`
	Nodes      []dummyXML       `xml:",any"`
}

type batchDummy struct {
	XMLName              xml.Name              `xml:"EFilingBatchXML"`
	SeqNum               fincen.SeqNumber      `xml:"SeqNum,attr"`
	StatusCode           string                `xml:"StatusCode,attr,omitempty" json:",omitempty"`
	TotalAmount          float64               `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
	PartyCount           int64                 `xml:"PartyCount,attr,omitempty" json:",omitempty"`
	ActivityCount        int64                 `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
	Attrs                []xml.Attr            `xml:",any,attr"`
	FormTypeCode         string                `xml:"FormTypeCode,omitempty" json:",omitempty"`
	Activity             []dummyXML            `xml:"Activity,omitempty" json:",omitempty"`
	EFilingSubmissionXML *EFilingSubmissionXML `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
}

func (r *EFilingBatchXML) copy(org batchDummy) {
	// copy object
	r.XMLName = org.XMLName
	r.Attrs = org.Attrs
	r.SeqNum = org.SeqNum
	r.StatusCode = org.StatusCode
	r.TotalAmount = org.TotalAmount
	r.PartyCount = org.PartyCount
	r.ActivityCount = org.ActivityCount
	r.FormTypeCode = org.FormTypeCode
	r.EFilingSubmissionXML = org.EFilingSubmissionXML
}

func (r *EFilingBatchXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var dummy batchDummy
	if err := d.DecodeElement(&dummy, &start); err != nil {
		return err
	}

	r.copy(dummy)

	for _, act := range dummy.Activity {
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

func (r EFilingBatchXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := struct {
		XMLName              xml.Name                 `xml:"EFilingBatchXML"`
		SeqNum               fincen.SeqNumber         `xml:"SeqNum,attr"`
		StatusCode           string                   `xml:"StatusCode,attr,omitempty" json:",omitempty"`
		TotalAmount          float64                  `xml:"TotalAmount,attr,omitempty" json:",omitempty"`
		PartyCount           int64                    `xml:"PartyCount,attr,omitempty" json:",omitempty"`
		ActivityCount        int64                    `xml:"ActivityCount,attr,omitempty" json:",omitempty"`
		Attrs                []xml.Attr               `xml:",any,attr"`
		FormTypeCode         string                   `xml:"FormTypeCode,omitempty" json:",omitempty"`
		Activity             []fincen.ElementActivity `xml:"Activity,omitempty" json:",omitempty"`
		EFilingSubmissionXML *EFilingSubmissionXML    `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
	}(r)

	for index := 0; index < len(a.Attrs); index++ {
		switch a.Attrs[index].Name.Local {
		case "schemaLocation", "xsi", "fc2":
			a.Attrs = append(a.Attrs[:index], a.Attrs[index+1:]...)
			index--
		}
	}

	a.Attrs = append(a.Attrs, xml.Attr{
		Name: xml.Name{
			Local: "xsi:schemaLocation",
		},
		Value: "www.fincen.gov/base https://www.fincen.gov/base https://www.fincen.gov/base/EFL_8300XBatchSchema.xsd",
	})

	a.Attrs = append(a.Attrs, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:xsi",
		},
		Value: "http://www.w3.org/2001/XMLSchema-instance",
	})

	a.Attrs = append(a.Attrs, xml.Attr{
		Name: xml.Name{
			Local: "xsi:fc2",
		},
		Value: "www.fincen.gov/base",
	})

	return e.EncodeElement(&a, start)
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

func (r *EFilingBatchXML) GenerateAttrs() error {
	// The count of all <Activity> elements in the batch
	r.ActivityCount = int64(len(r.Activity))

	// The sum of all <DetailTransactionAmountText> element amounts recorded in the batch

	// The count of all <Party> elements in the batch where the
	// <ActivityPartyTypeCode> is equal to 16, 23, 4, 3, and 8 (combined)

	return nil
}

func (r EFilingBatchXML) Validate(args ...string) error {

	if r.StatusCode == "A" {
		// FinCEN XML Acknowledgement
		if err := r.fieldInclusionSubmission(); err != nil {
			return err
		}
	} else {

		// FinCEN XML Batch Reporting
		if !fincen.CheckInvolved(r.FormTypeCode, "CTRX", "SARX", "DOEPX", "FBARX", "8300X") {
			return fincen.NewErrValueInvalid("FormTypeCode")
		}

		if err := r.fieldInclusionReport(); err != nil {
			return err
		}
	}

	return fincen.Validate(&r, args...)
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

type EFilingActivityXML struct {
	XMLName                 xml.Name                  `xml:"EFilingActivityXML"`
	SeqNum                  fincen.SeqNumber          `xml:"SeqNum,attr"`
	BSAID                   fincen.RestrictNumeric14  `xml:"BSAID"`
	EFilingActivityErrorXML []EFilingActivityErrorXML `xml:"EFilingActivityErrorXML"`
}

func (r EFilingActivityXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
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

type constructorFunc func() fincen.ElementActivity

var (
	activityConstructor = map[string]constructorFunc{
		"CTRX":  func() fincen.ElementActivity { return &currency_transaction.ActivityType{} },
		"SARX":  func() fincen.ElementActivity { return &suspicious_activity.ActivityType{} },
		"DOEPX": func() fincen.ElementActivity { return &exempt_designation.ActivityType{} },
		"FBARX": func() fincen.ElementActivity { return &financial_accounts.ActivityType{} },
		"8300X": func() fincen.ElementActivity { return &cash_payments.ActivityType{} },
	}
)
