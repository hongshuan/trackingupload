package main

var pageTpl = `
<html>
  <head>
  <title>Tracking Upload</title>
  <style type="text/css">
    body { width: 960px; margin: 0 auto; }
    pre { font-size: 16px; }
    input { width: 100px; }
  </style>
  </head>
  <body>
    <h1>Tracking Number Upload</h1>
    <form action="/" method="post">
	  <table>
	  <tr>
	    <td>Click "Upload" Button to Upload Tracking Numbers to Server</td>
        <td width="10%"></td>
        <td><input type="submit" name="btn" value="Upload"></td>
	  </tr>
	  <tr>
	    <td>Click "Download" Button to Download AddressBook from Server </td>
        <td width="10%"></td>
        <td><input type="submit" name="btn" value="Download"></td>
	  </tr>
	  </table>
    </form>
	<hr>
	<pre>{{range .}}
{{.}}{{end}}
    </pre>
  </body>
</html>`
