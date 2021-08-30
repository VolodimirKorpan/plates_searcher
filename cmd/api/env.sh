#!/bin/bash
export PORT=8080

# Postgres settings
export PG_ADDR=10.12.0.4
export PG_PORT=5432
export PG_USER=postgres
export PG_PASSWORD=postgres
export PG_DB=plates

# MySQL settings
export MYSQL_ADDR=127.0.0.1:3306
export MYSQL_USER=root
export MYSQL_PASSWORD=golang
export MYSQL_DB=kobi
export HASH_SALT=hash_satl
export SIGNING_KEY=signing_key
export TOKEN_TTL=86400
