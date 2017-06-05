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

package stats

import (
	"fmt"
	"os"
	"sync"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"

)

type Stats struct {
	StatFilename string
	Enable bool
	StatsRecord string
	mu sync.Mutex
}

// Function to create a stats jobs
func New() *Stats {
	statFilenName, isEnable := initStats()
	statPtr := &Stats {
		StatFilename: statFilenName,
		Enable: isEnable,
		StatsRecord:  "",
	}
	return statPtr
}

// Function to initialize the stats directoty and the stats file
func initStats() (string, bool) {
	// only if stats was enable
	if myGlobal.DefaultStats["stats"] == "false" {
		err := fmt.Errorf("Stats is disable.")
		myUtils.LogMsg(fmt.Sprintf("%s\n", err.Error()))
		return "", false
	}
	fmt.Printf("dir %s file %s\n\n", myGlobal.DefaultStats["statsdir"], myGlobal.DefaultStats["statsfile"])
	// make sure the both statsdir and statsfiel were set
	if len(myGlobal.DefaultStats["statsdir"]) == 0 ||
		len(myGlobal.DefaultStats["statsfile"]) == 0 {
		err := fmt.Errorf("Either statsdir or/and statsfile was not set, stats has been disabled.")
		myUtils.LogMsg(fmt.Sprintf("%s\n", err.Error()))
		myGlobal.DefaultStats["stats"] = "false"
		return "", false
	}
	// create the directory
	err := os.MkdirAll(myGlobal.DefaultLog["logdir"], 0755)
	if err != nil {
		fmt.Printf("Unable to create stats directory, stats has been disabled.\n")
		myUtils.LogMsg(fmt.Sprintf("%s\n", err.Error()))
		myGlobal.DefaultStats["stats"] = "false"
		return "", false
	}
	// create the full path name
	statsFile := fmt.Sprintf("%s/%s", myGlobal.DefaultStats["statsdir"], myGlobal.DefaultStats["statsfile"])
	return statsFile, true 
}

// Function to write given stats record to the stat file
func (statsPtr *Stats) WriteStat() error {
	// we return if forwhatever reason that stats is not enabled
	if  !statsPtr.Enable {
		return nil
	}
	statsFile, err := os.OpenFile(statsPtr.StatFilename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		myUtils.LogMsg(fmt.Sprintf("Unable to open the stats file, this record has been skipped.\n%s\n", statsPtr.StatsRecord))
		myUtils.LogMsg(fmt.Sprintf("%s\n", err.Error()))
		return err
	}
	defer statsFile.Close()
	statsPtr.mu.Lock()
	defer statsPtr.mu.Unlock()
	_, err = statsFile.WriteString(statsPtr.StatsRecord)
	if err != nil {
		fmt.Printf("Unable to open write stats record, this record has been skipped.\n")
		myUtils.LogMsg(fmt.Sprintf("%s\n", err.Error()))
		return err
	}
	return nil
}

// Function to create a start record based on the given string
func (statsPtr *Stats) CreateRecord() error {
	var err error = nil
	return err
}
