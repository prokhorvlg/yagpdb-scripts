{{/*
	DEARCHIVE: MESSAGE COMMAND
	Written by prokhorVLG
	
	This command posts a message and deletes it, then schedules a new message, in order to keep a thread de-archived.
	This command is meant to be run only by the DEARCHIVE COMMAND, not by the user.
	Trigger type: None
*/}}

{{/* fancy shmancy message color */}}
{{$color := 16733707}}

{{/* if dearchive is active, */}}
{{if (dbGet .Channel.ID "dearchive")}}
	{{/* get the delay duration from the channel metadata */}}
	{{$delay := (dbGet .Channel.ID "delayduration").Value}}
	{{/* post a dearchive message and delete it soon after */}}
	{{$message := sendMessageRetID nil (cembed 
		"description" "Pay no attention to this message - it will self-delete after several seconds."
		"color" $color
		"author" (sdict "name" (print "Dearchive") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
	)}}
	{{deleteMessage .Channel.ID $message 10}}
	{{/* start delayed message command to repeat the cycle */}}
	{{execCC 8 .Channel.ID $delay .Channel.ID "dearchive-delayed-message"}}
{{end}}