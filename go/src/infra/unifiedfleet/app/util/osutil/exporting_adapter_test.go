// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file

package osutil

import (
	"testing"

	"github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"

	"infra/libs/skylab/inventory"
	ufspb "infra/unifiedfleet/api/v1/models"
	device "infra/unifiedfleet/api/v1/models/chromeos/device"
	chromeosLab "infra/unifiedfleet/api/v1/models/chromeos/lab"
	manufacturing "infra/unifiedfleet/api/v1/models/chromeos/manufacturing"
)

var servoInV2 = chromeosLab.Servo{
	ServoHostname: "test_servo",
	ServoPort:     int32(9999),
	ServoSerial:   "test_servo_serial",
	ServoType:     "v3",
	ServoSetup:    chromeosLab.ServoSetupType_SERVO_SETUP_DUAL_V4,
	ServoTopology: &chromeosLab.ServoTopology{
		Main: &chromeosLab.ServoTopologyItem{
			Type:         "servo_v4",
			SysfsProduct: "Servo V4",
			Serial:       "C1903145591",
			UsbHubPort:   "6.4.1",
		},
		Children: []*chromeosLab.ServoTopologyItem{
			{
				Type:         "ccd_cr50",
				SysfsProduct: "Cr50",
				Serial:       "0681D03A-92DCCD64",
				UsbHubPort:   "6.4.2",
			},
			{
				Type:         "c2d2",
				SysfsProduct: "C2D2",
				Serial:       "0681D03A-YYYYYYYY",
				UsbHubPort:   "6.4.3",
			},
		},
	},
}

var machine = ufspb.Machine{
	Name:         "test_dut",
	SerialNumber: "test_serial",
	Device: &ufspb.Machine_ChromeosMachine{
		ChromeosMachine: &ufspb.ChromeOSMachine{
			Hwid: "test_hwid",
		},
	},
}

var lse = ufspb.MachineLSE{
	Name:     "test_host",
	Hostname: "test_host",
	Machines: []string{"test_dut"},
	Lse: &ufspb.MachineLSE_ChromeosMachineLse{
		ChromeosMachineLse: &ufspb.ChromeOSMachineLSE{
			ChromeosLse: &ufspb.ChromeOSMachineLSE_DeviceLse{
				DeviceLse: &ufspb.ChromeOSDeviceLSE{
					Device: &ufspb.ChromeOSDeviceLSE_Dut{
						Dut: &chromeosLab.DeviceUnderTest{
							Hostname: "test_host",
							Pools:    []string{"DUT_POOL_QUOTA", "hotrod"},
							Peripherals: &chromeosLab.Peripherals{
								Servo: &servoInV2,
								Chameleon: &chromeosLab.Chameleon{
									ChameleonPeripherals: []chromeosLab.ChameleonType{
										chromeosLab.ChameleonType_CHAMELEON_TYPE_BT_BLE_HID,
										chromeosLab.ChameleonType_CHAMELEON_TYPE_BT_PEER,
									},
									AudioBoard: true,
								},
								Rpm: &chromeosLab.OSRPM{
									PowerunitName:   "test_power_unit_name",
									PowerunitOutlet: "test_power_unit_outlet",
								},
								ConnectedCamera: []*chromeosLab.Camera{
									{
										CameraType: chromeosLab.CameraType_CAMERA_HUDDLY,
									},
									{
										CameraType: chromeosLab.CameraType_CAMERA_PTZPRO2,
									},
								},
								Audio: &chromeosLab.Audio{
									AudioBox:   true,
									AudioCable: true,
									Atrus:      true,
								},
								Wifi: &chromeosLab.Wifi{
									Wificell:    true,
									AntennaConn: chromeosLab.Wifi_CONN_CONDUCTIVE,
									Router:      chromeosLab.Wifi_ROUTER_802_11AX,
								},
								Touch: &chromeosLab.Touch{
									Mimo: true,
								},
								Carrier:   "att",
								Camerabox: true,
								CameraboxInfo: &chromeosLab.Camerabox{
									Facing: chromeosLab.Camerabox_FACING_BACK,
									Light:  chromeosLab.Camerabox_LIGHT_LED,
								},
								Chaos: true,
								Cable: []*chromeosLab.Cable{
									{
										Type: chromeosLab.CableType_CABLE_AUDIOJACK,
									},
									{
										Type: chromeosLab.CableType_CABLE_USBAUDIO,
									},
									{
										Type: chromeosLab.CableType_CABLE_USBPRINTING,
									},
									{
										Type: chromeosLab.CableType_CABLE_HDMIAUDIO,
									},
								},
								SmartUsbhub: true,
							},
							Licenses: []*chromeosLab.License{
								{
									Type:       chromeosLab.LicenseType_LICENSE_TYPE_WINDOWS_10_PRO,
									Identifier: "my-windows-identifier-A001",
								},
								{
									Type:       chromeosLab.LicenseType_LICENSE_TYPE_MS_OFFICE_STANDARD,
									Identifier: "my-office-identifier-B002",
								},
							},
						},
					},
				},
			},
		},
	},
}

