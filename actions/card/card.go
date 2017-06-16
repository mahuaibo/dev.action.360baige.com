package card

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/card"
)

type CardAction struct {
}

// 新增
func (*CardAction) Add(args *card.Card, reply *card.Card) error {
	o := orm.NewOrm()
	o.Using("card")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.PersonId = args.PersonId
		reply.Cardno = args.Cardno
		reply.Physicsno = args.Physicsno
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*CardAction) FindById(args *card.Card, reply *card.Card) error {
	o := orm.NewOrm()
	o.Using("card")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*CardAction) UpdateById(args *card.Card, reply *card.Card) error {
	o := orm.NewOrm()
	o.Using("card")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
