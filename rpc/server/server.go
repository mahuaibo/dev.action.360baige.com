package server

import (
	"github.com/astaxie/beego"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
	"github.com/astaxie/beego/logs"
	"dev.action.360baige.com/actions/account"
	"dev.action.360baige.com/actions/application"
	"dev.action.360baige.com/actions/attendance"
	"dev.action.360baige.com/actions/card"
	"dev.action.360baige.com/actions/company"
	"dev.action.360baige.com/actions/logger"
	"dev.action.360baige.com/actions/machine"
	"dev.action.360baige.com/actions/order"
	"dev.action.360baige.com/actions/user"
	"dev.action.360baige.com/actions/personnel"
	"time"
)

func init() {

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("server run start")
	// 服务
	services := map[string]interface{}{
		//
		"User":         &user.UserAction{},
		"UserPosition": &user.UserPositionAction{},
		////
		//"Account":     &account.AccountAction{},
		//"AccountItem": &account.AccountItemAction{},
		//"Transaction": &account.TransactionAction{},
		////
		//"Application":    &application.ApplicationAction{},
		//"ApplicationTpl": &application.ApplicationTplAction{},
		////
		//"AttendanceGroup":       &attendance.AttendanceGroupAction{},
		//"AttendanceRecord":      &attendance.AttendanceRecordAction{},
		//"AttendanceSetup":       &attendance.AttendanceSetupAction{},
		//"AttendanceShift":       &attendance.AttendanceShiftAction{},
		//"AttendanceShiftItem":   &attendance.AttendanceShiftItemAction{},
		//"AttendanceShiftRecord": &attendance.AttendanceShiftRecordAction{},
		////
		//"Card": &card.CardAction{},
		////
		//"Company": &company.CompanyAction{},
		////
		//"Logger": &logger.LoggerAction{},
		////
		//"Machine": &machine.MachineAction{},
		////
		//"Order": &order.OrderAction{},
		////
		//"Person":          &personnel.PersonAction{},
		//"PersonRelation":  &personnel.PersonRelationAction{},
		//"PersonStructure": &personnel.PersonStructureAction{},
		//"Structure":       &personnel.StructureAction{},
	}
	go register(services)
	log.Debug("server run end")
}

/**
 * 注册服务
 */
func register(services map[string]interface{}) {

	etcdServerRegisterAddr := beego.AppConfig.String("RpcEtcdURL")
	serverRegisterAddr := beego.AppConfig.String("RpcServer")
	servs := make([]string, len(services))
	for serName, _ := range services {
		servs = append(servs, serName)
	}
	rplugin := &plugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + serverRegisterAddr,
		EtcdServers:    []string{etcdServerRegisterAddr},
		BasePath:       "/rpcx",
		Metrics:        metrics.NewRegistry(),
		Services:       servs,
		UpdateInterval: time.Second,
	}
	rplugin.Start()
	server := rpcx.NewServer()
	server.PluginContainer.Add(rplugin)
	server.PluginContainer.Add(plugin.NewMetricsPlugin())

	//注册 s
	for name, serv := range services {
		server.RegisterName(name, serv, "weight=1&m=devops")
	}
	//注册 e
	server.Serve("tcp", serverRegisterAddr)
}