var devUFSState = chromeosLab.DutState{
	Id: &chromeosLab.ChromeOSDeviceID{
		Value: "test_dut",
	},
	Servo:                  chromeosLab.PeripheralState_BROKEN,
	Chameleon:              chromeosLab.PeripheralState_WORKING,
	AudioLoopbackDongle:    chromeosLab.PeripheralState_NOT_CONNECTED,
	WorkingBluetoothBtpeer: 3,
	Cr50Phase:              chromeosLab.DutState_CR50_PHASE_PVT,
	Cr50KeyEnv:             chromeosLab.DutState_CR50_KEYENV_PROD,
	StorageState:           chromeosLab.HardwareState_HARDWARE_NORMAL,
	ServoUsbState:          chromeosLab.HardwareState_HARDWARE_NEED_REPLACEMENT,
	RpmState:               chromeosLab.PeripheralState_WORKING,
}

var labstationMachine = ufspb.Machine{
	Name:         "test_labstation",
	SerialNumber: "labstation_serial",
	Device: &ufspb.Machine_ChromeosMachine{
		ChromeosMachine: &ufspb.ChromeOSMachine{
			Hwid: "labstation_hwid",
		},
	},
}

var labstationLSE = ufspb.MachineLSE{
	Name:     "test_labstation_host",
	Hostname: "test_labstation_host",
	Machines: []string{"test_labstation"},
	Lse: &ufspb.MachineLSE_ChromeosMachineLse{
		ChromeosMachineLse: &ufspb.ChromeOSMachineLSE{
			ChromeosLse: &ufspb.ChromeOSMachineLSE_DeviceLse{
				DeviceLse: &ufspb.ChromeOSDeviceLSE{
					Device: &ufspb.ChromeOSDeviceLSE_Labstation{
						Labstation: &chromeosLab.Labstation{
							Hostname: "test_labstation_host",
							Pools:    []string{"labstation_main"},
							Servos:   []*chromeosLab.Servo{},
							Rpm: &chromeosLab.OSRPM{
								PowerunitName:   "test_power_unit_name",
								PowerunitOutlet: "test_power_unit_outlet2",
							},
						},
					},
				},
			},
		},
	},
}

var labstationDevConfig = device.Config{
	Id: &device.ConfigId{
		PlatformId: &device.PlatformId{
			Value: "guado",
		},
		ModelId: &device.ModelId{
			Value: "test_model",
		},
		VariantId: &device.VariantId{
			Value: "",
		},
	},
}

var labstationManufacturingconfig = manufacturing.ManufacturingConfig{
	ManufacturingId: &manufacturing.ConfigID{
		Value: "labstation_hwid",
	},
}

var data = ufspb.ChromeOSDeviceData{
	LabConfig: &lse,
	DutState:  &devUFSState,
	DeviceConfig: &device.Config{
		Id: &device.ConfigId{
			PlatformId: &device.PlatformId{
				Value: "coral",
			},
			ModelId: &device.ModelId{
				Value: "test_model",
			},
			VariantId: &device.VariantId{
				Value: "test_variant",
			},
		},
		FormFactor: device.Config_FORM_FACTOR_CHROMEBASE,
		GpuFamily:  "test_gpu",
		Graphics:   device.Config_GRAPHICS_GLE,
		HardwareFeatures: []device.Config_HardwareFeature{
			device.Config_HARDWARE_FEATURE_BLUETOOTH,
			device.Config_HARDWARE_FEATURE_DETACHABLE_KEYBOARD,
			device.Config_HARDWARE_FEATURE_FINGERPRINT,
			device.Config_HARDWARE_FEATURE_TOUCHSCREEN,
		},
		Power:   device.Config_POWER_SUPPLY_AC_ONLY,
		Storage: device.Config_STORAGE_SSD,
		VideoAccelerationSupports: []device.Config_VideoAcceleration{
			device.Config_VIDEO_ACCELERATION_ENC_H264,
			device.Config_VIDEO_ACCELERATION_ENC_VP8,
			device.Config_VIDEO_ACCELERATION_ENC_VP9,
		},
		Cpu: device.Config_ARM64,
	},
	HwidData: &ufspb.HwidData{
		Sku:     "test_sku",
		Variant: "test_variant",
	},
	ManufacturingConfig: &manufacturing.ManufacturingConfig{
		ManufacturingId: &manufacturing.ConfigID{
			Value: "test_hwid",
		},
		DevicePhase: manufacturing.ManufacturingConfig_PHASE_DVT,
	},
	Machine: &machine,
}

