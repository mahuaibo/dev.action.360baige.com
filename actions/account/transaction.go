package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
)

type TransactionAction struct {
}

// 新增
func (*TransactionAction) Add(args *account.Transaction, reply *account.Transaction) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.FromAccountId = args.FromAccountId
		reply.ToAccountId = args.ToAccountId
		reply.Amount = args.Amount
		reply.OrderCode = args.OrderCode
		reply.Remark = args.Remark
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*TransactionAction) FindById(args *account.Transaction, reply *account.Transaction) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*TransactionAction) UpdateById(args *account.Transaction, reply *account.Transaction) error {
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
