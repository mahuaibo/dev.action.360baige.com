package schoolfee

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/schoolfee"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	. "dev.action.360baige.com/database"
	"time"
	"encoding/json"
)

type RecordAction struct {
}

// 1
func (*RecordAction) Add(args *schoolfee.Record, reply *schoolfee.Record) error {
	o := GetOrmer(DB_schoolfee)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*RecordAction) AddMultiple(args []*schoolfee.Record, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*RecordAction) FindById(args *schoolfee.Record, reply *schoolfee.Record) error {
	o := GetOrmer(DB_schoolfee)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*RecordAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("record").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*RecordAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("record").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*RecordAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("record").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*RecordAction) FindByCond(args *action.FindByCond, reply *schoolfee.Record) error {
	o := GetOrmer(DB_schoolfee)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("record").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*RecordAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("record").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*RecordAction) ListByCond(args *action.ListByCond, reply *[]schoolfee.Record) error {
	o := GetOrmer(DB_schoolfee)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("record").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*RecordAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := GetOrmer(DB_schoolfee)

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
	var replyList []schoolfee.Record
	reply.CurrentSize, err = o.QueryTable("record").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("record").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*RecordAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_schoolfee)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("record").SetCond(cond).Count()
	reply.Value = num
	return err
}
