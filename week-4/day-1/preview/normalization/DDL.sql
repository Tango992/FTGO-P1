-- Setelah Normalisasi
CREATE DATABASE p1_preview_week3;

CREATE TABLE customers (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    PRIMARY KEY (id)
);

CREATE TABLE authors (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    PRIMARY KEY (id)
);

CREATE TABLE types (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE books (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    author_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TABLE books_price (
    id INT AUTO_INCREMENT NOT NULL,
    book_id INT NOT NULL,
    type_id INT NOT NULL,
    price FLOAT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (type_id) REFERENCES types(id)
);

CREATE TABLE orders (
    id INT AUTO_INCREMENT NOT NULL,
    customer_id INT NOT NULL,
    books_price_id INT NOT NULL,
    quantity INT DEFAULT 1,
    grand_total FLOAT NOT NULL,
    order_date DATE DEFAULT CURRENT_DATE NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (books_price_id) REFERENCES books_price(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- Creating a trigger to automatically calculates the grand total
CREATE TRIGGER set_grand_total
BEFORE INSERT ON orders 
FOR EACH ROW
SET NEW.grand_total := 
    (SELECT NEW.quantity * bp.price
    FROM books_price bp 
    WHERE bp.id = NEW.books_price_id);