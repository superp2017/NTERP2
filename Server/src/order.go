/*
@Time : 2019/9/23 15:45
@Author : mp
@File : order
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/superp2017/golibs/Http"
	"github.com/superp2017/golibs/Logger"
)

//order 状态
const (
	Status_New          string = "Status_New"         //新建订单
	Status_Part_Produce string = "Status_PartProduce" //部分生产完成
	Status_Produce      string = "Status_Produce"     //全部生产完成
	Status_Part_Success string = "Status_PartSuccess" //部分出库
	Status_Success      string = "Status_Success"     //订单完成(出库)
	Status_Part_Part    string = "Status_Part_Part"   // 订单部分生产部分出库
	Status_Cancle       string = "Status_Cancle"      //订单取消
	//Status_Del          string = "Status_Del"         //订单删除
	Status_All string = "Status_All" //所有订单
)

//order 状态
const (
	Status_New          string = "Status_New"         //新建订单
	Status_Part_Produce string = "Status_PartProduce" //部分生产完成
	Status_Produce      string = "Status_Produce"     //全部生产完成
	Status_Part_Success string = "Status_PartSuccess" //部分出库
	Status_Success      string = "Status_Success"     //订单完成(出库)
	Status_Part_Part    string = "Status_Part_Part"   // 订单部分生产部分出库
	Status_Cancle       string = "Status_Cancle"      //订单取消
	//Status_Del          string = "Status_Del"         //订单删除
	Status_All string = "Status_All" //所有订单
)

type MaterialInfo struct {
	MaterielID      string //材料id
	MaterielDes     string //材料描述
	Plating         string //镀种
	Friction        string //摩擦系数
	Thickness       string //厚度
	Salt            string //盐度
	ComponentSolid  string //零件固号
	ComponentFormat string //零件规格
	Unit            string //单位
}

type RecordInfo struct {
	Flow  []OderFlow    //订单流程
	Print []PrintDetail //打印记录
}

type Order struct {
	OrderID   string //订单id
	OrderType string //订单类型（普通订单、试样订单、返工订单）

	UID      string //用户id
	UserName string //用户姓名

	MaterialInfo //物料基本信息
	RecordInfo   //订单各种打印记录

	CustomID    string //客户ID
	CustomName  string //客户姓名
	CustomBatch string //客户批次
	CustomNote  string //客户备注

	TotleMoney float64 //总价
	Money      float64 //单价

	OrderNum   float64 //订单数量
	ProduceNum float64 //生产完成数量
	SuccessNum float64 //出库数量

	ProduceTime string //出货时间
	SuccessTime string //完成时间
	CreatTime   string //创建时间
	CreatStamp  int64  //创建的时间戳
	LastTime    int64  //最后更新时间
}

func NewOrder(session *Http.Session) {
	st := &Order{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.UserName == "" || st.UID == "" {
		str := fmt.Sprintf("NewOrder failed,UserName = %s,UID = %s\n", st.UserName, st.UID)
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}
	if st.MaterielDes == "" {
		Logger.Error("NewOrder MaterielDes is empty\n!")
		session.Forward("1", "NewOrder MaterielDes is empty\n!", nil)
		return
	}
	cmd := "insert  order(ordertype,uid,username,materielid,customid,customname,custombatch,customnote,totlemoney,money,ordernum,producenum,sucessnum,producetime,sucesstime,creattime,creatstamp,lasttime) value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)

	if err != nil {
		errStr := "NewPlating Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.OrderType, st.UID, st.UserName, st.MaterielID, st.CustomID, st.CustomName, st.CustomBatch, st.CustomNote, st.TotleMoney, st.Money, st.OrderNum, st.ProduceNum, st.SuccessNum, st.ProduceTime, st.SuccessTime, st.CreatTime, st.CreatStamp, st.LastTime); err != nil {
		errStr := "NewPlating doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

}
