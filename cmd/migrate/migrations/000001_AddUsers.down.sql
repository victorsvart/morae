BEGIN;

-- Drop trigger if it exists
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM pg_trigger 
        WHERE tgname = 'trigger_update_timestamp'
    ) THEN
        DROP TRIGGER trigger_update_timestamp ON mre_users;
    END IF;
END $$;

-- Drop function if it exists
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM pg_proc 
        WHERE proname = 'set_updated_at_column'
    ) THEN
        DROP FUNCTION set_updated_at_column;
    END IF;
END $$;

-- Drop table if it exists
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM information_schema.tables 
        WHERE table_name = 'mre_users'
    ) THEN
        DROP TABLE mre_users;
    END IF;
END $$;

COMMIT;
