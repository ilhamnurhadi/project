package data

import (
	"database/sql"
	"day15/item/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository struct {
	DB *sql.DB
}

// untuk nilai return get all butuh struktur dari item
// 2.a buat model dari...
func GetAll(db *ItemRepository) []models.Item {
	fmt.Println(db.DB)
	result, err := db.DB.Query("Select ItemName From tblItem")
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	item := []models.Item{}
	for result.Next() {
		var i models.Item
		if err := result.Scan(&i.ItemName); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}
