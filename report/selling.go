package report

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (reportModule *ReportModule) GetSellingReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	record := []string{"ID-20171130-531154", "2017-12-01 1:18:53", "SSI-D00791015-LL-BWH", "Zalekia Plain Casual Blouse (L,Broken White)", "1", "Rp115,000", "Rp115,000", "Rp67,500", "Rp47,500"}

	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	// Header
	wr.Write([]string{"LAPORAN PENJUALAN"})
	wr.Write([]string{})
	wr.Write([]string{"Tanggal Cetak", "8 Januari 2018"})
	wr.Write([]string{"Tanggal", "1 December 2018 - 31 December 2018"})
	wr.Write([]string{"Total Omzet", "Rp85,690,000"})
	wr.Write([]string{"Total Laba Kotor", "Rp24,592,000"})
	wr.Write([]string{"Total Penjualan", "528"})
	wr.Write([]string{"Total Barang", "712"})
	wr.Write([]string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})

	for i := 0; i < 20; i++ {
		wr.Write(record)
	}
	wr.Flush()

	w.Header().Set("Content-Type", "text/csv")

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
	w.Write(b.Bytes())
}
