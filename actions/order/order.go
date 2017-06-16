package order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/order"
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