const dutTextProto = `
common {
	attributes {
		key: "HWID",
		value: "test_hwid",
	}
	attributes {
		key: "powerunit_hostname",
		value: "test_power_unit_name",
	}
	attributes {
		key: "powerunit_outlet",
		value: "test_power_unit_outlet",
	}
	attributes {
		key: "serial_number"
		value: "test_serial"
	}
	attributes {
		key: "servo_host"
		value: "test_servo"
	}
	attributes {
		key: "servo_port"
		value: "9999"
	}
	attributes {
		key: "servo_serial",
		value: "test_servo_serial",
	}
	attributes {
		key: "servo_type",
		value: "v3",
	}
	attributes {
		key: "servo_setup",
		value: "DUAL_V4",
	}
	hostname: "test_host"
	hwid: "test_hwid"
	id: "test_dut"
	serial_number: "test_serial"
	labels {
		arc: true
		board: "coral"
		brand: ""
		capabilities {
			atrus: true
			bluetooth: true
			carrier: CARRIER_ATT
			detachablebase: true
			fingerprint: true
			gpu_family: "test_gpu"
			graphics: "gles"
			power: "AC_only"
			storage: "ssd"
			touchscreen: true
			video_acceleration: VIDEO_ACCELERATION_ENC_H264
			video_acceleration: VIDEO_ACCELERATION_ENC_VP8
			video_acceleration: VIDEO_ACCELERATION_ENC_VP9
		}
		cr50_phase: CR50_PHASE_PVT
		cts_abi: CTS_ABI_ARM
		cts_cpu: CTS_CPU_ARM
		cr50_ro_keyid: "prod"
		ec_type: EC_TYPE_CHROME_OS
		hwid_sku: "test_sku"
		licenses: {
			type: 1
			identifier: "my-windows-identifier-A001"
		}
		licenses: {
			type: 2
			identifier: "my-office-identifier-B002"
		}
		model: "test_model"
		os_type: OS_TYPE_CROS
		sku: "test_variant"
		peripherals {
			audio_board: true
			audio_box: true
			audio_cable: true
			audio_loopback_dongle: false
			chameleon: true
			chameleon_type: CHAMELEON_TYPE_BT_BLE_HID
			chameleon_type: CHAMELEON_TYPE_BT_PEER
			conductive: true
			huddly: true
			mimo: true
			ptzpro2: true
			camerabox: true
			camerabox_facing: CAMERABOX_FACING_BACK
			camerabox_light: CAMERABOX_LIGHT_LED
			servo: true
			servo_topology: {
				main: {
					usb_hub_port: "6.4.1"
					serial: "C1903145591"
					type: "servo_v4"
					sysfs_product: "Servo V4"
				}
				children: {
					usb_hub_port: "6.4.2"
					serial: "0681D03A-92DCCD64"
					type: "ccd_cr50"
					sysfs_product: "Cr50"
				}
				children: {
					usb_hub_port: "6.4.3"
					serial: "0681D03A-YYYYYYYY"
					type: "c2d2"
					sysfs_product: "C2D2"
				}
			  }
			servo_state: BROKEN
			servo_type: "v3"
			rpm_state: WORKING
			smart_usbhub: true
			storage_state: HARDWARE_NORMAL,
			servo_usb_state: HARDWARE_NEED_REPLACEMENT,
			wificell: true
			router_802_11ax: true
			working_bluetooth_btpeer: 3
		}
		phase: PHASE_DVT
		platform: "coral"
		test_coverage_hints {
			chaos_dut: true
			hangout_app: true
			meet_app: true
			test_audiojack: true
			test_hdmiaudio: true
			test_usbaudio: true
			test_usbprinting: true
		}
		variant: "test_variant"
		critical_pools: DUT_POOL_QUOTA
		self_serve_pools: "hotrod"
		wifi_chip: ""
	}
}
`

