#!/bin/sh
#
# Finally execute stuff in the X environment

xset s off
xset -dpms
xset s noblank
xte 'sleep 60' 'mousemove 10000 10000' 'key F11' &
epiphany-browser https://www.bytemark.co.uk/
