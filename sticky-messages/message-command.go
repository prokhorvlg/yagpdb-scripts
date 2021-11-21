{{/*
	STICKY MESSAGES: Message Command
	Written by prokhorVLG
	Modified from Black Wolf's work, from https://github.com/BlackWolfWoof/yagpdb-cc/

	Trigger type: None
	
	This command is triggered by the regex command to actually post the sticky message.
*/}}

{{/* grab the sticky message from storage by key */}}
{{$db := (dbGet .Channel.ID "stickymessage").Value}}
{{/* set message value */}}
{{$message := $db.message}}
{{/* delete the old sticky message based on key */}}
{{if $db := dbGet .Channel.ID "smchannel"}}
	{{deleteMessage nil (toInt $db.Value) 0}}
{{end}}
{{/* post sticky to channel, and save id to variable */}}
{{$id := sendMessageRetID nil $message}}
{{/* save, to the channel's metadata, the key of the message */}}
{{dbSet .Channel.ID "smchannel" (str $id)}}
{{/* set delay to false for this channel */}}
{{dbSet .Channel.ID "indelay" false}}