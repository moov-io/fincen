package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/ach"
)

func parseFile(input string) (string, error) {
	r := strings.NewReader(input)
	file, err := ach.NewReader(r).Read()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func prettyPrintJSON() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		if json.Valid([]byte(args[0].String())) {
			return args[0].String()
		}

		parsed, err := parseFile(args[0].String())
		if err != nil {
			msg := fmt.Sprintf("unable to parse fincen file: %v", err)
			fmt.Println(msg)
			return msg
		}
		pretty, err := prettyJson(parsed)
		if err != nil {
			fmt.Printf("unable to convert fincen file to json %s\n", err)
			return "There was an error converting the json"
		}

		return pretty
	})
}

func prettyXml(input string) (string, error) {
	var raw interface{}
	if err := xml.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := xml.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func prettyPrintXML() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		if json.Valid([]byte(args[0].String())) {
			return args[0].String()
		}

		parsed, err := parseFile(args[0].String())
		if err != nil {
			msg := fmt.Sprintf("unable to parse fincen file: %v", err)
			fmt.Println(msg)
			return msg
		}
		pretty, err := prettyXml(parsed)
		if err != nil {
			fmt.Printf("unable to convert fincen file to json %s\n", err)
			return "There was an error converting the json"
		}

		return pretty
	})
}

func generateAttrs() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		file, err := parseFile(args[0].String())
		if err != nil {
			msg := fmt.Sprintf("unable to parse fincen file: %v", err)
			fmt.Println(msg)
			return msg
		}

		parsed, err := parseReadable(file)
		if err != nil {
			fmt.Printf("unable to convert fincen file to human-readable format %s\n", err)
			return "There was an error formatting the output"
		}

		return parsed
	})
}

func writeVersion() {
	span := js.Global().Get("document").Call("querySelector", "#version")
	span.Set("innerHTML", fmt.Sprintf("Version: %s", ach.Version))
}

func main() {
	js.Global().Set("parseXML", prettyPrintJSON())
	js.Global().Set("parseJSON", prettyPrintXML())
	js.Global().Set("generateAttrs", generateAttrs())

	writeVersion()

	<-make(chan bool)
}
