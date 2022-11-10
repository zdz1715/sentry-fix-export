package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	inputFile  string
	outputFile string
)

func init() {
	flag.StringVar(&inputFile, "i", "", "input json file")
	flag.StringVar(&outputFile, "o", "sentry_export.json", "output json file")
}

func errorPrint(msg string) {
	fmt.Printf("[Error] %s \n", msg)
}

func main() {
	flag.Parse()
	if inputFile == "" {
		errorPrint("-i is require")
		os.Exit(1)
	}

	data, _ := os.ReadFile(inputFile)

	var sentrySlice []interface{}

	if err := json.Unmarshal(data, &sentrySlice); err != nil {
		errorPrint(err.Error())
		os.Exit(1)
	}

	for _, v := range sentrySlice {
		if m, ok := v.(map[string]any); ok {

			if p, ok := m["pk"].(float64); ok {
				m["pk"] = strconv.FormatFloat(p, 'f', -1, 32)
			}
			if f, ok := m["fields"].(map[string]interface{}); ok {

				if fv, ok := f["value"].(float64); ok {
					f["value"] = strconv.FormatFloat(fv, 'f', -1, 32)
				}
				if fv, ok := f["value"].(bool); ok {
					if fv == true {
						f["value"] = "true"
					} else {
						f["value"] = "false"
					}
				}
				m["fields"] = f
			}
		}
	}

	res, _ := json.Marshal(sentrySlice)

	if _, err := os.Stat(outputFile); err != nil {
		if os.IsExist(err) {
			_ = os.Remove(outputFile)
		}
	}

	if err := os.WriteFile(outputFile, res, 755); err != nil {
		errorPrint(err.Error())
		os.Exit(1)
	}

	fmt.Printf("[Info] output file: %s", outputFile)
}
