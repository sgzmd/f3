#!/usr/bin/env sh

set -x

bash /app/scripts/docker_download.sh
/app/flibustier_server --flibusta_db=/var/local/flibudata/flibusta.db --datastore=/var/local/flibudata/datastore.badger --port=9000 --update_cmd=./docker_download.sh --update_every=24h

