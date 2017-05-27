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
// TODO:


package main

import (
	"fmt"
	"os"
	"time"

	myInit		"github.com/my10c/nagios-plugins-go/initialize"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
	myMySQL		"github.com/my10c/nagios-plugins-go/mysql"
	myGlobal	"github.com/my10c/nagios-plugins-go/global"
)

const (
	table = "MONITOR"
	field = "timestamp"
)

var (
	cfgRequired = []string{"username", "password", "database", "hostname", "port"}
)

func wrongMode() {
	fmt.Printf("%s", myGlobal.MyInfo)
	fmt.Printf("Wrong mode, supported mode:\n")
	fmt.Printf("\t basic       : check select/insert/delete\n")
	fmt.Printf("\t slavestatus : check if slave is running\n")
	fmt.Printf("\t slavelag    : check slave lag, requires the configs: lagwarning and lagcritical.\n")
	fmt.Printf("\t process     : check process count, requires the configs: processwarning and processcritical.\n")
	fmt.Printf("\t dropcreate  : check drop and create tables, requires the config: tablename.\n")
	os.Exit(3)
}

func main() {
	cfgFile, checkMode := myInit.InitArgs(cfgRequired)
	cfgDict := myInit.InitConfig(cfgRequired, cfgFile)
	myInit.InitLog(cfgDict)
	myUtils.SignalHandler()
	dbCheck := myMySQL.New(cfgDict)
	data := time.Now().String()
	switch checkMode {
		case "basic":
			dbCheck.BasisCheck(table, field, data)
		case "slavestatus":
			dbCheck.SlaveStatusCheck()
		case "slavelag":
			dbCheck.SlaveLagCheck()
		case "process":
			dbCheck.ProcessStatusCheck()
		case "dropcreate":
			dbCheck.DropCreateCheck()
		default:
			wrongMode()
	}
	os.Exit(0)
}
