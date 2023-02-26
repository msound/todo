{{ if . }}{{ template "task.html.tpl" . }}{{ end }}
<form class="add-task" data-hx-post="/task" data-hx-swap="outerHTML">
    <input type="text" name="newtask" size="40" style="line-height: 1.5em;" placeholder="Add a task"/>
</form>
