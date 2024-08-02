#!/usr/bin/env sh

set -x

if [ -n "$SKIP_FLIBUSTA_DOWNLOAD" ]; then
  echo "SKIP_FLIBUSTA_DOWNLOAD is set. Skipping the execution of the downloader."
  exit 0
fi

echo "PRODUCTION run of downloader"

/app/scripts/init-database.py \
  --telegram_api_key=${TELEGRAM_API_KEY} \
  --telegram_chat_id=${TELEGRAM_CHAT_ID} \
  --mysql_host=${MARIADB_HOST} \
  --mysql_user=${MARIADB_USER} \
  --mysql_password=${MARIADB_PASSWORD} \
  --mysql_database=${MARIADB_DATABASE} \
  --mysql_port=${MARIADB_PORT} \
  --flibusta_base_url=${FLIBUSTA_BASE_URL}
