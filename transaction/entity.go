package transaction

import (
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID               int       `json:"id"`
	Nama             string    `json:"nama"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	Keterangan       string    `json:"keterangan"`
	MingguKe         int       `json:"minggu_ke"`
	JumlahMasuk      int       `json:"jumlah_masuk"`
	Error            error
}

func (Transaction) TableName() string {
	return "transactions"
}

func (c Transaction) FormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.JumlahMasuk)
}

func (t Transaction) FormatTanggal() string {
	return t.TanggalTransaksi.Format("02-01-2006")
}
