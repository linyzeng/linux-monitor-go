// Copyright (c) 2017 - 2017 badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
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
// Date			:	June 4, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 4, 2017	LIS			First Go release
//
// TODO:

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	myGlobal "github.com/my10c/nagios-plugins-go/global"
	myInit "github.com/my10c/nagios-plugins-go/initialize"
	myMemory "github.com/my10c/nagios-plugins-go/memory"
	// myThreshold "github.com/my10c/nagios-plugins-go/threshold"
	myUtils "github.com/my10c/nagios-plugins-go/utils"
)

const (
	extraInfo    = "\tEmpty unit defaults to MB"
	CheckVersion = "0.1"
)

var (
	cfgRequired = []string{"unit"}
	err         error
	exitVal     int = 0
)

func wrongMode(modeSelect string) {
	fmt.Printf("%s", myGlobal.MyInfo)
	if modeSelect == "help" {
		fmt.Printf("Supported modes\n")
	} else {
		fmt.Printf("Wrong mode, supported modes:\n")
	}
	fmt.Printf("\t memory       : checks current memory usage, requires the configs: `memcritical` and `memwarning`.\n")
	fmt.Printf("\t swap         : checks current swap usage, requires the configs: `memcritical` and `memwarning`.\n")
	fmt.Printf("\t system       : show the current system memory status.\n")
	fmt.Printf("\t top-memory   : show top process memory usage, optional the config `topcount`\n")
	fmt.Printf("\t top-rss      : show top process memory usage, optional the config `topcount`\n")
	fmt.Printf("\t top-private  : show top process private memory usage, optional the config `topcount`\n")
	fmt.Printf("\t top-swap     : show top process swap memory usage, optional the config `topcount`\n")
	fmt.Printf("\t showconfig   : show the current configuration and then exit.\n")
	os.Exit(3)
}

func wrongUnit(confUnit string) {
	fmt.Printf("%s", myGlobal.MyInfo)
	fmt.Printf("Wrong unit %s, supported unit:\n", confUnit)
	fmt.Printf("\t KB	: KiloBytes, most accurate.\n")
	fmt.Printf("\t MB	: MegaBytes, good accuracy.\n")
	fmt.Printf("\t GB	: GigaBytes, less accurate.\n")
	fmt.Printf("\t TB	: TerraBytes, worst accuracy.\n")
	os.Exit(3)
}

func checkUnit(unit string) uint64 {
	// since everything is in KB
	var unitBytes uint64 = 1
	switch unit {
	case "", "KB":
		unitBytes = 1
	case "MB":
		unitBytes = myGlobal.KB
	case "GB":
		unitBytes = myGlobal.MB
	case "TB":
		unitBytes = myGlobal.GB
	default:
		wrongUnit(unit)
	}
	return unitBytes
}

func checkMode(givenMode string) {
	switch givenMode {
	case "memory", "swap", "system", "top-memory", "top-rss", "top-private", "top-swap", "show-config":
	default:
		wrongMode(givenMode)
	}
}

func main() {
	// working variables
	var usePercent bool = false
	var topCount int
	var exitVal int = 0
	var exitMsg string
	// create emtpy error message
	err = fmt.Errorf("")
	// need to be root since the config file wil have passwords
	myUtils.IsRoot()
	// get and setup phase
	myUtils.IsLinuxSystem()
	myGlobal.ExtraInfo = extraInfo
	myGlobal.MyVersion = CheckVersion
	cfgFile, givenMode := myInit.InitArgs(cfgRequired)
	switch givenMode {
	case "memory", "swap":
		cfgRequired = append(cfgRequired, "memcritical", "memwarning")
	case "top-memory", "top-rss", "top-private", "top-swap":
		cfgRequired = append(cfgRequired, "topcount")
	}
	cfgDict := myInit.InitConfig(cfgRequired, cfgFile)
	myInit.InitLog()
	myUtils.SignalHandler()
	//--> stats := myStats.New()
	givenUnit := checkUnit(cfgDict["unit"])
	checkMode(givenMode)
	//data := time.Now().Format(time.RFC3339)
	thresHold := fmt.Sprintf(" (W:%s C:%s Unit:%s)", cfgDict["memwarning"], cfgDict["memcritical"], cfgDict["unit"])
	if strings.HasSuffix(cfgDict["memwarning"], "%") {
		usePercent = true
	}
	iter, _ := strconv.Atoi(cfgDict["iter"])
	iterWait, _ := time.ParseDuration(cfgDict["iterwait"])
	if strings.HasPrefix(givenMode, "top") {
		topCount, _ = strconv.Atoi(cfgDict["topcount"])
	}
	// Get the memory infos
	systemMemInfo, processMemInfo := myMemory.New(givenUnit)
	for cnt := 0; cnt < iter; cnt++ {
		switch givenMode {
		case "memory":
			exitVal = systemMemInfo.CheckFreeMem(cfgDict["memwarning"], cfgDict["memcritical"])
			exitMsg = fmt.Sprintf("(Memory Total:%s, Free:%s, Usage:%s)%s",
				strconv.FormatUint(systemMemInfo.Total(), 10), strconv.FormatUint(systemMemInfo.RealFree(), 10),
				strconv.FormatUint(systemMemInfo.RealUsage(), 10), thresHold)
			if usePercent {
				exitMsg = fmt.Sprintf("%s (Usage %d%%)", exitMsg, systemMemInfo.UsagePercent())
			}
		case "swap":
			exitVal = systemMemInfo.CheckFreeSwap(cfgDict["memwarning"], cfgDict["memcritical"])
			exitMsg = fmt.Sprintf("(Swap Total:%s, Free:%s, Usage:%s)%s",
				strconv.FormatUint(systemMemInfo.Swap(), 10), strconv.FormatUint(systemMemInfo.FreeSwap(), 10),
				strconv.FormatUint(systemMemInfo.SwapUsage(), 10), thresHold)
			if usePercent {
				exitMsg = fmt.Sprintf("%s (Usage %d%%)", exitMsg, systemMemInfo.SwapUsagePercent())
			}
		case "system":
			exitMsg = systemMemInfo.Show()
		case "top-memory", "top-rss":
			exitMsg = myMemory.GetTop(topCount, "rss", processMemInfo)
			break
		case "top-private":
			exitMsg = myMemory.GetTop(topCount, "private", processMemInfo)
			break
		case "top-swap":
			exitMsg = myMemory.GetTop(topCount, "swap", processMemInfo)
			break
		case "showconfig":
			myUtils.ShowMap(cfgDict)
			myUtils.ShowMap(nil)
		default:
			wrongMode(givenMode)
		}
		// if we get an OK then exit no need to do all iterations
		if exitVal == myGlobal.OK {
			break
		}
		time.Sleep(iterWait * time.Second)
	}

	// add the check name
	exitMsg = fmt.Sprintf("%s %s - %s\n",
		strings.ToUpper(myGlobal.MyProgname), myGlobal.Result[exitVal], exitMsg)
	fmt.Printf("%s", exitMsg)
	myUtils.LogMsg(fmt.Sprintf("%s", exitMsg))
	os.Exit(exitVal)
}
