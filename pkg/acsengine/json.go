package acsengine

import (
	"encoding/json"
	"strings"
)

// PrettyPrintArmTemplate will pretty print the arm template ensuring ordered by params, vars, resources, and outputs
func PrettyPrintArmTemplate(template string) (string, error) {
	translateParams := [][]string{
		{"\"parameters\"", "\"dparameters\""},
		{"\"variables\"", "\"evariables\""},
		{"\"resources\"", "\"fresources\""},
		{"\"outputs\"", "\"zoutputs\""},
		// there is a bug in ARM where it doesn't correctly translate back '\u003e' (>)
		{">", "GREATERTHAN"},
		{"<", "LESSTHAN"},
		{"&", "AMPERSAND"},
	}

	template = translateJSON(template, translateParams, false)
	var err error
	if template, err = PrettyPrintJSON(template); err != nil {
		return "", err
	}
	template = translateJSON(template, translateParams, true)

	return template, nil
}

// PrettyPrintJSON will pretty print the json into
func PrettyPrintJSON(content string) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(content), &data); err != nil {
		return "", err
	}
	prettyprint, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyprint), nil
}

func translateJSON(content string, translateParams [][]string, reverseTranslate bool) string {
	for _, tuple := range translateParams {
		if len(tuple) != 2 {
			panic("string tuples must be of size 2")
		}
		a := tuple[0]
		b := tuple[1]
		if reverseTranslate {
			content = strings.Replace(content, b, a, -1)
		} else {
			content = strings.Replace(content, a, b, -1)
		}
	}
	return content
}
