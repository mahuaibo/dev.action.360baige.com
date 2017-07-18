package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/application"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type ApplicationAction struct {
}

// 1
func (*ApplicationAction) Add(args *application.Account, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*ApplicationAction) AddMultiple(args []*application.Application, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*ApplicationAction) FindById(args *application.Application, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*ApplicationAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("application").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*ApplicationAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("application").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*ApplicationAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("application").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*ApplicationAction) FindByCond(args *action.FindByCond, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("application").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*ApplicationAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("application").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*ApplicationAction) ListByCond(args *action.ListByCond, reply *[]application.Application) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("application").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*ApplicationAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("application")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = 20
	}
	if args.CurrentSize == 0 {
		args.CurrentSize = 1
	}

	if args.OrderBy == nil || len(args.OrderBy) == 0 {
		args.OrderBy = []string{"id"}
	}

	var err error
	var replyList []application.Application
	reply.CurrentSize, err = o.QueryTable("application").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("application").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}
