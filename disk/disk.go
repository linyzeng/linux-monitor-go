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

package disk

import (
	"io/ioutil"
//	"os"
	"path/filepath"
	"regexp"
	"strings"
//	"strconv"
	"syscall"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
	myUtils		"github.com/my10c/nagios-plugins-go/utils"
//	myThreshold	"github.com/my10c/nagios-plugins-go/threshold"
)

const (
	PROCMOUNT = "/proc/mounts"
)

var (
	// valid device we support
	devRegex = `^(/dev/)(xvd|sd|disk|mapper)`
	symRegex = `^(/dev/)(disk|mapper)`
)

type parStruct struct {
	device		string
	mountpoint	string
	fsType		string
}

type diskStruct struct {
	totalSpace	uint64	`json:"total"`
	totalUse	uint64	`json:"used"`
	totalFree	uint64	`json:"free"`
	totalInodes	uint64	`json:"inodes"`
	freeInodes	uint64	`json:"freeinodes"`
	mountPoint	string	`json:"mount"`
	device		string	`json:"device"`
	fsType		string	`json:"fstype"`
}

func getPartitions() map[string]parStruct {
	// working variable
	var currPartition parStruct
	// get disks info from proc
	contents, err := ioutil.ReadFile(PROCMOUNT)
	myUtils.ExitWithNagiosCode(myGlobal.UNKNOWN, err)
	// prep the regex, we ignore the errors
	expDev, _ :=  regexp.Compile(devRegex)
	expLogics, _ := regexp.Compile(symRegex)
	// create the return map
	detectedPartitions := make(map[string]parStruct)
	// get all lines and walk one at the time
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		if line != "" {
			// we need the first 3 fields : device, mountpoint, type
			currDevice := strings.Fields(line)[0]
			currMountPoint := strings.Fields(line)[1]
			currFSType := strings.Fields(line)[2]
			// we only want those matching parRegex
			match := expDev.MatchString(currDevice)
			if match {
				// check is we have a possible symlink or fullpath
				match = expLogics.MatchString(currDevice)
				if match {
					currDevice, _ = filepath.EvalSymlinks(currDevice)
				}
				// get the disk/partion info
				currPartition.device = currDevice
				currPartition.mountpoint = currMountPoint
				currPartition.fsType = currFSType
				detectedPartitions[currMountPoint] = currPartition
			}
		}
	}
	return detectedPartitions
}

// Function to get the given partition/mount point file system info
func getDiskinfo(path string) (diskStruct, error) {
	var disk diskStruct
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return disk, err
	}
	disk.totalSpace = fs.Blocks * uint64(fs.Bsize)
	disk.totalFree = fs.Bfree * uint64(fs.Bsize)
	disk.totalUse = disk.totalSpace - disk.totalFree
	disk.totalInodes = fs.Files
	disk.freeInodes = fs.Ffree
	disk.mountPoint = path
	return disk, nil
}

// Function to get the available disks information
func New() map[string]diskStruct {
	var err error
	var currDisk diskStruct
	// create the disk/partition map
	detectedPart := make(map[string]diskStruct)
	// will return empty map if no valid disk/partition was found
	for mntPoint, partInfo := range getPartitions() {
		currDisk, err = getDiskinfo(mntPoint)
		if err != nil {
			return nil
		}
		currDisk.device = partInfo.device
		currDisk.fsType = partInfo.fsType
		detectedPart[mntPoint] = currDisk
	}
	return detectedPart
}

// Functions to get disk/partitions element info
func (diskPtr *diskStruct) GetType() string {
	return diskPtr.fsType
}

func (diskPtr *diskStruct) GetSize() uint64 {
	return diskPtr.totalSpace
}

func (diskPtr *diskStruct) GetUse() uint64 {
	return diskPtr.totalUse
}

func (diskPtr *diskStruct) GetFree() uint64 {
	return diskPtr.totalFree
}

func (diskPtr *diskStruct) GetInodes() uint64 {
	return diskPtr.totalInodes
}

func (diskPtr *diskStruct) GetFreeInodes() uint64 {
	return diskPtr.freeInodes
}

func (diskPtr *diskStruct) GetMountPoint() string {
	return diskPtr.mountPoint
}

func (diskPtr *diskStruct) GetDev() string {
	return diskPtr.device
}
