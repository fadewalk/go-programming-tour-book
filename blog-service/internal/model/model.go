package model

import (
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"


	"github.com/fadewalk/go-programming-tour-book/blog-service/global"
	"github.com/fadewalk/go-programming-tour-book/blog-service/pkg/setting"
)
type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB,error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local")
	db, err := gorm.Open(databaseSetting.DBType,s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host, 
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,)

	if err != nil {
		return nil, err
	}
	
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil

}