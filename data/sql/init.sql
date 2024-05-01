SELECT 'CREATE DATABASE lockbin'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'lockbin')\gexec

\c lockbin;

CREATE TABLE IF NOT EXISTS lockbin (
    uuid UUID PRIMARY KEY NOT NULL,
    masterKey VARCHAR NOT NULL,
    unlockTime TIMESTAMP NOT NULL,
    deleteTime TIMESTAMP NOT NULL,
    message TEXT NOT NULL
);
