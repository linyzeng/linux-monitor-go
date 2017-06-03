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
// Date			:	June 2, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 2, 2017	LIS			First release
//

package alerts

import (
	"fmt"
	"log/syslog"
	"net/smtp"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
)

// Function to sent alerts
func SendAlert(message string) error {
	var err error = nil
	// Syslog : only if syslog tag was not set to off
	if myGlobal.DefaultSyslog["syslogtag"]  != "off" {
		alertSyslog(message)
		//fmt.Printf("Syslog %s\n", message)
	}
	// Email : only if emailto is not empty
	if len(myGlobal.DefaultValues["emailto"]) > 0 {
		fmt.Printf("Email %s\n", message)
	}
	// Pagerduty : only if key and service-name are not empty
	if len(myGlobal.DefaultPD["pdservicekey"]) > 0 &&
	   len(myGlobal.DefaultPD["pdservicename"]) > 0 {
		fmt.Printf("PD %s\n", message)
	}
	// Slack : only if key and channel are not empty
	if len(myGlobal.DefaultSlack["slackservicekey"]) > 0 &&
	   len(myGlobal.DefaultSlack["slackchannel"]) > 0 {
		fmt.Printf("Slack %s\n", message)
	}
	return err
}

// Function to create an syslog record
func alertSyslog(message string) error {
	// get the tag fro syslog
	tag := myGlobal.DefaultSyslog["syslogtag"]
	// get the int values
	priority, facility, err := myUtils.GetSyslog(myGlobal.DefaultSyslog["syslogpriority"],
			myGlobal.DefaultSyslog["syslogfacility"])
	if err != nil {
		return err
	}
	// create a syslog handler
	syslogHandler, err := syslog.New(syslog.Priority(priority|facility), tag)
	if err != nil {
		return err
	}
	_, err = syslogHandler.Write([]byte(message))
	return err
}

// Function to send an alert email
func alertEmail(message string) error {
}
