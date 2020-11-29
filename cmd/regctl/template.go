package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"text/template"
	"time"
)

var tmplFuncs = template.FuncMap{
	"json": func(v interface{}) string {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		enc.Encode(v)
		return buf.String()
	},
	"jsonPretty": func(v interface{}) string {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		enc.Encode(v)
		return buf.String()
	},
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
	"time":  TimeFunc,
}

func templateRun(out io.Writer, tmpl string, data interface{}) error {
	t, err := template.New("out").Funcs(tmplFuncs).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

// TimeFunc provides the "time" template, returning a struct with methods
func TimeFunc() *TimeFuncs {
	return &TimeFuncs{}
}

// TimeFuncs wraps all time based templates
type TimeFuncs struct{}

// Now returns current time
func (t *TimeFuncs) Now() time.Time {
	return time.Now()
}

// Parse parses the current time according to layout
func (t *TimeFuncs) Parse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
