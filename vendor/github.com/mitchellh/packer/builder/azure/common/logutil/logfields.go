// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in builder/azure for license information.

package logutil

import "fmt"

type Fields map[string]interface{}

func (f Fields) String() string {
	var s string
	for k, v := range f {
		if sv, ok := v.(string); ok {
			v = fmt.Sprintf("%q", sv)
		}
		s += fmt.Sprintf(" %s=%v", k, v)
	}
	return s
}