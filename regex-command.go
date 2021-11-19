{{/*
	Trigger type: Regex
	Trigger: \A

	Copyright (c): Black Wolf, 2021
	License: MIT
	Repository: https://github.com/BlackWolfWoof/yagpdb-cc/
	Modified by prokhorVLG
*/}}

{{/* delay constant in seconds, todo: turn this into a setting */}}
{{$delay := 5}}

{{/* if sticky message exists, */}}
{{if (dbGet .Channel.ID "stickymessage")}}
	{{/* if the message is currently NOT being delayed, */}}
	{{if (dbGet .Channel.ID "indelay"|not)}}
		{{/* set delay to true for this channel */}}
		{{dbSet .Channel.ID "indelay" true}}
		{{/* delay the message using a custom command */}}
		{{scheduleUniqueCC MESSAGE-COMMAND .Channel.ID $delay "stickydelayfunction" "input data"}}
	{{/* if the message IS currently being delayed, */}}
	{{else}}
		{{/* reset the execCC so the duration is reset */}}
		{{cancelScheduledUniqueCC MESSAGE-COMMAND "stickydelayfunction"}}
		{{scheduleUniqueCC MESSAGE-COMMAND .Channel.ID $delay "stickydelayfunction" "input data"}}
	{{end}}
{{end}}