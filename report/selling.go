package report

import (
	"bytes"
	"context"
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/mistikel/inventoy/model"

	"github.com/julienschmidt/httprouter"
)

func (reportModule *ReportModule) GetSellingReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	m, _ := strconv.Atoi(p.ByName("month"))
	m += 1
	dt1 := p.ByName("year") + "-" + p.ByName("month") + "-01"
	dt2 := p.ByName("year") + "-" + strconv.Itoa(m) + "-01"

	outboundModel := model.NewOutboundItemModel(ctx)
	record, omset, laba, selling, totalItem := outboundModel.GetRecord(ctx, dt1, dt2)

	year, month, day := time.Now().Date()

	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	// Header
	wr.Write([]string{"LAPORAN PENJUALAN"})
	wr.Write([]string{})
	wr.Write([]string{"Tanggal Cetak", strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)})
	wr.Write([]string{"Tanggal", dt1 + " - " + dt2})
	wr.Write([]string{"Total Omzet", "Rp" + strconv.Itoa(omset)})
	wr.Write([]string{"Total Laba Kotor", "Rp" + strconv.Itoa(laba)})
	wr.Write([]string{"Total Penjualan", strconv.Itoa(selling)})
	wr.Write([]string{"Total Barang", strconv.Itoa(totalItem)})
	wr.Write([]string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})

	for _, r := range record {
		wr.Write([]string{strconv.Itoa(r.Id), r.Time.String(), r.Sku, r.Name, strconv.Itoa(r.TotalItem), strconv.Itoa(r.SellingPrice), strconv.Itoa(r.TotalPrice), strconv.Itoa(r.BuyingPrice), strconv.Itoa(r.Laba)})
	}
	wr.Flush()

	w.Header().Set("Content-Type", "text/csv")

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=Laporan Penjualan.csv")
	w.Write(b.Bytes())
}
