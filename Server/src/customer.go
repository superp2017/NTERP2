/*
@Time : 2019/9/20 17:32
@Author : mp
@File : customer
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/superp2017/golibs/Http"
	"github.com/superp2017/golibs/Logger"
)

type Customer struct {
	CID             string //客户编号
	CName           string //客户公司名称
	Addr            string //客户公司地址
	Tel             string //公司电话
	ContactName     string //联系人
	ContactCell     string //联系人电话
	BankName        string //开户行
	BankNumber      string //银行卡号
	BankBranch      string //银行支行
	CertificatesNum string //税号
	Note            string //备注
	CreatTime       string //创建时间
	CreatStamp      int64  //创建的时间戳
	LastTime        int64  //最后更新时间
}

//新建一个客户
func NewCustomer(session *Http.Session) {
	st := &Customer{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.CName == "" {
		str := fmt.Sprintf("NewCustomer faild,Name = %s\n", st.CName)
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}

	st.CreatTime = CurTime()
	st.CreatStamp = CurStamp()
	st.LastTime = CurStamp()

	cmd := "insert  customer(cname,addr,tel,contactname,contactcell,bankname,banknumber,bankbranch,certificatesnum,note,creattime,creatstamp,lasttime)value(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m_db.Prepare(cmd)
	if err != nil {
		errStr := "NewCustomer Prepare failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.CName, st.Addr, st.Tel, st.ContactName, st.ContactCell, st.BankName, st.BankNumber, st.BankBranch, st.CertificatesNum, st.Note, st.CreatTime, st.CreatStamp, st.LastTime); err != nil {
		errStr := "NewCustomer doSqlExec failed" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", nil)
}

//修改客户信息
func ModCustomer(session *Http.Session) {
	type Para struct {
		CID             string //客户编号
		CName           string //客户公司名称
		Addr            string //客户公司地址
		Tel             string //公司电话
		ContactName     string //联系人
		ContactCell     string //联系人电话
		BankName        string //开户行
		BankNumber      string //银行卡号
		BankBranch      string //银行支行
		CertificatesNum string //税号
		Note            string //备注
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}
	if st.CName == "" {
		str := fmt.Sprintf("ModCustomer faild,CName = %s\n", st.CName)
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}

	stmt, err := m_db.Prepare("UPDATE customer SET cname=?,addr=?,tel=?,contactname=?,contactcell=?,bankname=?,banknumber=?,bankbranch=?,certificatesnum=?,note=? lasttime=? WHERE cid=?")
	if err != nil {
		errStr := "ModCustomer Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.CName, st.Addr, st.Tel, st.ContactName, st.ContactCell, st.BankName, st.BankNumber, st.BankBranch, st.CertificatesNum, st.Note, CurStamp(), st.CID); err != nil {
		errStr := "ModCustomer doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", nil)
}

//删除一个客户
func DelCustomer(session *Http.Session) {
	type Para struct {
		CID string
	}
	st := &Para{}
	if err := session.GetPara(st); err != nil {
		Logger.Error(err.Error())
		session.Forward("1", err.Error(), nil)
		return
	}

	if st.CID == "" {
		str := fmt.Sprintf("DelCustomer CID is empty\n")
		Logger.Error(str)
		session.Forward("1", str, nil)
		return
	}

	stmt, err := m_db.Prepare("DELETE  FROM customer WHERE cid=?")
	if err != nil {
		errStr := "DelCustomer Prepare:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	if err := doSqlExec(stmt, st.CID); err != nil {
		errStr := "DelCustomer doSqlExec:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}

	session.Forward("0", "success", nil)
}

//获取所有客户列表
func GetAllCustomer(session *Http.Session) {
	rows, err := m_db.Query("SELECT * FROM customer ")
	if err != nil {
		errStr := "GetAllCustomer Query:" + err.Error()
		Logger.Error(errStr)
		session.Forward("1", errStr, nil)
		return
	}
	defer rows.Close()

	list := make([]*Customer, 0)

	for rows.Next() {
		u := &Customer{}
		err := rows.Scan(&u.CID, &u.CName, &u.Addr, &u.Tel, &u.ContactName, &u.ContactCell, &u.BankName, &u.BankNumber, &u.BankBranch, &u.CertificatesNum, &u.Note, &u.CreatTime, &u.CreatStamp, &u.LastTime)
		if err != nil {
			Logger.Error(err.Error())
			continue
		}

		list = append(list, u)
	}
	session.Forward("0", "GetAllCustomer success", list)
}
