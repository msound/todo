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
        <div id="tasks">
        {{ range .Tasks }}{{ template "task.html.tpl" . }}{{ end }}
        {{ template "add-task.html.tpl" }}
        </div>
    </body>
</html>
