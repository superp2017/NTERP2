/*
@Time : 2019/9/20 17:04
@Author : mp
@File : list
@Software: GoLand
*/
package main

import (
	"github.com/superp2017/golibs/Http"
	"github.com/superp2017/golibs/Logger"
)

//添加一个部门
func NewDepartment(session *Http.Session) {
	type Para struct {
		Department string
	}

	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.Department == "" {
		Logger.Error("Department is empty\n")
		session.Forward("1", "Department is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("insert  departments(depart) value(?)")
	if err != nil {
		errStr := "NewDepartment Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Department); err != nil {
		errStr := "NewDepartment doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "NewDepartment success\n", nil)
}

//删除一个部门
func RemoveDepartment(session *Http.Session) {
	type Para struct {
		Department string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.Department == "" {
		Logger.Error("Department is empty\n")
		session.Forward("1", "Department is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM departments WHERE depart=?")
	if err != nil {
		errStr := "RemoveDepartment Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Department); err != nil {
		errStr := "RemoveDepartment doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "RemoveDepartment success\n", nil)
}

//获取所有的部门
func GetAllDepartment(session *Http.Session) {
	rows, err := m_db.Query("SELECT * FROM departments ")
	if err != nil {
		errStr := "GetAllDepartment Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]string, 0)
	for rows.Next() {
		de := ""
		err := rows.Scan(&de)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, de)
	}
	session.Forward("0", "GetAllDepartment success", list)
}

//添加一个单位
func NewUnit(session *Http.Session) {
	type Para struct {
		Unit string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.Unit == "" {
		Logger.Error("Unit is empty\n")
		session.Forward("1", "Unit is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("insert  units(unit) value(?)")
	if err != nil {
		errStr := "NewUnit Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Unit); err != nil {
		errStr := "NewUnit doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "NewUnit success\n", nil)
}

//删除一个单位
func RemoveUnit(session *Http.Session) {
	type Para struct {
		Unit string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.Unit == "" {
		Logger.Error("Unit is empty\n")
		session.Forward("1", "Unit is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM units WHERE unit=?")
	if err != nil {
		errStr := "RemoveUnit Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Unit); err != nil {
		errStr := "RemoveUnit doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "RemoveUnit success\n", nil)
}

//获取所有的单位
func GetAllUnit(session *Http.Session) {
	rows, err := m_db.Query("SELECT * FROM units ")
	if err != nil {
		errStr := "GetAllUnit Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]string, 0)
	for rows.Next() {
		de := ""
		err := rows.Scan(&de)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, de)
	}

	session.Forward("0", "GetAllUnit success", list)
}

//添加一个镀种
func NewPlating(session *Http.Session) {
	type Para struct {
		Plating string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.Plating == "" {
		Logger.Error("Plating is empty\n")
		session.Forward("1", "Plating is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("insert  platings(plating) value(?)")
	if err != nil {
		errStr := "NewPlating Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Plating); err != nil {
		errStr := "NewPlating doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "NewPlating success\n", nil)
}

//删除一个镀种
func RemovePlating(session *Http.Session) {
	type Para struct {
		Plating string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.Plating == "" {
		Logger.Error("Plating is empty\n")
		session.Forward("1", "Plating is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM platings WHERE plating=?")
	if err != nil {
		errStr := "RemovePlating Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.Plating); err != nil {
		errStr := "RemovePlating doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "RemovePlating success\n", nil)
}

//获取所有的镀种
func GetAllPlating(session *Http.Session) {
	rows, err := m_db.Query("SELECT * FROM platings ")
	if err != nil {
		errStr := "GetAllPlating Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]string, 0)
	for rows.Next() {
		de := ""
		err := rows.Scan(&de)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}
		list = append(list, de)
	}

	session.Forward("0", "GetAllPlating success", list)
}
