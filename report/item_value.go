package report

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (reportModule *ReportModule) GetItemValueReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	record := []string{"test1", "test2", "test3"} // just some test data to use for the wr.Writer() method below.

	b := &bytes.Buffer{}       // creates IO Writer
	wr := csv.NewWriter(b)     // creates a csv writer that uses the io buffer.
	for i := 0; i < 100; i++ { // make a loop formake 100 rows just for testing purposes
		wr.Write(record) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
	w.Write(b.Bytes())
}
