/*
 * Copyright 2026 The HAMi Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package brgpu

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Project-HAMi/biren-device-plugin/pkg/utils"

	"github.com/sirupsen/logrus"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	DeviceType     = "Biren"
	HandshakeAnnos = "hami.io/node-handshake-biren"
	RegisterAnnos  = "hami.io/node-biren-register"

	Memory = 65024
)

var (
	nodeName = flag.String("node_name", os.Getenv("NODE_NAME"), "node name")
)

func GetNodeName() string {
	return *nodeName
}

func RegisterHAMi(devs []*pluginapi.Device) error {
	apiDevices := make([]*utils.DeviceInfo, 0, len(devs))
	// hami currently believes that the index starts from 0 and is continuous.
	for i, dev := range devs {
		device := &utils.DeviceInfo{
			Index:  uint(i),
			ID:     dev.ID,
			Type:   DeviceType,
			Health: true,
			Count:  1,
			Devmem: Memory,
		}
		apiDevices = append(apiDevices, device)
	}
	return register(apiDevices)
}

func RegisterHAMiWithRawDevice(devs DevicesInfoList) error {
	apiDevices := make([]*utils.DeviceInfo, 0, len(devs))
	// hami currently believes that the index starts from 0 and is continuous.
	index := 0
	for i, dev := range devs {
		for _, ins := range dev.Instances {
			device := &utils.DeviceInfo{
				Index:  uint(i),
				ID:     ins.CardID,
				Type:   dev.Name,
				Health: true,
				Count:  1,
				Devmem: int32(ins.Memory / 1024 / 1024),
			}
			apiDevices = append(apiDevices, device)
			index++
		}
	}
	return register(apiDevices)
}

func register(apiDevices []*utils.DeviceInfo) error {
	annos := make(map[string]string)
	annos[RegisterAnnos] = utils.MarshalNodeDevices(apiDevices)
	annos[HandshakeAnnos] = "Reported_" + time.Now().Format("2026.01.02 15:04:05")
	node, err := utils.GetNode(*nodeName)
	if err != nil {
		return fmt.Errorf("get node %s error: %v", nodeName, err)
	}
	err = utils.PatchNodeAnnotations(node, annos)
	if err != nil {
		return fmt.Errorf("patch node %s annotations error: %v", nodeName, err)
	}
	logrus.Debugf("patch node %s annotations: %v", nodeName, annos)
	return nil
}
