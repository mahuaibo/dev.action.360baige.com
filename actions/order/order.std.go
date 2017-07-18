package order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/order"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type OrderAction struct {
}

// 1
func (*OrderAction) Add(args *order.Account, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*OrderAction) AddMultiple(args []*order.Order, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("order")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*OrderAction) FindById(args *order.Order, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*OrderAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("order").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*OrderAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("order").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*OrderAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("order").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*OrderAction) FindByCond(args *action.FindByCond, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("order").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*OrderAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("order").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*OrderAction) ListByCond(args *action.ListByCond, reply *[]order.Order) error {
	o := orm.NewOrm()
	o.Using("order")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("order").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*OrderAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("order")

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
	var replyList []order.Order
	reply.CurrentSize, err = o.QueryTable("order").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("order").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}
