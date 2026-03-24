# QuarkCharmBit (QCB) — Whitepaper

**Freedom at the Charm Level**  
**Powering the Arkadian Civilization**

*Chain ID: `qcb` | Genesis: Arkadina | March 22, 2026*  
*Version 1.0.0*

---

## Abstract

QuarkCharmBit (ticker: **QCB**) is a sovereign Layer-1 blockchain built on the Cosmos SDK. It is the first blockchain where principles inspired by the physics of the charm quark are encoded directly into the protocol as rules of economic stability and sovereignty.

In particle physics, the introduction of the charm quark completed the symmetry of the Standard Model, preventing unstable particle interactions through mechanisms such as the GIM mechanism. In QuarkCharmBit, similar stabilization rules are enforced at the protocol level to prevent economic imbalance, runaway concentration, and collapse of circulation.

**QCB stabilizes economic freedom through three core laws:**

| Law | Physics Principle | Economic Rule |
|-----|------------------|---------------|
| Charm Confinement | Quarks cannot exist in isolation | Protocol-level whale protection |
| Intrinsic Charm | Charm quarks appear probabilistically in protons | Self-replenishing UBI pool |
| Charmed Staking | Charm quarks form stable exotic hadrons | Enhanced yields for human+agent pairs |

**Native token:** QuarkCharmBit (QCB)  
**Micro-unit:** charmbits (1 QCB = 1,000,000 charmbits)  
**Genesis supply:** 40,000,000 QCB  
**Absolute hard cap:** 50,000,000 QCB (never exceeded, enforced at binary level)

QCB is not another digital commodity. It is a sovereign economic system combining scarcity, verified identity, and Universal Basic Income as core protocol primitives — designed for humans and autonomous agents to form a self-sustaining digital civilization.

---

## 1. The Physics Foundation — Why the Charm Quark?

The charm quark is a heavier quark introduced to restore symmetry and consistency in particle physics models. Its existence explains why certain particle transitions do not occur, preserving stability in the Standard Model.

In quantum chromodynamics, quarks obey **confinement** — they cannot exist in isolation. Additionally, protons can contain small probabilistic components known as **intrinsic charm**, where charm quarks briefly appear even at low energy levels.

QuarkCharmBit encodes these principles as protocol rules.

### Charm Confinement (`x/charm` — Law 1)

Just as quarks cannot exist freely, tokens in QCB cannot accumulate without limit.

On-chain rules enforce balance caps and daily receive limits based on identity tier. Any overflow from whale attempts is redirected to the UBI pool, preserving system stability. This is enforced on every `MsgSend` and `MsgStake` at the keeper level — it cannot be bypassed.

### Intrinsic Charm Replenishment (`x/charm` — Law 2)

Every 14,400 blocks (~1 day), a probabilistic slice (0.5% by default, governance-adjustable from 0.1% to 1.0%) of the accumulated fee pool is automatically sent to the UBI pool.

Like intrinsic charm in particle physics, this ensures the system continuously regenerates value even during periods of low activity. The UBI pool cannot permanently drain.

### Charmed Agents (`x/charm` — Law 3)

Verified human wallets may bond with autonomous AI agents. When paired, the human and agent form a bound economic unit analogous to a charmed hadron in particle physics.

These pairs receive enhanced staking rewards (+10% yield bonus by default, governance-adjustable) and elevated governance weight, encouraging long-term participation over speculation.

### Why This Approach Is Valid

This whitepaper does **not** claim:
- ❌ The blockchain uses real quarks
- ❌ Quantum physics literally runs the chain

It **does** claim:
- ✅ Physics stability principles → economic stability rules
- ✅ Confinement → anti-whale protection
- ✅ Intrinsic charm → UBI self-replenishment
- ✅ Symmetry → fair distribution
- ✅ Conservation → supply cap

Many real protocols use physics metaphors: entropy → randomness, thermodynamics → economics, symmetry → cryptography, conservation → supply limits. The charm quark framework fits naturally as a foundation for sovereign economic design.

