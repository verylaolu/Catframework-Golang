package Model

import (
	"encoding/json"
	"fmt"
	"sso/Common"
)

type User struct {
	Empno string
	Username string
	Title string
	Email string
}
type CasbinUser struct {
	Id            int    `gorm:"primary_key"`
	Empgrade      string `gorm:"type:varchar(32)"`
	Jobstatus     string `gorm:"type:varchar(32)"`
	Empno         string `gorm:"type:varchar(32)"`
	Orgid         int    `gorm:"type:int(11)"`
	Jobstatusdesc string `gorm:"type:varchar(50)"`
	Empkind       string `gorm:"type:varchar(32)"`
	Empid         string `gorm:"type:varchar(32)"`
	Createdate    string `gorm:"type:varchar(50)"`
	Genderdesc    string `gorm:"type:varchar(32)"`
	Postid        string `gorm:"type:varchar(32)"`
	Postname      string `gorm:"type:varchar(32)"`
	Empkinddesc   string `gorm:"type:varchar(32)"`
	Telno         string `gorm:"type:varchar(32)"`
	Gender        string `gorm:"type:varchar(32)"`
	Mobileno      string `gorm:"type:varchar(32)"`
	Email         string `gorm:"type:varchar(32)"`
	Workcity      string `gorm:"type:varchar(32)"`
	Empgradedesc  string `gorm:"type:varchar(32)"`
	Workcitydesc  string `gorm:"type:varchar(32)"`
	Title         string `gorm:"type:varchar(32)"`
	Create_time   int    `gorm:"type:int(11)"`
	Update_time   int    `gorm:"type:int(11)"`
}
func GetUserByEmpno(empno string) (* CasbinUser, error) {
	var user CasbinUser
	redis := Common.Redis

	result,_ := redis.Get(Common.StrEncode(empno)).Result()
	if(result != ""){
		err:=json.Unmarshal([]byte(result), &user)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("cache")
		return &user, nil
	}
	fmt.Println("key", result)

	db :=Common.DB

	if err := db.Table("casbin_user").Where("empno=?", empno).Find(&user).Error; err != nil {
		return nil, err
	}
	data,_ := json.Marshal(&user)
	err :=redis.Set(Common.StrEncode(empno), data,0).Err()
	if err != nil {
		fmt.Println("key", err)
	}
	return &user, nil
}
