package utils

import (
	"dev.model.360baige.com/action"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/context/param"
)

//字符串和字符串数组
func ConvertUint8ToString(cc interface{}) interface{} {
	if _, ok := cc.([]uint8); ok {
		cc = string(cc.([]uint8))
	} else {
		if _, ok := cc.([]interface{}); ok {
			var crr []interface{}
			crr = cc.([]interface{})
			if len(crr) > 0 {
				for key, value := range crr {
					if _, ok := value.([]uint8); ok {
						crr[key] = string(value.([]uint8))
					}
				}
				cc = crr
			}
		}

	}
	return cc
}

//
func ConvertCond(condList []action.CondValue) *orm.Condition {
	cond := orm.NewCondition()
	for _, item := range condList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, ConvertUint8ToString(item.Val))
		}
	}
	return cond
}

//
func ConvertValues(updateList []action.UpdateValue) orm.Params {
	values := orm.Params{}
	for _, item := range updateList {
		values[item.Key] = ConvertUint8ToString(item.Val)
	}
	return values
}
