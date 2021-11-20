{{/*
	Trigger type: Command
	Trigger: checkin

	prokhorVLG
*/}}

{{/* arguments that can be passed to this command */}}
{{$args := parseArgs 0 "-checkin <command>" (carg "string" "command")}}

{{/* fancy shmancy message color */}}
{{$color := 11559167}}

{{/* bot message self-delete delay */}}
{{$botdeletedelay := 10}}

{{/* delete the message that triggered this command */}}
{{deleteMessage .Channel.ID .Message.ID}}

{{/* if there is any message after the triggering command, do stuff */}}
{{if .StrippedMsg}}
	{{/* if the message is "start", start the checkin on this channel */}}
	{{if (eq ($args.Get 0) "start")}}
		{{/* set checkin to true */}}
		{{dbSet .Channel.ID "checkin" true}}
		{{/* inform the user */}}
		{{$message := sendMessageRetID nil (cembed 
			"description" "Checkin started in this channel."
			"color" $color
			"author" (sdict "name" (print "Checkin") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
		)}}
		{{deleteMessage .Channel.ID $message $botdeletedelay}}
	{{/* if the message is "end", end the checkin on this channel */}}
	{{else if (eq ($args.Get 0) "end")}}
		{{/* set checkin to true */}}
		{{dbDel .Channel.ID "checkin"}}
		{{/* inform the user */}}
		{{$message := sendMessageRetID nil (cembed 
			"description" "Checkin ended in this channel."
			"color" $color
			"author" (sdict "name" (print "Checkin") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
		)}}
		{{deleteMessage .Channel.ID $message $botdeletedelay}}
	{{end}}
{{/* if there is no message after the triggering command, */}}
{{else}}
	{{/* show the user the available commands */}}
	{{$message := sendMessageRetID nil (cembed 
		"description" (joinStr "" 
		"A utility that gives you a role if you comment in the channel/thread.\n"
		"\n"
		"Commands:\n"
		"**`-checkin start`**: Start checkin for this channel.\n"
		"**`-checkin end`**: End checkin for this channel.")
		"color" $color
		"author" (sdict "name" (print "Checkin") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
	)}}
	{{deleteMessage .Channel.ID $message $botdeletedelay}}
{{end}}