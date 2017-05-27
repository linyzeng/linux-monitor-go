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
	cnt int = 0
	warnThreshold int
	critThreshold int
)

// Function to check that the threshold are correct
func SanityCheck(warning string, critical string) (int, int, bool){
	if  strings.HasSuffix(warning, "%") {
		percent = true
		warnThreshold, _ = strconv.Atoi(warning[:len(warning) - 1])
		cnt++
	} else {
		warnThreshold , _ = strconv.Atoi(warning)
	}
	if  strings.HasSuffix(critical, "%") {
		percent = true
		critThreshold, _ = strconv.Atoi(critical[:len(critical) - 1])
		cnt++
	} else {
		critThreshold , _ = strconv.Atoi(critical)
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
	if warnThreshold >= critThreshold {
			fmt.Printf("%s", myGlobal.MyInfo)
		fmt.Printf("Warning threshold must be less than Critical threshold\n")
		os.Exit(1)
	}
	return warnThreshold, critThreshold, percent
}

// Function to check if the value is within threshold
func CalculateUsage(precent bool, warnThreshold int, critThreshold int, currValue int, totalValue int) int {
	// calculate based on %
	if precent == true {
		currValue = int((float64(currValue) * float64(100) ) / float64(totalValue))
	}
	if currValue >= critThreshold {
		return 2
	}
	if  currValue >= warnThreshold {
		return 1
	}
	return 0
}
