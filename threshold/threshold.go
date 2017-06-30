// Copyright (c) 2017 - 2017 badassops
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
// Date			:	June 4, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 4, 2017	LIS			First Go release
//
// TODO:

package threshold

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
)

var (
	percent bool = false
	warnThreshold uint64
	critThreshold uint64
)

// Function to check that the configured thresholds are correct
func SanityCheck(revert bool, warning string, critical string) (uint64, uint64, bool){
	var cnt int = 0
	if strings.HasSuffix(warning, "%") {
		percent = true
		warnThreshold, _ = strconv.ParseUint(warning[:len(warning) - 1], 10, 64)
		cnt++
	} else {
		warnThreshold , _ = strconv.ParseUint(warning, 10, 64)
	}
	if strings.HasSuffix(critical, "%") {
		percent = true
		critThreshold, _ = strconv.ParseUint(critical[:len(critical) - 1], 10, 64)
		cnt++
	} else {
		critThreshold , _ = strconv.ParseUint(critical, 10, 64)
	}
	if percent == true {
		if cnt != 2 {
			fmt.Printf("%s", myGlobal.MyInfo)
			fmt.Printf("Percentage was given but not both has the percent sign\n")
			os.Exit(1)
		}
		if warnThreshold < 0 || warnThreshold > 100 {
			fmt.Printf("%s", myGlobal.MyInfo)
			fmt.Printf("Warning threshold percentage must be between 0 and 100\n")
			os.Exit(1)
		}
		if critThreshold < 0 || critThreshold > 100 {
			fmt.Printf("%s", myGlobal.MyInfo)
			fmt.Printf("Critical threshold percentage must be between 0 and 100\n")
			os.Exit(1)
		}
	}
	if revert {
		if warnThreshold <= critThreshold {
			fmt.Printf("%s", myGlobal.MyInfo)
				fmt.Printf("Critical threshold must be less than Warning threshold\n")
			os.Exit(1)
		}
	} else {
		if warnThreshold >= critThreshold {
			fmt.Printf("%s", myGlobal.MyInfo)
				fmt.Printf("Warning threshold must be less than Critical threshold\n")
			os.Exit(1)
		}
	}
	return warnThreshold, critThreshold, percent
}

// Function to check if the value is within threshold
func CalculateUsage(revert bool, precent bool, warnThreshold uint64, critThreshold uint64, currValue uint64, totalValue uint64) int {
	// calculate based on %
	if precent == true {
		// need to use float to get correct division value
		// note the value will be down-rounded
		currValue = uint64((float64(currValue)/float64(totalValue)) * 100)
	}
	if revert {
		if currValue <= critThreshold {
			return 2
		}
		if currValue <= warnThreshold {
			return 1
		}
	} else {
		if currValue >= critThreshold {
			return 2
		}
		if currValue >= warnThreshold {
			return 1
		}
	}
	return 0
}
