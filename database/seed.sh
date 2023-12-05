#!/bin/sh
set -eux
cd `dirname $0`

mysql -u"$WEBAPI_DB_USER" \
  -p"$WEBAPI_DB_PASSWORD" \
  --host "$WEBAPI_DB_HOST" \
  --port "$WEBAPI_DB_PORT" \
  "$WEBAPI_DB_NAME" < seed.sql
