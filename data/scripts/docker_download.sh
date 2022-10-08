#!/bin/bash

echo "PRODUCTION run of downloader"
#/app/create_sqlite_file.sh -f /var/local/flibudata/flibusta.db -d /app -t
#${TELEGRAM_API_KEY} -c ${TELEGRAM_CHAT_ID}
/app/scripts/init-database.py \
  --telegram_api_key=${TELEGRAM_API_KEY} \
  --telegram_chat_id=${TELEGRAM_CHAT_ID} \
  --create_sqlite_file=/var/local/flibudata/flibusta.db