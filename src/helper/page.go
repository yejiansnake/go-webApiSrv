package helper

import (
	"errors"
	"github.com/yejiansnake/go-yedb"
)

type PageMeta struct {
	PageIndex int `json:"pageIndex"`
	PageSize int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	RowCount int64`json:"rowCount"`
}

type PageData struct {
	Items interface{} `json:"items"`
	Meta PageMeta `json:"meta"`
}

func GetPageDataEx(query yedb.IQuery, ptrRows interface{}, pageIndex string, pageSize string) (data PageData, err error)  {
	_pageIndex, _ := StrToInt(pageIndex)
	_pageSize, _ := StrToInt(pageSize)
	return GetPageData(query, ptrRows, _pageIndex, _pageSize)
}

func GetPageData(query yedb.IQuery, ptrRows interface{}, pageIndex int, pageSize int) (data PageData, err error)  {
	if query == nil {
		err = errors.New("params invalid")
		return
	}

	if pageIndex == 0 {
		pageIndex = 1
	}

	if pageSize == 0 {
		pageSize = 20
	}

	if pageIndex > 0 && pageSize > 0 {
		data.Meta.RowCount, _ = query.Limit(0).Offset(0).Count()
		data.Meta.PageIndex = pageIndex
		data.Meta.PageSize = pageSize

		if data.Meta.RowCount % int64(pageSize) != 0 {
			data.Meta.PageCount = int(data.Meta.RowCount / int64(pageSize) + 1)
		} else {
			data.Meta.PageCount = int(data.Meta.RowCount / int64(pageSize))
		}

		offset := (pageIndex - 1) * pageSize
		query.Limit(int64(pageSize)).Offset(int64(offset))
	}

	err = query.FillRows(ptrRows)
	if err != nil {
		return
	}

	data.Items = ptrRows
	return
}