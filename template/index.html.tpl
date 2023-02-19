<!DOCTYPE html>
<html>
    <head>
        <title>TODO</title>
        <style>
            {{ template "style.css" }}
        </style>
    </head>
    <body>
        <h1>TODO</h1>
        <p>List ID: {{ .ID }}</p>
        {{ range .Tasks }}{{ template "task.html.tpl" . }}{{ end }}
    </body>
</html>