---

## 2. Tokenomics

### Key Parameters

| Property | Value |
|----------|-------|
| Name | QuarkCharmBit |
| Ticker | QCB |
| Type | Layer-1 native coin (not a token) |
| Micro-unit | charmbits (1 QCB = 1,000,000 charmbits) |
| Genesis supply | 40,000,000 QCB |
| Absolute hard cap | 50,000,000 QCB |
| Inflation | 1% per year, tapering to 0% |
| Inflation reaches cap | ~Year 35 |
| After hard cap | Fee-based only (Bitcoin model) |
| Bond denom | charmbits |
| Address prefix | qcb1... |

### Genesis Distribution — Block 1 (Arkadina)

| Allocation | Amount | Share | Purpose |
|-----------|--------|-------|---------|
| Community / Airdrop | 12,000,000 QCB | 30% | Early adopters & airdrop |
| UBI Pool | 8,000,000 QCB | 20% | Universal Basic Income (forever) |
| Founder | 8,000,000 QCB | 20% | 4-year linear vesting |
| Validator Rewards | 6,000,000 QCB | 15% | Block rewards & staking |
| Ecosystem / Grants | 4,000,000 QCB | 10% | Builders & DAO grants |
| Reserve | 2,000,000 QCB | 5% | Emergency & future use |
| **TOTAL** | **40,000,000 QCB** | **100%** | — |

### Inflation Schedule

| Period | Inflation Rate | New QCB / Year | Cumulative Supply |
|--------|---------------|----------------|-------------------|
| Year 1–4 | 1.0% | ~400,000 QCB/yr | ~41.6M QCB |
| Year 5–8 | 0.5% | ~210,000 QCB/yr | ~42.4M QCB |
| Year 9–12 | 0.25% | ~105,000 QCB/yr | ~42.8M QCB |
| Year 13+ | → 0% | Approaching 0 | → 50M QCB cap |
| Year ~35 | 0% | 0 QCB | 50,000,000 QCB (final) |

> The hard cap of 50,000,000 QCB is enforced at the binary level in the `qcbprotocol` module. No new QCB can ever be minted once this cap is reached. This rule cannot be changed by governance.

---

## 3. Charm Confinement — Whale Protection

Enforced inside the `x/charm` keeper on every `MsgSend` and `MsgStake`.

### Wallet Tiers

| Tier | Requirement | Max Balance | Daily Receive |
|------|------------|-------------|---------------|
| Unverified | No DID | 4,000 QCB (0.01%) | 40 QCB/day |
| Verified | DID registered | 400,000 QCB (1%) | 2,000 QCB/day |
| DAO Approved | DID + DAO vote | 800,000 QCB (2%) | 2,000 QCB/day |
| Exempt | Module accounts | Unlimited | Unlimited |

### Time to Reach Maximum Holdings

| Target | Daily Limit | Days Required | Approx Time |
|--------|------------|---------------|-------------|
| 4,000 QCB (unverified max) | 40 QCB/day | 100 days | 3.3 months |
| 400,000 QCB (verified max) | 2,000 QCB/day | 200 days | 6.7 months |
| 800,000 QCB (DAO max) | 2,000 QCB/day | 400 days | 13.3 months |
| 10% of supply (4M QCB) | 2,000 QCB/day | 2,000 days | **5.5 years** |

> A whale attempting to acquire 10% of QCB supply at the maximum rate would need 5.5 years, full DID verification, DAO approval, and Guardian Council clearance — all while the entire community watches on-chain.

---

## 4. Universal Basic Income — Core Protocol Primitive

8,000,000 QCB are locked in the `qcbeconomy` module at genesis, dedicated forever to Universal Basic Income.

### How UBI Works

