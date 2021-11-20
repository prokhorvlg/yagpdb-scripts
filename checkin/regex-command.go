{{/*
	Trigger type: Regex
	Trigger: \A

	by prokhorVLG
*/}}

{{/* if checkin is enabled on this channel, */}}
{{if (dbGet .Channel.ID "checkin")}}
	{{/* give the role to the user who made the message */}}
	{{addRoleName "Checked-In"}}
{{end}}