package personnel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/personnel"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	. "dev.action.360baige.com/database"
	"time"
	"encoding/json"
)

type PersonRelationAction struct {
}

// 1
func (*PersonRelationAction) Add(args *personnel.PersonRelation, reply *personnel.PersonRelation) error {
	o := GetOrmer(DB_personnel)
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*PersonRelationAction) AddMultiple(args []*personnel.PersonRelation, reply *action.Num) error {
	o := GetOrmer(DB_personnel)
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*PersonRelationAction) FindById(args *personnel.PersonRelation, reply *personnel.PersonRelation) error {
	o := GetOrmer(DB_personnel)
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*PersonRelationAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := GetOrmer(DB_personnel)

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("person_relation").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*PersonRelationAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_personnel)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("person_relation").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*PersonRelationAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := GetOrmer(DB_personnel)

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("person_relation").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*PersonRelationAction) FindByCond(args *action.FindByCond, reply *personnel.PersonRelation) error {
	o := GetOrmer(DB_personnel)

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("person_relation").SetCond(cond).One(reply, args.Fileds...)
	if err == orm.ErrNoRows {
		return nil
	} else {
		return err
	}
}

// 8
func (*PersonRelationAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := GetOrmer(DB_personnel)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("person_relation").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*PersonRelationAction) ListByCond(args *action.ListByCond, reply *[]personnel.PersonRelation) error {
	o := GetOrmer(DB_personnel)

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("person_relation").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*PersonRelationAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := GetOrmer(DB_personnel)

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
	var replyList []personnel.PersonRelation
	reply.CurrentSize, err = o.QueryTable("person_relation").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("person_relation").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*PersonRelationAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := GetOrmer(DB_personnel)

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("person_relation").SetCond(cond).Count()
	reply.Value = num
	return err
}
