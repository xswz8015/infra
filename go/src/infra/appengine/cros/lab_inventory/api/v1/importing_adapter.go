// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package api

// Adapts the data defined by proto
// https://chromium.googlesource.com/infra/infra/+/refs/heads/master/go/src/infra/libs/skylab/inventory/device.proto
// to data defined by
// https://chromium.googlesource.com/chromiumos/infra/proto/src/lab/device.proto
import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	dev_proto "go.chromium.org/chromiumos/infra/proto/go/device"
	"go.chromium.org/chromiumos/infra/proto/go/lab"
	"go.chromium.org/chromiumos/infra/proto/go/manufacturing"
	"go.chromium.org/luci/common/errors"

	"infra/libs/skylab/inventory"
)

// A mapping from servo host name to servo host proto message.
type servoHostRegister map[string]*lab.ChromeOSDevice

func (r servoHostRegister) addServo(servo *lab.Servo) {
	hostname := servo.GetServoHostname()
	if hostname == "" {
		return
	}
	// FIXME(guocb) Try to load the labstation from datastore first. Otherwise
	// it may be overwritten and lost servos.
	if _, existing := r[hostname]; !existing {
		r[hostname] = &lab.ChromeOSDevice{
			Device: &lab.ChromeOSDevice_Labstation{
				Labstation: &lab.Labstation{
					Hostname: hostname,
				},
			},
		}
	}
	servoHost := r[hostname].GetLabstation()
	servoHost.Servos = append(servoHost.Servos, servo)
}

func (r servoHostRegister) getAllLabstations() []*lab.ChromeOSDevice {
	labstations := make([]*lab.ChromeOSDevice, 0, len(r))
	for _, v := range r {
		labstations = append(labstations, v)
	}
	return labstations
}

// ImportFromV1DutSpecs adapts v1 inventory data to v2 format.
func ImportFromV1DutSpecs(oldSpecs []*inventory.CommonDeviceSpecs) (devices []*lab.ChromeOSDevice, labstations []*lab.ChromeOSDevice, dutStates []*lab.DutState, err error) {
	servoHostRegister := servoHostRegister{}
	errs := errors.MultiError{}
	for _, olddata := range oldSpecs {
		if err := createCrosDevice(&devices, servoHostRegister, olddata); err != nil {
			errs = append(errs, errors.Annotate(err, "import spec for %s", olddata.GetHostname()).Err())
		}
		createDutState(&dutStates, olddata)
	}
	if len(errs) != 0 {
		err = errs
	}
	return devices, servoHostRegister.getAllLabstations(), dutStates, err
}

func createCrosDevice(results *[]*lab.ChromeOSDevice, servoHostRegister servoHostRegister, olddata *inventory.CommonDeviceSpecs) error {
	if osType := olddata.GetLabels().GetOsType(); osType == inventory.SchedulableLabels_OS_TYPE_LABSTATION {
		if err := createLabstation(servoHostRegister, olddata); err != nil {
			return err
		}
	} else {
		// Convert all other os_type (INVALID, ANDROID, CROS, MOBLAB, JETSTREAM)
		// to DUT.
		if err := createDut(results, servoHostRegister, olddata); err != nil {
			return err
		}
	}
	return nil
}

func importServo(servo *lab.Servo, key string, value string) error {
	switch key {
	case "servo_host":
		servo.ServoHostname = value
		if value == "" {
			return errors.Reason("invalid servo hostname: '%s'", value).Err()
		}
	case "servo_port":
		port, err := strconv.Atoi(value)
		if err != nil {
			return errors.Reason("invalid servo port: %s", value).Err()
		}
		servo.ServoPort = int32(port)
	case "servo_serial":
		servo.ServoSerial = value
	case "servo_type":
		servo.ServoType = value
	case "servo_setup":
		servoSetup := lab.ServoSetupType_SERVO_SETUP_REGULAR
		if ss, ok := lab.ServoSetupType_value["SERVO_SETUP_"+value]; ok {
			servoSetup = lab.ServoSetupType(ss)
		}
		servo.ServoSetup = servoSetup
	case "servo_topology":
		var topology *lab.ServoTopology
		if value != "" {
			jsonBytes, err := base64.StdEncoding.DecodeString(value)
			if err == nil {
				topology = &lab.ServoTopology{}
				json.Unmarshal(jsonBytes, topology)
			}
		}
		servo.ServoTopology = topology
	case "servo_fw_channel":
		servoFwChannel := lab.ServoFwChannel_SERVO_FW_STABLE
		if ss, ok := lab.ServoFwChannel_value["SERVO_FW_"+value]; ok {
			servoFwChannel = lab.ServoFwChannel(ss)
		}
		servo.ServoFwChannel = servoFwChannel
	}
	return nil
}

