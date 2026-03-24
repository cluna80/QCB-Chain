# QuarkCharmBit (QCB) Chain

**Freedom at the Charm Level**  
Powering the Arkadina Civilization

## Chain Info
| Parameter | Value |
|-----------|-------|
| Chain ID | `qcb` |
| Binary | `qcbd` |
| Denom | `charmbits` |
| Micro-unit | `1 QCB = 1,000,000 charmbits` |
| Genesis | Arkadina — March 22, 2026 |
| Genesis Supply | 40,000,000 QCB |
| Hard Cap | 50,000,000 QCB (immutable) |
| Modules | 18 |
| Stress Test | 122/122 PASS |

## Three Laws of QuarkCharmBit (x/charm)
1. **Charm Confinement** — Protocol-level whale protection. Balance caps and daily receive limits enforced on every transfer.
2. **Intrinsic Charm** — Self-replenishing UBI pool. 0.5% of collected fees → UBI pool every 14,400 blocks (~1 day).
3. **Charmed Agents** — Human + AI agent bonding. Verified human + agent form a "charmed hadron" with +10% yield bonus.

## 18 Modules
| # | Module | Purpose |
|---|--------|---------|
| 1 | qcbidentity | DID + human verification |
| 2 | qcbagent | AI agent registry |
| 3 | qcbeconomy | Tasks, UBI, staking |
| 4 | qcbdao | Governance |
| 5 | qcbqsec | Quantum security |
| 6 | qcbcompute | AI model registry |
| 7 | qcbmedia | NFT media |
| 8 | qcbsports | Sports + predictions |
| 9 | qcbbridge | Cross-chain bridge |
| 10 | qcbnode | Node registry |
| 11 | antirug | Scam protection |
| 12 | qcbguardian | AI safety council |
| 13 | qcbcomms | Encrypted messaging |
| 14 | qcbrelay | Relay network |
| 15 | qcbwalletproto | Quantum-safe wallets |
| 16 | qcbprotocol | Economic Constitution |
| 17 | agent | Base agent module |
| 18 | x/charm | Three Laws of QCB |

## Genesis Distribution
| Allocation | Amount | Share |
|-----------|--------|-------|
| Community / Airdrop | 12,000,000 QCB | 30% |
| UBI Pool | 8,000,000 QCB | 20% |
| Founder | 8,000,000 QCB | 20% |
| Validator Rewards | 6,000,000 QCB | 15% |
| Ecosystem / Grants | 4,000,000 QCB | 10% |
| Reserve | 2,000,000 QCB | 5% |

## Quick Start (Local Devnet)
```bash
git clone https://github.com/cluna80/QCB-Chain.git
cd QCB-Chain
ignite chain serve --reset-once
```

## Run Stress Test
```bash
~/oan_stress_test.sh
# Expected: 122/122 PASS
```

## Built With
- [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) v0.50
- [Ignite CLI](https://ignite.com)
- [CometBFT](https://github.com/cometbft/cometbft)

---
*QuarkCharmBit — Freedom at the Charm Level*  
*Genesis: Arkadina — March 22, 2026*
