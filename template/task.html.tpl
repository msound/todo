<p data-hx-post="/task/{{ .ID.Hex }}/{{ if .Done }}undo{{ else }}done{{ end }}" data-hx-swap="outerHTML" class="task{{ if .Done }}-done{{ end }}">{{ .Title }}</p>
