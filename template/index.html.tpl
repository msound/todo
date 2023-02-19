<!DOCTYPE html>
<html>
    <head>
        <title>TODO</title>
        <script src="https://unpkg.com/htmx.org@1.8.5"></script>
        <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,400,0,0" />
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
