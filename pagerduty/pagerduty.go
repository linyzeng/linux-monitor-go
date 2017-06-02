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
// Date			:	June 1, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 1, 2014		LIS			First release
//

package mysql

import (
	//"encoding/json"
	"fmt"
	//"strconv"

	//myGlobal	"github.com/my10c/nagios-plugins-go/global"
	//myUtils		"github.com/my10c/nagios-plugins-go/utils"
	//myThreshold	"github.com/my10c/nagios-plugins-go/threshold"

	PD			"github.com/PagerDuty/go-pagerduty"
)

type PagerDuty struct {
	ServiceKey string
	LoggerTag string
	ServiceName string
	PostUrl string
	Event PD.Event
}

func New(PDCfg map[string]string) *PagerDuty {
	:wq!
	return nil
}

func (pd *PagerDuty) TriggerIncident(tag string, description string) error {

	// create the header for post
	headers := fmt.Sprintf("{'Authorization': 'Token token={0}'.format(%s), 'Content-type': 'application/json', }", pd.ServiceName)

	_, error := PD.CreateEvent(pd.Event)
	return error
}