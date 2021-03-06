// Copyright 2016-2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smbios

import (
	"errors"
	"fmt"
	"strings"
)

// Much of this is auto-generated. If adding a new type, see README for instructions.

// SystemInformation is defined in DSP0134 7.2.
type SystemInformation struct {
	Table
	Manufacturer string     // 04h
	ProductName  string     // 05h
	Version      string     // 06h
	SerialNumber string     // 07h
	UUID         UUID       // 08h
	WakeupType   WakeupType // 18h
	SKUNumber    string     // 19h
	Family       string     // 1Ah
}

// UUID is defined in DSP0134 7.2.1.
type UUID [16]byte

// WakeupType is defined in DSP0134 7.2.2.
type WakeupType uint8

// NewSystemInformation parses a generic Table into SystemInformation.
func NewSystemInformation(t *Table) (*SystemInformation, error) {
	if t.Len() < 8 {
		return nil, errors.New("required fields missing")
	}
	si := &SystemInformation{Table: *t}
	if _, err := parseStruct(t, 0 /* off */, false /* complete */, si); err != nil {
		return nil, err
	}
	return si, nil
}

// ParseField parses UUD field within a table.
func (u *UUID) ParseField(t *Table, off int) (int, error) {
	ub, err := t.GetBytesAt(off, 16)
	if err != nil {
		return off, err
	}
	copy(u[:], ub)
	return off + 16, nil
}

func (si *SystemInformation) String() string {
	lines := []string{
		si.Header.String(),
		fmt.Sprintf("\tManufacturer: %s", si.Manufacturer),
		fmt.Sprintf("\tProduct Name: %s", si.ProductName),
		fmt.Sprintf("\tVersion: %s", si.Version),
		fmt.Sprintf("\tSerial Number: %s", si.SerialNumber),
	}
	if si.Len() >= 8 { // 2.1+
		lines = append(lines,
			fmt.Sprintf("\tUUID: %s", si.UUID),
			fmt.Sprintf("\tWake-up Type: %s", si.WakeupType),
		)
	}
	if si.Len() >= 0x19 { // 2.4+
		lines = append(lines,
			fmt.Sprintf("\tSKU Number: %s", si.SKUNumber),
			fmt.Sprintf("\tFamily: %s", si.Family),
		)
	}
	return strings.Join(lines, "\n")
}

func (u UUID) String() string {
	return fmt.Sprintf("%02X%02X%02X%02X-%02X%02X-%02X%02X-%02X%02X-%02X%02X%02X%02X%02X%02X",
		u[3], u[2], u[1], u[0],
		u[5], u[4],
		u[7], u[6],
		u[8], u[9],
		u[10], u[11], u[12], u[13], u[14], u[15],
	)
}

// WakeupType values are defined in DSP0134 7.2.2.
const (
	WakeupTypeReserved        WakeupType = 0x00 // Reserved
	WakeupTypeOther                      = 0x01 // Other
	WakeupTypeUnknown                    = 0x02 // Unknown
	WakeupTypeAPMTimer                   = 0x03 // APM Timer
	WakeupTypeModemRing                  = 0x04 // Modem Ring
	WakeupTypeLANRemote                  = 0x05 // LAN Remote
	WakeupTypePowerSwitch                = 0x06 // Power Switch
	WakeupTypePCIPME                     = 0x07 // PCI PME#
	WakeupTypeACPowerRestored            = 0x08 // AC Power Restored
)

func (v WakeupType) String() string {
	switch v {
	case WakeupTypeReserved:
		return "Reserved"
	case WakeupTypeOther:
		return "Other"
	case WakeupTypeUnknown:
		return "Unknown"
	case WakeupTypeAPMTimer:
		return "APM Timer"
	case WakeupTypeModemRing:
		return "Modem Ring"
	case WakeupTypeLANRemote:
		return "LAN Remote"
	case WakeupTypePowerSwitch:
		return "Power Switch"
	case WakeupTypePCIPME:
		return "PCI PME#"
	case WakeupTypeACPowerRestored:
		return "AC Power Restored"
	}
	return fmt.Sprintf("%d", v)
}
