{{/*
	Trigger type: Regex
	Trigger: \A

	Copyright (c): Black Wolf, 2021
	License: MIT
	Repository: https://github.com/BlackWolfWoof/yagpdb-cc/
	Modified by prokhorVLG
*/}}

{{/* if sticky message exists, */}}
{{if (dbGet .Channel.ID "stickymessage")}}
	{{/* grab the sticky message from storage by key */}}
	{{$db := (dbGet .Channel.ID "stickymessage").Value}}
	{{/* set message value */}}
	{{$message := $db.message}}
	{{/* delete the old sticky message based on key */}}
	{{if $db := dbGet .Channel.ID "smchannel"}}
		{{deleteMessage nil (toInt $db.Value) 0}}
	{{end}}
	{{/* send message to channel, and saves the id for the message in variable */}}
	{{$id := sendMessageRetID nil $message}}
	{{/* save, to the channel's metadata, the key of the message */}}
	{{dbSet .Channel.ID "smchannel" (str $id)}}
{{end}}