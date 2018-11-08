package validate

import (
	"WxSpider/pkg/logging"

	"github.com/astaxie/beego/validation"
)

func Check(mp interface{}) []string {
	valid := validation.Validation{}
	var errs []string
	b, err := valid.Valid(mp)
	if err != nil {
		logging.Warn("valid error:", err)
	}

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			errs = append(errs, err.Key+":"+err.Message)
		}
	}

	return errs
}
