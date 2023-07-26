#!/bin/bash

# Load environment variables from .env file
# while read -r line; do
#   if [[ $line =~ ^[[:space:]]*# || $line =~ ^[[:space:]]*$ ]]; then
#     continue
#   fi
#   export "$line"
# done < .env

# Print the environment variables
export APP_NAME= "SIR"
export SERVER_HOST="localhost"
export SERVER_PORT="8080"
export DB_USERNAME="freedb_gobabak"
export DB_PASSWORD="AyaRn8%ymVPH!uA"
export DB_HOST="sql.freedb.tech"
export DB_PORT="3306"
export DB_NAME="freedb_gobabak"