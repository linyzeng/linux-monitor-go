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
// TODO: process error

package alerts

import (
	"fmt"
	"log/syslog"
	"net/mail"
	"net/smtp"
	"os"
	"strings"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
	mytag		"github.com/my10c/nagios-plugins-go/tag"

	"github.com/nlopes/slack"
)

// Function to sent alerts
func SendAlert(exitVal int, checkMode string, checkErr string) error {
	var hostName string
	var message string
	var err error = nil
	// create the full message and subject
	errWord := myGlobal.Result[exitVal]
	hostName, hostOK := os.Hostname()
	if hostOK != nil {
		hostName = "Unable to get hostname"
	}
	hostName = strings.TrimSpace(hostName)
	tagInfo, tagOK := mytag.GetTagInfo()
	if tagOK != nil {
		message = fmt.Sprintf("TAG: no tag found\nHost: %s\n%s: %s\nCheck running mode: %s\nError: %s\n",
				hostName, myGlobal.MyProgname, errWord, checkMode, checkErr)
	} else {
		message = fmt.Sprintf("TAG: %s\nHost: %s\n%s: %s\nCheck running mode: %s\nError: %s\n",
			strings.TrimSpace(tagInfo), hostName, myGlobal.MyProgname, errWord, checkMode, checkErr)
	}
	errSubject := fmt.Sprintf("%s : MONITOR ALERT : %s : %s ", errWord, hostName, myGlobal.MyProgname)
	// Syslog : only if syslog tag was not set to of
	if myGlobal.DefaultSyslog["syslogtag"] != "off" {
		alertSyslog(message)
	}
	// Email : only if emailto is not empty
	if len(myGlobal.DefaultValues["emailto"]) > 0 {
		alertEmail(message, errSubject)
	}
	// Pagerduty : only if key and service-name are not empty
	// if len(myGlobal.DefaultPD["pdservicekey"]) > 0 &&
	//    len(myGlobal.DefaultPD["pdservicename"]) > 0 {
	// 	fmt.Printf("\nPD %s\n\n", message)
	// }
	// Slack : only if key and channel are not empty
	if len(myGlobal.DefaultSlack["slackservicekey"]) > 0 &&
	   len(myGlobal.DefaultSlack["slackchannel"]) > 0 {
		alertSlack(message)
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
	syslogHandler, err := syslog.New(syslog.Priority(facility|priority), tag)
	if err != nil {
		return err
	}
	_, err = syslogHandler.Write([]byte(message))
	return err
}

// Function to send an alert email
func alertEmail(message string, subject string) error {
	// if authEmail - empty then no authentication is required
	// for now only support PlainAuth
	authEmail := smtp.PlainAuth("",
			myGlobal.DefaultValues["emailuser"],
			myGlobal.DefaultValues["emailpass"],
			myGlobal.DefaultValues["emailhost"],
	)
	// build the email component
	emailTo := mail.Address{myGlobal.DefaultValues["emailtoname"], myGlobal.DefaultValues["emailto"]}
	emailFrom := mail.Address{myGlobal.DefaultValues["emailfromname"], myGlobal.DefaultValues["emailfrom"]}
	emailSubject := subject
	emailHost := fmt.Sprintf("%s:%s", myGlobal.DefaultValues["emailhost"], myGlobal.DefaultValues["emailhostport"])
	fromAndBody := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n",
			emailTo.String(), emailSubject, message)
	// send the email
	err := smtp.SendMail(emailHost, authEmail, emailFrom.String(), []string{emailTo.String()}, []byte(fromAndBody))
	return err
}

// Function to post an alert in slack
func alertSlack(message string) error {
	slackAPI := slack.New(myGlobal.DefaultSlack["slackservicekey"])
	slackMsg := fmt.Sprintf(":imp: `- %s error: %s` :disappointed:\n", myGlobal.MyProgname, message)
	// remove all carriage return
	slackMsg = strings.TrimSuffix(strings.Replace(slackMsg, "\n", " - ", -1), " - ")
	// need to build a minimum config
	slackMsgConfig := slack.PostMessageParameters{
		Username:		"MONITOR",
		AsUser:			false,
		Parse:			"",
		LinkNames:		0,
		Attachments:	nil,
		UnfurlLinks:	false,
		UnfurlMedia:	true,
		IconURL:		"",
		IconEmoji:		myGlobal.DefaultSlack["iconemoji"],
		Markdown:		true,
		EscapeText:		true,
	}
	_, _, err := slackAPI.PostMessage(myGlobal.DefaultSlack["slackchannel"], slackMsg, slackMsgConfig)
	return err
}
