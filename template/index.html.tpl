<!DOCTYPE html>
<html>
    <head>
        <title>TODO</title>
        <script src="https://unpkg.com/htmx.org@1.8.5"></script>
        <style>
            {{ template "style.css" }}
        </style>
    </head>
    <body>
        <h1>TODO</h1>
        {{ range .Tasks }}{{ template "task.html.tpl" . }}{{ end }}
    </body>
</html>
