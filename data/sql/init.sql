CREATE DATABASE lockbin;

\c lockbin;

CREATE TABLE lockbin (
    uuid UUID PRIMARY KEY,
    masterKey VARCHAR,
    unlockTime TIMESTAMP,
    deleteTime TIMESTAMP,
    message TEXT
);
