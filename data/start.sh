#!/usr/bin/env bash
/app/create_sqlite_file.sh -f /var/local/flibudata/flibusta.db -d /app -t $telegram_api_key -c $telegram_chat_id
/app/flibustier_server --flibusta_db=/var/local/flibudata/flibusta.db --datastore=/var/local/flibudata/datastore.badger --port=9000 --update_cmd=./docker_download.sh --update_every=24h
