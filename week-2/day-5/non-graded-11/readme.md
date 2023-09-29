# Entity Relationship Diagram

## Perusahaan

Setiap `karyawan` memiliki NIK, nama depan, nama belakang, jenis kelamin, email, phone numberSetiap `karyawan` menempati satu `departemen`Setiap `Departemen` memiliki `manager` dan hanya boleh memiliki satu `manager` sajaDi `perusahaan` tersebut memiliki beberapa `projek`, setiap `karyawan` dapat mengerjakan lebih dari satu `projek` dan `projek` tersebut setidaknya dikerjakan minimal oleh satu `karyawan`

![ERD Perusahaan](perusahaan.svg)

## Address Book

Contact memiliki property/column nama, perusahaan, nomor telepon dan emailGroup memiliki property/column namaSetiap Contact dapat memiliki lebih dari 1 Group begitupun sebaliknya.

![ERD Address Book](address-book.svg)

## Ecommerce

Terdapat 2 jenis user yang bisa log in ke sistem, yaitu admin dan customer.
Admin dapat input, edit, dan delete item, sedangkan customer tidak boleh ada akses ini.
Customer dapat melakukan transaksi, dalam 1 transaksi, customer dapat membeli lebih dari 1 item.
Terdapat tags seperti : kulit, suede, satin, cotton, small, medium, large.
Setiap item, bisa memiliki lebih dari 1 tag.

![ERD Ecommerce](ecommerce.svg)