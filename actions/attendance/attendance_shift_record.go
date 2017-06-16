package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
)

type AttendanceShiftRecordAction struct {
}

// 新增
func (*AttendanceShiftRecordAction) Add(args *attendance.AttendanceShiftRecord, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.DateTime = args.DateTime
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.ShiftItemId = args.ShiftItemId
		reply.AttendanceInRecordId = args.AttendanceInRecordId
		reply.AttendanceOutRecordId = args.AttendanceOutRecordId
		reply.AttendanceShiftItemId = args.AttendanceShiftItemId

	}
	return err
}

// 查询 by Id
func (*AttendanceShiftRecordAction) FindById(args *attendance.AttendanceShiftRecord, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AttendanceShiftRecordAction) UpdateById(args *attendance.AttendanceShiftRecord, reply *attendance.AttendanceShiftRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
