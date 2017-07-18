package city

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/city"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type CityAction struct {
}

// 1
func (*CityAction) Add(args *city.City, reply *city.City) error {
	o := orm.NewOrm()
	o.Using("city")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*CityAction) AddMultiple(args []*city.City, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*CityAction) FindById(args *city.City, reply *city.City) error {
	o := orm.NewOrm()
	o.Using("city")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*CityAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("city").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*CityAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("city").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*CityAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("city").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*CityAction) FindByCond(args *action.FindByCond, reply *city.City) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("city").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*CityAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("city").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*CityAction) ListByCond(args *action.ListByCond, reply *[]city.City) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("city").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*CityAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("city")

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
	var replyList []city.City
	reply.CurrentSize, err = o.QueryTable("city").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("city").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*CityAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("city")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("city").SetCond(cond).Count()
	reply.Value = num
	return err
}
