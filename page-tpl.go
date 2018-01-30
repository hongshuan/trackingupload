package main

var pageTpl = `
<html>
  <head>
  <title>Tracking Upload</title>
  <style type="text/css">
    body { width: 960px; margin: 0 auto; }
    pre { font-size: 16px; }
  </style>
  </head>
  <body>
    <h1>Tracking Number Upload</h1>
    <form action="/" method="post">
	  <label>Click "Upload" Button to Upload Tracking Numbers to Server </label>
      <input type="submit" value="Upload">
    </form>
	<pre>{{range .}}
{{.}}{{end}}
    </pre>
  </body>
</html>`
