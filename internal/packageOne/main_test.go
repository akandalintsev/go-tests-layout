package packageOne

import (
	"os"
)

func init() {
	// runs once per package
}

func SetupTest() {

}

func TearDownTest()
func TestMain(m *testing.M) {
	fmt.Println("packageOne TestMain")
	os,Exit(m.Run())
}



// setUp function, add a number to numbers slice
func setUp(key string, value int) {
    numbers[key] = value
}

// tearDown function, delete a number to numbers slice
func tearDown(key string) {
    delete(numbers, key)
}
