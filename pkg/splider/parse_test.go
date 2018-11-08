package splider

import (
	"fmt"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func TestGetArticle(t *testing.T) {
	article, _ := GetArticle("https://mp.weixin.qq.com/s?__biz=MjM5MjAwODM4MA==&amp;mid=2650707746&amp;idx=1&amp;sn=f3370b5efada029c5f4a51ceb7de861b&amp;chksm=bea6e6f189d16fe7a6f703711a3dd16a3816ec5a29f2cff76099d97eb508c43ef2985f6e4b2d&amp;scene=27#wechat_redirect")
	fmt.Printf("%+v", article)
}
