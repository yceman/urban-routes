package db

func db(){
	db, err := sql.DB("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close
}