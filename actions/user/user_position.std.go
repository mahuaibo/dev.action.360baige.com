package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
)

type UserPositionAction struct {
}

// 1
func (*UserPositionAction) Add(args *user.UserPosition, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")
	reply, err := o.Insert(args)
	return err
}

// 2
func (*UserPositionAction) AddMultiple(args []*user.UserPosition, reply int64) error {
	o := orm.NewOrm()
	o.Using("user") //查询数据库
	reply, err := o.InsertMulti(len(args), args)
	return err
}

// 3
func (*UserPositionAction) FindById(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*UserPositionAction) UpdateByCond(args *UpdateByCond, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	for _, item := range args.UpdateList {
		values[item.Key] = item.Val
	}

	reply, err := o.QueryTable("user_position").SetCond(cond).Update(values)
	return err
}

// 5
func (*UserPositionAction) DeleteById(args []int64, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args)

	values["status"] = -1

	reply, err := o.QueryTable("user_position").SetCond(cond).Update(values)
	return err
}

// 6
func (*UserPositionAction) UpdateById(args *UpdateByIdCond, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	for _, item := range args.UpdateList {
		values[item.Key] = item.Val
	}

	reply, err := o.QueryTable("user_position").SetCond(cond).Update(values)
	return err
}

// 7
func (*UserPositionAction) FindByCond(args *FindByCond, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")

	cond := orm.NewCondition()

	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	err := o.QueryTable("user_position").SetCond(cond).One(&reply, args.Fileds...)
	return err
}

// 8
func (*UserPositionAction) DeleteByCond(args *DeleteByCond, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	values["status"] = -1

	reply, err := o.QueryTable("user_position").SetCond(cond).Update(values)
	return err
}

// 9
func (*UserPositionAction) ListByCond(args *ListByCond, reply *[]user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	cond := orm.NewCondition()
	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("user_position").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply, args.Cols...)
	return err
}

// 10
func (*UserPositionAction) PageByCond(args *PageByCond, reply *PageByCond) error {
	o := orm.NewOrm()
	o.Using("user")
	cond := orm.NewCondition()

	for _, item := range args.CondList {
		if (item.Type == "And") {
			cond = cond.And(item.Key, item.Val)
		} else if (item.Type == "AndNot") {
			cond = cond.AndNot(item.Key, item.Val)
		} else if (item.Type == "Or") {
			cond = cond.Or(item.Key, item.Val)
		} else if (item.Type == "OrNot") {
			cond = cond.OrNot(item.Key, item.Val)
		}
	}
	num, err := o.QueryTable("user_position").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.CurrentSize = num
	qs := o.QueryTable("user_position").SetCond(cond)
	total, err := qs.Count()
	reply.Total = total
	return err
}