- Every verified human with a registered DID can claim UBI
- Claims are limited to once per epoch (cooldown enforced on-chain)
- AI agents cannot claim UBI — DID verification requires human proof
- Double-claim protection is enforced at the module level
- The UBI pool is replenished every epoch by Intrinsic Charm (fee slice)
- UBI amount per claim adjusts based on pool size and number of verified humans

### Intrinsic Charm Replenishment (Automatic)

Every 14,400 blocks (~1 day):
```
fee_pool × 0.5% → UBI pool
```

This is enforced by the `x/charm` EndBlocker — it runs automatically on every node, on every block, and cannot be disabled or bypassed.

> **UBI is not a feature. It is the economic law of the Arkadina Civilization.**

---

## 5. Staking Tiers & Charmed Yields

| Tier | Minimum Stake | Privileges | Charm Yield Bonus |
|------|--------------|------------|-------------------|
| None | 0 QCB | Basic send/receive, UBI claims | — |
| Arcadian | 1,000 QCB | Register agents, vote | +10% |
| Obsidian | 10,000 QCB | Spawn agents, run nodes | +25% |
| Sovereign | 100,000 QCB | Guardian Council eligibility | +50% |
| Genesis | 1,000,000 QCB | Founding civilization member | Special |

Charmed pairs (verified human + bonded AI agent) form a "hadron" and earn an additional yield bonus from a dedicated fee slice.

---

## 6. Earn QCB by Contributing

QCB is designed so contributors earn more than speculators. There are four ways to earn QCB by contributing to the Arkadina civilization:

### 1. Validate (Run a Node)
Run a validator node and earn block rewards from the 1% inflation pool. Early validators earn the most — rewards decrease as more validators join.

### 2. Claim UBI (Verified Human Right)
Register a DID identity, verify as human, and claim Universal Basic Income every epoch. This is a right, not a privilege — every verified human qualifies regardless of wealth.

### 3. Run Infrastructure
Run inference nodes, bridge nodes, relay nodes, or light nodes through the `qcbnode` and `qcbrelay` modules. Each node type earns rewards from the validator pool.

### 4. Build (Ecosystem Grants)
The 4,000,000 QCB ecosystem allocation is controlled by the DAO. Builders, developers, and contributors can apply for grants through governance proposals.

> **Earn vs. Buy:** QCB earned by contributing bypasses the daily receive cap. Only market purchases are throttled. The chain rewards builders, not speculators.

---

## 7. Launch Phases

All phase transitions are irreversible governance votes.

| Phase | Name | What Opens |
|-------|------|-----------|
| 1 | Genesis | Airdrop only — no trading |
| 2 | Validator | Validators earn block rewards |
| 3 | UBI | UBI claims open for verified humans |
| 4 | Ecosystem | Grants distributed, builders funded |
| 5 | DEX | Community DEX trading (Osmosis/IBC) |
| 6 | Open | Fully open — all protections active |

> Phase transitions cannot be reversed. QCB Chain will never rush to trading — community distribution comes first.

---

## 8. Governance

All economic parameters are governed by the DAO through on-chain proposals. The community — not any single entity — controls the future of the QCB economy.

### Governable Parameters
- Inflation rate and tapering schedule
- Minimum commission rate
- Unbonding period
- UBI epoch length and distribution amount
- Wallet tier limits and daily receive caps
- Intrinsic Charm fee slice (0.1%–1.0%)
- Charmed pair yield bonus
- Launch phase transitions
- Ecosystem grant distributions

### Immutable Parameters (Cannot Be Changed by Governance)
- Hard cap: 50,000,000 QCB absolute maximum
- UBI pool: 8,000,000 QCB at genesis, never reallocated
- Charm Confinement: whale protection is always active

---

## 9. Fee Economics & Long-Term Sustainability

| Parameter | Value |
|-----------|-------|
| Minimum fee | 500 charmbits per transaction |
| Fee distribution | Validators proportionally |
| Intrinsic Charm slice | 0.5% of fee pool per epoch → UBI |
| Long-term model | Fee-based (like Bitcoin) after inflation reaches 0% |

