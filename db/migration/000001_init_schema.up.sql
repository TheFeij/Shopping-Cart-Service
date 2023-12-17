CREATE TYPE BASKET_STATE AS ENUM ('COMPLETED', 'PENDING');


CREATE TABLE baskets (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    data VARCHAR(2048) NOT NULL,
    state BASKET_STATE NOT NULL
);