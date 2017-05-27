// Copyright (c) 2016 - 2017 Poynt Co
// All rights reserved.
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	* Proprietary and confidential *
//
// Author		:	Luc Suryo <luc@poynt.co>
//
// Version		:	0.1
//
// Date			:	Feb 4, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Feb 4, 2017		LIS			First release
//
// TODO:

package main

import (
	"fmt"
	"os"
	myInit		"github.com/my10c/nagios-plugins-go/initialize"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
	myGlobal	"github.com/my10c/nagios-plugins-go/global"
)

var (
	requiredCfg = []string{"username", "password", "database", "hostname", "port"}
)

func main() {
	cfgFile := myInit.InitArgs(requiredCfg)
	// dictCfg := myInit.InitConfig(requiredCfg, cfgFile)
	myInit.InitConfig(requiredCfg, cfgFile)
	// for k, v := range dictCfg {
	// 	fmt.Printf("%s %s\n", k, v)
	// }
	for k, v := range myGlobal.DefaultValues {
		fmt.Printf("%s %s\n", k, v)
	}
	myUtils.SignalHandler()
	os.Exit(0)
}
