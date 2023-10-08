package handler

import (
	"database/sql"
	"fmt"
	"p1-latihan-lc3/entity"
)

func DisplayBorrow(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT p.peminjaman_id, a.nama, b.judul, p.tanggal_pinjam, p.tanggal_kembali
		FROM peminjaman p
		JOIN anggota a ON p.anggota_id = a.anggota_id
		JOIN detail_peminjaman dp ON dp.peminjaman_id = p.peminjaman_id
		JOIN buku b ON dp.buku_id = b.buku_id`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var borrows []entity.DisplayBorrow
	for rows.Next() {
		var borrow entity.DisplayBorrow
		if err := rows.Scan(&borrow.PeminjamanId, &borrow.Nama, &borrow.Judul, &borrow.TanggalPinjam, &borrow.TanggalKembali); err != nil {
			return err
		}
		borrows = append(borrows, borrow)
	}
	BorrowsParser(borrows)
	return nil
}

func BorrowsParser(borrows []entity.DisplayBorrow) {
	fmt.Println("Daftar Peminjaman:")
	for _, row :=  range borrows {
		fmt.Printf("ID: %v, Pelanggan: %v, Buku: %v, Tanggal Pinjam: %v, Tanggal Kembali: %v\n", row.PeminjamanId, row.Nama, row.Judul, row.TanggalPinjam, row.TanggalKembali)
	}
}