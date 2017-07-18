package logger

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/logger"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type LoggerAction struct {
}

// 1
func (*LoggerAction) Add(args *logger.Account, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*LoggerAction) AddMultiple(args []*logger.Logger, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("logger")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*LoggerAction) FindById(args *logger.Logger, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*LoggerAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("logger").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*LoggerAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("logger").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*LoggerAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("logger").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*LoggerAction) FindByCond(args *action.FindByCond, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("logger").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*LoggerAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("logger").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*LoggerAction) ListByCond(args *action.ListByCond, reply *[]logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("logger").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*LoggerAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("logger")

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
	var replyList []logger.Logger
	reply.CurrentSize, err = o.QueryTable("logger").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("logger").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}
