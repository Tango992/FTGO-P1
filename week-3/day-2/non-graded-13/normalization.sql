-- School Database
CREATE TABLE IF NOT EXISTS Student (
    student_id INT AUTO_INCREMENT NOT NULL,
    student_name VARCHAR(100),
    PRIMARY KEY (student_id)
);

CREATE TABLE IF NOT EXISTS Course (
    course_id INT AUTO_INCREMENT NOT NULL,
    course_title VARCHAR(100),
    PRIMARY KEY (course_id)
);

CREATE TABLE IF NOT EXISTS Instructor (
    instructor_id INT AUTO_INCREMENT NOT NULL,
    instructor_name VARCHAR(100),
    PRIMARY KEY (instructor_id)
);

CREATE TABLE IF NOT EXISTS StudentInformation (
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

-- Ecommerce 
CREATE TABLE IF NOT EXISTS Customer (
    customer_id INT AUTO_INCREMENT NOT NULL,
    customer_name VARCHAR(100) NOT NULL,
    customer_address VARCHAR(255) NOT NULL,
    PRIMARY KEY (customer_id)
);

CREATE TABLE IF NOT EXISTS Category (
    category_id INT AUTO_INCREMENT NOT NULL,
    category_name VARCHAR(100),
    PRIMARY KEY (category_id)
);

CREATE TABLE IF NOT EXISTS Product (
    product_id INT AUTO_INCREMENT NOT NULL,
    product_name INT AUTO_INCREMENT NOT NULL,
    price FLOAT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (product_id),
    FOREIGN KEY (category_id) REFERENCES Category(category_id)
);

CREATE TABLE IF NOT EXISTS Order (
    order_id INT AUTO_INCREMENT NOT NULL,
    customer_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (order_id),
    FOREIGN KEY (customer_id) REFERENCES Customer(customer_id),
    FOREIGN KEY (product_id) REFERENCES Product(product_id)
);

-- Employee
CREATE TABLE IF NOT EXISTS Department (
    department_id INT AUTO_INCREMENT NOT NULL,
    department_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (department_id)
);

CREATE TABLE IF NOT EXISTS Employee (
    employee_id INT AUTO_INCREMENT NOT NULL,
    employee_name VARCHAR(100) NOT NULL,
    salary FLOAT NOT NULL,
    PRIMARY KEY (employee_id)
);

CREATE TABLE IF NOT EXISTS Supervisor (
    supervisor_id INT AUTO_INCREMENT NOT NULL,
    supervisor_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (supervisor_id)
);

CREATE TABLE IF NOT EXISTS EmployeeSupervisor (
    employee_supervisor_id INT AUTO_INCREMENT NOT NULL,
    employee_id INT NOT NULL,
    supervisor_id INT NOT NULL,
    PRIMARY KEY (employee_supervisor_id),
    FOREIGN KEY (employee_id) REFERENCES Employee(employee_id),
    FOREIGN KEY (supervisor_id) REFERENCES Supervisor(supervisor_id)
);