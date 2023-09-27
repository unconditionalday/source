#!/bin/bash

calver=$(date +%Y-%m-%d)

git tag -a "$calver" -m "Version $calver"
git push origin "$calver"

OWNER="unconditionalday"
REPO="source-test"

TOKEN=$1

response=$(curl -X POST "https://api.github.com/repos/$OWNER/$REPO/releases" \
  -H "Authorization: token $TOKEN" \
  -d '{
    "tag_name": "'"$calver"'",
    "name": "'"$calver"'",
    "body": "'"$calver"'"
  }')

release_id=$(echo "$response" | jq -r '.id')

if [ "$release_id" != "null" ]; then
  echo "Release $calver creata con successo su GitHub (ID: $release_id)"
else
  echo "Errore nella creazione della release su GitHub."
  echo "$response"
  exit 1
fi

FILE_PATH="source.json"

upload_url=$(echo "$response" | jq -r '.upload_url' | sed 's/{?name,label}//')

curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" \
  --data-binary @"$FILE_PATH" \
  "$upload_url?name=source.json"

echo "Allegato aggiunto con successo alla release."
