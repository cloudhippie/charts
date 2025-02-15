{{/*
Expand the name of the chart.
*/}}
{{- define "ansible-semaphore.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ansible-semaphore.fullname" -}}
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
{{- define "ansible-semaphore.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create basic labels
*/}}
{{- define "ansible-semaphore.labels" -}}
helm.sh/chart: "{{ include "ansible-semaphore.chart" . }}"
app.kubernetes.io/name: "{{ include "ansible-semaphore.name" . }}"
app.kubernetes.io/instance: "{{ .Release.Name }}"
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
{{- with .Values.labels }}
{{ toYaml . }}
{{- end }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "ansible-semaphore.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "ansible-semaphore.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "ansible-semaphore.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ansible-semaphore.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Database options
*/}}
{{- define "ansible-semaphore.databaseOptions" -}}
{{- $options := .Values.database.options | default dict }}
{{- if and (eq .Values.database.type "postgres") (not (hasKey $options "sslmode")) }}
{{- $options = merge (dict "sslmode" "disable") $options }}
{{- end }}
{{- toYaml $options }}
{{- end }}
