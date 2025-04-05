BEGIN;

CREATE EXTENSION IF NOT EXISTS citext;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_name = 'mre_users'
    ) THEN
        CREATE TABLE mre_users (
            id bigserial PRIMARY KEY,
            full_name varchar(255) NOT NULL,
            email_address citext UNIQUE NOT NULL,
            password bytea NOT NULL,
            created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
            updated_at timestamp(0)
        );
    END IF;
END $$;

CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_trigger
        WHERE tgname = 'trigger_update_timestamp'
    ) THEN
        CREATE TRIGGER trigger_update_timestamp
        BEFORE UPDATE ON mre_users
        FOR EACH ROW
        EXECUTE FUNCTION set_updated_at_column();
    END IF;
END $$;

COMMIT;
