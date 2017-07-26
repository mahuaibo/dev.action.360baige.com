package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	. "dev.action.360baige.com/database"
	"time"
	"encoding/json"
)

type TransactionAction struct {
}

// 1
func (*TransactionAction) Add(args *account.Transaction, reply *account.Transaction) error {
	o := GetOrmer(DB_account)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*TransactionAction) AddMultiple(args []*account.Transaction, reply *action.Num) error {
	o := GetOrmer(DB_account)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*TransactionAction) FindById(args *account.Transaction, reply *account.Transaction) error {
	o := GetOrmer(DB_account)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*TransactionAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_account)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("transaction").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*TransactionAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_account)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("transaction").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*TransactionAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_account)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("transaction").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*TransactionAction) FindByCond(args *action.FindByCond, reply *account.Transaction) error {
	o := GetOrmer(DB_account)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("transaction").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*TransactionAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_account)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("transaction").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*TransactionAction) ListByCond(args *action.ListByCond, reply *[]account.Transaction) error {
	o := GetOrmer(DB_account)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("transaction").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*TransactionAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := GetOrmer(DB_account)

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
	var replyList []account.Transaction
	reply.CurrentSize, err = o.QueryTable("transaction").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("transaction").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*TransactionAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_account)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("transaction").SetCond(cond).Count()
	reply.Value = num
	return err
}
