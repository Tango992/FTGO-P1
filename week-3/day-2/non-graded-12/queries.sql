CREATE TABLE IF NOT EXISTS country (
    country_id INT AUTO_INCREMENT NOT NULL,
    country_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (country_id)
);

CREATE TABLE IF NOT EXISTS authors (
    author_id INT AUTO_INCREMENT NOT NULL,
    author_name VARCHAR(100) NOT NULL UNIQUE,
    country_id INT NOT NULL,
    PRIMARY KEY (author_id),
    FOREIGN KEY (country_id) REFERENCES country(country_id)
);

CREATE TABLE IF NOT EXISTS book (
    book_id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL UNIQUE,
    author_id INT NOT NULL,
    publication_year INT NOT NULL,
    available_quantity INT,
    PRIMARY KEY (book_id),
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

CREATE TABLE IF NOT EXISTS book_loans (
    loan_id INT AUTO_INCREMENT NOT NULL,
    book_id INT NOT NULL,
    borrower_name VARCHAR(100) NOT NULL,
    loan_date DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    return_date DATETIME,
    PRIMARY KEY (loan_id),
    FOREIGN KEY (book_id) REFERENCES book(book_id)
);

CREATE INDEX idx_auth_id ON authors (author_id);

CREATE INDEX idx_return_date ON book_loans (return_date, loan_date, book_id);