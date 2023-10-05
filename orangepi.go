// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Orange Pi pin out.

package orangepi

import (
	"errors"
	"fmt"
	"strings"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/host/v3/distro"
)

// Present return true if a Orange Pi board is detected.
func Present() bool {

	if isArm {
		// This works for the Orange Pi Zero, not sure if other Orange Pi boards
		// match the same DTModel prefix.
		println("GET distro.DTModel() -> ", distro.DTModel())
		// return strings.HasPrefix(distro.DTModel(), "OrangePi")
		return strings.Contains(strings.ToLower(distro.DTModel()), "orangepi") || strings.Contains(strings.ToLower(distro.DTModel()), "orange pi")
	}

	return false
}

const (
	boardZero   string = "Orange Pi Zero"    // + LTS (H2/H3 have identical pinouts)
	boardPC     string = "Orange Pi PC"      // H3
	boardPCPlus string = "Orange Pi PC Plus" // H3 Xunlong Orange Pi PC Plus
)

func registerHeaders(model string) error {
	// http://www.orangepi.org/html/hardWare/computerAndMicrocontrollers/details/Orange-Pi-Zero.html
	if strings.Contains(model, boardZero) {
		// registerHeaders_PiZero()
	} else if strings.Contains(model, boardPC) || strings.Contains(model, boardPCPlus) {
		RegisterHeaders_PiPC()
	}

	return nil
}

// driver implements periph.Driver.
type driver struct {
}

// String is the text representation of the board.
func (d *driver) String() string {
	return "orangepi"
}

// Prerequisites load drivers before the actual driver is loaded. For
// these boards, we do not need any prerequisites.
func (d *driver) Prerequisites() []string {
	return nil
}

// After this driver is loaded, we need to load generic Allwinner drivers
// for the GPIO pins which are identical on all Allwinner CPUs.
func (d *driver) After() []string {
	return []string{"allwinner-gpio", "allwinner-gpio-pl"}
}

// Init initializes the driver by checking its presence and if found, the
// driver will be registered.
func (d *driver) Init() (bool, error) {
	if !Present() {
		return false, errors.New("borad Orange Pi not detected")
	}

	model := distro.DTModel()
	if model == "<unknown>" {
		return true, fmt.Errorf("orangepi: failed to obtain model")
	}

	err := registerHeaders(model)
	return true, err
}

// init register the driver.
func init() {
	if isArm {
		driverreg.MustRegister(&drv)
	}
}

var drv driver
