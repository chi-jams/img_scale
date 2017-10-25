#!/usr/bin/env bash

tmpbg='/tmp/screen.jpeg'
outbg='/tmp/out.png'

scrot "$tmpbg"
./img_scale "$tmpbg" "$outbg"

# convert "$tmpbg" "$icon" -gravity center -composite -matte "$tmpbg"
#i3lock -u -i "$tmpbg"
#i3lock -i "$outbg"

rm "$tmpbg"
rm "$outbg"
