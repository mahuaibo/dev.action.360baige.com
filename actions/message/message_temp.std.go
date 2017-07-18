package message

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/message"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type MessageTempAction struct {
}

// 1
func (*MessageTempAction) Add(args *message.MessageTemp, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*MessageTempAction) AddMultiple(args []*message.MessageTemp, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("message")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*MessageTempAction) FindById(args *message.MessageTemp, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*MessageTempAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("message_temp").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*MessageTempAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("message_temp").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*MessageTempAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("message_temp").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*MessageTempAction) FindByCond(args *action.FindByCond, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("message_temp").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*MessageTempAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("message_temp").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*MessageTempAction) ListByCond(args *action.ListByCond, reply *[]message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("message_temp").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*MessageTempAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("message")

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
	var replyList []message.MessageTemp
	reply.CurrentSize, err = o.QueryTable("message_temp").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("message_temp").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}
