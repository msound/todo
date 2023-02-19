<p
    data-hx-post="/task/{{ .ID.Hex }}/{{ if .Done }}undo{{ else }}done{{ end }}"
    data-hx-swap="outerHTML"
    class="task{{ if .Done }}-done{{ end }}"
>
    <span class="material-symbols-outlined">
    {{ if .Done }}check_box{{ else }}check_box_outline_blank{{ end }}
    </span>
    {{ .Title }}
    <span class="htmx-indicator">&nbsp;*</span>
</p>
