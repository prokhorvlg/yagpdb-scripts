{{/*
	Trigger type: Command
	Trigger: sticky

	Copyright (c): Black Wolf, 2021
	License: MIT
	Repository: https://github.com/BlackWolfWoof/yagpdb-cc/
	Modified by prokhorVLG
*/}}

{{/* permissions that may access the functionality of the bot
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
{{$perms := "ManageMessages"}}

{{/* arguments that can be passed to this command */}}
{{$args := parseArgs 0 "-sticky <command> <body>" (carg "string" "command") (carg "string" "body")}}

{{/* fancy shmancy message color */}}
{{$color := 6356832}}

{{/* if you have the permissions, */}}
{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{/* if there is any message after the triggering command, do stuff */}}
	{{if .StrippedMsg}}
		{{/* if the message is "reset", reset the sticky to nothing */}}
		{{if (eq ($args.Get 0) "reset")}}
			{{/* delete the sticky from channel's metadata by key */}}
			{{dbDel .Channel.ID "stickymessage"}}
			{{/* delete the old sticky message based on last key */}}
			{{if $db := dbGet .Channel.ID "smchannel"}}
				{{deleteMessage nil (toInt $db.Value) 0}}
			{{end}}
			{{/* delete the key metadata once we're done with it */}}
			{{dbDel .Channel.ID "smchannel"}}
			{{/* inform the user */}}
			{{sendMessage nil (cembed 
				"description" "Sticky message has been removed and reset."
				"color" $color
				"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
			)}}
		{{/* if the message is "set", set the message to the second arg */}}
		{{else if (eq ($args.Get 0) "set")}}
			{{/* if there is a second arg for body, */}}
			{{if ($args.Get 1)}}
				{{/* set the sticky! */}}
				{{/* set image to empty */}}
				{{$img := ""}}
				{{/* set text to second argument */}}
				{{$text := $args.Get 1}}
				{{with reFindAllSubmatches `(?:(?P<TxtSnip1>(?:.*[\r\n]?){0,}))?(?:-img\s(?P<Link>(?:https?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+))(?P<TxTSnip2>(?:.*[\r\n]?){0,})` $text}}
					{{$img = index . 0 2}}
					{{$text = print (index . 0 1) (index . 0 3)}}
				{{end}}
				{{with reFindAllSubmatches `\A((?:.|[\r\n])*)(-d\s(?P<Duration>(?:(?:\d+)?(?:months?|mo|minutes?|s|seconds?|m|hours?|h|days?|d|weeks?|w|years?|y|permanent|p)){1,}))(\s(?:.|[\r\n])*)?\z` $text}}
					{{$text = print (index . 0 1) (index . 0 4)}}
				{{end}}
				{{/* save message to key under the channel id */}}
				{{dbSet .Channel.ID "stickymessage" (sdict "message" $text "author" .User.String "img" $img)}}
				{{/* inform the user */}}
				{{sendMessage nil (cembed 
					"description" "Sticky message has been enabled and set!"
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
			{{/* if there is no second arg for body, */}}
			{{else}}
				{{/* inform the user */}}
				{{sendMessage nil (cembed 
					"description" "You must add a message body to set a sticky message."
					"color" $color
					"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
				)}}
			{{end}}
		{{end}}
	{{/* if there is no message after the triggering command, */}}
	{{else}}
		{{/* show the user the available commands */}}
		{{sendMessage nil (cembed 
			"description" "Commands include:\n- **sticky set <body>**: Set a sticky message for this channel\n- **sticky reset**: Remove the sticky for this channel"
			"color" $color
			"author" (sdict "name" (print "Sticky Messages") "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png")
		)}}
	{{end}}
{{/* if you do not have the permissions, */}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}