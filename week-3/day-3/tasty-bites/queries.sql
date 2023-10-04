-- Normalized DDL
CREATE TABLE Position (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE Employees (
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    position_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (position_id) REFERENCES Position (id)
);

CREATE TABLE Category (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);
    
CREATE TABLE MenuItems (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(30) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    category_id INT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES Category (id)
);

CREATE TABLE Status (
    id INT AUTO_INCREMENT NOT NULL,
    status VARCHAR(20) NOT NULL,   
    PRIMARY KEY (id)
);

CREATE TABLE Orders (
    id INT AUTO_INCREMENT NOT NULL,
    table_number INT NOT NULL,
    employee_id INT NOT NULL,
    order_date DATE NOT NULL DEFAULT CURRENT_DATE,
    status_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (status_id) REFERENCES Status (id)
);

CREATE TABLE OrderItems (
    id INT AUTO_INCREMENT NOT NULL,
   	order_id INT NOT NULL,
    item_id INT NOT NULL,
    quantity INT,
    subtotal FLOAT,
    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES Orders (id),
    FOREIGN KEY (item_id) REFERENCES MenuItems (id)
);

CREATE TABLE Method (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(20) NOT NULL,   
    PRIMARY KEY (id)
);

CREATE TABLE Payments (
    id INT AUTO_INCREMENT NOT NULL,
    order_id INT NOT NULL,
    payment_date DATE DEFAULT CURRENT_DATE,
    method_id INT NOT NULL,
    total_amount FLOAT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES Orders (id),
    FOREIGN KEY (method_id) REFERENCES Method (id)
);

-- DML 
INSERT INTO Position (name)
VALUES ('Waiter'), ('Cashier'), ('Chief'), ('Manager');

INSERT INTO Employees (first_name, last_name, position_id)
VALUES ('John', 'Doe', 1);

INSERT INTO Category (name)
VALUES ('Main Course');

INSERT INTO MenuItems (name, description, price, category_id)
VALUES 
    ('Steak', 'Grilled sirloin steak', 25.99, 1),
    ('Carbonara', 'Spaghetti carbonara with mushroom', 20.00, 1);

INSERT INTO Status (status)
VALUES ('Pending'), ('Completed'), ('Canceled');

INSERT INTO Orders (table_number, employee_id, order_date, status_id)
VALUES 
    (5, 1, '2023-08-04', 1),
    (6, 1, '2023-08-05', 2);

INSERT INTO OrderItems (order_id, item_id, quantity, subtotal)
VALUES 
    (1, 1, 2, 51.98),
    (2, 2, 5, 100.00);

INSERT INTO Method (name)
VALUES ('Cash'), ('Credit Card'), ('QRIS');

INSERT INTO Payments (order_id, payment_date, method_id, total_amount)
VALUES 
    (1, '2023-08-04', 2, 51.98),
    (2, '2023-08-05', 3, 90.00);


-- Retrieve all orders with their applied discounts:
SELECT p.order_id, o.subtotal, p.total_amount AS GrandTotal, ((o.subtotal - p.total_amount) / o.subtotal) * 100 AS DiscountPercentage
FROM Payments p
JOIN OrderItems o ON p.order_id = o.order_id;

-- Calculate the total revenue (including discounts) for a specific day:
SELECT ROUND(SUM(total_amount),2) AS TotalRevenue
FROM Payments
WHERE payment_date = '2023-08-04';