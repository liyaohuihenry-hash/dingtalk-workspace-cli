---
category: Fixed
---

- **AI Table form sharing** now returns the server-verified `shareFormUuid`, status, cover, and CP synchronization state from share reads and updates. Updates report partial failures when readback or CP projection is incomplete, and DWS no longer treats an unverified share mutation as a completed sharing workflow.