const labstationProtoFromV2 = `
common {
	attributes {
		key: "HWID",
		value: "labstation_hwid",
	}
	attributes {
		key: "powerunit_hostname",
		value: "test_power_unit_name",
	}
	attributes {
		key: "powerunit_outlet",
		value: "test_power_unit_outlet2",
	}
	attributes {
		key: "serial_number"
		value: "labstation_serial"
	}
	id: "test_labstation"
	hostname: "test_labstation_host"
	serial_number: "labstation_serial"
	labels {
		arc: false
		os_type: OS_TYPE_LABSTATION
		self_serve_pools: "labstation_main"
		model: "test_model"
		board: "guado"
		brand: ""
		sku: ""
        capabilities {
          atrus: false
          bluetooth: false
          carrier: CARRIER_INVALID
          detachablebase: false
          fingerprint: false
          flashrom: false
          gpu_family: ""
          graphics: ""
          hotwording: false
          internal_display: false
          lucidsleep: false
          modem: ""
          power: "AC_only"
          storage: ""
          telephony: ""
          webcam: false
          touchpad: false
          touchscreen: false
        }
        cr50_phase: CR50_PHASE_INVALID
        cr50_ro_keyid: ""
        cr50_ro_version: ""
        cr50_rw_keyid: ""
        cr50_rw_version: ""
        ec_type: EC_TYPE_INVALID
        hwid_sku: ""
		peripherals {
          audio_board: false
          audio_box: false
          audio_loopback_dongle: false
          chameleon: false
          chameleon_type: CHAMELEON_TYPE_INVALID
          conductive: false
          huddly: false
          mimo: false
          servo: false
          servo_state: UNKNOWN
          smart_usbhub: false
          stylus: false
          camerabox: false
          wificell: false
          router_802_11ax: false
		}
		platform:""
        test_coverage_hints {
            chaos_dut: false
            chaos_nightly: false
            chromesign: false
            hangout_app: false
            meet_app: false
            recovery_test: false
            test_audiojack: false
            test_hdmiaudio: false
            test_usbaudio: false
            test_usbprinting: false
            usb_detect: false
            use_lid: false
        }
        wifi_chip: ""
	}
}
`

