-- DDL
CREATE DATABASE p1_livecode_library;

CREATE TABLE buku (
    buku_id INT AUTO_INCREMENT PRIMARY KEY,
    judul VARCHAR(255) NOT NULL,
    pengarang VARCHAR(255) NOT NULL,
    tahun_terbit YEAR,
    stok INT DEFAULT 0
);

CREATE TABLE anggota (
    anggota_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    alamat TEXT,
    tanggal_daftar DATE DEFAULT CURRENT_DATE
);

CREATE TABLE peminjaman (
    peminjaman_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    anggota_id INT,
    tanggal_pinjam DATE DEFAULT CURRENT_DATE,
    tanggal_kembali DATE,
    FOREIGN KEY (anggota_id) REFERENCES anggota(anggota_id)
);

CREATE TABLE detail_peminjaman (
    detail_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    peminjaman_id INT,
    buku_id INT,
    jumlah_buku INT DEFAULT 0,
    denda_keterlambatan DECIMAL(10,2) DEFAULT 0,
    FOREIGN KEY (peminjaman_id) REFERENCES peminjaman(peminjaman_id),
    FOREIGN KEY (buku_id) REFERENCES buku(buku_id)
);


-- DML
INSERT INTO anggota (nama, alamat)
VALUES 
    ("John Doe", "Jalan Sukawati"), 
    ("Jane Smith", "Jalan Ligma");

INSERT INTO buku (judul, pengarang, stok, tahun_terbit)
VALUES
	("Harry Potter", "J.K. Rowling", 1, 2000),
	("The Hobbit", "J.R.R. Tolkien", 1, 1937);