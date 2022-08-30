package batch

import (
	"encoding/xml"
	"github.com/moov-io/fincen"
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
	StatusCode           string                   `xml:"StatusCode,attr"`
	TotalAmount          float64                  `xml:"TotalAmount,attr"`
	PartyCount           int64                    `xml:"PartyCount,attr"`
	ActivityCount        int64                    `xml:"ActivityCount,attr"`
	FormTypeCode         string                   `xml:"FormTypeCode,omitempty" json:",omitempty"`
	Activity             []fincen.ElementActivity `xml:"Activity,omitempty" json:",omitempty"`
	EFilingSubmissionXML *EFilingSubmissionXML    `xml:"EFilingSubmissionXML,omitempty" json:",omitempty"`
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

func (r EFilingBatchXML) Validate(args ...string) error {

	if r.StatusCode == "A" {
		// FinCEN XML Acknowledgement
		if err := r.fieldInclusionSubmission(); err != nil {
			return err
		}
	} else {
		// FinCEN XML Batch Reporting
		if fincen.CheckInvolved(r.FormTypeCode, "CTRX", "SARX", "DOEPX", "FBARX", "8300X") {
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