func TestAdaptToV1DutSpec(t *testing.T) {
	t.Parallel()

	Convey("Verify V2 => V1", t, func() {
		var d1 inventory.DeviceUnderTest
		err := proto.UnmarshalText(dutTextProto, &d1)
		So(err, ShouldBeNil)
		s1, err := inventory.WriteLabToString(&inventory.Lab{
			Duts: []*inventory.DeviceUnderTest{&d1},
		})
		So(err, ShouldBeNil)
		dataCopy := proto.Clone(&data).(*ufspb.ChromeOSDeviceData)

		Convey("empty input", func() {
			_, err := AdaptToV1DutSpec(&ufspb.ChromeOSDeviceData{})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "chromeosdevicedata is nil to adapt")
		})
		Convey("empty hwid data", func() {
			dataCopy.HwidData = nil
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(d.GetCommon().GetHostname(), ShouldEqual, "test_host")
		})
		Convey("empty device config", func() {
			dataCopy.DeviceConfig = nil
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(d.GetCommon().GetHostname(), ShouldEqual, "test_host")
		})
		Convey("empty manufacturing config", func() {
			dataCopy.ManufacturingConfig = nil
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(d.GetCommon().GetHostname(), ShouldEqual, "test_host")
		})
		Convey("may os_type", func() {
			board := "fizz-moblab"
			osType := inventory.SchedulableLabels_OS_TYPE_MOBLAB
			d := proto.Clone(&d1).(*inventory.DeviceUnderTest)
			d.GetCommon().GetLabels().Board = &board
			d.GetCommon().GetLabels().Platform = &board
			d.GetCommon().GetLabels().OsType = &osType
			d.GetCommon().GetLabels().Arc = &falseValue
			s1, err := inventory.WriteLabToString(&inventory.Lab{
				Duts: []*inventory.DeviceUnderTest{d},
			})
			So(err, ShouldBeNil)

			//dataCopy.LabConfig = proto.Clone(data.LabConfig).(*ufspb.MachineLSE)
			dataCopy.GetDeviceConfig().GetId().GetPlatformId().Value = board
			d2, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			s2, err := inventory.WriteLabToString(&inventory.Lab{
				Duts: []*inventory.DeviceUnderTest{d2},
			})
			So(s1, ShouldEqual, s2)
		})
		Convey("servo_state is UNKNOWN/false by default", func() {
			dataCopy.DutState = &chromeosLab.DutState{}
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(*d.GetCommon().GetLabels().GetPeripherals().ServoState, ShouldEqual, inventory.PeripheralState_UNKNOWN)
			So(*d.GetCommon().GetLabels().GetPeripherals().Servo, ShouldBeFalse)
		})
		Convey("servo_state is broken", func() {
			dataCopy.DutState = &chromeosLab.DutState{}
			dataCopy.DutState.Servo = chromeosLab.PeripheralState_BROKEN
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(*d.GetCommon().GetLabels().GetPeripherals().ServoState,
				ShouldEqual,
				inventory.PeripheralState_BROKEN)
			So(*d.GetCommon().GetLabels().GetPeripherals().Servo, ShouldEqual, true)
		})
		Convey("servo_state is wrong_config", func() {
			dataCopy.DutState = &chromeosLab.DutState{}
			dataCopy.DutState.Servo = chromeosLab.PeripheralState_WRONG_CONFIG
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(*d.GetCommon().GetLabels().GetPeripherals().ServoState,
				ShouldEqual,
				inventory.PeripheralState_WRONG_CONFIG)
			So(*d.GetCommon().GetLabels().GetPeripherals().Servo, ShouldEqual, true)
		})
		Convey("servo_state is working", func() {
			dataCopy.DutState = &chromeosLab.DutState{}
			dataCopy.DutState.Servo = chromeosLab.PeripheralState_WORKING
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(*d.GetCommon().GetLabels().GetPeripherals().ServoState,
				ShouldEqual,
				inventory.PeripheralState_WORKING)
			So(*d.GetCommon().GetLabels().GetPeripherals().Servo, ShouldEqual, true)
		})
		Convey("servo_state is not_connected", func() {
			dataCopy.DutState = &chromeosLab.DutState{}
			dataCopy.DutState.Servo = chromeosLab.PeripheralState_NOT_CONNECTED
			d, err := AdaptToV1DutSpec(dataCopy)
			So(err, ShouldBeNil)
			So(*d.GetCommon().GetLabels().GetPeripherals().ServoState,
				ShouldEqual,
				inventory.PeripheralState_NOT_CONNECTED)
			So(*d.GetCommon().GetLabels().GetPeripherals().Servo, ShouldEqual, false)
		})
		Convey("happy path", func() {
			d, err := AdaptToV1DutSpec(&data)
			So(err, ShouldBeNil)
			s, err := inventory.WriteLabToString(&inventory.Lab{
				Duts: []*inventory.DeviceUnderTest{d},
			})
			So(err, ShouldBeNil)
			So(proto.Equal(&d1, d), ShouldBeTrue)
			So(s1, ShouldEqual, s)
		})
	})

	Convey("Verify labstation v2 => v1", t, func() {
		var labstation inventory.DeviceUnderTest
		err := proto.UnmarshalText(labstationProtoFromV2, &labstation)
		So(err, ShouldBeNil)

		extLabstaion := ufspb.ChromeOSDeviceData{
			LabConfig:           &labstationLSE,
			Machine:             &labstationMachine,
			DeviceConfig:        &labstationDevConfig,
			ManufacturingConfig: &labstationManufacturingconfig,
		}
		d, err := AdaptToV1DutSpec(&extLabstaion)
		So(err, ShouldBeNil)

		s, err := inventory.WriteLabToString(&inventory.Lab{
			Duts: []*inventory.DeviceUnderTest{d},
		})
		So(err, ShouldBeNil)
		strLabstation, err := inventory.WriteLabToString(&inventory.Lab{
			Duts: []*inventory.DeviceUnderTest{&labstation},
		})
		So(err, ShouldBeNil)
		So(s, ShouldEqual, strLabstation)
	})
}
