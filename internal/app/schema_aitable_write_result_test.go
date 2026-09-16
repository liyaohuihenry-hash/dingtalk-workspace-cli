// Copyright 2026 Alibaba Group
// SPDX-License-Identifier: Apache-2.0
package app

import "testing"

func TestCrossPlatformCoverageAITableWriteResultPublishesReadContract(t *testing.T) {
	leaf := executeShortcutSchemaQuery(t, "--cli-path", "aitable +record-write-result")
	result, _ := leaf["result"].(map[string]any)
	if result == nil {
		t.Fatal("reconciliation is missing its declared result contract")
	}
	dataSchema, _ := result["data_schema"].(map[string]any)
	properties := schemaContractMap(dataSchema["properties"])
	for _, key := range []string{"baseId", "tableId", "clientToken", "state", "recordIds"} {
		if properties[key] == nil {
			t.Errorf("read reconciliation result lacks %s: %#v", key, result)
		}
	}
	parameters := schemaContractMap(leaf["parameters"])
	for _, key := range []string{"base-id", "table-id", "client-token"} {
		if parameters[key] == nil {
			t.Errorf("read reconciliation parameter lacks %s", key)
		}
	}
}
