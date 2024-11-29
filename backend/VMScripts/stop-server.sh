#!/bin/bash

SESSION_NAME="minecraft"
STOP_COMMAND="screen -S $SESSION_NAME -X stuff 'stop'`echo -ne '\015'`"

# Check if the screen session is running
if screen -ls | grep -q "$SESSION_NAME"; then
  # Send the stop command to the screen session
  /bin/bash -c "$STOP_COMMAND"

  # Wait until the screen session is no longer running
  while screen -ls | grep -q "$SESSION_NAME"; do
    echo "Waiting for Minecraft server to stop..."
    sleep 6 # Check every 6 seconds if the session is still running
  done

  echo "Minecraft server has stopped."
  exit 0
fi

echo "Minecraft server not already running."
exit 0