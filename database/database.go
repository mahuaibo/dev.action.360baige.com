package database

import (
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
	"dev.model.360baige.com/models/application"
	"dev.model.360baige.com/models/attendance"
	"dev.model.360baige.com/models/card"
	"dev.model.360baige.com/models/company"
	"dev.model.360baige.com/models/logger"
	"dev.model.360baige.com/models/machine"
	"dev.model.360baige.com/models/order"
	"dev.model.360baige.com/models/personnel"
	"dev.model.360baige.com/models/user"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_user?charset=utf8", 30)

	orm.RegisterDataBase("account", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_account?charset=utf8", 30)
	orm.RegisterModel(&account.Account{}, &account.AccountItem{}, &account.Transaction{})

	orm.RegisterDataBase("application", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_application?charset=utf8", 30)
	orm.RegisterModel(&application.Application{}, &application.ApplicationTpl{})

	orm.RegisterDataBase("attendance", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_attendance?charset=utf8", 30)
	orm.RegisterModel(&attendance.AttendanceGroup{}, &attendance.AttendanceRecord{}, &attendance.AttendanceSetup{}, &attendance.AttendanceShift{}, &attendance.AttendanceShiftItem{}, &attendance.AttendanceShiftRecord{})

	orm.RegisterDataBase("card", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_card?charset=utf8", 30)
	orm.RegisterModel(&card.Card{})

	orm.RegisterDataBase("company", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_company?charset=utf8", 30)
	orm.RegisterModel(&company.Company{})

	orm.RegisterDataBase("logger", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_logger?charset=utf8", 30)
	orm.RegisterModel(&logger.Logger{})

	orm.RegisterDataBase("machine", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_machine?charset=utf8", 30)
	orm.RegisterModel(&machine.Machine{})

	orm.RegisterDataBase("order", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_order?charset=utf8", 30)
	orm.RegisterModel(&order.Order{})

	orm.RegisterDataBase("personnel", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_personnel?charset=utf8", 30)
	orm.RegisterModel(&personnel.Person{}, &personnel.PersonRelation{}, &personnel.Structure{}, &personnel.PersonStructure{})

	orm.RegisterDataBase("user", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_user?charset=utf8", 30)
	orm.RegisterModel(&user.User{}, &user.UserPosition{})
}
