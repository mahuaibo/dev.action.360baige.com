package personnel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/personnel"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type StructureAction struct {
}

// 1
func (*StructureAction) Add(args *personnel.Account, reply *personnel.Structure) error {
	o := orm.NewOrm()
	o.Using("personnel")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*StructureAction) AddMultiple(args []*personnel.Structure, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("personnel")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*StructureAction) FindById(args *personnel.Structure, reply *personnel.Structure) error {
	o := orm.NewOrm()
	o.Using("personnel")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*StructureAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("structure").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*StructureAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("structure").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*StructureAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("structure").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*StructureAction) FindByCond(args *action.FindByCond, reply *personnel.Structure) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("structure").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*StructureAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("structure").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*StructureAction) ListByCond(args *action.ListByCond, reply *[]personnel.Structure) error {
	o := orm.NewOrm()
	o.Using("personnel")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("structure").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*StructureAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("personnel")

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
	var replyList []personnel.Structure
	reply.CurrentSize, err = o.QueryTable("structure").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("structure").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}
