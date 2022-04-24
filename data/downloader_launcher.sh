#!/usr/bin/env bash

/app/create_sqlite_file.sh -f /var/local/flibudata/flibusta.db -d /app -t $telegram_api_key -c $telegram_chat_id
/app/flibustier_server --flibusta_db=/var/local/flibudata/flibusta.db --port=9000