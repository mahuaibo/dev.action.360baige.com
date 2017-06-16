package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
)

type AccountItemAction struct {
}

// 新增
func (*AccountItemAction) Add(args *account.AccountItem, reply *account.AccountItem) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.TransactionId = args.TransactionId
		reply.AccountId = args.AccountId
		reply.Amount = args.Amount
		reply.Balance = args.Balance
		reply.Remark = args.Remark

	}
	return err
}

// 查询 by Id
func (*AccountItemAction) FindById(args *account.AccountItem, reply *account.AccountItem) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AccountItemAction) UpdateById(args *account.AccountItem, reply *account.AccountItem) error {
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
