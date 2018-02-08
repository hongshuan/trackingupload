# Using Layouts and Nested Templates

File: templates/layout.html

<html>
  <head>
    <title>{{ template "title" . }}</title>
  </head>
  <body>
    {{ template "content" . }}
  </body>
</html>

File: templates/index.html

{{ define "title" }}An example layout{{ end }}

{{ define "content" }}
<h1>Hello {{ .Name }}</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
{{ end }}

File: main.go

package main

import (
    "html/template"
    "net/http"
    "path"
)

type Profile struct {
    Name    string
    Hobbies []string
}

func main() {
    http.HandleFunc("/", foo)
    http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}

    lp := path.Join("templates", "layout.html")
    fp := path.Join("templates", "index.html")

    // Note that the layout file must be the first parameter in ParseFiles
    tmpl, err := template.ParseFiles(lp, fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, profile); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

