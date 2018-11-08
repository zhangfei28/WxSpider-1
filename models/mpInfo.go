package models

//"github.com/astaxie/beego/validation"
import (
	"errors"
	"time"
)

type Mp struct {
	Model
	MpName string `form:"mp_name" valid:"Required"`
	Biz    string `form:"biz" valid:"Required"`
	Photo  string `form:"photo" valid:"Required"`
}

//新增一个公众号
func (mp *Mp) AddMp() bool {
	mp.CreatedOn = time.Now().Unix()
	mp.ModifiedOn = time.Now().Unix()
	db.Create(mp)

	return true
}

//通过biz 判断公众号是否存在
func (mp *Mp) IsExist() bool {
	var mp1 Mp
	db.Select("id").Where("biz = ?", mp.Biz).First(&mp1)

	if mp1.ID > 0 {
		return true
	}

	return false
}

//通过biz获取公众号的mpid
func (mp *Mp) GetMpIdByBiz() (mpid int64, err error) {
	var mp1 Mp
	db.Select("id").Where("biz = ?", mp.Biz).First(&mp1)

	if mp1.ID > 0 {
		return mp1.ID, nil
	}

	return 0, errors.New("mp not exist")
}
