package packageOne

import "testing"

func createUser(t *testing.T, db *gorm.DB) {
	user := User{ID: "1", Username: "jane", Password: "doe123"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Delete(&user)
	})
}
func createOrder(t *testing.T, db *gorm.DB) {
	order := Order{UserID: "1"}
	if err := db.Create(&order).Error; err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Delete(&order)
	})
}

func TestServer(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./cleanups_test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	createUser(t, db)

	r := Router(db)

	// Snipped for brewity...
}
