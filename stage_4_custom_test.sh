#!/bin/sh
# I also added a sleep time wbefore writing the response back within the script
# to examine if all these background processes are really spawning individual routines
i=0
while [ $i -ne 100 ]
do
  i=$(($i+1))
  echo -e "ping\nping" | redis-cli &
done
wait