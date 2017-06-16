package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
)

type AttendanceRecordAction struct {
}

// 新增
func (*AttendanceRecordAction) Add(args *attendance.AttendanceRecord, reply *attendance.AttendanceRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.EquipmentId = args.EquipmentId
		reply.UserId = args.UserId
		reply.Type = args.Type

	}
	return err
}

// 查询 by Id
func (*AttendanceRecordAction) FindById(args *attendance.AttendanceRecord, reply *attendance.AttendanceRecord) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AttendanceRecordAction) UpdateById(args *attendance.AttendanceRecord, reply *attendance.AttendanceRecord) error {
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
