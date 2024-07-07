package main

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	PowerOffline = iota
	PowerOnline
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	getSystemPowerFunc = kernel32.NewProc("GetSystemPowerStatus")
)

type SystemPowerStatus struct {
	ACLineStatus       byte
	BatteryFlag        byte
	BatteryLifePercent byte
	SystemStatusFlag   byte
}

func IsPowerConnected() (bool, error) {
	var status SystemPowerStatus
	err := getSystemPowerStatus(&status)
	if err != nil {
		return false, err
	}
	if status.ACLineStatus == PowerOffline {
		return false, nil
	} else if status.ACLineStatus == PowerOnline {
		return true, nil
	} else {
		return false, errors.New("received unknown power status")
	}
}

func getSystemPowerStatus(status *SystemPowerStatus) error {
	r, _, err := getSystemPowerFunc.Call(uintptr(unsafe.Pointer(status)))
	if r == 0 {
		return err
	}
	return nil
}
