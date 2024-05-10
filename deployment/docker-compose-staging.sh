#!/usr/bin/env sh
set -x

docker context use default

tag=`date --iso-8601`T`date +%H-%M`

echo "Creating git tag $tag"

git tag $tag
git push --tags

echo $tag > version.txt

echo "Starting docker compose"

docker compose --env-file deployment/staging/staging.env \
  -f deployment/docker-compose.yml up --build -d