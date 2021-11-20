
{{/*
	Trigger type: None
	by prokhorVLG
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