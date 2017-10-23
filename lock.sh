#!/usr/bin/env bash

tmpbg='/tmp/screen.jpeg'
outbg='/tmp/out.png'

scrot "$tmpbg"
convert "$tmpbg" -scale 4% -scale 2500% "$outbg"

# convert "$tmpbg" "$icon" -gravity center -composite -matte "$tmpbg"
#i3lock -u -i "$tmpbg"
#i3lock -i "$tmpbg"
rm "$tmpbg"
rm "$outbg"
