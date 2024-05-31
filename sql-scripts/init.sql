CREATE TABLE users
(
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "login" VARCHAR(255) UNIQUE NOT NULL,
    "password" TEXT NOT NULL UNIQUE,
    "phone_number" VARCHAR(15) NOT NULL
);

CREATE TABLE contact
(
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(15) NOT NULL
);

CREATE TABLE product
(
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "serial_number" VARCHAR(15) NOT NULL,
    "seller_id" INTEGER NOT NULL REFERENCES contact(id)
);

CREATE TABLE orders
(
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "date" DATE NOT NULL

);

CREATE TABLE order_product
(
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "order_id" BIGINT NOT NULL REFERENCES "orders" (id),
    "product_id" BIGINT NOT NULL REFERENCES "product" (id),
    "amount" BIGINT NOT NULL

);

INSERT INTO users (name, role, login, password, phone_number) VALUES
                                                                  ('John Doe', 'admin', 'johndoe', 'password123', '+380971234567'),
                                                                  ('Jane Smith', 'admin', 'janesmith', 'qwerty', '+380981234567'),
                                                                  ('Bob Johnson', 'admin', 'bobjohnson', 'password456', '+380991234567'),
                                                                  ('Alice Williams', 'admin', 'alicewilliams', 'abc123', '+380961234567'),
                                                                  ('Tom Brown', 'admin', 'tombrown', 'pass123', '+380951234567');

INSERT INTO contact (name, role, phone_number) VALUES
                                                   ('Acme Corp', 'seller', '+380971111111'),
                                                   ('Globex Inc', 'seller', '+380972222222'),
                                                   ('Stark Industries', 'seller', '+380973333333'),
                                                   ('Umbrella Corp', 'buyer', '+380974444444'),
                                                   ('Cyberdyne Systems', 'buyer', '+380975555555');

INSERT INTO product (name, description, serial_number, seller_id) VALUES
                                                                      ('iPhone 13', 'Apple smartphone', 'A123456789', 1),
                                                                      ('Samsung Galaxy S22', 'Samsung smartphone', 'B987654321', 2),
                                                                      ('Dell XPS 13', 'Dell laptop', 'C456789123', 3),
                                                                      ('Sony PlayStation 5', 'Sony gaming console', 'D789123456', 4),
                                                                      ('Microsoft Surface Pro', 'Microsoft tablet', 'E321654987', 5);
INSERT INTO orders (date) VALUES
                              ('2023-05-01'),
                              ('2023-05-10'),
                              ('2023-05-15'),
                              ('2023-05-20'),
                              ('2023-05-25');

INSERT INTO order_product (order_id, product_id, amount) VALUES
                                                             (1, 1, 2),
                                                             (1, 3, 1),
                                                             (2, 2, 1),
                                                             (3, 4, 1),
                                                             (4, 5, 3),
                                                             (5, 1, 1),
                                                             (5, 3, 2);




