package account

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/http/window"
	"strconv"
	"strings"
)

//统计收入、支出数量
func (*AccountItemAction) AccountItemStatistics(args *window.AccountItemStatisticsCond, reply *window.AccountItemStatisticsCond) error {
	o := orm.NewOrm()
	o.Using("account")
	sql := "select SUM(amount) count FROM account_item WHERE  account_id = '" + strconv.FormatInt(args.AccountId, 10) + "' "
	if args.StartTime > 0 && args.EndTime > 0 {
		sql = sql + " and create_time>='" + strconv.FormatInt(args.StartTime, 10) + "' and create_time<'" + strconv.FormatInt(args.EndTime, 10) + "' "
	}
	var result1 []orm.Params
	num1, err := o.Raw(sql + " and amount<0 ").Values(&result1)
	if err == nil && num1 > 0 {
		if result1[0]["count"] == nil {

		} else {
			f, _ := strconv.ParseFloat(strings.Replace(result1[0]["count"].(string), "-", "", 1), 64)
			reply.Income = f
		}

	}
	var result2 []orm.Params
	num2, err := o.Raw(sql + " and amount>0 ").Values(&result2)
	if err == nil && num2 > 0 {
		if result2[0]["count"] == nil {

		} else {
			f, _ := strconv.ParseFloat(result2[0]["count"].(string), 64)
			reply.Pay = f
		}

	}
	return err
}

//分页list
func (*AccountItemAction) PageBy(args *window.AccountItemListPaginator, reply *window.AccountItemListPaginator) error {
	o := orm.NewOrm()
	o.Using("account")
	cond := orm.NewCondition()
	for _, c := range args.Cond {
		if (c.Type == "And") {
			cond = cond.And(c.Exprs, c.Args)
		} else if (c.Type == "AndNot") {
			cond = cond.AndNot(c.Exprs, c.Args)
		} else if (c.Type == "Or") {
			cond = cond.Or(c.Exprs, c.Args)
		} else if (c.Type == "OrNot") {
			cond = cond.OrNot(c.Exprs, c.Args)
		}
	}
	num, err := o.QueryTable("account_item").SetCond(cond).OrderBy(args.OrderBy...).Limit(args.PageSize, (args.Current-1)*args.PageSize).All(&reply.List, args.Cols...)
	reply.CurrentSize = num
	qs := o.QueryTable("account_item").SetCond(cond)
	total, err := qs.Count()
	reply.Total = total
	return err
}
