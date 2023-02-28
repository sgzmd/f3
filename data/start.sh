#!/usr/bin/env sh

set -x

bash /app/scripts/docker_download.sh
/app/flibustier_server --flibusta_db=/var/local/flibudata/flibusta.db \
  --datastore=/var/local/flibudata/datastore.badger \
  --port=9000 \
  --update_cmd=/app/scripts/docker_download.sh \
  --update_every=24h \
  --mysql_host=${MARIADB_HOST} \
  --mysql_user=${MARIADB_USER} \
  --mysql_pass=${MARIADB_PASSWORD} \
  --mysql_db=${MARIADB_DATABASE} \
  --mysql_port=${MARIADB_PORT}


