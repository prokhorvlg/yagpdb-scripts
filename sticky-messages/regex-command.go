{{/*
	STICKY MESSAGES: Regex Command
	Written by prokhorVLG
	Modified from Black Wolf's work, from https://github.com/BlackWolfWoof/yagpdb-cc/

	Trigger type: Regex
	Trigger: \A
	
	This command detects any message in a channel and starts a countdown to posting the sticky message.
*/}}

{{/* if sticky message exists, */}}
{{if (dbGet .Channel.ID "stickymessage")}}
	{{/* if the message is currently NOT being delayed, */}}
	{{if (dbGet .Channel.ID "indelay"|not)}}
		{{/* set delay to true for this channel */}}
		{{dbSet .Channel.ID "indelay" true}}
		{{/* get the delay duration from the channel metadata */}}
		{{$delay := (dbGet .Channel.ID "delayduration").Value}}
		{{/* delay the message using a custom command */}}
		{{scheduleUniqueCC 4 .Channel.ID $delay .Channel.ID "input data"}}
	{{/* if the message IS currently being delayed, */}}
	{{else}}
		{{/* cancel the previous operation to stop it from happening */}}
		{{cancelScheduledUniqueCC 4 .Channel.ID}}
		{{/* get the delay duration from the channel metadata */}}
		{{$delay := (dbGet .Channel.ID "delayduration").Value}}
		{{/* start a new operation so the delay is reset */}}
		{{scheduleUniqueCC 4 .Channel.ID $delay .Channel.ID "input data"}}
	{{end}}
{{end}}