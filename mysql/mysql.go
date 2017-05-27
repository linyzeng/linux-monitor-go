// Copyright (c) 2014 - 2017 badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//   * Redistributions of source code must retain the above copyright
//   notice, this list of conditions and the following disclaimer.
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//   * Neither the name of the <organization> nor the
//   names of its contributors may be used to endorse or promote products
//   derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Version		:	0.1
//
// Date			:	May 18, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Mar 3, 2014		LIS			First release
//	May 18, 2017	LIS			Convert from bash/python/perl to Go
//

package mysql

import (
	"fmt"
	myUtils	"github.com/my10c/nagios-plugins-go/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type dbMysql struct {
	*sqlx.DB
}

func New(mysqlCfg map[string]string) *dbMysql {
	// set the username and password
	mysql_user :=  mysqlCfg["username"] + ":" +  mysqlCfg["password"] + "@"
	// need to create the string tcp(fqdn:port) according to the docs
	mysql_host := "tcp(" + mysqlCfg["hostname"] + ":" + mysqlCfg["port"] + ")"
	// set the database to use
	mysql_db := "/" +  mysqlCfg["database"]
	// set the full authentication string
	auth_string :=  mysql_user + mysql_host + mysql_db + "?parseTime=true"
	db, err := sqlx.Open("mysql", auth_string)
	myUtils.ExitIfError(err)
	// check we can make the connection
	err = db.Ping()
	myUtils.ExitIfError(err)
	// make sure the connection get close
	return &dbMysql{db}
}

func (db *dbMysql) CheckWrite(table string, field string, data string) error {
	// Prepare statement for inserting data
	stmtWrite, err := db.Prepare("INSERT INTO ? (?) VALUES(?)")
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtWrite.Close()
	_, err = stmtWrite.Exec(table, field, data)
	return err
}

func (db *dbMysql) CheckRead(table string, data string) error {
	// Prepare statement for reading the data
	stmtRead, err := db.Prepare("SELECT * FROM ?")
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtRead.Close()
	_, err = stmtRead.Exec(table, data)
	return err
}

func (db *dbMysql) CheckDelete(table string, field string, data string) error {
	// Prepare statement for delete the data
	stmtDelete, err := db.Prepare("DELETE FROM ? WHERE ? = ?")
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtDelete.Close()
	_, err = stmtDelete.Exec(table, field, data)
	return err
}

func (db *dbMysql) BasisCheck(table string, field string,  data string) error{
	if err := db.CheckWrite(table, field, data); err != nil {
		db.Close()
		return err
	}

	if err := db.CheckRead(table, data); err != nil {
		db.Close()
		return err
	}

	if err := db.CheckDelete(table, field, data); err != nil {
		db.Close()
		return err
	}
	return nil
}

func (db *dbMysql) SlaveStatusCheck() {
	// TODO:
	fmt.Printf("SlaveStatusCheck not implemented yet\n")
	return
}

func (db *dbMysql) SlaveLagCheck() {
	// TODO: need threshold
	fmt.Printf("SlaveLagCheck not implemented yet\n")
	return
}

func (db *dbMysql) ProcessStatusCheck() {
	// TODO: need threshold
	fmt.Printf("ProcessStatusCheck not implemented yet\n")
	return
}

func (db *dbMysql) DropCreateCheck() {
	// TODO: need table name
	fmt.Printf("DropCreateCheck not implemented yet\n")
	return
}
