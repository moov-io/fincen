// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/moov-io/fincen/pkg/batch"
)

func Reformat(paths []string, generate bool, newType string) ([]byte, error) {
	fmt.Fprintf(os.Stdout, "\nFormating report file...\n")
	r, err := batch.CreateReportWithFile(paths[0])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to create report from file: %v\n", err)
		return nil, err
	}

	if generate {
		err = r.GenerateAttrs()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed to generate report attributes: %v\n", err)
			return nil, err
		}
		err = r.GenerateSeqNumbers()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed to generate sequence numbers: %v\n", err)
			return nil, err
		}
	}

	if strings.ToLower(newType) == "json" {
		return json.MarshalIndent(r, "", "  ")
	}

	return xml.MarshalIndent(r, "", "  ")
}

func Validate(paths []string) error {
	fmt.Fprintf(os.Stdout, "\nValidating report file...\n")
	r, err := batch.CreateReportWithFile(paths[0])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to create report from file: %v\n", err)
		return err
	}

	return r.Validate()
}

func Summary(paths []string) error {
	fmt.Fprintf(os.Stdout, "\n")
	for _, path := range paths {
		fmt.Fprintf(os.Stdout, "%s\n", path)
		summaryForm(path)
		fmt.Fprintf(os.Stdout, "\n")
	}

	return nil
}

func summaryForm(path string) error {

	r, err := batch.CreateReportWithFile(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to create report from file: %v\n", err)
		fmt.Fprintf(os.Stdout, "Trying to summary file anyway...\n")
		return err
	}

	if r.StatusCode == "A" {
		fmt.Fprintf(os.Stdout, "Acknowledgement\n")
		if r.EFilingSubmissionXML != nil {
			for _, act := range r.EFilingSubmissionXML.EFilingActivityXML {
				fmt.Fprintf(os.Stdout, "BSAID\t: %s\n", act.BSAID)
				for _, actErr := range act.EFilingActivityErrorXML {
					if actErr.ErrorTypeCode != nil {
						fmt.Fprintf(os.Stdout, "ErrorTypeCode\t: %s\n", *actErr.ErrorTypeCode)
					}
					if actErr.ErrorLevelText != nil {
						fmt.Fprintf(os.Stdout, "ErrorLevelText\t: %s\n", *actErr.ErrorLevelText)
					}
					if actErr.ErrorText != nil {
						fmt.Fprintf(os.Stdout, "ErrorText\t: %s\n", *actErr.ErrorText)
					}
					if actErr.ErrorElementNameText != nil {
						fmt.Fprintf(os.Stdout, "ErrorElementNameText\t: %s\n", *actErr.ErrorElementNameText)
					}
				}
			}
		}
	} else {
		switch r.FormTypeCode {
		case "8300X":
			fmt.Fprintf(os.Stdout, "Report of Cash Payments Over $10,000 Received\n")
			fmt.Fprintf(os.Stdout, "ActivityCount\t: %v\n", r.ActivityCount)
			fmt.Fprintf(os.Stdout, "TotalAmount\t: %v\n", r.TotalAmount)
			fmt.Fprintf(os.Stdout, "PartyCount\t: %v\n", r.PartyCount)
		case "CTRX":
			fmt.Fprintf(os.Stdout, "FinCEN Currency Transaction Report\n")
			fmt.Fprintf(os.Stdout, "ActivityCount\t: %v\n", r.ActivityCount)
			fmt.Fprintf(os.Stdout, "TotalAmount\t: %v\n", r.TotalAmount)
			fmt.Fprintf(os.Stdout, "PartyCount\t: %v\n", r.PartyCount)
		case "DOEPX":
			fmt.Fprintf(os.Stdout, "FinCEN Designation of Exempt Person\n")
			fmt.Fprintf(os.Stdout, "ActivityCount\t: %v\n", r.ActivityCount)
			fmt.Fprintf(os.Stdout, "PartyCount\t: %v\n", r.PartyCount)
		case "FBARX":
			fmt.Fprintf(os.Stdout, "FinCEN Report of Foreign Bank and Financial Accounts\n")
			fmt.Fprintf(os.Stdout, "ActivityCount\t: %v\n", r.ActivityCount)
			fmt.Fprintf(os.Stdout, "AccountCount\t: %v\n", r.AccountCount)
			fmt.Fprintf(os.Stdout, "JointlyOwnedOwnerCount\t: %v\n", r.JointlyOwnedOwnerCount)
			fmt.Fprintf(os.Stdout, "NoFIOwnerCount\t: %v\n", r.NoFIOwnerCount)
			fmt.Fprintf(os.Stdout, "ConsolidatedOwnerCount\t: %v\n", r.ConsolidatedOwnerCount)
		case "SARX":
			fmt.Fprintf(os.Stdout, "FinCEN Suspicious Activity Report\n")
			fmt.Fprintf(os.Stdout, "ActivityCount\t: %v\n", r.ActivityCount)
			fmt.Fprintf(os.Stdout, "TotalAmount\t: %v\n", r.TotalAmount)
			fmt.Fprintf(os.Stdout, "PartyCount\t: %v\n", r.PartyCount)
			fmt.Fprintf(os.Stdout, "ActivityAttachmentCount\t: %v\n", 0)
			fmt.Fprintf(os.Stdout, "AttachmentCount\t: %v\n", 0)
		default:
			fmt.Fprintf(os.Stdout, "Invalid form type code\n")
		}
	}

	return nil
}