func importRpm(rpm *lab.RPM, key string, value string) {
	switch key {
	case "powerunit_hostname":
		rpm.PowerunitName = value
	case "powerunit_outlet":
		rpm.PowerunitOutlet = value
	}
}

// importAttributes imports the "Attributes" section of inventory v1 specs. It
// returns HWID, servo and rpm if they are in the section, ortherwise returns
// empty string or nil.
func importAttributes(attrs []*inventory.KeyValue) (string, *lab.Servo, *lab.RPM) {
	skipServo := false
	var hwid string
	var servo *lab.Servo
	var rpm *lab.RPM
	for _, attr := range attrs {
		value := attr.GetValue()
		switch key := attr.GetKey(); key {
		case "HWID":
			hwid = value
		case "servo_host", "servo_port", "servo_serial", "servo_type", "servo_setup", "servo_fw_channel":
			if servo == nil {
				servo = new(lab.Servo)
			}
			if err := importServo(servo, key, value); err != nil {
				skipServo = true
			}
		case "powerunit_hostname", "powerunit_outlet":
			if rpm == nil {
				rpm = new(lab.RPM)
			}
			importRpm(rpm, key, value)
		}
	}
	if skipServo {
		return hwid, nil, rpm
	}
	return hwid, servo, rpm
}

func getChameleonType(oldperi *inventory.Peripherals) []lab.ChameleonType {
	oldtypes := oldperi.GetChameleonType()
	newtype := make([]lab.ChameleonType, len(oldtypes))
	for i, v := range oldtypes {
		newtype[i] = lab.ChameleonType(v)
	}
	return newtype
}

func getAntennaConn(peri *inventory.Peripherals) lab.Wifi_AntennaConnection {
	if peri.GetConductive() {
		return lab.Wifi_CONN_CONDUCTIVE
	}
	return lab.Wifi_CONN_OTA
}

func getRouter(peri *inventory.Peripherals) lab.Wifi_Router {
	if peri.GetRouter_802_11Ax() {
		return lab.Wifi_ROUTER_802_11AX
	}
	return lab.Wifi_ROUTER_UNSPECIFIED
}

func getConnectedCamera(peri *lab.Peripherals, oldPeri *inventory.Peripherals) {
	if oldPeri.GetHuddly() {
		peri.ConnectedCamera = append(peri.ConnectedCamera, &lab.Camera{
			CameraType: lab.CameraType_CAMERA_HUDDLY,
		})
	}
	if oldPeri.GetPtzpro2() {
		peri.ConnectedCamera = append(peri.ConnectedCamera, &lab.Camera{
			CameraType: lab.CameraType_CAMERA_PTZPRO2,
		})
	}
}

func getDeviceConfigID(labels *inventory.SchedulableLabels) *dev_proto.ConfigId {
	return &dev_proto.ConfigId{
		PlatformId: &dev_proto.PlatformId{
			Value: strings.ToLower(labels.GetBoard()),
		},
		ModelId: &dev_proto.ModelId{
			Value: strings.ToLower(labels.GetModel()),
		},
		VariantId: &dev_proto.VariantId{
			// Use sku (an integer) instead of HwidSKU (a string).
			Value: strings.ToLower(labels.GetSku()),
		},
	}
}

