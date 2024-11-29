#!/bin/bash

SESSION_NAME="minecraft"

START_COMMAND="/usr/lib/jvm/java-21-openjdk-amd64/bin/java -Xms2G -Xmx4G -jar purpur.jar"

if screen -ls | grep -q "$SESSION_NAME"; then
  echo "A screen session named '$SESSION_NAME' is already running."
  exit 0
fi

screen -dmS "$SESSION_NAME" /bin/bash -c "$START_COMMAND"

echo "Minecraft server started in a new screen session named '$SESSION_NAME'."