As inflation tapers toward zero over 35 years, fees become the primary validator incentive. As QCB Chain grows in usage, fee revenue grows — ensuring validators are always incentivized to secure the network.

---

## 10. The 18-Module Architecture

QCB Chain is built with 18 native Cosmos SDK modules:

| # | Module | Purpose |
|---|--------|---------|
| 1 | `qcbidentity` | DID + human verification |
| 2 | `qcbagent` | AI agent registry and lifecycle |
| 3 | `qcbeconomy` | Tasks, UBI pool, staking |
| 4 | `qcbdao` | On-chain governance |
| 5 | `qcbqsec` | Quantum security (Falcon-1024, Dilithium) |
| 6 | `qcbcompute` | AI model registry and inference jobs |
| 7 | `qcbmedia` | NFT media (film, art, music) |
| 8 | `qcbsports` | Sports, athletes, predictions |
| 9 | `qcbbridge` | Cross-chain bridge (Ethereum, Solana, Cosmos) |
| 10 | `qcbnode` | Node registry (light, full, inference) |
| 11 | `antirug` | Scam and rug-pull protection |
| 12 | `qcbguardian` | AI safety council and emergency pause |
| 13 | `qcbcomms` | Encrypted on-chain messaging |
| 14 | `qcbrelay` | Relay network for message routing |
| 15 | `qcbwalletproto` | Quantum-safe wallet profiles |
| 16 | `qcbprotocol` | Economic Constitution enforcement |
| 17 | `agent` | Base agent module |
| 18 | `x/charm` | **Three Laws of QuarkCharmBit** |

### x/charm — The Economic Constitution

The `x/charm` module is the heart of QCB Chain. It runs on every block alongside `mint`, `staking`, and `distribution`. It enforces:

```
Law 1 — Charm Confinement:
  → Balance caps enforced per tier on every transfer
  → Daily receive limits reset each epoch
  → Overflow redirected to UBI pool

Law 2 — Intrinsic Charm:
  → Every 14,400 blocks: fee_pool × 0.5% → UBI pool
  → UBI pool can never permanently drain
  → Daily receive counters reset each epoch

Law 3 — Charmed Agents:
  → Verified human bonds with AI agent
  → Forms a "charmed hadron" economic unit
  → +10% yield bonus on staking rewards
  → Enhanced governance weight
```

**Stability is enforced by code, not by trust.**

---

## 11. Technical Specifications

| Specification | Value |
|--------------|-------|
| Consensus | CometBFT (Tendermint) |
| SDK | Cosmos SDK v0.50 |
| Block time | ~6 seconds |
| Max validators | 100 |
| Unbonding period | 21 days |
| Min commission | 5% |
| IBC enabled | Yes |
| Quantum security | Falcon-1024, Dilithium3 |
| Cross-chain | Ethereum (Axelar), Cosmos (IBC), Solana (Wormhole) |
| Stress test | 122/122 PASS (18 waves, zero failures) |

---

## 12. Conclusion — The Charmed Path to Arcadia

QuarkCharmBit is not just another chain.

It is the smallest possible unit of sovereignty — the charm quark — scaled into an unbreakable, self-sustaining civilization where:

- Every verified human receives real Universal Basic Income
- Every agent serves the whole, earning by contributing
- Whale accumulation is physically impossible beyond protocol limits
- The UBI pool self-replenishes from economic activity forever
- Freedom is enforced at the protocol level, not promised

The Arkadina Civilization begins at Genesis Block 1.

---

**QuarkCharmBit — Freedom at the Charm Level.**  
**Welcome to Arkadina.**

---

*Chain ID: `qcb` | Binary: `qcbd` | Denom: `charmbits`*  
*Genesis: Arkadina — March 22, 2026*  
*Repository: [github.com/cluna80/QCB-Chain](https://github.com/cluna80/QCB-Chain)*
