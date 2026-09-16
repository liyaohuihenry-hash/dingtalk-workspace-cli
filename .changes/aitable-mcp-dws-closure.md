---
category: Fixed
---

- **AI Table write recovery**: `+table-copy` and `+record-batch-create` now generate one UUID v4 `clientToken` per create batch, submit it once, and reconcile uncertain receipts through MCP `get_record_write_result`. A fully recovered ID set is independently checked against record values before continuing. Unknown or partial results preserve the token/IDs and stop without replaying creation. The new read-only `+record-write-result` command exposes the recovery path in Help and Schema.
- **AI field readiness**: `+field-create` recognizes MCP `CREATE_FIELD_READBACK_PENDING`, keeps the created IDs, and uses the existing bounded exact-ID readback/resume path. `+field-run-ai` requires readable `aiConfig` before mutation and reports `submitted` only for complete per-field task receipts; submission is not computation completion.
- **Review follow-up**: duplicate-token errors and explicit `get_record_write_result` recovery hints no longer count as input rejections, even with `retryable=false`. They use the original token for read-only reconciliation and value verification; ordinary input rejections still stop immediately.
- Deployment dependency: the paired MCP branch must be deployed and the new read-only tool registered in the gateway. No `lippi-doc-notable` changes are included. Local mock/contract tests do not establish live service readiness; real-environment evaluation and long test suites remain pending user approval.
