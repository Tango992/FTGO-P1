# Entity Relationship Diagram

## Perusahaan

Setiap `karyawan` memiliki NIK, nama depan, nama belakang, jenis kelamin, email, phone number.\
Setiap `karyawan` menempati satu `departemen`.\
Setiap `Departemen` memiliki `manager` dan hanya boleh memiliki satu `manager` saja.\
Di `perusahaan` tersebut memiliki beberapa `projek`, setiap `karyawan` dapat mengerjakan lebih dari satu `projek` dan `projek` tersebut setidaknya dikerjakan minimal oleh satu `karyawan`.

![ERD Perusahaan](perusahaan.svg)

## Address Book

Contact memiliki property/column nama, perusahaan, nomor telepon dan email.

Group memiliki property/column nama.

Setiap Contact dapat memiliki lebih dari 1 Group begitupun sebaliknya.

![ERD Address Book](address-book.svg)

## Ecommerce

Terdapat 2 jenis user yang bisa log in ke sistem, yaitu admin dan customer.\
Admin dapat input, edit, dan delete item, sedangkan customer tidak boleh ada akses ini.\
Customer dapat melakukan transaksi, dalam 1 transaksi, customer dapat membeli lebih dari 1 item.\
Terdapat tags seperti : kulit, suede, satin, cotton, small, medium, large.\
Setiap item, bisa memiliki lebih dari 1 tag.

![ERD Ecommerce](ecommerce.svg)

## Poll

Sistem dapat menyimpan data mengenai kandidat2 pejabat, terdiri dari nama, partai, lokasi dan grade_current.\
grade_current merupakan angka decimal.\
Sistem juga harus menyimpan data mengenai orang yang akan ikut dalam voting tsb, \
di antaranya :nama depan, nama belakang, jenis kelamin, umur.\
Sistem dapat mencatat siapa saja yang vote pejabat. \
1 orang dapat melakukan vote lebih dari 1 pejabat.\
Kandidat penjabat bisa di vote oleh banyak orang.

