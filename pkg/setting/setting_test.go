package setting

import (
	"fmt"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func clearConfig() {
	DbConfig = &Db{}
}

func TestDbConfig(t *testing.T) {
	InitSetUp("../../conf/app.ini")
	defer clearConfig()

	fmt.Printf("%+v", AppConfig)
}
