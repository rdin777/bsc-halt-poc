# BSC Geth: Deterministic Network Halt (PoC)

*If this research helped you, please consider giving it a ⭐ Star.*


## 🚀 Stay Updated
Found this research useful?
* **Star ⭐** this repo to keep track of it.
* **Follow me** on GitHub for more DeFi security research.
* **Fork** it if you want to run your own experiments.

### ☕ Support the Research
If you appreciate the work and want to support further security research:

<img src="456.PNG" alt="Donate QR" width="200"/>

**Wallet Address (ETH/EVM):** 0xBDDD7973D0DE27B715A4A5cbdb87d0DF78757b3A 


### Overview
This repository contains a Proof of Concept for a **Critical** vulnerability found in the Binance Smart Chain (BSC) execution engine (Geth). 

### Vulnerability: Nil Pointer Dereference
The vulnerability exists in the `state` package during contract creation. When `CreateContract` is called for an address that does not exist in the current state database, the system fails to validate the account pointer, leading to a fatal runtime panic.

### Impact: Deterministic DoS
Since blockchain state transitions must be identical across all nodes, any malicious transaction triggering this path will cause **every node** in the network to crash simultaneously. This results in a full **Network Halt**.

### Reproduction
1. Integrate `state_panic_test.go` into the `core/state` directory of the BSC source code.
2. Run the test:
   ```bash
   go test -v -run TestBSCCreateContractPanic
   ```

### Status
Reported via Bugcrowd (Binance Program). Marked as "Not Applicable" due to being out of Web App scope. Published for educational and security awareness purposes.
