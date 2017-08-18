package website

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/website"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	. "dev.action.360baige.com/database"
	"time"
	"encoding/json"
)

type MaterialAction struct {
}

// 1
func (*MaterialAction) Add(args *website.Material, reply *website.Material) error {
	o := GetOrmer(DB_website)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*MaterialAction) AddMultiple(args []*website.Material, reply *action.Num) error {
	o := GetOrmer(DB_website)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*MaterialAction) FindById(args *website.Material, reply *website.Material) error {
	o := GetOrmer(DB_website)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*MaterialAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_website)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("material").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*MaterialAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_website)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("material").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*MaterialAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_website)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("material").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*MaterialAction) FindByCond(args *action.FindByCond, reply *website.Material) error {
	o := GetOrmer(DB_website)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("material").SetCond(cond).One(reply, args.Fileds...)
	if err == orm.ErrNoRows {
		return nil
	} else {
		return err
	}
}

// 8
func (*MaterialAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_website)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("material").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*MaterialAction) ListByCond(args *action.ListByCond, reply *[]website.Material) error {
	o := GetOrmer(DB_website)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("material").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*MaterialAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := GetOrmer(DB_website)

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
	var replyList []website.Material
	reply.CurrentSize, err = o.QueryTable("material").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("material").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*MaterialAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_website)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("material").SetCond(cond).Count()
	reply.Value = num
	return err
}
