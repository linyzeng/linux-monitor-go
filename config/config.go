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
// Date			:	May 14, 2017
//
// History	:
// 	Date:			Author:		Info:
//	May 24, 2017		LIS			First release
//

package config

import (
	"bufio"
	//"fmt"
	//"flag"
	"log"
	"os"
	"strconv"
	"strings"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	// myHelp		"github.com/my10c/nagios-plugins-go/help"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"

	"gopkg.in/natefinch/lumberjack.v2"
)

func getKeyValue(arg string) map[string]string {
}

func GetConfig(argv ...string) map[string]string {
	// the dictionary that will hold the config name-value pairs
	// we get 2 section in the yaml, common and app-name
	configMap := make(map[string]string)

	configFile := argv[0]
	// open the config file and read one line at the fime
	if yamlFile, err :=  os.Open(configFile); err == nil {
		// make sure it gets closed
		defer yamlFile.Close()

		// create a new scanner and read the file line by line
		var common_hit int = 0
		var app_hit  int = 0
		scanner := bufio.NewScanner(yamlFile)
		for scanner.Scan() {
			// skip comments, starts with #
			if !strings.HasPrefix(scanner.Text(), "#") {
				// look for the common section
				if strings.HasPrefix(scanner.Text(), "common:") {
					common_hit = 1
				}
				// look for the app section
				if strings.HasPrefix(scanner.Text(), myGlobal.MyProgname + ":") {
					app_hit = 1
				}
			}
		}
	} else {
		myUtils.ExitIfError(err)
	}
	return configMap
}

// Function to initialize logging
func InitLog(logSettings map[string]string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	MaxSize, _		:= strconv.Atoi(logSettings["LogMaxSize"])
	MaxBackups, _	:= strconv.Atoi(logSettings["LogMaxBackups"])
	MaxAge, _		:= strconv.Atoi(logSettings["LogMaxAge"])
	log.SetOutput(&lumberjack.Logger{
		Filename: 	logSettings["LogFile"],
		MaxSize:	MaxSize,
		MaxBackups:	MaxBackups,
		MaxAge:		MaxAge,
	})
}

// Function to process the given args
// func InitArgs() map[string]string {
// 	argsMap := make(map[string]string)
// 	flag.Usage = func() {
// 		fmt.Fprintf(os.Stderr, "Usage of %s:\n", myGlobal.MyProgname)
// 		flag.PrintDefaults()
// 	}
// 	version := flag.Bool("version", false, "Prints current version and exit.")
// 	setup := flag.Bool("setup", false, "Show the setup information and exit.")
// 	flag.Var(&myYamlFile, "yaml_file", "Yaml Configuration file to us.")
// 	flag.Parse()
// 	if *version {
// 		fmt.Printf("%s\n", myGlobal.MyVersion)
// 		os.Exit(0)
// 	}
// 	return argsMap
// }
