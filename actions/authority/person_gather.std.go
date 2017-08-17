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

type PersonGatherAction struct {
}

// 1
func (*PersonGatherAction) Add(args *authority.PersonGather, reply *authority.PersonGather) error {
	o := GetOrmer(DB_authority)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*PersonGatherAction) AddMultiple(args []*authority.PersonGather, reply *action.Num) error {
	o := GetOrmer(DB_authority)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*PersonGatherAction) FindById(args *authority.PersonGather, reply *authority.PersonGather) error {
	o := GetOrmer(DB_authority)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*PersonGatherAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("person_gather").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*PersonGatherAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("person_gather").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*PersonGatherAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("person_gather").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*PersonGatherAction) FindByCond(args *action.FindByCond, reply *authority.PersonGather) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("person_gather").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*PersonGatherAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("person_gather").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*PersonGatherAction) ListByCond(args *action.ListByCond, reply *[]authority.PersonGather) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("person_gather").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*PersonGatherAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
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
	var replyList []authority.PersonGather
	reply.CurrentSize, err = o.QueryTable("person_gather").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("person_gather").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*PersonGatherAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_authority)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("person_gather").SetCond(cond).Count()
	reply.Value = num
	return err
}
