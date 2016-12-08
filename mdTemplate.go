package MarkGen

import (
	"fmt"
	"net/http"
	"text/template"
)

func Template(w http.ResponseWriter, filepath string) {
	var style string
	if css, err := CustomCSS(); err == nil {
		style = *css
	} else {
		style = "<style>" + DefaultStyle + "</style>"
	}

	templateStr := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset='UTF-8' />
  <title>%[1]s</title>
  %[2]s
</head>
<body>
  <div id='md' class='markdown-body'></div>
  <script>
    (function () {
      var markdown = document.getElementById("md");
      var conn = new WebSocket("ws://" + location.host + "/%[1]s");
      conn.onmessage = function (evt) {
        markdown.innerHTML = evt.data;
      };
    })();
  </script>
</body>`, filepath, style)

	var (
		t   *template.Template
		err error
	)

	if t, err = template.New("template").Parse(templateStr); err != nil {
		panic(err)
	}

	if err = t.Execute(w, nil); err != nil {
		panic(err)
	}
}
