package v1

import (
	"log"
	"testing"
	"time"

	//"github.com/Unknwon/com"
)

func TestGetLogFilePath(t *testing.T) {
	var n int64
	n = 100
	i := time.Time().Second()
	log.Println("test...")
	log.Println(i)
}
