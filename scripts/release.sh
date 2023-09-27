#!/bin/bash

calver=$(date +%Y-%m-%d)

git tag -a "$calver" -m "Version $calver"
git push origin "$calver"

OWNER=$1
REPO=$2
TOKEN=$3

response=$(curl -X POST "https://api.github.com/repos/$OWNER/$REPO/releases" \
  -H "Authorization: token $TOKEN" \
  -d '{
    "tag_name": "'"$calver"'",
    "name": "'"$calver"'",
    "body": "'"$calver"'"
  }')

release_id=$(echo "$response" | jq -r '.id')

if [ "$release_id" != "null" ]; then
  echo "Release $calver created successfully on GitHub (ID: $release_id)"
else
  echo "Error creating the release on GitHub."
  echo "$response"
  exit 1
fi

FILE_PATH="source.json"

upload_url=$(echo "$response" | jq -r '.upload_url' | sed 's/{?name,label}//')

curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" \
  --data-binary @"$FILE_PATH" \
  "$upload_url?name=source.json"

echo "Attachment added successfully to the release."
