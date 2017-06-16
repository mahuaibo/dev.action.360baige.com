package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
)

type AccountAction struct {
}

// 新增
func (*AccountAction) Add(args *account.Account, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.UserId = args.UserId
		reply.Type = args.Type
		reply.Unit = args.Unit
		reply.Balance = args.Balance
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*AccountAction) FindById(args *account.Account, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AccountAction) UpdateById(args *account.Account, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
