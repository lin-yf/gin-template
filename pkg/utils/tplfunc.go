package utils

import (
	htmpl "html/template"
	"strings"
	"text/template"
	"time"
)

var TplFuncMap = make(template.FuncMap)

func init() {
	TplFuncMap["dateformat"] = DateFormat
	TplFuncMap["str2html"] = Str2html
	TplFuncMap["join"] = Join
	TplFuncMap["isnotzero"] = IsNotZero
}

func Str2html(raw string) htmpl.HTML {
	return htmpl.HTML(raw)
}

// DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"
func DateFormat(t time.Time, layout string) string {
	return t.Format(layout)
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func IsNotZero(t time.Time) bool {
	return !t.IsZero()
}
