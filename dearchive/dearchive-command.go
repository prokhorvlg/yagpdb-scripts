{{/*
	!!! CURRENTLY BROKEN IN THREADS, MAKING THIS USELESS !!!
	DEARCHIVE: DEARCHIVE COMMAND
	Written by prokhorVLG
	
	This command starts the dearchiver, which prevents a thread from being archived.
	This command is meant to be run by the user to start and stop the process.
	Trigger type: Command
	Trigger: dearchive
*/}}

{{/* arguments that can be passed to this command */}}
{{$args := parseArgs 0 "-dearchive <command>" (carg "string" "command")}}

{{/* fancy shmancy message color */}}
{{$color := 16733707}}

{{/* bot message self-delete delay */}}
{{$botdeletedelay := 10}}

{{/* bot dearchive message delay in seconds */}}
{{/* 1hr = 3600s, so post every 3400s */}}
{{/* 1week = 604800s, so post every 604600s */}}
{{$delay := 10}}

{{/* delete the message that triggered this command */}}
{{deleteMessage .Channel.ID .Message.ID}}

{{/* if there is any message after the triggering command, do stuff */}}
{{if .StrippedMsg}}
	{{/* if the message is "start", start the dearchive on this channel */}}
	{{if (eq ($args.Get 0) "start")}}
		{{/* set checkin to true */}}
		{{dbSet .Channel.ID "dearchive" true}}
		{{/* save delay duration for use by custom command */}}
		{{dbSet .Channel.ID "delayduration" $delay}}
		{{/* inform the user */}}
		{{$message := sendMessageRetID nil (cembed 
			"description" "Dearchive started in this channel."
			"color" $color
			"author" (sdict "name" (print "Checkin") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
		)}}
		{{deleteMessage .Channel.ID $message $botdeletedelay}}
		{{/* start delayed message command */}}
		{{sendMessage nil .Channel.ID}}
		{{scheduleUniqueCC 8 .Channel.ID $delay .Channel.ID "input data"}}
	{{/* if the message is "end", end the dearchive on this channel */}}
	{{else if (eq ($args.Get 0) "end")}}
		{{/* set checkin to true */}}
		{{dbDel .Channel.ID "dearchive"}}
		{{/* end delayed message command */}}
		{{cancelScheduledUniqueCC 8 .Channel.ID}}
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
		"A utility that prevents this thread from being archived.\n"
		"\n"
		"Commands:\n"
		"**`-dearchive start`**: Start dearchive for this thread. Remember to set archive to 1 week. \n"
		"**`-dearchive end`**: Stop dearchive for this thread.")
		"color" $color
		"author" (sdict "name" (print "Dearchive") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
	)}}
	{{deleteMessage .Channel.ID $message $botdeletedelay}}
{{end}}