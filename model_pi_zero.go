// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Orange Pi pin out.

package orangepi

// import (
// 	"periph.io/x/conn/v3/gpio"
// 	"periph.io/x/conn/v3/pin"
// 	"periph.io/x/conn/v3/pin/pinreg"
// 	"periph.io/x/host/v3/allwinner"
// )

// var (
// 	PA1_1  pin.Pin    = pin.DC_IN // VCC 3v3 Ext
// 	PA1_2  pin.Pin    = pin.V5
// 	PA1_3  gpio.PinIO = allwinner.PA12
// 	PA1_4  pin.Pin    = pin.V5
// 	PA1_5  gpio.PinIO = allwinner.PA11
// 	PA1_6  pin.Pin    = pin.GROUND
// 	PA1_7  gpio.PinIO = allwinner.PA6
// 	PA1_8  gpio.PinIO = allwinner.PG6
// 	PA1_9  pin.Pin    = pin.GROUND
// 	PA1_10 gpio.PinIO = allwinner.PG7
// 	PA1_11 gpio.PinIO = allwinner.PA1
// 	PA1_12 gpio.PinIO = allwinner.PA7
// 	PA1_13 gpio.PinIO = allwinner.PA0
// 	PA1_14 pin.Pin    = pin.GROUND
// 	PA1_15 gpio.PinIO = allwinner.PA3
// 	PA1_16 gpio.PinIO = allwinner.PA19
// 	PA1_17 pin.Pin    = pin.DC_IN // VCC 3v3 Ext
// 	PA1_18 gpio.PinIO = allwinner.PA18
// 	PA1_19 gpio.PinIO = allwinner.PA15
// 	PA1_20 pin.Pin    = pin.GROUND
// 	PA1_21 gpio.PinIO = allwinner.PA16
// 	PA1_22 gpio.PinIO = allwinner.PA2
// 	PA1_23 gpio.PinIO = allwinner.PA14
// 	PA1_24 gpio.PinIO = allwinner.PA13
// 	PA1_25 pin.Pin    = pin.GROUND
// 	PA1_26 gpio.PinIO = allwinner.PA10

// 	FUN1_1  pin.Pin = pin.V5
// 	FUN1_2  pin.Pin = pin.GROUND
// 	FUN1_3  pin.Pin = gpio.INVALID       // USB-DM2
// 	FUN1_4  pin.Pin = gpio.INVALID       // USB-DP2
// 	FUN1_5  pin.Pin = gpio.INVALID       // USB-DM3
// 	FUN1_6  pin.Pin = gpio.INVALID       // USB-DP3
// 	FUN1_7  pin.Pin = allwinner.HP_RIGHT // LINEOUTR
// 	FUN1_8  pin.Pin = allwinner.HP_LEFT  // LINEOUTL
// 	FUN1_9  pin.Pin = gpio.INVALID       // TVOUT
// 	FUN1_10 pin.Pin = gpio.INVALID       // MBIAS, Bias Voltage output for mic
// 	FUN1_11 pin.Pin = allwinner.MIC_IN   // INPUT Analog Microphone pin (+)
// 	FUN1_12 pin.Pin = allwinner.MIC_GND  // INPUT Analog Microphone pin (-)
// 	FUN1_13 pin.Pin = allwinner.PL11     // IR-RX
// )

// // registerHeaders registers the headers for various Orange Pi boards. Currently
// // only Orange Pi Zero is supported.
// func registerHeaders_PiZero() error {
// 	// http://www.orangepi.org/html/hardWare/computerAndMicrocontrollers/details/Orange-Pi-Zero.html

// 	// 26pin expansion port
// 	if err := pinreg.Register("PA", [][]pin.Pin{
// 		{PA1_1, PA1_2},
// 		{PA1_3, PA1_4},
// 		{PA1_5, PA1_6},
// 		{PA1_7, PA1_8},
// 		{PA1_9, PA1_10},
// 		{PA1_11, PA1_12},
// 		{PA1_13, PA1_14},
// 		{PA1_15, PA1_16},
// 		{PA1_17, PA1_18},
// 		{PA1_19, PA1_20},
// 		{PA1_21, PA1_22},
// 		{PA1_23, PA1_24},
// 		{PA1_25, PA1_26},
// 	}); err != nil {
// 		return err
// 	}

// 	// 13pin function interface
// 	if err := pinreg.Register("FUN", [][]pin.Pin{
// 		{FUN1_1},
// 		{FUN1_2},
// 		{FUN1_3},
// 		{FUN1_4},
// 		{FUN1_5},
// 		{FUN1_6},
// 		{FUN1_7},
// 		{FUN1_8},
// 		{FUN1_9},
// 		{FUN1_10},
// 		{FUN1_11},
// 		{FUN1_12},
// 		{FUN1_13},
// 	}); err != nil {
// 		return err
// 	}

// 	return nil
// }
