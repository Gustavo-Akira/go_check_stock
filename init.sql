CREATE TABLE stock (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price FLOAT,
    target_price FLOAT,
    links TEXT[]
);