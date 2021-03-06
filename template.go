package oauthmw

import (
	"html/template"
	"net/http"
	"strings"
)

// DefaultProtectedPageTpl is the default protected page template.
const DefaultProtectedPageTpl = `<!DOCTYPE html>
<html>
<head>
  <title>Login Required</title>
</head>
<body>
{{range $provName, $prov := .}}
  <a href="{{$prov}}">Login with {{$provName | title}}</a><br/>
{{else}}
  Sorry, no login options are currently available.
{{end}}
</body>
</html>`

// protectedPageTpl is the parsed DefaultProtectedPageTpl html/template instance.
var protectedPageTpl = template.Must(template.New("oauthmw").Funcs(
	template.FuncMap{
		"title": strings.Title,
	},
).Parse(DefaultProtectedPageTpl))

// defaultTemplateFn writes the DefaultProtectedPageTpl to the http.ResponseWriter.
func defaultTemplateFn(res http.ResponseWriter, req *http.Request, params map[string]interface{}) {
	protectedPageTpl.Execute(res, params)
}

// defaultErrorFn handles the default behavior when errors are encountered.
func defaultErrorFn(code int, msg string, res http.ResponseWriter, req *http.Request) {
	http.Error(res, msg, code)
}
