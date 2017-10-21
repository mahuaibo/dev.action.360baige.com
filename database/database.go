package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
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
	"dev.model.360baige.com/models/website"
	"dev.model.360baige.com/models/authority"
	"dev.model.360baige.com/models/card"
	"dev.model.360baige.com/models/machine"
	"dev.model.360baige.com/models/personnel"

	//"dev.model.360baige.com/models/attendance"
	"fmt"
)

const (
	DB_default     = "db_user"
	DB_user        = "db_user"
	DB_city        = "db_city"
	DB_company     = "db_company"
	DB_logger      = "db_logger"
	DB_account     = "db_account"
	DB_application = "db_application"
	DB_order       = "db_order"
	DB_schoolfee   = "db_schoolfee"
	DB_message     = "db_message"
	DB_personnel   = "db_personnel"
	DB_card        = "db_card"
	DB_machine     = "db_machine"
	DB_authority   = "db_authority"
	DB_website     = "db_website"
	// 暂未启用 TODO
	DB_attendance = "db_attendance"
)

var (
	DB_default_driver = beego.AppConfig.String("DB_default_driver")
	DB_default_ip     = beego.AppConfig.String("DB_default_ip")
	DB_default_user   = beego.AppConfig.String("DB_default_user")
	DB_default_pwd    = beego.AppConfig.String("DB_default_pwd")
	DB_default_port   = beego.AppConfig.String("DB_default_port")
	dataSource        = DB_default_user + ":" + DB_default_pwd + "@tcp(" + DB_default_ip + ":" + DB_default_port + ")/"
)

func init() {
	fmt.Println(DB_default_driver, DB_default_ip)
	fmt.Println("数据库注册开始")

	registerDefault(DB_default, dataSource, true)

	registerDBM(DB_user, dataSource, &user.User{}, &user.UserPosition{})

	registerDBM(DB_city, dataSource, &city.City{})

	registerDBM(DB_company, dataSource, &company.Company{})

	registerDBM(DB_logger, dataSource, &logger.Logger{})

	registerDBM(DB_account, dataSource, &account.Account{}, &account.AccountItem{}, &account.Transaction{})

	registerDBM(DB_application, dataSource, &application.Application{}, &application.ApplicationTpl{})

	registerDBM(DB_order, dataSource, &order.Order{})

	registerDBM(DB_schoolfee, dataSource, &schoolfee.Project{}, &schoolfee.Record{})

	registerDBM(DB_message, dataSource, &message.MessageTemp{}, &message.MessageTotal{}, &message.MessageSend{})

	registerDBM(DB_card, dataSource, &card.Card{})

	registerDBM(DB_machine, dataSource, &machine.Machine{})

	registerDBM(DB_authority, dataSource, &authority.Gather{}, &authority.GatherItem{}, &authority.PersonGather{})

	registerDBM(DB_website, dataSource, &website.Menu{}, &website.Material{})

	registerDBM(DB_personnel, dataSource, &personnel.Person{}, &personnel.PersonRelation{}, &personnel.Structure{}, &personnel.PersonStructure{})

	//registerDBM(DB_attendance, dataSource,&attendance.AttendanceGroup{}, &attendance.AttendanceRecord{}, &attendance.AttendanceSetup{}, &attendance.AttendanceShift{}, &attendance.AttendanceShiftItem{}, &attendance.AttendanceShiftRecord{})
	fmt.Println("数据库注册完成")
}

func registerDefault(dbName, dataSource string, debug bool) {
	orm.Debug = debug
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", DB_default_driver, dataSource+dbName+"?charset=utf8", 30)
}

func registerDBM(dbName, dataSource string, models ...interface{}) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbName, DB_default_driver, dataSource+dbName+"?charset=utf8", 30)
	orm.RegisterModel(models...)
}

func GetOrmer(dbName string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(dbName)
	return o
}
