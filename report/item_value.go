package report

import (
	"bytes"
	"context"
	"encoding/csv"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mistikel/inventoy/model"
)

func (reportModule *ReportModule) GetItemValueReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	record := []string{"SSI-D00791015-LL-BWH", "Zalekia Plain Casual Blouse (L,Broken White)", "53", "Rp70,448", "Rp3,733,735"}

	itemDatamodel := model.NewItemModel(ctx)
	TotalSku, _ := itemDatamodel.GetMany(ctx) // len(totalsku)

	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	// Header
	wr.Write([]string{"LAPORAN NILAI BARANG"})
	wr.Write([]string{})
	wr.Write([]string{"Tanggal Cetak", "8 Januari 2018"})
	wr.Write([]string{"Jumlah SKU", "31"})
	wr.Write([]string{"Jumlah Total Barang", "4086"})
	wr.Write([]string{"Total Nilai", "Rp286,272,941"})
	wr.Write([]string{})
	wr.Write([]string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})

	for i := 0; i < 20; i++ {
		wr.Write(record)
	}
	wr.Flush()

	w.Header().Set("Content-Type", "text/csv")

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
	w.Write(b.Bytes())
}
