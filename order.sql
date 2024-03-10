CREATE DATABASE order;

-- Create orders table
CREATE TABLE orders (
  id BIGINT PRIMARY KEY,
  customer_name VARCHAR(50),
  ordered_at TIMESTAMP
);

-- Create items table
CREATE TABLE items (
  id BIGINT PRIMARY KEY,
  code VARCHAR(10),
  description VARCHAR(50),
  quantity BIGINT,
  order_id BIGINT REFERENCES orders(id)
);