func getPeripherals(l *inventory.SchedulableLabels) *lab.Peripherals {
	peripherals := l.GetPeripherals()
	capabilities := l.GetCapabilities()
	testHints := l.GetTestCoverageHints()
	p := lab.Peripherals{
		Chameleon: &lab.Chameleon{
			AudioBoard:           peripherals.GetAudioBoard(),
			ChameleonPeripherals: getChameleonType(peripherals),
		},
		Audio: &lab.Audio{
			AudioBox:   peripherals.GetAudioBox(),
			AudioCable: peripherals.GetAudioCable(),
			Atrus:      capabilities.GetAtrus(),
		},
		Wifi: &lab.Wifi{
			Wificell:    peripherals.GetWificell(),
			AntennaConn: getAntennaConn(peripherals),
			Router:      getRouter(peripherals),
		},
		Touch: &lab.Touch{
			Mimo: peripherals.GetMimo(),
		},
		Carrier:   parseCarrier(capabilities.GetCarrier()),
		Camerabox: peripherals.GetCamerabox(),
		CameraboxInfo: &lab.Camerabox{
			Facing: lab.Camerabox_Facing(peripherals.GetCameraboxFacing()),
			Light:  lab.Camerabox_Light(peripherals.GetCameraboxLight()),
		},
		Chaos:       testHints.GetChaosDut(),
		SmartUsbhub: peripherals.GetSmartUsbhub(),
	}
	getCables(&p, testHints)
	getConnectedCamera(&p, peripherals)
	return &p
}

func parseCarrier(c inventory.HardwareCapabilities_Carrier) string {
	return strings.ToLower(c.String()[len("CARRIER_"):])
}

func getCables(p *lab.Peripherals, testHints *inventory.TestCoverageHints) {
	if testHints.GetTestAudiojack() {
		p.Cable = append(p.Cable, &lab.Cable{
			Type: lab.CableType_CABLE_AUDIOJACK,
		})
	}
	if testHints.GetTestUsbaudio() {
		p.Cable = append(p.Cable, &lab.Cable{
			Type: lab.CableType_CABLE_USBAUDIO,
		})
	}
	if testHints.GetTestUsbprinting() {
		p.Cable = append(p.Cable, &lab.Cable{
			Type: lab.CableType_CABLE_USBPRINTING,
		})
	}
	if testHints.GetTestHdmiaudio() {
		p.Cable = append(p.Cable, &lab.Cable{
			Type: lab.CableType_CABLE_HDMIAUDIO,
		})
	}
}

func getPools(l *inventory.SchedulableLabels) []string {
	var pools []string
	for _, p := range l.GetCriticalPools() {
		pools = append(pools, inventory.SchedulableLabels_DUTPool_name[int32(p)])
	}
	for _, p := range l.GetSelfServePools() {
		pools = append(pools, p)
	}
	return pools
}

func getLicenses(l *inventory.SchedulableLabels) []*lab.License {
	var licenses []*lab.License
	for _, v := range l.Licenses {
		var t lab.LicenseType
		switch v.GetType() {
		case inventory.LicenseType_LICENSE_TYPE_WINDOWS_10_PRO:
			t = lab.LicenseType_LICENSE_TYPE_WINDOWS_10_PRO
		case inventory.LicenseType_LICENSE_TYPE_MS_OFFICE_STANDARD:
			t = lab.LicenseType_LICENSE_TYPE_MS_OFFICE_STANDARD
		default:
			t = lab.LicenseType_LICENSE_TYPE_UNSPECIFIED
		}
		licenses = append(licenses, &lab.License{
			Type:       t,
			Identifier: v.GetIdentifier(),
		})
	}
	return licenses
}

func createDut(devices *[]*lab.ChromeOSDevice, servoHostRegister servoHostRegister, olddata *inventory.CommonDeviceSpecs) error {
	hwid, servo, rpm := importAttributes(olddata.GetAttributes())

	peri := getPeripherals(olddata.GetLabels())
	if servo != nil {
		servo.ServoType = olddata.GetLabels().GetPeripherals().GetServoType()
		servo.ServoTopology = getServoTopology(olddata.GetLabels().GetPeripherals().GetServoTopology())
		peri.Servo = servo
		servoHostRegister.addServo(servo)
	}
	if rpm != nil {
		peri.Rpm = rpm
	}

	pools := getPools(olddata.GetLabels())
	licenses := getLicenses(olddata.GetLabels())
	newDut := lab.ChromeOSDevice{
		Id:              &lab.ChromeOSDeviceID{Value: olddata.GetId()},
		SerialNumber:    olddata.GetSerialNumber(),
		ManufacturingId: &manufacturing.ConfigID{Value: hwid},

		DeviceConfigId: getDeviceConfigID(olddata.GetLabels()),
		Device: &lab.ChromeOSDevice_Dut{
			Dut: &lab.DeviceUnderTest{
				Hostname:    olddata.GetHostname(),
				Peripherals: peri,
				Pools:       pools,
				Licenses:    licenses,
			},
		},
	}
	*devices = append(*devices, &newDut)
	return nil
}

