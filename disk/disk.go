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
// Date			:	May 30, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Mar 3, 2014		LIS			First release
//	May 30, 2017	LIS			Convert from bash/python/perl to Go
//

package disk

import (
	"fmt"
	"io/ioutil"
//	"regexp"
	"strings"
//	"strconv"
//	"syscall"
//
	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
//	myThreshold	"github.com/my10c/nagios-plugins-go/threshold"
)

const (
	PROCMOUNT = "/proc/mounts"
)

var (
	// valid partiton and disk we support
	parRegex = `^(/dev/)(xvd|sd|disk|mapper)`
	symRegex = `^(/dev/)(disk|mapper)`
)

type LnxDisk struct {
	totalSpace uint64
	totalUse uint64
	totalFree uint64
	mountPoint string
}

func New() *LnxDisk {
	contents, err := ioutil.ReadFile(PROCMOUNT)
	myUtils.ExitWithNagiosCode(myGlobal.UNKNOWN, err)
	// get all lines and walk one at the time
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fmt.Printf("- %s - ", strings.Fields(line))
	}
	return nil
}
