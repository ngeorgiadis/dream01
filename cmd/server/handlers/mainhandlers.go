package handlers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"dream01/internal/intlog"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("./ui/dist/index.html")
	if err != nil {
		intlog.Error(err.Error())
	}
	w.Write(b)

}

func getRecordsHandler(w http.ResponseWriter, r *http.Request) {

	result, err := getAllRecords()
	if err != nil {
		intlog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	//get file from the client
	f, h, err := r.FormFile("xlsfile")
	if err != nil {
		intlog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//read the xls file
	recs, err := xlsreader.ReadXLS(f, h)
	if err != nil {
		intlog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	intlog.Infof("total rows: %v\n", len(recs))

	// backup files
	backFile := fmt.Sprintf("db_%v_%v.zip", time.Now().Format("20060102"), time.Now().UTC().UnixNano())
	stor.BackupTo(backFile)

	for _, rec := range recs {

		k := md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v%v", rec.Name, rec.Color, rec.Size, rec.Supplier, rec.OrderDate, time.Now().UTC().UnixNano())))
		key := fmt.Sprintf("%x", k)

		// CHECK FOR DUPLICATE KEY
		//
		// if _, ok := db[key]; ok {
		// 	fmt.Printf("rec exists in DB: %v, %v, %v, %v, %v\n", rec.Name, rec.Color, rec.Size, rec.Supplier, rec.OrderDate)
		// }

		fmt.Printf("rec: %v, %v, %v, %v, %v\n", rec.Name, rec.Color, rec.Size, rec.Supplier, rec.OrderDate)

		rec.ID = key

		b, _ := json.Marshal(rec)
		stor.Put(key, b)
	}

	result, err := getAllRecords()
	if err != nil {
		intlog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//clear undo list
	undo := map[int][]xlsreader.XLSRow{}
	b, _ := json.Marshal(undo)
	stor.Put("undolist", b)

	//retrun result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

// func saveCommentHandler(w http.ResponseWriter, r *http.Request) {

// 	rec := r.FormValue("record")
// 	//comment := r.FormValue("comment")
// 	//fmt.Println(rec, comment)

// 	var record xlsRow
// 	err := json.Unmarshal([]byte(rec), &record)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func updateRecordHandler(w http.ResponseWriter, r *http.Request) {

// 	var record xlsreader.XLSRow
// 	err := json.Unmarshal([]byte(r.FormValue("record")), &record)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	db := map[string]xlsreader.XLSRow{}
// 	stor.get("db", &db)
// 	db[record.ID] = record
// 	stor.put("db", db)

// 	w.WriteHeader(http.StatusOK)

// }

// func updateRecordsHandler(w http.ResponseWriter, r *http.Request) {

// 	var recs []xlsreader.XLSRow
// 	err := json.Unmarshal([]byte(r.FormValue("recs")), &recs)

// 	if err != nil {
// 		log.Error(err)
// 	}

// 	db := map[string]xlsreader.XLSRow{}
// 	stor.get("db", &db)
// 	for _, rec := range recs {
// 		db[rec.ID] = rec
// 	}
// 	stor.put("db", db)
// 	w.WriteHeader(http.StatusOK)
// }

// func deleteRecordsHandler(w http.ResponseWriter, r *http.Request) {

// 	ids := []string{}
// 	err := json.Unmarshal([]byte(r.FormValue("ids")), &ids)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	//load
// 	db := map[string]xlsreader.XLSRow{}
// 	undoBatch := []xlsreader.XLSRow{}

// 	stor.get("db", &db)

// 	for _, id := range ids {
// 		if _, ok := db[id]; ok {

// 			//save to undo list
// 			undoBatch = append(undoBatch, db[id])

// 			//delete
// 			delete(db, id)
// 		}
// 	}

// 	//save undo list
// 	undo := map[int][]xlsRow{}
// 	stor.get("undo", &undo)
// 	idx := len(undo)
// 	undo[idx] = undoBatch
// 	stor.put("undo", undo)

// 	//delete
// 	stor.put("db", db)

// 	w.WriteHeader(http.StatusOK)
// }

// func undoDeleteHandler(w http.ResponseWriter, r *http.Request) {

// 	undo := map[int][]xlsreader.XLSRow{}
// 	stor.get("undo", &undo)
// 	idx := len(undo)

// 	db := map[string]xlsreader.XLSRow{}
// 	stor.get("db", &db)

// 	if undo != nil {
// 		undoBatch := []xlsreader.XLSRow{}
// 		undoBatch = undo[idx-1]
// 		for _, item := range undoBatch {
// 			db[item.ID] = item
// 		}

// 		delete(undo, idx-1)
// 	}

// 	stor.put("db", db)
// 	stor.put("undo", undo)

// 	w.WriteHeader(http.StatusOK)
// }

func getAllRecords() ([]xlsreader.XLSRow, error) {

	result := []xlsreader.XLSRow{}
	files, err := stor.GetIndex()
	if err != nil {
		return nil, err
	}

	for _, f := range files {

		buff, err := stor.Get(f.Name())
		if err != nil {
			return nil, err
		}

		row := xlsreader.XLSRow{}

		json.Unmarshal(buff, &row)
		result = append(result, row)
	}

	return result, nil
}
