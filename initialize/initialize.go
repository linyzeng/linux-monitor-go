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

package initialize

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myHelp		"github.com/my10c/nagios-plugins-go/help"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"

	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/my10c/simpleyaml"
)

// type used for flags in initArgs
type stringFlag struct {
	value   string
	set     bool
}

// Function for the stringFlag struct, set the values
func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

// Function for the stringFlag struct, get the values
func (sf *stringFlag) String() string {
	return sf.value
}

// Function to return the yaml value, nil if error or nil if not found
func getYamlValue(yamFile *simpleyaml.Yaml, section string, key string) (string, error) {
	// Check if section exist and/or key, no point to go further if it doesn't exist
	keyExist := yamFile.GetPath(section, key)
	if keyExist.IsFound() == false {
		err := fmt.Errorf("Section %s and/or key %s not found\n", section, key)
		return "", err
	}
	// We need to ge the value and since we do not know what it is, we check
	// against the 3 supported type
	// check if value is a string
	if value, err := yamFile.Get(section).Get(key).String(); err == nil {
		return value, err
	}
	// check if value is a int
	if value, err := yamFile.Get(section).Get(key).Int(); err == nil {
		return strconv.Itoa(value), err
	}
	// check if value is a boolean
	if value, err := yamFile.Get(section).Get(key).Bool(); err == nil {
		return strconv.FormatBool(value), err
	}
	err := fmt.Errorf("Unsupported value for section %s and key %s, suported are: string, int and bool\n", section, key)
	return "", err
}

// Function to get the configuration
func InitConfig(cfgList []string, argv...string) map[string]string {
	// working variable
	var missingKeys []string
	dictCfg := make(map[string]string)
	// open given file and check that is a correct yaml file
	cfgFile, err := ioutil.ReadFile(argv[0])
	myUtils.ExitIfError(err)
	yamlFile, err := simpleyaml.NewYaml(cfgFile)
	myUtils.ExitIfError(err)
	// first check if the default common values need to be modify
	for defaultKey, _ := range myGlobal.DefaultValues {
		if newValue, err := getYamlValue(yamlFile, "common", defaultKey); err == nil {
			// replace the default value
			myGlobal.DefaultValues[defaultKey] = newValue
		}
	}
	// for Syslog
	for defaultSyslog, _ := range myGlobal.DefaultSyslog {
		if newValue, err := getYamlValue(yamlFile, "syslog", defaultSyslog); err == nil {
			// replace the default value
			myGlobal.DefaultSyslog[defaultSyslog] = newValue
		}
	}
	// for Pagerduty
	for defaultPD, _ := range myGlobal.DefaultPD {
		if newValue, err := getYamlValue(yamlFile, "pagerduty", defaultPD); err == nil {
			// replace the default value
			myGlobal.DefaultPD[defaultPD] = newValue
		}
	}
	// for Slack
	for defaultSlack, _ := range myGlobal.DefaultSlack {
		if newValue, err := getYamlValue(yamlFile, "slack", defaultSlack); err == nil {
			// replace the default value
			myGlobal.DefaultSlack[defaultSlack] = newValue
		}
	}
	// set the config value
	for cnt := range cfgList {
		keyName := cfgList[cnt]
		if cfgValue, err := getYamlValue(yamlFile, myGlobal.MyProgname, keyName); err == nil {
			// assign the value
			dictCfg[keyName] = cfgValue
		} else {
			missingKeys = append(missingKeys, keyName)
		}
	}
	// make sure we have all required configs
	if len(missingKeys) != 0 {
		fmt.Printf("Following keys are missing in the configration files: %s\n", missingKeys)
		os.Exit(2)
	}
	return dictCfg
}

// Function to initialize logging
func InitLog(logSettings map[string]string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	MaxSize, _		:= strconv.Atoi(logSettings["logmaxsize"])
	MaxBackups, _	:= strconv.Atoi(logSettings["logmaxbackups"])
	MaxAge, _		:= strconv.Atoi(logSettings["logmaxage"])
	log.SetOutput(&lumberjack.Logger{
		Filename:	logSettings["logfile"],
		MaxSize:	MaxSize,
		MaxBackups:	MaxBackups,
		MaxAge:		MaxAge,
	})
}

// Function to process the given args
func InitArgs(cfg []string) (string, string) {
	var myConfigFile stringFlag
	var myMode stringFlag
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", myGlobal.MyProgname)
		flag.PrintDefaults()
	}
	version := flag.Bool("version", false, "Prints current version and exit.")
	setup := flag.Bool("setup", false, "Show the setup information and exit.")
	noalert := flag.Bool("noalert", false, "Send no alert.")
	nolog := flag.Bool("nolog", false, "Do not log result.")
	flag.Var(&myConfigFile, "config", "Configuration file to be used.")
	flag.Var(&myMode, "mode", "check mode, use `-mode help` to see available modes.")
	flag.Parse()
	if *version {
		fmt.Printf("%s\n", myGlobal.MyVersion)
		os.Exit(0)
	}
	if *setup {
		myHelp.SetupHelp(cfg)
	}
	if !myMode.set {
		myHelp.Help(1)
	}
	// if not set we use the default
	if !myConfigFile.set{
		myConfigFile.Set(myGlobal.DefaultConfigFile)
	}
	// set the noalert and nolog
	myGlobal.DefaultValues["noalert"] = strconv.FormatBool(*noalert)
	myGlobal.DefaultValues["nolog"] = strconv.FormatBool(*nolog)
	return myConfigFile.value, myMode.value
}
