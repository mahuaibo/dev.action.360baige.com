package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/application"
	"dev.model.360baige.com/models/paginator"
	"dev.model.360baige.com/models/batch"
	"strings"
	"encoding/json"
	"time"
	"dev.model.360baige.com/action"
)

type ApplicationTplAction struct {
}

// 1
func (*ApplicationTplAction) Add(args *application.ApplicationTpl, reply *application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*ApplicationTplAction) AddMultiple(args []*application.ApplicationTpl, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application") //查询数据库
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*ApplicationTplAction) FindById(args *application.ApplicationTpl, reply *application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*ApplicationTplAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	var values orm.Params
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	for _, item := range args.UpdateList {
		values[item.Key] = item.Val
	}

	num, err := o.QueryTable("application_tpl").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*ApplicationTplAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	values["status"] = -1

	num, err := o.QueryTable("application_tpl").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 6
func (*ApplicationTplAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	for _, item := range args.UpdateList {
		values[item.Key] = item.Val
	}

	num, err := o.QueryTable("application_tpl").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*ApplicationTplAction) FindByCond(args *action.FindByCond, reply *application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := orm.NewCondition()

	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	err := o.QueryTable("application_tpl").SetCond(cond).One(&reply, args.Fileds...)
	return err
}

// 8
func (*ApplicationTplAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	var values orm.Params
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	values["status"] = -1

	num, err := o.QueryTable("application_tpl").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 9
func (*ApplicationTplAction) ListByCond(args *action.ListByCond, reply *[]application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}
	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("application_tpl").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply, args.Cols...)
	return err
}

// 10
func (*ApplicationTplAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("application")
	cond := orm.NewCondition()

	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}
	var err error
	reply.CurrentSize, err = o.QueryTable("application_tpl").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.Total, err = o.QueryTable("application_tpl").SetCond(cond).Count()
	return err
}
