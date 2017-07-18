package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
)

type AccountAction struct {
}

// 1
func (*AccountAction) Add(args *account.Account, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*AccountAction) AddMultiple(args []*account.Account, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("account") //查询数据库
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*AccountAction) FindById(args *account.Account, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*AccountAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
		}
	}

	values := orm.Params{}
	for _, item := range args.UpdateList {
		values[item.Key] = utils.ConvertUint8ToString(item.Val)
	}

	num, err := o.QueryTable("account").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*AccountAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	values := orm.Params{}
	values["status"] = -1

	num, err := o.QueryTable("account").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 6
func (*AccountAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := orm.Params{}
	for _, item := range args.UpdateList {
		values[item.Key] = utils.ConvertUint8ToString(item.Val)
	}

	num, err := o.QueryTable("account").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*AccountAction) FindByCond(args *action.FindByCond, reply *account.Account) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
		}
	}

	err := o.QueryTable("account").SetCond(cond).One(&reply, args.Fileds...)
	return err
}

// 8
func (*AccountAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
		}
	}

	values := orm.Params{}
	values["status"] = -1

	num, err := o.QueryTable("account").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 9
func (*AccountAction) ListByCond(args *action.ListByCond, reply *[]account.Account) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
		}
	}

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("account").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply, args.Cols...)
	return err
}

// 10
func (*AccountAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("account")

	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
		}
	}

	var err error
	reply.CurrentSize, err = o.QueryTable("account").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.Total, err = o.QueryTable("account").SetCond(cond).Count()
	return err
}
