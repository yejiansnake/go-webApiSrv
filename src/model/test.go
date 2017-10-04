package model

import (
	"../common"
	"github.com/yejiansnake/go-yedb"
)

type Test struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Type      uint8     `json:"type"`
	Create_at string `json:"createAt"`
	Update_at string `json:"updateAt"`
}

func GetTestModel() *yedb.DbModel {
	return yedb.ModelNew(common.DB_NAME_TEST, "t_test")
}
