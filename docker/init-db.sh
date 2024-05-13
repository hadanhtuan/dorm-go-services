#!/usr/bin/env bash
echo "** Creating DB"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d "$POSTGRES_DB"  <<-EOSQL
      CREATE DATABASE ${DB_PROPERTY_SERVICE};
      CREATE DATABASE ${DB_MEMBER_SERVICE};
      CREATE DATABASE ${DB_LOCATION_SERVICE};
      CREATE DATABASE ${DB_CHAT_SERVICE};
      CREATE DATABASE ${DB_PAYMENT_SERVICE};
EOSQL

echo "** Finished creating default DB"