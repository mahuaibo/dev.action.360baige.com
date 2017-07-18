package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
	"dev.model.360baige.com/action"
	"dev.action.360baige.com/utils"
	"time"
	"encoding/json"
)

type AttendanceSetupAction struct {
}

// 1
func (*AttendanceSetupAction) Add(args *attendance.AttendanceSetup, reply *attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*AttendanceSetupAction) AddMultiple(args []*attendance.AttendanceSetup, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*AttendanceSetupAction) FindById(args *attendance.AttendanceSetup, reply *attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*AttendanceSetupAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("attendance_setup").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*AttendanceSetupAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("attendance_setup").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*AttendanceSetupAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("attendance_setup").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*AttendanceSetupAction) FindByCond(args *action.FindByCond, reply *attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("attendance_setup").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*AttendanceSetupAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("attendance_setup").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*AttendanceSetupAction) ListByCond(args *action.ListByCond, reply *[]attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("attendance_setup").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*AttendanceSetupAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
	o := orm.NewOrm()
	o.Using("attendance")

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
	var replyList []attendance.AttendanceSetup
	reply.CurrentSize, err = o.QueryTable("attendance_setup").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("attendance_setup").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*AttendanceSetupAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("attendance_setup").SetCond(cond).Count()
	reply.Value = num
	return err
}
