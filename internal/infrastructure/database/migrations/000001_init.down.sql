-- Rollback Script

-- Drop the orders table if it exists
DROP TABLE IF EXISTS orders;

-- Drop the enum types if they exist
DO $$ 
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        DROP TYPE order_status;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method') THEN
        DROP TYPE payment_method;
    END IF;
END $$;

-- Drop the update_updated_at_column function if it exists
DROP FUNCTION IF EXISTS order_updated_at_column();
