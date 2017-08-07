package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/account"
	"dev.model.360baige.com/models/application"
	"dev.model.360baige.com/models/city"
	"dev.model.360baige.com/models/company"
	"dev.model.360baige.com/models/logger"
	"dev.model.360baige.com/models/message"
	"dev.model.360baige.com/models/order"
	"dev.model.360baige.com/models/user"
	"dev.model.360baige.com/models/schoolfee"

	//"dev.model.360baige.com/models/attendance"
	//"dev.model.360baige.com/models/card"
	//"dev.model.360baige.com/models/machine"
	"dev.model.360baige.com/models/personnel"
	"fmt"
)

const (
	DB_default_driver = "mysql"
	DB_default_ip     = "182.92.163.192"
	DB_default_port   = "3306"
	DB_default_user   = "demo2015"
	DB_default_pwd    = "baige.2016"

	DB_default     = "default"
	DB_user        = "db_user"
	DB_city        = "db_city"
	DB_company     = "db_company"
	DB_logger      = "db_logger"
	DB_account     = "db_account"
	DB_application = "db_application"
	DB_order       = "db_order"
	DB_schoolfee   = "db_schoolfee"
	DB_message     = "db_message"
	DB_personnel  = "db_personnel"

	// 暂未启用 TODO
	DB_attendance = "db_attendance"
	DB_card       = "db_card"
	DB_machine    = "db_machine"
)

func init() {
	fmt.Println("数据库注册开始")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	orm.RegisterDataBase(DB_default, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_user+"?charset=utf8", 30)

	orm.RegisterDataBase(DB_user, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_user+"?charset=utf8", 30)
	orm.RegisterModel(&user.User{}, &user.UserPosition{})

	orm.RegisterDataBase(DB_city, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_city+"?charset=utf8", 30)
	orm.RegisterModel(&city.City{})

	orm.RegisterDataBase(DB_company, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_company+"?charset=utf8", 30)
	orm.RegisterModel(&company.Company{})

	orm.RegisterDataBase(DB_logger, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_logger+"?charset=utf8", 30)
	orm.RegisterModel(&logger.Logger{})

	orm.RegisterDataBase(DB_account, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_account+"?charset=utf8", 30)
	orm.RegisterModel(&account.Account{}, &account.AccountItem{}, &account.Transaction{})

	orm.RegisterDataBase(DB_application, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_application+"?charset=utf8", 30)
	orm.RegisterModel(&application.Application{}, &application.ApplicationTpl{})

	orm.RegisterDataBase(DB_order, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_order+"?charset=utf8", 30)
	orm.RegisterModel(&order.Order{})

	orm.RegisterDataBase(DB_schoolfee, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_schoolfee+"?charset=utf8", 30)
	orm.RegisterModel(&schoolfee.Project{}, &schoolfee.Record{})

	orm.RegisterDataBase(DB_message, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_message+"?charset=utf8", 30)
	orm.RegisterModel(&message.MessageTemp{}, &message.MessageTotal{}, &message.MessageSend{})

	// 暂未启用 TODO
	//orm.RegisterDataBase(DB_attendance, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_attendance+"?charset=utf8", 30)
	//orm.RegisterModel(&attendance.AttendanceGroup{}, &attendance.AttendanceRecord{}, &attendance.AttendanceSetup{}, &attendance.AttendanceShift{}, &attendance.AttendanceShiftItem{}, &attendance.AttendanceShiftRecord{})

	//orm.RegisterDataBase(DB_card, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_card+"?charset=utf8", 30)
	//orm.RegisterModel(&card.Card{})
	//
	//orm.RegisterDataBase(DB_machine, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_machine+"?charset=utf8", 30)
	//orm.RegisterModel(&machine.Machine{})

	orm.RegisterDataBase(DB_personnel, DB_default_driver, DB_default_user+":"+DB_default_pwd+"@tcp("+DB_default_ip+":"+DB_default_port+")/"+DB_personnel+"?charset=utf8", 30)
	orm.RegisterModel(&personnel.Person{}, &personnel.PersonRelation{}, &personnel.Structure{}, &personnel.PersonStructure{})

	fmt.Println("数据库注册完成")
}

func GetOrmer(dbName string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(dbName)
	return o
}
