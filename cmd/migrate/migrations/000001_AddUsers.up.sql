CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS mre_users (
  id bigserial PRIMARY KEY,
  full_name varchar(255) NOT NULL,
  email_address citext UNIQUE NOT NULL,
  password bytea NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0)
  );


CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_timestamp
BEFORE UPDATE ON mre_users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_column();
