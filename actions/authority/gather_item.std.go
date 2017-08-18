package authority

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/authority"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	. "dev.action.360baige.com/database"
	"time"
	"encoding/json"
)

type GatherItemAction struct {
}

// 1
func (*GatherItemAction) Add(args *authority.GatherItem, reply *authority.GatherItem) error {
	o := GetOrmer(DB_authority)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*GatherItemAction) AddMultiple(args []*authority.GatherItem, reply *action.Num) error {
	o := GetOrmer(DB_authority)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*GatherItemAction) FindById(args *authority.GatherItem, reply *authority.GatherItem) error {
	o := GetOrmer(DB_authority)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*GatherItemAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("gather_item").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*GatherItemAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("gather_item").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*GatherItemAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("gather_item").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*GatherItemAction) FindByCond(args *action.FindByCond, reply *authority.GatherItem) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("gather_item").SetCond(cond).One(reply, args.Fileds...)
	if err == orm.ErrNoRows {
		return nil
	} else {
		return err
	}
}

// 8
func (*GatherItemAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("gather_item").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*GatherItemAction) ListByCond(args *action.ListByCond, reply *[]authority.GatherItem) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("gather_item").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*GatherItemAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := GetOrmer(DB_authority)

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
	var replyList []authority.GatherItem
	reply.CurrentSize, err = o.QueryTable("gather_item").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("gather_item").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*GatherItemAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("gather_item").SetCond(cond).Count()
	reply.Value = num
	return err
}
