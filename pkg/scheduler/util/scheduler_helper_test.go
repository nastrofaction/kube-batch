/*
Copyright 2019 The Kubernetes Authors.

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

package util

import (
	"github.com/kubernetes-sigs/kube-batch/pkg/scheduler/api"
	"reflect"
	"testing"
)

func TestSelectBestNode(t *testing.T) {
	cases := []struct {
		NodeScores map[float64][]*api.NodeInfo
		// Expected node is one of ExpectedNodes
		ExpectedNodes []*api.NodeInfo
	}{
		{
			NodeScores: map[float64][]*api.NodeInfo{
				1.0: {&api.NodeInfo{Name: "node1"}, &api.NodeInfo{Name: "node2"}},
				2.0: {&api.NodeInfo{Name: "node3"}, &api.NodeInfo{Name: "node4"}},
			},
			ExpectedNodes: []*api.NodeInfo{{Name: "node3"}, {Name: "node4"}},
		},
		{
			NodeScores: map[float64][]*api.NodeInfo{
				1.0: {&api.NodeInfo{Name: "node1"}, &api.NodeInfo{Name: "node2"}},
				3.0: {&api.NodeInfo{Name: "node3"}},
				2.0: {&api.NodeInfo{Name: "node4"}, &api.NodeInfo{Name: "node5"}},
			},
			ExpectedNodes: []*api.NodeInfo{{Name: "node3"}},
		},
	}

	oneOf := func(node *api.NodeInfo, nodes []*api.NodeInfo) bool {
		for _, v := range nodes {
			if reflect.DeepEqual(node, v) {
				return true
			}
		}
		return false
	}
	for i, test := range cases {
		result := SelectBestNode(test.NodeScores)
		if !oneOf(result, test.ExpectedNodes) {
			t.Errorf("Failed test case #%d, expected: %#v, got %#v", i, test.ExpectedNodes, result)
		}
	}
}
