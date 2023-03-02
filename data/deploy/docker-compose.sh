#!/usr/bin/env sh
set -x

tag=`date --iso-8601`T`date +%H-%M`

echo "Creating git tag $tag"

git tag $tag
git push --tags

echo $tag > version.txt

echo "Starting docker compose"

docker compose up --build -d