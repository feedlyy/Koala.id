CREATE TABLE customers (
    customer_id VARCHAR(64) NOT NULL PRIMARY KEY,
    customer_name VARCHAR(80) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    email VARCHAR(50) NOT NULL,
    dob DATE NOT NULL,
    sex VARCHAR(20) NOT NULL,
    salt VARCHAR(80) NOT NULL,
    password VARCHAR(400) NOT NULL,
    create_date TIMESTAMP NOT NULL
);

CREATE TABLE tokens (
    token_id VARCHAR(100) NOT NULL PRIMARY KEY,
    token VARCHAR(100),
    refresh_type VARCHAR(100),
    customer_id VARCHAR(64) NOT NULL REFERENCES customers (customer_id) ON DELETE CASCADE
);

CREATE TABLE products (
  product_id VARCHAR(64) NOT NULL PRIMARY KEY,
  product_name VARCHAR(80),
  basic_price money,
  created_date TIMESTAMP
);

CREATE TABLE paymentMethods (
    payment_method_id VARCHAR(64) NOT NULL PRIMARY KEY,
    method_name VARCHAR(70) NOT NULL,
    created_date TIMESTAMP
);

CREATE TABLE orders (
  order_id VARCHAR(64) NOT NULL PRIMARY KEY,
  customer_id VARCHAR(80) NOT NULL REFERENCES customers (customer_id) ON DELETE CASCADE,
  order_number VARCHAR(40) NOT NULL,
  payment_method_id VARCHAR(64) NOT NULL REFERENCES paymentMethods (payment_method_id) ON DELETE CASCADE
);

CREATE TABLE orderDetails (
    order_detail_id VARCHAR(64) NOT NULL PRIMARY KEY,
    order_id VARCHAR(64) NOT NULL REFERENCES orders (order_id) ON DELETE CASCADE,
    product_id VARCHAR(64) NOT NULL REFERENCES products (product_id) ON DELETE  CASCADE,
    qty int,
    created_date TIMESTAMP
);