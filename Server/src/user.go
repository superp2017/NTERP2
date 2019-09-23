/*
@Time : 2019/9/20 10:15
@Author : mp
@File : user
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/superp2017/golibs/Http"
	"github.com/superp2017/golibs/Logger"
)

type User struct {
	UID        string //用户id
	UName      string //用户姓名
	Sex        string //性别
	Age        int    //年龄
	Cell       string //联系方式
	Department string //部门
	Salary     int    //薪水
	LoginName  string //账号
	LoginCode  string //密码
	Author     int    //用户权限
	InTime     string //入职时间
	OutTime    string //离职时间
	CreatTime  string //创建时间
	CreatStamp int64  //创建的时间戳
	LastTime   int64  //最后更新时间
}

//新增一个员工
func NewUser(session *Http.Session) {
	st := &User{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.UName == "" {
		session.Forward("1", "NewUser failed,UName is empty\n", nil)
		return
	}
	if st.LoginName == "" {
		session.Forward("1", "NewUser failed,Account is empty\n", nil)
		return
	}
	if st.LoginCode == "" {
		session.Forward("1", "NewUser failed,Code is empty\n", nil)
		return
	}

	st.CreatTime = CurTime()
	st.CreatStamp = CurStamp()
	st.LastTime = CurStamp()

	cmd := "insert  user(uname,sex,age,cell,department,salary,loginname,logincode,author,intime,outtime,creattime,creatstamp,lasttime)value(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)
	if err != nil {
		errStr := "NewUser Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.UName, st.Sex, st.Age, st.Cell, st.Department, st.Salary, st.LoginName, st.LoginCode, st.Author, st.InTime, st.OutTime, st.CreatTime, st.CreatStamp, st.LastTime); err != nil {
		errStr := "NewUser doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	session.Forward("0", "success", st)
}

//修改员工信息
func ModUser(session *Http.Session) {
	type Para struct {
		UID        string //用户id
		UName      string //用户姓名
		Sex        string //性别
		Cell       string //联系方式
		Department string //部门
		Author     int    //用户权限
		LoginName  string //账号
		LoginCode  string //密码
		Age        int    //年龄
		Salary     int    //薪水
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.UID == "" || st.LoginName == "" || st.LoginCode == "" || st.UName == "" {
		errStr := fmt.Sprintf("ModUser faild,UID = %s,LoginName = %s,LoginCode = %s,UName = %s\n", st.UID, st.LoginName, st.LoginCode, st.UName)
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	stmt, err := m_db.Prepare("UPDATE user SET uname=?,sex=?,cell=?,department=?,author=?,loginname=?,logincode=?,age=?,salary=? lasttime=? WHERE uid=?")
	if err != nil {
		errStr := "ModUser Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.UName, st.Sex, st.Cell, st.Department, st.Author, st.LoginName, st.LoginCode, st.Age, st.Salary, CurStamp(), st.UID); err != nil {
		errStr := "ModUser doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	session.Forward("0", "success", nil)
}

//删除某一个员工
func DelUser(session *Http.Session) {
	type Para struct {
		UID string //用户id
	}

	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.UID == "" {
		session.Forward("1", "DelUser failed,UID is empty\n", nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM user WHERE uid=?")
	if err != nil {
		errStr := "DelUser Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.UID); err != nil {
		errStr := "DelUser doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", nil)
}

// 获取所有的员工信息
func GetAllUser(session *Http.Session) {
	rows, err := m_db.Query("SELECT * FROM user ")
	if err != nil {
		errStr := "GetAllUser Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]*User, 0)
	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.UID, &u.UName, &u.Sex, &u.Age, &u.Cell, &u.Department, &u.Salary, &u.LoginName, &u.LoginCode, &u.Author, &u.InTime, &u.OutTime, &u.CreatTime, &u.CreatStamp, &u.LastTime)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}

		list = append(list, u)
	}
	session.Forward("0", "success", list)
}
