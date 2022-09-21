// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package batch

import (
	"encoding/xml"
	"os"
	"path"
	"testing"

	"github.com/moov-io/fincen"
	"github.com/stretchr/testify/require"
)

// validating elements
var (
	_ fincen.Element = (*EFilingBatchXML)(nil)
	_ fincen.Element = (*EFilingSubmissionXML)(nil)
	_ fincen.Element = (*EFilingActivityXML)(nil)
	_ fincen.Element = (*EFilingActivityErrorXML)(nil)
)

type emptyElementActivity struct{}

func (r *emptyElementActivity) Validate(args ...string) error {
	return nil
}

func (r *emptyElementActivity) FormTypeCode() string {
	return ""
}

func (r *emptyElementActivity) TotalAmount() float64 {
	return 0
}

func (r *emptyElementActivity) PartyCount(args ...string) int64 {
	return 0
}

func TestAcknowledgement(t *testing.T) {
	t.Run("FinCEN SAR XML Acknowledgement", func(t *testing.T) {
		buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "sar_acknowledgement.xml"))
		require.NoError(t, err)

		batch := EFilingBatchXML{}
		err = xml.Unmarshal(buf, &batch)
		require.NoError(t, err)
		require.NotNil(t, batch.EFilingSubmissionXML)
		require.Equal(t, 3, len(batch.EFilingSubmissionXML.EFilingActivityXML))

		require.Equal(t, fincen.SeqNumber(1), batch.SeqNum)
		require.Equal(t, "A", batch.StatusCode)

		require.Equal(t, fincen.SeqNumber(1), batch.EFilingSubmissionXML.SeqNum)
		require.Equal(t, "A", batch.EFilingSubmissionXML.StatusCode)

		activity := batch.EFilingSubmissionXML.EFilingActivityXML[0]

		require.Equal(t, fincen.SeqNumber(1), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053784"), activity.BSAID)
		require.Equal(t, 2, len(activity.EFilingActivityErrorXML))

		errorXML := activity.EFilingActivityErrorXML[0]
		require.Equal(t, fincen.RestrictString4000("EFilingBatchXML/Activity[1]/Party[25]/PartyIdentification[49]/PartyIdentificationNumberText"), *errorXML.ErrorContextText)
		require.Equal(t, fincen.RestrictString512("PartyIdentificationNumberText"), *errorXML.ErrorElementNameText)
		require.Equal(t, fincen.RestrictString50("WARN"), *errorXML.ErrorLevelText)
		require.Equal(t, fincen.RestrictString525("The value recorded for the EIN or SSN/ITIN is an invalid number string."), *errorXML.ErrorText)
		require.Equal(t, fincen.RestrictString50("C21"), *errorXML.ErrorTypeCode)

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[1]

		require.Equal(t, fincen.SeqNumber(2), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053785"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[2]

		require.Equal(t, fincen.SeqNumber(3), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053786"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		require.Nil(t, batch.Validate())
	})

	t.Run("FinCEN CTR XML Acknowledgement", func(t *testing.T) {
		buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "ctr_acknowledgement.xml"))
		require.NoError(t, err)

		batch := EFilingBatchXML{}
		err = xml.Unmarshal(buf, &batch)
		require.NoError(t, err)
		require.NotNil(t, batch.EFilingSubmissionXML)
		require.Equal(t, 3, len(batch.EFilingSubmissionXML.EFilingActivityXML))

		require.Equal(t, fincen.SeqNumber(1), batch.SeqNum)
		require.Equal(t, "A", batch.StatusCode)

		require.Equal(t, fincen.SeqNumber(1), batch.EFilingSubmissionXML.SeqNum)
		require.Equal(t, "A", batch.EFilingSubmissionXML.StatusCode)

		activity := batch.EFilingSubmissionXML.EFilingActivityXML[0]

		require.Equal(t, fincen.SeqNumber(1), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000001"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		errorXML := activity.EFilingActivityErrorXML[0]
		require.Equal(t, fincen.RestrictString4000("EFilingBatchXML/Activity[1]/Party[218]/Address[232]/RawZIPCode"), *errorXML.ErrorContextText)
		require.Equal(t, fincen.RestrictString512("RawZIPCode"), *errorXML.ErrorElementNameText)
		require.Equal(t, fincen.RestrictString50("WARN"), *errorXML.ErrorLevelText)
		require.Equal(t, fincen.RestrictString525("The value recorded contains non-numeric characters or is in an invalid format and the country is US."), *errorXML.ErrorText)
		require.Equal(t, fincen.RestrictString50("F18"), *errorXML.ErrorTypeCode)

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[1]

		require.Equal(t, fincen.SeqNumber(2), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000002"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[2]

		require.Equal(t, fincen.SeqNumber(3), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000003"), activity.BSAID)
		require.Equal(t, 2, len(activity.EFilingActivityErrorXML))

		require.Nil(t, batch.Validate())
	})

	t.Run("FinCEN 8300 XML Acknowledgement", func(t *testing.T) {
		buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "8300_acknowledgement.xml"))
		require.NoError(t, err)

		batch := EFilingBatchXML{}
		err = xml.Unmarshal(buf, &batch)
		require.NoError(t, err)
		require.NotNil(t, batch.EFilingSubmissionXML)
		require.Equal(t, 3, len(batch.EFilingSubmissionXML.EFilingActivityXML))

		require.Equal(t, fincen.SeqNumber(1), batch.SeqNum)
		require.Equal(t, "A", batch.StatusCode)

		require.Equal(t, fincen.SeqNumber(1), batch.EFilingSubmissionXML.SeqNum)
		require.Equal(t, "A", batch.EFilingSubmissionXML.StatusCode)

		activity := batch.EFilingSubmissionXML.EFilingActivityXML[0]

		require.Equal(t, fincen.SeqNumber(1), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053784"), activity.BSAID)
		require.Equal(t, 2, len(activity.EFilingActivityErrorXML))

		errorXML := activity.EFilingActivityErrorXML[0]
		require.Equal(t, fincen.RestrictString4000("EFilingBatchXML/Activity[1]/Party[25]/PartyIdentification[49]/OtherIssuerStateText"), *errorXML.ErrorContextText)
		require.Equal(t, fincen.RestrictString512("PartyIdentificationNumberText"), *errorXML.ErrorElementNameText)
		require.Equal(t, fincen.RestrictString50("WARN"), *errorXML.ErrorLevelText)
		require.Equal(t, fincen.RestrictString525("The element is not recorded or does not contain a value and the PartyIdentificationTypeCode is equal to 5 (Driver’s License/State ID), 6 (Passport), 7 (Alien registration), or 999 (Other)."), *errorXML.ErrorText)
		require.Equal(t, fincen.RestrictString50("P05"), *errorXML.ErrorTypeCode)

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[1]

		require.Equal(t, fincen.SeqNumber(2), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053785"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[2]

		require.Equal(t, fincen.SeqNumber(3), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053786"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		require.Nil(t, batch.Validate())
	})

	t.Run("FinCEN DOEP XML Acknowledgement", func(t *testing.T) {
		buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "doep_acknowledgement.xml"))
		require.NoError(t, err)

		batch := EFilingBatchXML{}
		err = xml.Unmarshal(buf, &batch)
		require.NoError(t, err)
		require.NotNil(t, batch.EFilingSubmissionXML)
		require.Equal(t, 3, len(batch.EFilingSubmissionXML.EFilingActivityXML))

		require.Equal(t, fincen.SeqNumber(1), batch.SeqNum)
		require.Equal(t, "A", batch.StatusCode)

		require.Equal(t, fincen.SeqNumber(1), batch.EFilingSubmissionXML.SeqNum)
		require.Equal(t, "A", batch.EFilingSubmissionXML.StatusCode)

		activity := batch.EFilingSubmissionXML.EFilingActivityXML[0]

		require.Equal(t, fincen.SeqNumber(1), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053784"), activity.BSAID)
		require.Equal(t, 2, len(activity.EFilingActivityErrorXML))

		errorXML := activity.EFilingActivityErrorXML[0]
		require.Equal(t, fincen.RestrictString4000("EFilingBatchXML/Activity[1]/Party[25]/PartyName[30]/RawIndividualFirstName"), *errorXML.ErrorContextText)
		require.Equal(t, fincen.RestrictString512("RawIndividualFirstName"), *errorXML.ErrorElementNameText)
		require.Equal(t, fincen.RestrictString50("WARN"), *errorXML.ErrorLevelText)
		require.Equal(t, fincen.RestrictString525("The exempt person is an individual and first name is not provided."), *errorXML.ErrorText)
		require.Equal(t, fincen.RestrictString50("P02"), *errorXML.ErrorTypeCode)

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[1]

		require.Equal(t, fincen.SeqNumber(2), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053785"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[2]

		require.Equal(t, fincen.SeqNumber(3), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000055053786"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		require.Nil(t, batch.Validate())
	})

	t.Run("FinCEN FBAR XML Acknowledgement", func(t *testing.T) {
		buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "fbar_acknowledgement.xml"))
		require.NoError(t, err)

		batch := EFilingBatchXML{}
		err = xml.Unmarshal(buf, &batch)
		require.NoError(t, err)
		require.NotNil(t, batch.EFilingSubmissionXML)
		require.Equal(t, 3, len(batch.EFilingSubmissionXML.EFilingActivityXML))

		require.Equal(t, fincen.SeqNumber(1), batch.SeqNum)
		require.Equal(t, "A", batch.StatusCode)

		require.Equal(t, fincen.SeqNumber(1), batch.EFilingSubmissionXML.SeqNum)
		require.Equal(t, "A", batch.EFilingSubmissionXML.StatusCode)

		activity := batch.EFilingSubmissionXML.EFilingActivityXML[0]

		require.Equal(t, fincen.SeqNumber(1), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000001"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		errorXML := activity.EFilingActivityErrorXML[0]
		require.Equal(t, fincen.RestrictString4000("EFilingBatchXML/Activity[1]/Party[3]/Address[1]/RawStateCodeText"), *errorXML.ErrorContextText)
		require.Equal(t, fincen.RestrictString512("RawStateCodeText"), *errorXML.ErrorElementNameText)
		require.Equal(t, fincen.RestrictString50("WARN"), *errorXML.ErrorLevelText)
		require.Equal(t, fincen.RestrictString525("The element is not recorded or does not contain a value and the associated country is US, CA, MX, or a U.S. Territory."), *errorXML.ErrorText)
		require.Equal(t, fincen.RestrictString50("C33"), *errorXML.ErrorTypeCode)

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[1]

		require.Equal(t, fincen.SeqNumber(2), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000002"), activity.BSAID)
		require.Equal(t, 1, len(activity.EFilingActivityErrorXML))

		activity = batch.EFilingSubmissionXML.EFilingActivityXML[2]

		require.Equal(t, fincen.SeqNumber(3), activity.SeqNum)
		require.Equal(t, fincen.RestrictNumeric14("31000000000003"), activity.BSAID)
		require.Equal(t, 2, len(activity.EFilingActivityErrorXML))

		require.Nil(t, batch.Validate())
	})

}

func TestBatch(t *testing.T) {
	samples := map[string]string{
		"8300_batch.xml": "FinCEN 8300 XML Batch",
		"ctr_batch.xml":  "FinCEN CTR XML Batch",
		"doep_batch.xml": "FinCEN DOEP XML Batch",
		"fbar_batch.xml": "FinCEN FBAR XML Batch",
		"sar_batch.xml":  "FinCEN SAR XML Batch",
	}

	for name, description := range samples {
		t.Run(description, func(t *testing.T) {
			buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", name))
			require.NoError(t, err)

			batch := EFilingBatchXML{}
			err = xml.Unmarshal(buf, &batch)
			require.NoError(t, err)

			err = batch.Validate()
			require.NoError(t, err)

			err = batch.GenerateAttrs()
			require.NoError(t, err)

			err = batch.GenerateSeqNumbers()
			require.NoError(t, err)

			err = batch.Validate()
			require.NoError(t, err)
		})
	}

}

func TestElements(t *testing.T) {

	t.Run("EFilingBatchXML", func(t *testing.T) {

		sample := EFilingBatchXML{}
		require.Equal(t, "The FormTypeCode has invalid value", sample.Validate().Error())

		sample.FormTypeCode = "CTRX"
		require.Equal(t, "The Activity has invalid value", sample.Validate().Error())

		var emptyActivity fincen.ElementActivity = (*emptyElementActivity)(nil)
		sample.Activity = append(sample.Activity, emptyActivity)
		sample.ActivityCount = 1

		require.NotNil(t, sample.Validate())
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate().Error())

		sample.StatusCode = "A"
		require.Equal(t, "The EFilingSubmissionXML has invalid value", sample.Validate().Error())
	})
}
