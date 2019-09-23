/*
@Time : 2019/9/23 16:10
@Author : mp
@File : orderrecord
@Software: GoLand
*/
package main

import "github.com/superp2017/golibs/Logger"

type OderFlow struct {
	OrderID    string //订单ID
	UserName   string //用户姓名
	OpreatTime string //操作时间
	Action     string //动作
	Status     string //状态
}
type PrintDetail struct {
	OrderID   string //订单ID
	UserID    string //用户ID
	UserName  string //用户姓名
	PrintDate string //打印时间
}

//追加一条订单流程
func appendOrderFlow(OrderID, UserName, OpreatTime, Action, Status string) error {
	cmd := "insert  orderflow(orderid,username,opreatime,opreat,cstatus)value(?,?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)
	if err != nil {
		errStr := "NewOrderFlow Prepare failed" + err.Error()
		Logger.Error(errStr)
		return err
	}

	if err := doSqlExec(stmt, OrderID, UserName, OpreatTime, Action, Status); err != nil {
		errStr := "NewOrderFlow doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		return err
	}
	return nil
}

//获取某个订单的所有流程
func getOrderFlow(OrderID string) []*OderFlow {
	list := make([]*OderFlow, 0)

	rows, err := m_db.Query("SELECT * FROM orderflow where orderid=?", OrderID)
	if err != nil {
		errStr := "getOrderFlow Query:" + err.Error()
		Logger.Error(errStr)
		return list
	}

	defer rows.Close()

	for rows.Next() {
		de := &OderFlow{}
		err := rows.Scan(&de.OrderID, &de.UserName, &de.OpreatTime, &de.Action, &de.Status)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, de)
	}
	return list
}

//追加一条订单打印记录
func appendOrderPrint(OrderID, UserID, UserName, PrintDate string) error {
	cmd := "insert  orderprint(orderid,userid,username,printdate)value(?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)
	if err != nil {
		errStr := "NewOrderFlow Prepare failed" + err.Error()
		Logger.Error(errStr)
		return err
	}

	if err := doSqlExec(stmt, OrderID, UserID, UserName, PrintDate); err != nil {
		errStr := "NewOrderFlow doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		return err
	}

	return nil
}

//获取某个订单所有打印记录
func getOrderPrint(OrderID string) []*PrintDetail {
	list := make([]*PrintDetail, 0)

	rows, err := m_db.Query("SELECT * FROM orderprint where orderid=?", OrderID)
	if err != nil {
		errStr := "getOrderPrint Query:" + err.Error()
		Logger.Error(errStr)
		return list
	}

	defer rows.Close()

	for rows.Next() {
		de := &PrintDetail{}
		err := rows.Scan(&de.OrderID, &de.UserID, &de.UserName, &de.PrintDate)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, de)
	}
	return list
}
