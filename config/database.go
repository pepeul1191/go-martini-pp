package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DB() (*gorm.DB) {
    db, err := gorm.Open("sqlite3", "/home/pepe/Documentos/go/src/github.com/pepe/accesos/db/db_accesos.db")
    if err != nil {
      panic("failed to connect database")
    }
    //defer db.Close()
    return db
}