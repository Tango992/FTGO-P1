-- DML
INSERT INTO customers (name, email)
VALUES 
	("John Doe", "john.doe@email.com"),
    ("Alice Bob", "alice.bob@email.com"),
    ("Ferdi Samsak", "ferdi.samsak@mail.com");

INSERT INTO authors (name, email)
VALUES 
	("Jane Smith", "jane.smith@email.com"),
    ("Tom Brown", "tom.brown@email.com");

INSERT INTO types (name)
VALUES ("Physical"), ("E-Book");

INSERT INTO books (name, author_id)
VALUES
	("Database Design", 1),
    ("Web Development", 2);

INSERT INTO books_price (book_id, type_id, price)
VALUES 
	(1, 1, 25.99),
    (1, 2, 20.99),
    (2, 1, 24.99),
    (2, 2, 19.99);

INSERT INTO orders (customer_id, books_price_id, order_date)
VALUES 
	(1, 1, "2023-08-10"),
    (1, 3, "2023-08-11"),
    (2, 2, "2023-08-12"),
    (3, 1, CURRENT_DATE),
    (3, 4, CURRENT_DATE);


-- List all books by 'Jane Smith'.
SELECT b.name
FROM books b
JOIN authors a ON b.author_id = a.id
WHERE a.name = "Jane Smith";

-- Find the total sales (in terms of price) for each book type. 
SELECT t.name AS Type, ROUND(SUM(o.grand_total),2) As GrandTotal
FROM orders o 
JOIN books_price bp ON o.books_price_id = bp.id
JOIN types t ON t.id = bp.type_id
GROUP BY t.id;

-- Identify customers who have ordered more than one book.
WITH all_transactions AS 
    (SELECT c.name, COUNT(c.id) AS Transactions
    FROM orders o
    JOIN customers c ON c.id = o.customer_id
    GROUP BY c.id)
SELECT *
FROM all_transactions
WHERE Transactions > 1;

-- Display the author who has the highest earnings from their books.
SELECT a.name, ROUND(SUM(o.grand_total), 2) AS Earnings
FROM orders o
JOIN books_price bp ON o.books_price_id = bp.id
JOIN books b ON b.id = bp.book_id
JOIN authors a ON a.id = b.author_id
GROUP BY a.name
ORDER BY Earnings DESC LIMIT 1;