func createLabstation(servoHostRegister servoHostRegister, olddata *inventory.CommonDeviceSpecs) error {
	hostname := olddata.GetHostname()
	hwid, _, rpm := importAttributes(olddata.GetAttributes())
	servoHost := &lab.ChromeOSDevice{
		Id:              &lab.ChromeOSDeviceID{Value: olddata.GetId()},
		SerialNumber:    olddata.GetSerialNumber(),
		DeviceConfigId:  getDeviceConfigID(olddata.GetLabels()),
		ManufacturingId: &manufacturing.ConfigID{Value: hwid},

		Device: &lab.ChromeOSDevice_Labstation{
			Labstation: &lab.Labstation{
				Hostname: hostname,
				Rpm:      rpm,
				Servos:   []*lab.Servo{},
				Pools:    getPools(olddata.GetLabels()),
			},
		},
	}
	// The one in servoHostRegister may have some servos registered.
	if s, existing := servoHostRegister[hostname]; existing {
		servoHost.GetLabstation().Servos = s.GetLabstation().GetServos()
	}
	servoHostRegister[hostname] = servoHost
	return nil
}

func getServoTopology(st *inventory.ServoTopology) *lab.ServoTopology {
	var t *lab.ServoTopology
	if st != nil {
		stString := proto.MarshalTextString(st)
		t = &lab.ServoTopology{}
		proto.UnmarshalText(stString, t)
	}
	return t
}

func boolToDutState(state bool) lab.PeripheralState {
	if state {
		return lab.PeripheralState_WORKING
	}
	return lab.PeripheralState_NOT_CONNECTED
}

func getServoState(peri *inventory.Peripherals) lab.PeripheralState {
	if peri == nil {
		return lab.PeripheralState_NOT_CONNECTED
	}
	if peri.GetServoState() != inventory.PeripheralState_UNKNOWN {
		return lab.PeripheralState(peri.GetServoState())
	}

	return boolToDutState(peri.GetServo())
}

func getCr50Phase(l *inventory.SchedulableLabels) lab.DutState_CR50Phase {
	switch l.GetCr50Phase() {
	case inventory.SchedulableLabels_CR50_PHASE_PVT:
		return lab.DutState_CR50_PHASE_PVT
	case inventory.SchedulableLabels_CR50_PHASE_PREPVT:
		return lab.DutState_CR50_PHASE_PREPVT
	}
	return lab.DutState_CR50_PHASE_INVALID
}

func getCr50Env(l *inventory.SchedulableLabels) lab.DutState_CR50KeyEnv {
	switch l.GetCr50RoKeyid() {
	case "prod":
		return lab.DutState_CR50_KEYENV_PROD
	case "dev":
		return lab.DutState_CR50_KEYENV_DEV
	}
	return lab.DutState_CR50_KEYENV_INVALID
}

func createDutState(states *[]*lab.DutState, olddata *inventory.CommonDeviceSpecs) {
	labels := olddata.GetLabels()
	if labels == nil {
		return
	}
	if ostype := labels.GetOsType(); ostype == inventory.SchedulableLabels_OS_TYPE_LABSTATION {
		return
	}
	peri := labels.GetPeripherals()
	*states = append(*states, &lab.DutState{
		Id:                     &lab.ChromeOSDeviceID{Value: olddata.GetId()},
		Servo:                  getServoState(peri),
		Chameleon:              boolToDutState(peri.GetChameleon()),
		AudioLoopbackDongle:    boolToDutState(peri.GetAudioLoopbackDongle()),
		WorkingBluetoothBtpeer: peri.GetWorkingBluetoothBtpeer(),
		Cr50Phase:              getCr50Phase(labels),
		Cr50KeyEnv:             getCr50Env(labels),
		StorageState:           lab.HardwareState(int32(peri.GetStorageState())),
		ServoUsbState:          lab.HardwareState(int32(peri.GetServoUsbState())),
		BatteryState:           lab.HardwareState(int32(peri.GetBatteryState())),
		RpmState:               lab.PeripheralState(int32(peri.GetRpmState())),
	})
}
