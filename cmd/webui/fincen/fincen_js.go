package main

import (
	"encoding/xml"
	"fmt"
	"syscall/js"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
)

func prettyXml(input *batch.EFilingBatchXML) (string, error) {
	pretty, err := xml.MarshalIndent(input, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func generateAttrs() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		r, err := batch.CreateReportWithBuffer([]byte(args[0].String()))
		if err != nil {
			fmt.Println(err)
			return "Unable to parse report file"
		}

		err = r.GenerateAttrs()
		if err != nil {
			fmt.Println(err)
			return "Unable to generate report attributes"
		}
		err = r.GenerateSeqNumbers()
		if err != nil {
			fmt.Println(err)
			return "Unable to generate sequence numbers"
		}

		pretty, err := prettyXml(r)
		if err != nil {
			fmt.Printf("unable to convert fincen file to xml %s\n", err)
			return "There was an error converting the xml"
		}

		return pretty
	})
}

func validateForm() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		r, err := batch.CreateReportWithBuffer([]byte(args[0].String()))
		if err != nil {
			fmt.Println(err)
			return "Unable to parse report file"
		}

		err = r.Validate()
		if err != nil {
			fmt.Println(err)
			return fmt.Sprintf("The report form is invalid\n %s\n", err.Error())
		}

		return "The report form is valid"
	})
}

func writeVersion() {
	span := js.Global().Get("document").Call("querySelector", "#version")
	span.Set("innerHTML", fmt.Sprintf("Version: %s", fincen.Version))
}

func main() {
	js.Global().Set("generateAttrs", generateAttrs())
	js.Global().Set("validateForm", validateForm())

	writeVersion()

	<-make(chan bool)
}
