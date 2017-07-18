package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
)

type UserAction struct {
}

// 1
func (*UserAction) Add(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*UserAction) AddMultiple(args []*user.User, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("user") //查询数据库
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*UserAction) FindById(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*UserAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("user")

	//cond := orm.NewCondition()
	//for _, item := range args.CondList {
	//	if (item.Type == "And") {
	//		cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "AndNot") {
	//		cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "Or") {
	//		cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "OrNot") {
	//		cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	}
	//}
	cond := utils.ConvertCond(args.CondList)

	//values := orm.Params{}
	//for _, item := range args.UpdateList {
	//	values[item.Key] = utils.ConvertUint8ToString(item.Val)
	//}
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("user").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*UserAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("user")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("user").SetCond(cond).Update(orm.Params{"status": -1})
	reply.Value = num
	return err
}

// 6
func (*UserAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("user")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	//values := orm.Params{}
	//for _, item := range args.UpdateList {
	//	values[item.Key] = utils.ConvertUint8ToString(item.Val)
	//}
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("user").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*UserAction) FindByCond(args *action.FindByCond, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")

	//cond := orm.NewCondition()
	//for _, item := range args.CondList {
	//	if (item.Type == "And") {
	//		cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "AndNot") {
	//		cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "Or") {
	//		cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "OrNot") {
	//		cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	}
	//}
	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("user").SetCond(cond).One(&reply, args.Fileds...)
	return err
}

// 8
func (*UserAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("user")

	//cond := orm.NewCondition()
	//for _, item := range args.CondList {
	//	if (item.Type == "And") {
	//		cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "AndNot") {
	//		cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "Or") {
	//		cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "OrNot") {
	//		cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	}
	//}
	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("user").SetCond(cond).Update(orm.Params{"status": -1})
	reply.Value = num
	return err
}

// 9
func (*UserAction) ListByCond(args *action.ListByCond, reply *[]user.User) error {
	o := orm.NewOrm()
	o.Using("user")

	//cond := orm.NewCondition()
	//for _, item := range args.CondList {
	//	if (item.Type == "And") {
	//		cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "AndNot") {
	//		cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "Or") {
	//		cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "OrNot") {
	//		cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	}
	//}

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("user").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply, args.Cols...)
	return err
}

// 10
func (*UserAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("user")

	//cond := orm.NewCondition()
	//for _, item := range args.CondList {
	//	if (item.Type == "And") {
	//		cond = cond.And(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "AndNot") {
	//		cond = cond.AndNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "Or") {
	//		cond = cond.Or(item.Key, utils.ConvertUint8ToString(item.Val))
	//	} else if (item.Type == "OrNot") {
	//		cond = cond.OrNot(item.Key, utils.ConvertUint8ToString(item.Val))
	//	}
	//}
	cond := utils.ConvertCond(args.CondList)

	var err error
	reply.CurrentSize, err = o.QueryTable("user").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.Total, err = o.QueryTable("user").SetCond(cond).Count()
	return err
}
