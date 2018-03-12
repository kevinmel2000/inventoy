package report

import (
	"bytes"
	"context"
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mistikel/inventoy/model"
)

func (reportModule *ReportModule) GetItemValueReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	year, month, day := time.Now().Date()
	itemDatamodel := model.NewItemModel(ctx)
	record := itemDatamodel.GetRecord(ctx)
	TotalSku, _ := itemDatamodel.GetMany(ctx)
	totalItem := itemDatamodel.GetTotalItem(ctx)
	totalValue := itemDatamodel.GetTotalValue(ctx)
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	// Header
	wr.Write([]string{"LAPORAN NILAI BARANG"})
	wr.Write([]string{})
	wr.Write([]string{"Tanggal Cetak", strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)})
	wr.Write([]string{"Jumlah SKU", strconv.Itoa(len(TotalSku))})
	wr.Write([]string{"Jumlah Total Barang", strconv.Itoa(totalItem)})
	wr.Write([]string{"Total Nilai", "Rp" + strconv.Itoa(totalValue)})
	wr.Write([]string{})
	wr.Write([]string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})

	for _, r := range record {
		wr.Write([]string{r.Sku, r.Name, strconv.Itoa(r.TotalItem), "Rp" + strconv.Itoa(r.Avarage), "Rp" + strconv.Itoa(r.TotalValue)})
	}
	wr.Flush()

	w.Header().Set("Content-Type", "text/csv")

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=Laporan Nilai Barang.csv")
	w.Write(b.Bytes())
}
