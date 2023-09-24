package transaction

type RegisterData struct {
	Nama             string `json:"nama" binding:"required"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	Keterangan       string `json:"keterangan" binding:"required"`
	MingguKe         int    `json:"minggu_ke" binding:"required"`
	JumlahMasuk      int    `json:"jumlah_masuk" binding:"required"`
	Error            error
}

type FormUpdateDataInput struct {
	ID               int    `json:"id"`
	Nama             string `json:"nama binding:"required"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	Keterangan       string `json:"keterangan"`
	MingguKe         int    `json:"minggu_ke binding:"required"`
	JumlahMasuk      int    `json:"jumlah_masuk binding:"required"`
	Error            error
}

type FormCreateDataInput struct {
	ID               int    `json:"id"`
	Nama             string `json:"nama" binding:"required"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	Keterangan       string `json:"keterangan" binding:"required"`
	MingguKe         int    `json:"minggu_ke" binding:"required"`
	JumlahMasuk      int    `json:"jumlah_masuk binding:"required"`
	Error            error
}
