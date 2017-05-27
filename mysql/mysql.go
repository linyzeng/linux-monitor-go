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
	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUE ('%s')", table, field, data)
	stmtWrite, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtWrite.Close()
	_, err = stmtWrite.Exec()
	return err
}

func (db *dbMysql) CheckRead(table string, field string, data string) error {
	// Prepare statement for reading the data
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE %s = '%s'", table, field, data)
	stmtRead, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtRead.Close()
	_, err = stmtRead.Exec()
	return err
}

func (db *dbMysql) CheckDelete(table string, field string, data string) error {
	// Prepare statement for delete the data
	stmt := fmt.Sprintf("DELETE FROM %s WHERE %s = '%s'", table, field, data)
	stmtDelete, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtDelete.Close()
	_, err = stmtDelete.Exec()
	return err
}

func (db *dbMysql) CreateTable(table string) error {
	// Prepare statement to create a table
	stmt := fmt.Sprintf("CREATE TABLE %s (timestamp varchar(128))", table)
	stmtCreate, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtCreate.Close()
	_, err = stmtCreate.Exec()
	return err
}

func (db *dbMysql) DropTable(table string) error {
	// Prepare statement to drop a table
	stmt := fmt.Sprintf("DROP TABLE %s", table)
	stmtCreate, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtCreate.Close()
	_, err = stmtCreate.Exec()
	return err
}

func (db *dbMysql) BasisCheck(table string, field string,  data string)  (int, error) {
	if err := db.CheckWrite(table, field, data); err != nil {
		db.Close()
		return 2, err
	}
	if err := db.CheckRead(table, field, data); err != nil {
		db.Close()
		return 2, err
	}
	if err := db.CheckDelete(table, field, data); err != nil {
		db.Close()
		return 2, err
	}
	return 0, nil
}

func (db *dbMysql) SlaveStatusCheck()  (int, error) {
	stmt := fmt.Sprintf("SHOW SLAVE STATUS")
	stmtStatus, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtStatus.Close()
	rows, err := stmtStatus.Query()
	// make sure to close the query
	defer rows.Close()
	if err != nil {
		db.Close()
		return 3, err
	}
	fmt.Printf("\n %s \n\n", rows)
	return 0, nil
}

func (db *dbMysql) SlaveLagCheck(warning int, critical int)  (int, error) {
	stmt := fmt.Sprintf("SHOW SLAVE STATUS")
	stmtStatus, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtStatus.Close()
	rows, err := stmtStatus.Query()
	// make sure to close the query
	defer rows.Close()
	if err != nil {
		db.Close()
		return 3, err
	}
	return 0, nil
}

func (db *dbMysql) ProcessStatusCheck(warning int, critical int)  (int, error) {
	stmt := fmt.Sprintf("SHOW PROCESSLIST")
	stmtStatus, err := db.Prepare(stmt)
	myUtils.ExitIfError(err)
	// make sure to close the statement
	defer stmtStatus.Close()
	rows, err := stmtStatus.Query()
	// make sure to close the query
	defer rows.Close()
	if err != nil {
		db.Close()
		return 3, err
	}
	fmt.Printf("\n %s \n\n", rows)
	return 0, nil
}

func (db *dbMysql) DropCreateCheck(tablename string) (int, error) {
	if err := db.CreateTable(tablename); err != nil {
		db.Close()
		return 2, err
	}
	if err := db.DropTable(tablename); err != nil {
		db.Close()
		return 2, err
	}
	return 0, nil
}
