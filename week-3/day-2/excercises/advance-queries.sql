-- SUBSTRING
SELECT author_name, SUBSTRING(suthor_name, 1, 5) AS short_name FROM authors;

-- CONCAT
SELECT CONCAT(author_name, 'from', author_country) AS Author_info FROM authors;

-- Menggunakan DATE_ADD untuk menambahkan hari ke tanggal.
SELECT book_id, DATE_ADD(loan_date, INTERVAL 7 DAY) AS due_date FROM book_loans;

-- Menggunakan DATEDIFF untuk menghitung selisih hari.
SELECT book_id, DATEDIFF(return_date, loan_date) as loan_duration FROM book_loans;

-- Menggunakan CASE untuk memberi label status buku berdasarkan ketersediaan
SELECT title,
    CASE
        WHEN available_quantity > 0 THEN 'Available'
        ELSE 'Out of stock'
    END AS Book_status
FROM books;

-- Menggunakan CASE untuk mengelompokkan buku berdasarkan tahun publikasi: 
-- If >= 2020 THEN 'RECENT'
SELECT title, publication_year,
    CASE
        WHEN publication_year >= 2020 THEN 'RECENT'
        WHEN publication_year >= 2010 AND publication_year < 2020 THEN 'MODERATE'
        ELSE 'OLD'
    END AS categories
FROM books;