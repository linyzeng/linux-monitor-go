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
// Version		:	0.2
//
// Date			:	May 18, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Mar 3, 2014		LIS			First release
//	May 18, 2017	LIS			Convert from bash/python/perl to Go
//

package help

import (
	"fmt"
	"os"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
)

// Function to show how to setup the aws credentials and the simple-aws-lb config
func SetupHelp(cfg []string) {
	fmt.Printf("%s", myGlobal.MyInfo)
	fmt.Printf("Setup the configuration file:\n")
	fmt.Printf("\t# Create a configuration file, any name would do, as long its in yaml fornmat.\n")
	fmt.Printf("\t# Default to %s\n", myGlobal.DefaultConfigFile)
	fmt.Printf("\t# *** Note that the key most be all lowercase! ***\n")
	fmt.Printf("\t# Add the following key/pair values, these are required:\n")
	fmt.Printf("%s:\n", myGlobal.MyProgname)
	for cnt := range cfg {
		fmt.Printf("  %s:\n", cfg[cnt])
	}
	fmt.Printf("\t# Optional add these values in the common section.\n")
	fmt.Printf("\t# Values shown are the default values. If either emailfrom or emailto is empty then no email will be sent.\n")
	fmt.Printf("\t# tagfile and tagkeyname are use to get the tag info by looking for the key tagkeyname in the\n")
	fmt.Printf("\t# configured file tagfile, the format need to be just 'keyname value' nothing fancy!\n")
	fmt.Printf("common:\n")
	for defaultKey, defaultValue := range myGlobal.DefaultValues {
		fmt.Printf("  %s: %s\n", defaultKey, defaultValue)
	}
	fmt.Printf("\n\t# Syslog support, to disable set tag value to off, syslogtag default to %s,\n", myGlobal.MyProgname)
	fmt.Printf("syslog:\n")
	for defaultKey, defaultValue := range myGlobal.DefaultSyslog {
		fmt.Printf("  %s: %s\n", defaultKey, defaultValue)
	}
	fmt.Printf("\n\t# Optional for pagerduty support, if any of these keys are empty then pagerduty is not used.\n")
	fmt.Printf("pagerduty:\n")
	for defaultKey, defaultValue := range myGlobal.DefaultPD {
		fmt.Printf("  %s: %s\n", defaultKey, defaultValue)
	}
	fmt.Printf("\n\t# Optional for slack support, if slackservicekey and/or slackchannel is empty then slack is not used.\n")
	fmt.Printf("slack:\n")
	for defaultKey, defaultValue := range myGlobal.DefaultSlack {
		fmt.Printf("  %s: %s\n", defaultKey, defaultValue)
	}
	fmt.Printf("\n\nNOTE\n")
	fmt.Printf("\t* Any key that has any of these charaters: ':#[]()*' in their value must be double quoted!\n") 
	fmt.Printf("\t* Syslog Valid Priority: ")
	for keyPriority, _ := range myUtils.SyslogPriority {
		fmt.Printf("%s ", keyPriority)
	}
	cnt := 0
	fmt.Printf("\n\t* Syslog Valid Facility: ")
	for keyFacility, _ := range myUtils.SyslogFacility {
		if cnt > 5 {
			fmt.Printf("\n\t\t%s ", keyFacility)
			cnt = 0
		} else {
			fmt.Printf("%s ", keyFacility)
			cnt += 1
		}
	}
	fmt.Printf("\n")
	os.Exit(0)
}

// Function to show the help information
func Help(exitVal int) {
	fmt.Printf("%s", myGlobal.MyInfo)
	optionList := "<--config config file> <--check mode> <--setup> <--version> <--help>"
	fmt.Printf("\nUsage : %s\n\tflags: %s\n", myGlobal.MyProgname, optionList)
	fmt.Printf("\t*config: the configuration file to use, should be full path, use --setup for more information.\n")
	fmt.Printf("\t*check: mode, this is defined per check, use 'mode help' to see valid modes.\n")
	fmt.Printf("\tsetup: show the setup guide.\n")
	fmt.Printf("\tversion: print %s version.\n", myGlobal.MyProgname)
	fmt.Printf("\thelp: short version of this help page.\n")
	fmt.Printf("\n\t* == required flag.\n")
	os.Exit(exitVal)
}
