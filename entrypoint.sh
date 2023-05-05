#!/bin/sh -e

# exec env file if exists, .env might be in /.env or /config/.env
FILE=/.env ; [ -f $FILE ] && . $FILE
FILE=/config/.env ; [ -f $FILE ] && . $FILE

# start main program
#cd app
#ls -la
#exec ./app -profiler=true
exec /app/app.bin -profiler=true