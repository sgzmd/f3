#!/bin/bash

echo "PRODUCTION run of downloader"
/app/create_sqlite_file.sh -f /var/local/flibudata/flibusta.db -d /app -t $telegram_api_key -c $telegram_chat_id