// Copyright (c) 2017 - 2017 badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
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
// Date			:	July 1, 2017
//
// History	:
// 	Date:			Author:		Info:
//	July 1, 2017	LIS			First Go release
//
// TODO:

package memory

import (
//	"fmt"
//	"io/ioutil"
//	"path/filepath"
//	"regexp"
//	"strings"
//	"syscall"
//
//	myGlobal	"github.com/my10c/nagios-plugins-go/global"
//	myUtils		"github.com/my10c/nagios-plugins-go/utils"
//	myThreshold	"github.com/my10c/nagios-plugins-go/threshold"
)

const (
	PROCMEM = "/proc/meminfo"
	PROCESSCOM = "comm"
)

type memStruct struct {
	memTotal		uint64	`json:"memTotal"`
	memFree			uint64	`json:"memFree"`
	memAvailable	uint64	`json:"memAvailable"`
	buffers			uint64	`json:"buffers"`
	cached			uint64	`json:"cachedcached"`
	swapTotal		uint64	`json:"swapTotal"`
	swapFree		uint64	`json:"swapTotal"`
}

// Rss: resident memory usage, all memory the process uses,
//		including all memory this process shares with other processes. It does not include swap;
// Shared: memory that this process shares with other processes;
// Private: private memory used by this process, you can look for memory leaks here;
// Swap: swap memory used by the process;
// Pss: Proportional Set Size, a good overall memory indicator.
//		It is the Rss adjusted for sharing: if a process has 1MiB private and 20MiB shared
//		between other 10 processes, Pss is 1 + 20/10 = 3MiB

type processMemStruct struct {
	processName		string	`json:"procname"`
	rrsTotal		uint64	`json:"rss"`
	sharedTotal		uint64	`json:"shared"`
	privateTotal	uint64	`json:"private"`
	swapTotal		uint64	`json:"swap"`
	pssTptal		uint64	`json:"pss"`
}

var (
	memRegex = `^(MemTotal|MemFree|Cached)`
	procMemRegex = `^(Rss:|Shared:|Private:|Pss:)`
)

//func geProcMem() map[string]memStruct {
//	// working variable
//	contents, err := ioutil.ReadFile(PROCMOUNT)
//	myUtils.ExitWithNagiosCode(myGlobal.UNKNOWN, err)
//	// create the return map
//	detectedPartitions := make(map[string]parStruct)
//	// get all lines and walk one at the time
//	lines := strings.Split(string(contents), "\n")
//	for _, line := range(lines) {
//		if line != "" {
//			// we need the first 3 fields : device, mountpoint, type and first word of mount (rw or ro)
//			currDevice := strings.Fields(line)[0]
//			currMountPoint := strings.Fields(line)[1]
//			currFSType := strings.Fields(line)[2]
//			currState := strings.Split(strings.Fields(line)[3], ",")[0]
//			// we only want those matching parRegex
//			match := expDev.MatchString(currDevice)
//			if match {
//				// check is we have a possible symlink or fullpath
//				match = expLogics.MatchString(currDevice)
//				if match {
//					currDevice, _ = filepath.EvalSymlinks(currDevice)
//				}
//				// get the disk/partion info
//				currPartition.device = currDevice
//				currPartition.mountpoint = currMountPoint
//				currPartition.fsType = currFSType
//				currPartition.mountState = currState
//				detectedPartitions[currMountPoint] = currPartition
//			}
//		}
//	}
//	return detectedPartitions
//}

// Function to get the given partition/mount point file system info
// func getDiskinfo(path string) *diskType {
// 	fs := syscall.Statfs_t{}
// 	err := syscall.Statfs(path, &fs)
// 	if err != nil {
// 		myUtils.ExitWithNagiosCode(myGlobal.UNKNOWN, err)
// 	}
// 	disk := &diskType {
// 		totalSpace	: fs.Blocks * uint64(fs.Bsize),
// 		totalFree	: fs.Bfree * uint64(fs.Bsize),
// 		totalUse	: (fs.Blocks * uint64(fs.Bsize)) - (fs.Bfree * uint64(fs.Bsize)),
// 		totalInodes	: fs.Files,
// 		freeInodes	: fs.Ffree,
// 		mountPoint	: path,
// 	}
// 	return disk
// }

// Function to get the available disks information
// func New() map[string]*diskType {
// 	// create the disk/partition map
// 	detectedPart := make(map[string]*diskType)
// 	// will return empty map if no valid disk/partition was found
// 	for mntPoint, partInfo := range getPartitions() {
// 		currDisk := getDiskinfo(mntPoint)
// 		if currDisk == nil {
// 			return nil
// 		}
// 		currDisk.device = partInfo.device
// 		currDisk.fsType = partInfo.fsType
// 		currDisk.mountState = partInfo.mountState
// 		detectedPart[mntPoint] = currDisk
// 	}
// 	return detectedPart
// }
