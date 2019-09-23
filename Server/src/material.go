/*
@Time : 2019/9/23 10:07
@Author : mp
@File : material
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/superp2017/golibs/Http"
	"github.com/superp2017/golibs/Logger"
)

type Material struct {
	MaterID         string  //物料编号
	MaterDes        string  //物料描述
	CID             string  //客户ID
	CustomName      string  //客户姓名
	Plating         string  //镀种
	Friction        string  //摩擦系数
	Thickness       string  //厚度
	Salt            string  //盐度
	ComponentSolid  string  //零件固号
	ComponentFormat string  //零件规格
	Factory         string  //分厂名称
	ProductionLine  string  //产线名称
	Unit            string  //单位
	Money           float64 //未税单价
	CreatTime       string  //创建时间
	CreatStamp      int64   //创建的时间戳
	LastTime        int64   //最后更新时间
}

//新建物料
func NewMaterial(session *Http.Session) {
	st := &Material{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.CID == "" || st.CustomName == "" || st.MaterDes == "" || st.ComponentSolid == "" {
		str := fmt.Sprintf("NewMaterial MaterialInfo 部分参数为空：%v", st)
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}

	st.CreatTime = CurTime()
	st.CreatStamp = CurStamp()
	st.LastTime = CurStamp()

	cmd := "insert  material(materdes,cid,customname,plating,friction,thickness,salt,componentsolid,componentformat,factory,productionline,unit,money,creattime,creatstamp,lastTime)value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)
	if err != nil {
		errStr := "NewMaterial Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.MaterID, st.CID, st.CustomName, st.Plating, st.Friction, st.Thickness, st.Salt, st.ComponentSolid, st.ComponentFormat, st.Factory, st.ProductionLine, st.Unit, st.Money, st.CreatTime, st.CreatStamp, st.LastTime); err != nil {
		errStr := "NewCustomer doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", st)
}

//修改物料
func ModMaterial(session *Http.Session) {
	st := &Material{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.CID == "" || st.CustomName == "" || st.MaterDes == "" || st.ComponentSolid == "" {
		str := fmt.Sprintf("ModMaterial ,MaterialInfo 部分参数为空：%v", st)
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}
	st.LastTime = CurStamp()

	stmt, err := m_db.Prepare("UPDATE material SET materdes=?,cid=?,customname=?,plating=?,friction=?,thickness=?,salt=?,componentsolid=?,componentformat=?,factory=?,productionline=?,unit=? money=?,lastTime=? WHERE materid=?")
	if err != nil {
		errStr := "ModMaterial Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.MaterDes, st.CID, st.CustomName, st.Plating, st.Friction, st.Thickness, st.Salt, st.ComponentSolid, st.ComponentFormat, st.Factory, st.ProductionLine, st.Unit, st.Money, st.LastTime, st.MaterID); err != nil {
		errStr := "ModMaterial doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", st)
}

//删除物料
func DelMaterial(session *Http.Session) {
	type Para struct {
		MaterID string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.MaterID == "" {
		Logger.Error("DelMaterial failed,MaterID is empty!\n")
		session.Forward("1", "DelMaterial failed,MaterID is empty!\n", nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM material WHERE materid=?")
	if err != nil {
		errStr := "DelMaterial Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.MaterID); err != nil {
		errStr := "DelMaterial doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success\n", nil)
}

//获取所有材料
func GetAllMaterial(session *Http.Session) {

	rows, err := m_db.Query("SELECT * FROM material ")
	if err != nil {
		errStr := "GetAllMaterial Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]*Material, 0)

	for rows.Next() {
		u := &Material{}
		err := rows.Scan(&u.MaterID, &u.MaterDes, &u.CID, &u.CustomName, &u.Plating, &u.Friction, &u.Thickness, &u.Salt, &u.ComponentSolid, &u.ComponentFormat, &u.Factory, &u.ProductionLine, &u.Unit, &u.Money, &u.CreatTime, &u.CreatStamp, &u.LastTime)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, u)
	}

	session.Forward("0", "GetAllMaterial success", list)
}
