#!/usr/bin/env sh

set -x
echo "PRODUCTION run of downloader"

/app/scripts/init-database.py \
  --telegram_api_key=${TELEGRAM_API_KEY} \
  --telegram_chat_id=${TELEGRAM_CHAT_ID} \
  --mysql_host=${MARIADB_HOST} \
  --mysql_user=${MARIADB_USER} \
  --mysql_password=${MARIADB_PASSWORD} \
  --mysql_database=${MARIADB_DATABASE} \
  --mysql_port=${MARIADB_PORT}