// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmdhelp

var (
	// ListPageSizeDesc description for List PageSize
	ListPageSizeDesc string = `number of items to get. The service may return fewer than this value.If unspecified, at most 100 items will be returned.
The maximum value is 1000; values above 1000 will be coerced to 1000.`

	// AddSwitchLongDesc long description for AddSwitchCmd
	AddSwitchLongDesc string = `add switch by name.
./shivas switch add -j -f switch.json
Adds a Switch by reading a JSON file input.

./shivas switch add -i
Adds a Switch by reading input through interactive mode.`

	// UpdateSwitchLongDesc long description for UpdateSwitchCmd
	UpdateSwitchLongDesc string = `update switch by name.
./shivas switch set -j -f switch.json
Adds a Switch by reading a JSON file input.

./shivas switch set -i
Adds a Switch by reading input through interactive mode.`

	// ListSwitchLongDesc long description for ListSwitchCmd
	ListSwitchLongDesc string = `list all Switches

./shivas switch ls
Fetches 100 items and prints the output in table format

./shivas switch ls -n 50
Fetches 50 items and prints the output in table format

./shivas switch ls -json
Fetches 100 items and prints the output in JSON format

./shivas switch ls -n 50 -json
Fetches 50 items and prints the output in JSON format
`

	// SwitchFileText description for switch file input
	SwitchFileText string = `Path to a file containing Switch specification in JSON format.
This file must contain one Switch JSON message

Switch Example :
{
    "name": "switch-test-example",
    "capacityPort": 456
}

The protobuf definition of Switch is part of
https://chromium.googlesource.com/infra/infra/+/refs/heads/master/go/src/infra/unifiedfleet/api/v1/proto/peripherals.proto#71`

	// AddMachineLongDesc long description for AddMachineCmd
	AddMachineLongDesc string = `add Machine by name.
./shivas machine add -j -f machine.json
Adds a Machine by reading a JSON file input.

./shivas machine add -i
Adds a Machine by reading input through interactive mode.`

	// UpdateMachineLongDesc long description for UpdateMachineCmd
	UpdateMachineLongDesc string = `update Machine by name.
./shivas machine set -j -f machine.json
Adds a Machine by reading a JSON file input.

./shivas machine set -i
Adds a Machine by reading input through interactive mode.`

	// ListMachineLongDesc long description for ListMachineCmd
	ListMachineLongDesc string = `list all Machines

./shivas machine ls
Fetches 100 items and prints the output in table format

./shivas machine ls -n 50
Fetches 50 items and prints the output in table format

./shivas machine ls -json
Fetches 100 items and prints the output in JSON format

./shivas machine ls -n 50 -json
Fetches 50 items and prints the output in JSON format
`

	// MachineFileText description for machine file input
	MachineFileText string = `Path to a file containing Machine specification in JSON format.
This file must contain one Machine JSON message

Example Browser Lab Machine:
{
	"name": "machine-BROWSERLAB-example",
	"location": {
		"lab": "LAB_DATACENTER_MTV97",
		"rack": "RackName"
	},
	"chromeBrowserMachine": {
		"displayName": "ax105-34-230",
		"chromePlatform": "Dell R230",
		"nics": ["ax105-34-230-eth0"],
		"kvmInterface": {
			"kvm": "kvm.mtv97",
			"port": 34
		},
		"rpmInterface": {
			"rpm": "rpm.mtv97",
			"port": 65
		},
		"drac": "ax105-34-230-drac",
		"deploymentTicket": "846026"
	},
	"realm": "Browserlab"
}

Example OS Lab Machine:
{
	"name": "machine-OSLAB-example",
	"location": {
		"lab": "LAB_CHROME_ATLANTA",
		"aisle": "1",
		"row": "2",
		"rack": "Rack-42",
		"rackNumber": "42",
		"shelf": "3",
		"position": "5"
	},
	"chromeosMachine": {},
	"realm": "OSlab"
}

The protobuf definition of Machine is part of
https://chromium.googlesource.com/infra/infra/+/refs/heads/master/go/src/infra/unifiedfleet/api/v1/proto/machine.proto#19`

	// AddMachinelseLongDesc long description for AddMachinelseCmd
	AddMachinelseLongDesc string = `add MachineLSE by name.
./shivas machinelse add -j -f machinelse.json
Adds a MachineLSE by reading a JSON file input.

./shivas machinelse add -i
Adds a MachineLSE by reading input through interactive mode.`

	// UpdateMachinelseLongDesc long description for UpdateMachinelseCmd
	UpdateMachinelseLongDesc string = `update MachineLSE by name.
./shivas machinelse set -j -f machinelse.json
Updates a MachineLSE by reading a JSON file input.

./shivas machinelse set -i
Updates a MachineLSE by reading input through interactive mode.`

	// MachinelseFileText description for machinelse file input
	MachinelseFileText string = `Path to a file containing MachineLSE specification in JSON format.
This file must contain one MachineLSE JSON message

Example Browser Lab MachineLSE:
{
	"name": "A-Browser-MachineLSE-1",
	"machineLsePrototype": "browser-lab:vm",
	"hostname": "Dell Server 3200",
	"chromeBrowserMachineLse": {
		"vms": [
			{
				"name": "Windows8.0",
				"osVersion": {
					"value": "8.0",
					"description": "Windows Server"
				},
				"macAddress": "2.44.65.23",
				"hostname": "Windows8-lab1"
			},
			{
				"name": "Linux3.4",
				"osVersion": {
					"value": "3.4",
					"description": "Ubuntu Server"
				},
				"macAddress": "32.45.12.32",
				"hostname": "Ubuntu-lab2"
			}
		],
		"vmCapacity": 3
	}
	"machines": [
		"machine-DellServer-123"
	]
}

Example OS Lab MachineLSE:
{
	"name": "test-machinelse-OS",
	"machineLsePrototype": "acs-lab:qwer",
	"hostname": "ChromeOSMachineLSE",
	"chromeosMachineLse": {
		"deviceLse": {
			"chromeosDevice": {
				"id": {
					"value": "22565c49-93d2-43f4-ab4f-f9171e3f252d"
				},
				"manufacturingId": {
					"value": "samus"
				},
				"deviceConfigId": {
					"platformId": {
						"value": "samus"
					},
					"modelId": {
						"value": "veyron"
					},
					"variantId": {
						"value": "2.0"
					},
					"brandId": {
						"value": "chrome"
					}
				},
				"dut": {
					"hostname": "ChromeOSMachineLSE",
					"peripherals": {
						"servo": {
							"servoHostname": "v3",
							"servoPort": 23,
							"servoSerial": "3456",
							"servoType": "v3"
						},
						"chameleon": {
							"chameleonPeripherals": [
								"CHAMELEON_TYPE_VGA",
								"CHAMELEON_TYPE_HDMI"
							],
							"audioBoard": true
						},
						"rpm": {
							"powerunitOutlet": "0"
						},
						"connectedCamera": [
							{
								"cameraType": "CAMERA_HUDDLY"
							}
						],
						"audio": {
							"audioBox": true,
							"atrus": true
						},
						"wifi": {
							"wificell": true,
							"antennaConn": "CONN_CONDUCTIVE",
							"router": "ROUTER_802_11AX"
						},
						"touch": {
							"mimo": true
						},
						"carrier": "ATT",
						"camerabox": true,
						"chaos": true,
						"cable": [
							{
								"type": "CABLE_USBPRINTING"
							},
							{
								"type": "CABLE_HDMIAUDIO"
							}
						],
						"cameraboxInfo": {
							"facing": "FACING_BACK"
						}
					},
					"pools": [
						"ACS_LAB"
					]
				}
			},
			"rpmInterface": {
				"rpm": "rpm-1",
				"port": 23
			},
			"networkDeviceInterface": {
				"switch": "switch-1",
				"port": 23
			}
		}
	},
	"machines": [
		"test-machine-1"
	],
	"updateTime": "2020-06-18T17:25:48.123623554Z"
}

The protobuf definition of MachineLSE is part of
https://chromium.googlesource.com/infra/infra/+/refs/heads/master/go/src/infra/unifiedfleet/api/v1/proto/machine_lse.proto#24`

	// AddMachinelsePrototypeLongDesc long description for AddMachinelsePrototypeCmd
	AddMachinelsePrototypeLongDesc string = `add MachineLSEPrototype by name.
./shivas machinelseprototype add -j -f machinelseprototype.json
Adds a MachineLSEPrototype by reading a JSON file input.

./shivas machinelseprototype add -i
Adds a MachineLSEPrototype by reading input through interactive mode.`

	// UpdateMachinelsePrototypeLongDesc long description for UpdateMachinelsePrototypeCmd
	UpdateMachinelsePrototypeLongDesc string = `update MachineLSEPrototype by name.
./shivas machinelseprototype set -j -f machinelseprototype.json
Updates a MachineLSEPrototype by reading a JSON file input.

./shivas machinelseprototype set -i
Updates a MachineLSEPrototype by reading input through interactive mode.`

	// ListMachinelsePrototypeLongDesc long description for ListMachinelsePrototypeCmd
	ListMachinelsePrototypeLongDesc string = `list all MachineLSEPrototypes

./shivas machinelseprototype ls
Fetches 100 items and prints the output in table format

./shivas machinelseprototype ls -n 50
Fetches 50 items and prints the output in table format

./shivas machinelseprototype ls -lab acs
Fetches only ACS lab items and prints the output in table format

./shivas machinelseprototype ls -json
Fetches 100 items and prints the output in JSON format
`

	// MachinelsePrototypeFileText description for machinelseprototype file input
	MachinelsePrototypeFileText string = `Path to a file containing MachineLSEPrototype specification in JSON format.
This file must contain one MachineLSEPrototype JSON message

Example MachineLSEPrototype:
{
	"name": "browser-lab:vm",
	"peripheralRequirements": [{
		"peripheralType": "PERIPHERAL_TYPE_SWITCH",
		"min": 5,
		"max": 7
	}],
	"occupiedCapacityRu": 32,
	"virtualRequirements": [{
		"virtualType": "VIRTUAL_TYPE_VM",
		"min": 3,
		"max": 4
	}]
}

The protobuf definition of MachineLSEPrototype is part of
https://chromium.googlesource.com/infra/infra/+/refs/heads/master/go/src/infra/unifiedfleet/api/v1/proto/lse_prototype.proto#29`
)
