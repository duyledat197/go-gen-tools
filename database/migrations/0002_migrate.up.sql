CREATE TABLE IF NOT EXISTS users (
  id text PRIMARY KEY,
  name text NOT NULL,
  bio text,
  updated_at timestamptz NOT NULL
);

CREATE OR REPLACE FUNCTION public.test()
RETURNS TABLE (id text)
LANGUAGE sql
STABLE
AS $$
SELECT
    id
FROM
    users
$$;