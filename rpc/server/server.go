package server

import (
	"github.com/astaxie/beego"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
	"github.com/astaxie/beego/logs"
	"dev.action.360baige.com/actions/user"
	"dev.action.360baige.com/actions/company"
	"dev.action.360baige.com/actions/logger"
	"dev.action.360baige.com/actions/city"
	"dev.action.360baige.com/actions/account"
	"dev.action.360baige.com/actions/order"
	"dev.action.360baige.com/actions/application"
	"dev.action.360baige.com/actions/message"
	"dev.action.360baige.com/actions/card"
	"dev.action.360baige.com/actions/machine"
	"dev.action.360baige.com/actions/personnel"
	"dev.action.360baige.com/actions/schoolfee"
	"dev.action.360baige.com/actions/website"
	"dev.action.360baige.com/actions/authority"
	//"dev.action.360baige.com/actions/attendance"
)

func init() {

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("server run start")
	// 服务
	services := map[string]map[string]interface{}{
		"User": map[string]interface{}{
			"User":         &user.UserAction{},
			"UserPosition": &user.UserPositionAction{},
		},
		"Company": map[string]interface{}{
			"Company": &company.CompanyAction{},
		},
		"Logger": map[string]interface{}{
			"Logger": &logger.LoggerAction{},
		},
		"City": map[string]interface{}{
			"City": &city.CityAction{},
		},
		"Account": map[string]interface{}{
			"Account":     &account.AccountAction{},
			"AccountItem": &account.AccountItemAction{},
			"Transaction": &account.TransactionAction{},
		},
		"Order": map[string]interface{}{
			"Order": &order.OrderAction{},
		},
		"Application": map[string]interface{}{
			"Application":    &application.ApplicationAction{},
			"ApplicationTpl": &application.ApplicationTplAction{},
		},
		"Message": map[string]interface{}{
			"MessageSend":  &message.MessageSendAction{},
			"MessageTemp":  &message.MessageTempAction{},
			"MessageTotal": &message.MessageTotalAction{},
		},
		"Schoolfee": map[string]interface{}{
			"Project": &schoolfee.ProjectAction{},
			"Record":  &schoolfee.RecordAction{},
		},
		"Personnel": map[string]interface{}{
			"Person":          &personnel.PersonAction{},
			"PersonRelation":  &personnel.PersonRelationAction{},
			"PersonStructure": &personnel.PersonStructureAction{},
			"Structure":       &personnel.StructureAction{},
		},
		"Card": map[string]interface{}{
			"Card": &card.CardAction{},
		},
		"Machine": map[string]interface{}{
			"Machine": &machine.MachineAction{},
		},
		"Website": map[string]interface{}{
			"Menu":     &website.MenuAction{},
			"Material": &website.MaterialAction{},
		},
		"Authority": map[string]interface{}{
			"Gather":       &authority.GatherAction{},
			"GatherItem":   &authority.GatherItemAction{},
			"PersonGather": &authority.PersonGatherAction{},
		},
	}
	go register(services)
	log.Debug("server run end")
}

/**
 * 注册服务
 */
func register(services map[string]map[string]interface{}) {

	etcdServerRegisterAddr := beego.AppConfig.String("RpcEtcdURL")
	serverRegisterAddr := beego.AppConfig.String("RpcServer")

	//var servs []string
	//for _, group_services := range services {
	//	for name, _ := range group_services {
	//		servs = append(servs, name)
	//	}
	//}
	var length = 0
	for _, group_services := range services {
		length += len(group_services)
	}
	servs := make([]string, length)

	//EtcdV3RegisterPlugin EtcdRegisterPlugin time.Minute
	rplugin := &plugin.EtcdV3RegisterPlugin{
		ServiceAddress:      "tcp@" + serverRegisterAddr,
		EtcdServers:         []string{etcdServerRegisterAddr },
		BasePath:            "/rpcx",
		Metrics:             metrics.NewRegistry(),
		Services:            servs,
		UpdateIntervalInSec: 60,
	}

	rplugin.Start()
	server := rpcx.NewServer()
	server.PluginContainer.Add(rplugin)
	server.PluginContainer.Add(plugin.NewMetricsPlugin())

	//注册 s
	for group, group_services := range services {
		for name, serv := range group_services {
			server.RegisterName(name, serv, "g="+group+"&weight=1&m=devops")
		}
	}
	//注册 e
	server.Serve("tcp", serverRegisterAddr)
}
