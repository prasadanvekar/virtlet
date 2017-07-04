/*
Copyright 2016 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cni

import (
	"github.com/containernetworking/cni/pkg/types/current"
)

func BytesToResult(data []byte) (*current.Result, error) {
	if len(data) == 0 {
		return nil, nil
	}
	r, err := current.NewResult(data)
	if err != nil {
		return nil, err
	}
	return r.(*current.Result), err
}

// GetPodIP retrieves the IP address of the pod as a string. It uses
// the first IPv4 address if finds. If it fails to determine the pod
// IP or the result argument is nil, it returns an empty string.
func GetPodIP(result *current.Result) string {
	if result == nil {
		return ""
	}
	for _, ip := range result.IPs {
		if ip.Version == "4" {
			return ip.Address.IP.String()
		}
	}
	return ""
}
