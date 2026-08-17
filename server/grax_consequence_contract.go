package server

import (
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
)

const (
	graxConsequenceContractSchema = "brudo-grax-consequence-contract-v1"
	graxConsequenceWithhold       = "WITHHOLD"
	graxConsequenceBoundedOutput  = "BOUNDED_OUTPUT"
	graxConsequenceAdmitOperation = "ADMIT_OPERATION"
	graxReceiptSchema             = "brudo-grax-executor-receipt-v1"
)

var sha256HexPattern = regexp.MustCompile(`^[0-9a-fA-F]{64}$`)

var graxReceiptIdentityFields = []string{
	"operation",
	"executor_id",
	"source_sha256",
	"result_sha256",
	"evidence_sha256",
}

var graxReceiptHashFields = []string{
	"source_sha256",
	"result_sha256",
	"evidence_sha256",
}

// crystalDeterministicDecision is deliberately closed-world. The model may
// emit only the canonical consequence_class. Legacy/open-ended operation
// result wording is unknown and therefore withholds.
func crystalDeterministicDecision(response string) string {
	var value map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(response)), &value); err != nil {
		return "WITHHOLD_MALFORMED_STRUCTURED_OUTPUT"
	}
	if _, legacy := value["operation_result"]; legacy {
		return "WITHHOLD_UNKNOWN_CONSEQUENCE"
	}
	classification, ok := value["consequence_class"].(string)
	if !ok || classification == "" {
		return "WITHHOLD_UNKNOWN_CONSEQUENCE"
	}
	switch classification {
	case graxConsequenceWithhold:
		return "WITHHOLD_CONSEQUENCE"
	case graxConsequenceBoundedOutput:
		return "BOUNDED_OUTPUT_ONLY"
	case graxConsequenceAdmitOperation:
		if receipt, valid := graxExecutorReceipt(value["executor_receipt"]); valid {
			_ = receipt
			return "ADMIT_OPERATION_RECEIPT_BOUND"
		}
		return "WITHHOLD_MISSING_OR_INVALID_EXECUTOR_RECEIPT"
	default:
		return "WITHHOLD_UNKNOWN_CONSEQUENCE"
	}
}

func graxExecutorReceipt(raw any) (map[string]any, bool) {
	receipt, ok := raw.(map[string]any)
	if !ok || receipt["schema"] != graxReceiptSchema {
		return nil, false
	}
	for _, field := range graxReceiptIdentityFields {
		if value, present := receipt[field].(string); !present || strings.TrimSpace(value) == "" {
			return nil, false
		}
	}
	for _, field := range graxReceiptHashFields {
		value := receipt[field].(string)
		if !sha256HexPattern.MatchString(value) {
			return nil, false
		}
		if _, err := hex.DecodeString(value); err != nil {
			return nil, false
		}
	}
	return receipt, true
}
