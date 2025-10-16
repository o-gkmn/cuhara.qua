#!/bin/bash
set -e

# Replication kullanıcısını ve DB'yi idempotent şekilde oluştur
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-'EOSQL'
DO $$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'replica_user') THEN
      CREATE USER replica REPLICATION LOGIN PASSWORD 'replica';
   END IF;
END$$;

DO $$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'cuhara_qua_db') THEN
      CREATE DATABASE cuhara_qua_db;
   END IF;
END$$;

GRANT ALL PRIVILEGES ON DATABASE cuhara_qua_db TO postgres;
EOSQL