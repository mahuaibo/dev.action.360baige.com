package order

import (
	_ "github.com/go-sql-driver/mysql"
)

//type OrderAction struct {
//}

//// 1
//func (*OrderAction) Add(args *order.Order, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order")
//	reply, err := o.Insert(args)
//	return err
//}
//
//// 2
//func (*OrderAction) AddMultiple(args []*order.Order, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order") //查询数据库
//	reply, err := o.InsertMulti(len(args), args)
//	return err
//}
//
//// 3
//func (*OrderAction) FindById(args *order.Order, reply *order.Order) error {
//	o := orm.NewOrm()
//	o.Using("order")
//	reply.Id = args.Id
//	err := o.Read(reply)
//	return err
//}
//
//// 4
//func (*OrderAction) UpdateByCond(args *action.UpdateByCond, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order")
//
//	var values orm.Params
//	cond := orm.NewCondition()
//	for _, item := range args.CondList {
//		if (item.Type == "And") {
//			cond = cond.And(item.Key, item.Val)
//		} else if (item.Type == "AndNot") {
//			cond = cond.AndNot(item.Key, item.Val)
//		} else if (item.Type == "Or") {
//			cond = cond.Or(item.Key, item.Val)
//		} else if (item.Type == "OrNot") {
//			cond = cond.OrNot(item.Key, item.Val)
//		}
//	}
//
//	for _, item := range args.UpdateList {
//		values[item.Key] = item.Val
//	}
//
//	reply, err := o.QueryTable("order").SetCond(cond).Update(values)
//	return err
//}
//
//// 5
//func (*OrderAction) DeleteById(args []int64, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order")
//
//	var values orm.Params
//	cond := orm.NewCondition()
//	cond = cond.And("id__in", args)
//
//	values["status"] = -1
//
//	reply, err := o.QueryTable("order").SetCond(cond).Update(values)
//	return err
//}
//
//// 6
//func (*OrderAction) UpdateById(args *action.UpdateByIdCond, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order")
//
//	var values orm.Params
//	cond := orm.NewCondition()
//	cond = cond.And("id__in", args.Id)
//
//	for _, item := range args.UpdateList {
//		values[item.Key] = item.Val
//	}
//
//	reply, err := o.QueryTable("order").SetCond(cond).Update(values)
//	return err
//}
//
//// 7
//func (*OrderAction) FindByCond(args *action.FindByCond, reply *order.Order) error {
//	o := orm.NewOrm()
//	o.Using("order")
//
//	cond := orm.NewCondition()
//
//	for _, item := range args.CondList {
//		if (item.Type == "And") {
//			cond = cond.And(item.Key, item.Val)
//		} else if (item.Type == "AndNot") {
//			cond = cond.AndNot(item.Key, item.Val)
//		} else if (item.Type == "Or") {
//			cond = cond.Or(item.Key, item.Val)
//		} else if (item.Type == "OrNot") {
//			cond = cond.OrNot(item.Key, item.Val)
//		}
//	}
//
//	err := o.QueryTable("order").SetCond(cond).One(&reply, args.Fileds...)
//	return err
//}
//
//// 8
//func (*OrderAction) DeleteByCond(args *action.DeleteByCond, reply int64) error {
//	o := orm.NewOrm()
//	o.Using("order")
//
//	var values orm.Params
//	cond := orm.NewCondition()
//	for _, item := range args.CondList {
//		if (item.Type == "And") {
//			cond = cond.And(item.Key, item.Val)
//		} else if (item.Type == "AndNot") {
//			cond = cond.AndNot(item.Key, item.Val)
//		} else if (item.Type == "Or") {
//			cond = cond.Or(item.Key, item.Val)
//		} else if (item.Type == "OrNot") {
//			cond = cond.OrNot(item.Key, item.Val)
//		}
//	}
//
//	values["status"] = -1
//
//	reply, err := o.QueryTable("order").SetCond(cond).Update(values)
//	return err
//}
//
//// 9
//func (*OrderAction) ListByCond(args *action.ListByCond, reply *[]order.Order) error {
//	o := orm.NewOrm()
//	o.Using("order")
//	cond := orm.NewCondition()
//	for _, item := range args.CondList {
//		if (item.Type == "And") {
//			cond = cond.And(item.Key, item.Val)
//		} else if (item.Type == "AndNot") {
//			cond = cond.AndNot(item.Key, item.Val)
//		} else if (item.Type == "Or") {
//			cond = cond.Or(item.Key, item.Val)
//		} else if (item.Type == "OrNot") {
//			cond = cond.OrNot(item.Key, item.Val)
//		}
//	}
//
//	if args.PageSize == 0 {
//		args.PageSize = -1
//	}
//	_, err := o.QueryTable("order").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply, args.Cols...)
//	return err
//}
//
//// 10
//func (*OrderAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
//	o := orm.NewOrm()
//	o.Using("order")
//	cond := orm.NewCondition()
//
//	for _, item := range args.CondList {
//		if (item.Type == "And") {
//			cond = cond.And(item.Key, item.Val)
//		} else if (item.Type == "AndNot") {
//			cond = cond.AndNot(item.Key, item.Val)
//		} else if (item.Type == "Or") {
//			cond = cond.Or(item.Key, item.Val)
//		} else if (item.Type == "OrNot") {
//			cond = cond.OrNot(item.Key, item.Val)
//		}
//	}
//	num, err := o.QueryTable("order").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
//	reply.CurrentSize = num
//	qs := o.QueryTable("order").SetCond(cond)
//	total, err := qs.Count()
//	reply.Total = total
//	return err
//}
