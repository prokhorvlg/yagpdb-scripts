{{/*
	STICKY MESSAGES: Sticky Command
	Written by prokhorVLG
	Modified from Black Wolf's work, from https://github.com/BlackWolfWoof/yagpdb-cc/

	Trigger type: Command
	Trigger: sticky
	
	This command allows the user to save or delete the sticky on a channel.
*/}}

{{/* permissions that may access the functionality of the bot
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
{{$perms := "ManageMessages"}}

{{/* arguments that can be passed to this command */}}
{{$args := parseArgs 0 "-sticky <command> <body>" (carg "string" "command") (carg "string" "body")}}

{{/* fancy shmancy message color */}}
{{$color := 6356832}}

{{/* default duration of delay in seconds */}}
{{$defaultduration := 10}}

{{/* bot message self-delete delay */}}
{{$botdeletedelay := 10}}

{{/* delete the message that triggered this command */}}
{{deleteMessage .Channel.ID .Message.ID}}

{{/* if you have the permissions, */}}
{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{/* if there is any message after the triggering command, do stuff */}}
	{{if .StrippedMsg}}
		{{/* if the message is "reset", reset the sticky to nothing */}}
		{{if (eq ($args.Get 0) "reset")}}
			{{/* delete the sticky from channel's metadata by key */}}
			{{dbDel .Channel.ID "stickymessage"}}
			{{/* delete the old duration */}}
			{{dbDel .Channel.ID "delayduration"}}
			{{/* delete the old sticky message based on last key */}}
			{{if $db := dbGet .Channel.ID "smchannel"}}
				{{deleteMessage nil (toInt $db.Value) 0}}
			{{end}}
			{{/* delete the key metadata once we're done with it */}}
			{{dbDel .Channel.ID "smchannel"}}
			{{/* inform the user */}}
			{{$message := sendMessageRetID nil (cembed 
				"description" "Sticky message has been removed and reset."
				"color" $color
				"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
			)}}
			{{deleteMessage .Channel.ID $message $botdeletedelay}}
		{{/* if the message is "delay", set the delay to the second arg */}}
		{{else if (eq ($args.Get 0) "delay")}}
			{{/* if there is a second arg for body, */}}
			{{if ($args.Get 1)}}
				{{/* set the sticky! */}}
				{{/* set the duration to whatever was given in the argument */}}
				{{dbSet .Channel.ID "delayduration" (toInt ($args.Get 1))}}
				{{/* inform the user */}}
				{{$message := sendMessageRetID nil (cembed 
					"description" "Delay has been set."
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
				{{deleteMessage .Channel.ID $message $botdeletedelay}}
			{{/* if there is no second arg for body, */}}
			{{else}}
				{{/* inform the user */}}
				{{$message := sendMessageRetID nil (cembed 
					"description" "You must add a duration in seconds (eg. `5` for five seconds)."
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
				{{deleteMessage .Channel.ID $message $botdeletedelay}}
			{{end}}
		{{/* if the message is "set", set the message to the second arg */}}
		{{else if (eq ($args.Get 0) "set")}}
			{{/* if there is a second arg for body, */}}
			{{if ($args.Get 1)}}
				{{/* set the sticky! */}}
				{{/* set the indelay flag to false */}}
				{{dbSet .Channel.ID "indelay" false}}
				{{/* set the duration to our default */}}
				{{dbSet .Channel.ID "delayduration" $defaultduration}}
				{{/* set image to empty */}}
				{{$img := ""}}
				{{/* set text to second argument */}}
				{{$text := $args.Get 1}}
				{{/* save message to key under the channel id */}}
				{{dbSet .Channel.ID "stickymessage" (sdict "message" $text "author" .User.String "img" $img)}}
				{{/* inform the user */}}
				{{$message := sendMessageRetID nil (cembed 
					"description" "Sticky message has been enabled and set!"
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
				{{deleteMessage .Channel.ID $message $botdeletedelay}}
			{{/* if there is no second arg for body, */}}
			{{else}}
				{{/* inform the user */}}
				{{$message := sendMessageRetID nil (cembed 
					"description" "You must add a message body to set a sticky message."
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
				{{deleteMessage .Channel.ID $message $botdeletedelay}}
			{{end}}
		{{end}}
	{{/* if there is no message after the triggering command, */}}
	{{else}}
		{{/* show the user the available commands */}}
		{{$message := sendMessageRetID nil (cembed 
			"description" (joinStr "" 
			"A utility that lets you 'stick' a custom message to the bottom of a channel. Useful for a world intro, perhaps?\n"
			"\n"
			"Commands:\n"
			"**`-sticky set <body>`**: Set a sticky message for this channel. This will overwrite an existing sticky.\n"
			"**`-sticky reset`**: Remove the sticky message for this channel.\n"
			"**`-sticky delay <seconds>`**: Customize the timed delay for the sticky message. Must be an integer 1 or above.\n"
			"**`-sticky pause`**: [TO-DO] Pause the sticky effect without removing the sticky message.\n"
			"**`-sticky play`**: [TO-DO] Restart the sticky effect that was previously paused.\n"
			"**`-sticky set-embed <json>`**: [TO-DO] Set a fancy embed sticky message for this channel using a json structure.\n")
			"color" $color
			"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
		)}}
		{{deleteMessage .Channel.ID $message $botdeletedelay}}
	{{end}}
{{/* if you do not have the permissions, */}}
{{else}}
	{{$message := sendMessageRetID nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
	{{deleteMessage .Channel.ID $message $botdeletedelay}}
{{end}}