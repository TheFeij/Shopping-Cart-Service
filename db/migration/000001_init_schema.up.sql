CREATE TYPE shopping_cart_state AS ENUM ('COMPLETED', 'PENDING');


CREATE TABLE shopping_carts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    data VARCHAR(2048) NOT NULL,
    state SHOPPING_CART_STATE NOT NULL
);