package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
)

type UserAction struct {
}

// 1
func (*UserAction) Add(args *user.User, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")
	reply, err := o.Insert(args)
	return err
}

// 2
func (*UserAction) AddMultiple(args []*user.User, reply int64) error {
	o := orm.NewOrm()
	o.Using("user") //查询数据库
	reply, err := o.InsertMulti(len(args), args)
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

type UpdateByCond struct {
	CondList   []CondValue   // 更新条件
	UpdateList []UpdateValue // 更新内容
}

type CondValue struct {
	Type string
	Key  string
	Val  interface{}
}

type UpdateValue struct {
	Key string
	Val interface{}
}

// 4
func (*UserAction) UpdateByCond(args *UpdateByCond, reply int64) error {
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

	reply, err := o.QueryTable("user").SetCond(cond).Update(values)
	return err
}

// 5
func (*UserAction) DeleteById(args []int64, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args)

	values["status"] = -1

	reply, err := o.QueryTable("user").SetCond(cond).Update(values)
	return err
}

type UpdateByIdCond struct {
	Id         []int64       // 更新条件
	UpdateList []UpdateValue // 更新内容
}

// 6
func (*UserAction) UpdateById(args *UpdateByIdCond, reply int64) error {
	o := orm.NewOrm()
	o.Using("user")

	var values orm.Params
	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	for _, item := range args.UpdateList {
		values[item.Key] = item.Val
	}

	reply, err := o.QueryTable("user").SetCond(cond).Update(values)
	return err
}

type FindByCond struct {
	CondList []CondValue // 更新条件
	Fileds   []string
}

// 7
func (*UserAction) FindByCond(args *FindByCond, reply *user.User) error {
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

	err := o.QueryTable("user").SetCond(cond).One(&reply, args.Fileds...)
	return err
}

type DeleteByCond struct {
	CondList []CondValue // 更新条件
}

// 8
func (*UserAction) DeleteByCond(args *DeleteByCond, reply int64) error {
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

	reply, err := o.QueryTable("user").SetCond(cond).Update(values)
	return err
}

type ListByCond struct {
	CondList []CondValue
	Cols     []string
	OrderBy  []string
	PageSize int64 //每页数量
}

// 9
func (*UserAction) ListByCond(args *ListByCond, reply *[]user.UserPosition) error {
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

type PageByCond struct {
	CondList    []CondValue
	Cols        []string
	OrderBy     []string
	List        []user.UserPosition
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}

// 10
func (*UserAction) PageByCond(args *PageByCond, reply *PageByCond) error {
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
	num, err := o.QueryTable("account_item").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.CurrentSize = num
	qs := o.QueryTable("account_item").SetCond(cond)
	total, err := qs.Count()
	reply.Total = total
	return err
}
