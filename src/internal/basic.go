/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:36
 * @File  : basic.go
 * @Description: ...
 */
package internal

import (
	"github.com/mitchellh/mapstructure"
	"go-lion/utils"
)

//响应返回参数
type Result struct {
	State int `json:"state"`
}

//分页参数
type PageInput struct {
	ID       uint   `json:"id"`        //列表项id
	Page     int    `json:"page"`      //第几页
	PageSize int    `json:"page_size"` //每页数量
	Filter   string `json:"filter"`    //关键字搜索
}

//分页返回数据
type PageResult struct {
	Total int         `json:"total"` //总数
	Items interface{} `json:"items"` //数据
}

//通用返回方法
func LionResult(state int, err error) (Result, error) {
	if state == 200 {
		return Result{State: state}, nil
	} else {
		return Result{State: state}, err
	}
}

//通用分页返回方法
func LionPageResult(count int, data interface{}) PageResult {
	return PageResult{count, data}
}

func HandlePageInput(data map[string]interface{}) (PageInput, error) {
	var input PageInput
	err := mapstructure.Decode(data, &input)
	if err != nil {
		return input, err
	}
	if input.Page <= 0 {
		input.Page = 1
	}
	if input.PageSize <= 0 {
		input.PageSize = 10
	} else if input.PageSize >= 1000 {
		input.PageSize = 1000
	}

	if input.Filter != "" && len(input.Filter) > 0 {
		input.Filter = utils.StitchingStr("%", input.Filter, "%")
	}
	input.Page = (input.Page - 1) * input.PageSize
	return input, nil
}
