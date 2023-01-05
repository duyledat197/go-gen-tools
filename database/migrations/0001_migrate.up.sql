CREATE TABLE IF NOT EXISTS users (
  id text PRIMARY KEY,
  name text NOT NULL,
  type text NOT NULL,
  team_id text NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz NULL,
  deleted_at timestamptz 
);

CREATE TABLE IF NOT EXISTS teams (
    id text PRIMARY KEY,
    name text NOT NULL,
    type text NOT NULL,
    hub_id text NOT NULL,
    location_id text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    deleted_at timestamptz 
);

CREATE TABLE IF NOT EXISTS hubs (
    id text PRIMARY KEY,
    name text NOT NULL,
    location_id text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    deleted_at timestamptz 
);

ALTER TABLE users ADD CONSTRAINT users_team_id_fk FOREIGN KEY (team_id) REFERENCES teams(id);
ALTER TABLE teams ADD CONSTRAINT teams_hub_id_fk FOREIGN KEY (hub_id) REFERENCES hubs(id);

