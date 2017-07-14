package order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/order"
	"dev.model.360baige.com/models/batch"
	"dev.model.360baige.com/http/window"
	"strings"
	"time"
)

type OrderAction struct {
}

// 新增
func (*OrderAction) Add(args *order.Order, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.UserPositionType = args.UserPositionType
		reply.UserPositionId = args.UserPositionId
		reply.Code = args.Code
		reply.Price = args.Price
		reply.Type = args.Type
		reply.PayType = args.PayType
		reply.Brief = args.Brief
		reply.Status = args.Status
		
	}
	return err
}

// 查询 by Id
func (*OrderAction) FindById(args *order.Order, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*OrderAction) UpdateById(args *order.Order, reply *order.Order) error {
	o := orm.NewOrm()
	o.Using("order")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

// 1. AddMultiple 增加多个
func (*OrderAction) AddMultiple(args []*order.Order, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("order") //查询数据库
	num, err := o.InsertMulti(100, args)
	reply.Num = num
	return err
}

// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
func (*OrderAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("order")            //查询数据库
	qs := o.QueryTable("order") //查询表名
	if (args.UpdateTime == 0) {
		args.UpdateTime = time.Now().UnixNano() / 1e6
	}
	if (args.Status == 0) {
		args.Status = -1
	}
	idsArg := strings.Split(args.Ids, ",")
	qs = qs.Filter("id__in", idsArg)
	num, err := qs.Update(orm.Params{
		"status":      args.Status,
		"update_time": args.UpdateTime,
	})
	reply.Num = num
	return err
}

//分页list
func (*OrderAction) PageBy(args *window.OrderListPaginator, reply *window.OrderListPaginator) error {
	o := orm.NewOrm()
	o.Using("order")
	cond := orm.NewCondition()
	for _, c := range args.Cond {
		if (c.Type == "And") {
			cond = cond.And(c.Exprs, c.Args)
		} else if (c.Type == "AndNot") {
			cond = cond.AndNot(c.Exprs, c.Args)
		} else if (c.Type == "Or") {
			cond = cond.Or(c.Exprs, c.Args)
		} else if (c.Type == "OrNot") {
			cond = cond.OrNot(c.Exprs, c.Args)
		}
	}
	num, err := o.QueryTable("order").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.CurrentSize = num
	qs := o.QueryTable("order").SetCond(cond)
	total, err := qs.Count()
	reply.Total = total
	return err
}