-- Create enum
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        CREATE TYPE order_status AS ENUM ('PAID', 'FAILED', 'EXPIRED', 'PENDING');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method') THEN
        CREATE TYPE payment_method AS ENUM ('APP_BANKING', 'CREDIT_CARD', 'QR_CODE', 'TRUEMONEY');
    END IF;
END $$;

-- Create order table
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  charge_id VARCHAR(255) UNIQUE,
  invoice_no VARCHAR(255) UNIQUE,
  vat NUMERIC(5, 2),
  service NUMERIC(5, 2),
  amount NUMERIC(10, 2) NOT NULL,
  total_amount NUMERIC(10, 2) NOT NULL,
  order_status order_status,
  payment_method payment_method,
  seat_label VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create a trigger function to update the updated_at column
CREATE OR REPLACE FUNCTION order_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

 -- Create a trigger to call the above function before each update
CREATE TRIGGER update_orders_updated_at
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE FUNCTION order_updated_at_column();
