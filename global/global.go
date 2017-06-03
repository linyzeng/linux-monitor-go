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
// Version		:	0.3
//
// Date			:	June 1, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Mar 3, 2014		LIS			First release
//	May 18, 2017	LIS			Convert from bash/python/perl to Go
//	June 1, 2017	LIS			Added defaults for syslog, pagerduty, slack and email
//

package global

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	HR = "__________________________________________________"
	OK = 0
	WARNING = 1
	CRITICAL = 2
	UNKNOWN = 3

	// [KMG]Bytes units
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

)

var (
	MyVersion	= "0.1"
	now			= time.Now()
	MyProgname	= path.Base(os.Args[0])
	myAuthor	= "Luc Suryo"
	myCopyright	= "Copyright 2014 - " + strconv.Itoa(now.Year()) + " ©badassops"
	myLicense	= "License BSD, http://www.freebsd.org/copyright/freebsd-license.html ♥"
	myEmail		= "<luc@badassops.com>"
	MyInfo		= fmt.Sprintf("%s %s\n%s\n%s\nWritten by %s %s\n",
					MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)

	// Global variables
	Logfile			string
	ConfFile		string

	// require configuration keys
	RequiredKeys		[]string
	BabyshipKeys		[]string
	ServerKeys			[]string
	DirectoriesKeys		[]string
	LoggingKeys			[]string
	DirectoryKeys		[]string
	DefaultValues		map[string]string

	// defaults
	DefaultConfDir		= "/etc/nagios-plugins-go"
	DefaultLogsDir		= "/var/log/nagios-plugins-go"
	DefaultConfigFile	= fmt.Sprintf("%s/nagios-plugins-go.yaml", DefaultConfDir)

	// for logging
	DefaultLogFile			= fmt.Sprintf("%s/%s.log", DefaultLogsDir, MyProgname)
	DefaultLogMaxSize		= 128	// megabytes
	DefaultLogMaxBackups	= 3		// 3 files
	DefaultLogMaxAge		= 10	// days

	// debuging mode
	DefaultDebug		= "false"

	// email
	DefaultEmailTo		= ""

	// syslog
	DefaultSyslog			map[string]string
	DefaultSyslogTag		= fmt.Sprintf("[%s]", MyProgname)
	DefaultSyslogPriority	= "LOG_INFO"
	DefaultSyslogFacility	= "LOG_SYSLOG"

	// pagerdutry
	DefaultPD				map[string]string
	DefaultPDServiceKey		= ""
	DefaultPDServiceName	= ""

	// slack
	DefaultSlack			map[string]string
	DefaultSlackServiceKey	= ""
	DefaultSlackChannel		= ""

	// result wording
    Result = []string{ "OK", "WARNING", "CRITICAL", "UNKNOWN" }
)

func init() {
	// setup the default value, these are hardcoded.
	DefaultValues = make(map[string]string)
	DefaultValues["debug"]				=	DefaultDebug
	DefaultValues["logdir"]				=	DefaultLogsDir
	DefaultValues["logfile"]			=	DefaultLogFile
	DefaultValues["logmaxsize"]			=	strconv.Itoa(DefaultLogMaxSize)
	DefaultValues["logmaxbackups"]		=	strconv.Itoa(DefaultLogMaxBackups)
	DefaultValues["logmaxage"]			=	strconv.Itoa(DefaultLogMaxAge)
	DefaultValues["emailto"]			=	DefaultEmailTo
	//these are for getting a instance/system tag
	DefaultValues["tagfile"]			=	 ""
	DefaultValues["tagkeyname"]			=	""
	// for syslog
	DefaultSyslog = make(map[string]string)
	DefaultSyslog["syslogtag"]			=	DefaultSyslogTag
	DefaultSyslog["syslogpriority"]		=	DefaultSyslogPriority
	DefaultSyslog["syslogfacility"]		=	DefaultSyslogFacility
	// for pagerduty
	DefaultPD = make(map[string]string)
	DefaultPD["pdservicekey"]		=	DefaultPDServiceKey
	DefaultPD["pdservicename"]		=	DefaultPDServiceName
	// for slack
	DefaultSlack = make(map[string]string)
	DefaultSlack["slackservicekey"]	=	DefaultSlackServiceKey
	DefaultSlack["slackchannel"]		=	DefaultSlackChannel
}
