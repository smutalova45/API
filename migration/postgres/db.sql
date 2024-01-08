CREATE TABLE users(
    id uuid primary key,
    firstname varchar(15),
    email varchar(16),
    phone varchar(15)
);

CREATE TABLE orders(
    id uuid primary key,
    amount int,
    created_at timestamp default current_timestamp,
    user_id uuid references users(id)
);

CREATE TABLE products(
    id uuid primary key,
    productname varchar(15),
    price int
);

CREATE TABLE orderproducts(
    id uuid primary key,
    order_id uuid references orders(id),
    product_id uuid references products(id),
    quantity int,
    price int
);