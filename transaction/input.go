package user

import "time"

type RegisterData struct {
	Nama             string    `json:"nama" binding:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" binding:"required"`
	Keterangan       string    `json:"keterangan" binding:"required"`
	MingguKe         int       `json:"minggu_ke" binding:"required"`
	JumlahMasuk      int       `json:"jumlah_masuk" binding:"required"`
}

type FormUpdateDataInput struct {
	ID               int       `json:"id"`
	Nama             string    `json:"nama binding:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" binding:"required"`
	Keterangan       string    `json:"keterangan" binding:"required"`
	MingguKe         int       `json:"minggu_ke binding:"required"`
	JumlahMasuk      int       `json:"jumlah_masuk binding:"required"`
	Error            error
}

type FormCreateDataInput struct {
	ID               int       `json:"id"`
	Nama             string    `json:"nama" binding:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" binding:"required"`
	Keterangan       string    `json:"keterangan" binding:"required"`
	MingguKe         int       `json:"minggu_ke" binding:"required"`
	JumlahMasuk      int       `json:"jumlah_masuk binding:"required"`
	Error            error
}
