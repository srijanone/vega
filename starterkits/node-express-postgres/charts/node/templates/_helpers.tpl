{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "node.appName" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "node.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "node.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "node.labels" -}}
app: {{ include "node.appName" . }}
chart: {{ include "node.chart" . }}
release: {{ .Release.Name }}
heritage: {{ .Release.Service }}
{{- end -}}

{{/*
Common labels for cron-jobs to be used for deletion
*/}}
{{- define "node.cronLabels" -}}
app: {{ include "node.appName" . }}
chart: {{ include "node.chart" . }}
release: {{ .Release.Name }}
heritage: {{ .Release.Service }}
workload: cronjob
{{- end -}}

{{/*
Labels to use on deploy.spec.selector.matchLabels and svc.spec.selector
*/}}
{{- define "node.matchLabels" -}}
app: {{ include "node.appName" . }}
release: {{ .Release.Name }}
{{- end -}}

{{/*
Return the proper php-fpm image name
*/}}
{{- define "node.image" -}}
{{- $registryName := .Values.node.image.registry -}}
{{- $repositoryName := .Values.node.image.repository -}}
{{- $tag := .Values.node.image.tag | toString -}}
{{/*
Helm 2.11 supports the assignment of a value to a variable defined in a different scope,
but Helm 2.9 and 2.10 doesn't support it, so we need to implement this if-else logic.
Also, we can't use a single if because lazy evaluation is not an option
*/}}
{{- if .Values.global }}
    {{- if .Values.global.imageRegistry }}
        {{- printf "%s/%s" .Values.global.imageRegistry $repositoryName -}}
    {{- else -}}
        {{- printf "%s/%s" $registryName $repositoryName -}}
    {{- end -}}
{{- else -}}
    {{- printf "%s/%s" $registryName $repositoryName -}}
{{- end -}}
{{- end -}}



{{/*
Renders a value that contains template.
Usage:
{{ include "node.tplValue" (dict "value" .Values.path.to.the.Value "context" $) }}
*/}}
{{- define "node.tplValue" -}}
    {{- if typeIs "string" .value }}
        {{- tpl .value .context }}
    {{- else }}
        {{- tpl (.value | toYaml) .context }}
    {{- end }}
{{- end -}}
