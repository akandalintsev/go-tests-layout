package subpackageone

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("subpackageOne TestMain")
	os.Exit(m.Run())
}
