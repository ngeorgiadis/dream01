package handlers

// func adminHandler(w http.ResponseWriter, r *http.Request) {

// 	b, err := ioutil.ReadFile("./ui/dist/admin.html")
// 	if err != nil {
// 		intlog.Error(err.Error())
// 	}
// 	w.Write(b)

// }

// func adminImportHandler(w http.ResponseWriter, r *http.Request) {

// 	//get file from the client
// 	f, h, err := r.FormFile("xlsxfile")
// 	if err != nil {
// 		intlog.Error(err.Error())
// 	}

// 	//read the xls file
// 	recs, errors, err := xlsreader.ReadXLSX(f, h)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(errors)
// 		return
// 	}

// 	fmt.Printf("total rows: %v\n", len(recs))

// 	db := map[string]xlsreader.XLSRow{}
// 	stor.get("db", &db)

// 	for _, rec := range recs {

// 		k := md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v%v", rec.Name, rec.Color, rec.Size, rec.Supplier, rec.OrderDate, time.Now().UTC().UnixNano())))
// 		key := fmt.Sprintf("%x", k)

// 		if _, ok := db[key]; ok {
// 			fmt.Printf("rec exists in DB: %v, %v, %v, %v, %v\n", rec.Name, rec.Color, rec.Size, rec.Supplier, rec.OrderDate)
// 		}

// 		rec.ID = key
// 		db[key] = rec
// 	}

// 	result := []xlsreader.XLSRow{}
// 	for _, v := range db {
// 		result = append(result, v)
// 	}

// 	backFile := fmt.Sprintf("db_%v_%v.zip", time.Now().Format("20060102"), time.Now().UTC().UnixNano())
// 	stor.backup(backFile)

// 	stor.put("db", db)
// 	w.WriteHeader(http.StatusOK)
// }

// func deleteCompletedHandler(w http.ResponseWriter, r *http.Request) {

// 	dateFrom := r.FormValue("dateFrom")
// 	d, err := time.Parse("02-01-2006", dateFrom)
// 	if err != nil {
// 		intlog.Error(err.Error())
// 	}

// 	db := map[string]xlsreader.XLSRow{}
// 	stor.Get("db", &db)

// 	ids := []string{}

// 	for _, item := range db {
// 		itemDate, err := time.Parse("2006-01-02", item.OrderDate)
// 		if err == nil && itemDate.Unix() <= d.Unix() && item.Status == "OK" {
// 			//delete
// 			ids = append(ids, item.ID)
// 		}
// 	}

// 	for _, id := range ids {
// 		delete(db, id)
// 	}

// 	stor.put("db", db)

// 	w.WriteHeader(http.StatusOK)
// }
