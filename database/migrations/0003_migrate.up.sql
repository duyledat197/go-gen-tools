-- Just make sure you already enabled PGroonga Extension
CREATE EXTENSION IF NOT EXISTS pgroonga;

CREATE INDEX IF NOT EXISTS pgroonga_name_index ON users USING pgroonga (name);

