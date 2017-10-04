package model

import (
	"../common"
	"github.com/yejiansnake/go-yedb"
)

type Test struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Type      uint8     `json:"type"`
	Create_at string `json:"create_at"`
	Update_at string `json:"update_at"`
}

func GetTestModel() *yedb.DbModel {
	return yedb.ModelNew(common.DB_NAME_TEST, "t_test")
}
