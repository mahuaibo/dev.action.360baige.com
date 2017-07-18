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

type AttendanceShiftRecordAction struct {
}

// 1
func (*AttendanceShiftRecordAction) Add(args *attendance.AttendanceShiftRecord, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	reply.Id = id
	return err
}

// 2
func (*AttendanceShiftRecordAction) AddMultiple(args []*attendance.AttendanceShiftRecord, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")
	num, err := o.InsertMulti(len(args), args)
	reply.Value = num
	return err
}

// 3
func (*AttendanceShiftRecordAction) FindById(args *attendance.AttendanceShiftRecord, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 4
func (*AttendanceShiftRecordAction) UpdateByCond(args *action.UpdateByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)
	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("attendance_shift_record").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 5
func (*AttendanceShiftRecordAction) DeleteById(args *action.DeleteByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Value)

	num, err := o.QueryTable("attendance_shift_record").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 6
func (*AttendanceShiftRecordAction) UpdateById(args *action.UpdateByIdCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := orm.NewCondition()
	cond = cond.And("id__in", args.Id)

	values := utils.ConvertValues(args.UpdateList)

	num, err := o.QueryTable("attendance_shift_record").SetCond(cond).Update(values)
	reply.Value = num
	return err
}

// 7
func (*AttendanceShiftRecordAction) FindByCond(args *action.FindByCond, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	err := o.QueryTable("attendance_shift_record").SetCond(cond).One(reply, args.Fileds...)
	return err
}

// 8
func (*AttendanceShiftRecordAction) DeleteByCond(args *action.DeleteByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("attendance_shift_record").SetCond(cond).Update(orm.Params{"update_time": time.Now().UnixNano() / 1e6, "status": -1})
	reply.Value = num
	return err
}

// 9
func (*AttendanceShiftRecordAction) ListByCond(args *action.ListByCond, reply *[]attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	if args.PageSize == 0 {
		args.PageSize = -1
	}
	_, err := o.QueryTable("attendance_shift_record").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(reply, args.Cols...)
	return err
}

// 10
func (*AttendanceShiftRecordAction) PageByCond(args *action.PageByCond, reply *action.PageByCond) error {
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
	var replyList []attendance.AttendanceShiftRecord
	reply.CurrentSize, err = o.QueryTable("attendance_shift_record").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&replyList, args.Cols...)
	reply.Total, err = o.QueryTable("attendance_shift_record").SetCond(cond).Count()
	reply.Json, _ = json.Marshal(replyList)
	return err
}

// 11
func (*AttendanceShiftRecordAction) CountByCond(args *action.CountByCond, reply *action.Num) error {
	o := orm.NewOrm()
	o.Using("attendance")

	cond := utils.ConvertCond(args.CondList)

	num, err := o.QueryTable("attendance_shift_record").SetCond(cond).Count()
	reply.Value = num
	return err
}
