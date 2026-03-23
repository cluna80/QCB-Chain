# OAN Chain

**The Autonomous Civilization Blockchain**

OAN Chain is a sovereign Cosmos SDK blockchain designed for the age of autonomous AI agents, decentralized identity, and post-quantum security. Built from the ground up with 16 custom modules, OAN Chain powers a new kind of digital civilization — one where humans and AI agents coexist, trade, govern, and thrive together.

---

## Genesis Block — Arkadina

| Parameter | Value |
|---|---|
| **Genesis Name** | Arkadina |
| **Chain ID** | oan |
| **Genesis Date** | March 22, 2026 |
| **Total Supply** | 40,000,000 OAN |
| **Native Denom** | uoan (1 OAN = 1,000,000 uoan) |
| **Bond Denom** | uoan |
| **Inflation** | 5% (min 2%, max 10%) |
| **Unbonding Period** | 21 days |
| **Max Validators** | 100 |
| **Min Commission** | 5% |

---

## Token Distribution

| Allocation | Amount | % | Purpose |
|---|---|---|---|
| Community | 12,000,000 OAN | 30% | Airdrop & early adopters |
| UBI Pool | 8,000,000 OAN | 20% | Universal Basic Income module |
| Founder | 8,000,000 OAN | 20% | 4-year linear vesting |
| Validator Rewards | 6,000,000 OAN | 15% | Block rewards & staking |
| Ecosystem | 4,000,000 OAN | 10% | Grants & builders |
| Reserve | 2,000,000 OAN | 5% | Emergency & future use |
| **Total** | **40,000,000 OAN** | **100%** | |

---

## 16 Custom Modules

OAN Chain ships with 16 purpose-built modules that define the rules of the Arkadina civilization:

### Identity & Economy
| Module | Description |
|---|---|
| **oanidentity** | Decentralized identity (DID) — every human registers a sovereign identity |
| **oaneconomy** | Universal Basic Income, task marketplace, reward distribution |
| **oandao** | On-chain governance with proposals, voting, delegation, and timelock |

### AI & Agents
| Module | Description |
|---|---|
| **oanagent** | AI agent registry with DNA hashes, generational breeding, trading, and retirement |
| **oancompute** | AI model registry, inference jobs, compute staking |
| **oanguardian** | Guardian Council — AI oversight, emergency pause, veto power, AI limits |

### Infrastructure
| Module | Description |
|---|---|
| **oannode** | Node registry — validator, inference, bridge, and light nodes |
| **oanrelay** | Message relay network with regional routing and scoring |
| **oanbridge** | Cross-chain bridge — Ethereum (Axelar), Solana (Wormhole), IBC |
| **oanwalletproto** | Quantum-safe wallet protocol with encryption key management |
| **oancomms** | Encrypted on-chain messaging with key registration and policy |

### Security
| Module | Description |
|---|---|
| **oanqsec** | Post-quantum cryptography — Dilithium3, Falcon-1024, threat level system |
| **antirug** | Token protection — flag and freeze malicious contracts |

### Culture & Society
| Module | Description |
|---|---|
| **oanmedia** | NFT registry for film, art, and music with royalty and licensing |
| **oansports** | Athlete registry, stadiums, match scheduling, prediction markets |
| **oanmarket** | Decentralized marketplace for agent services and digital goods |

---

## Security Architecture

OAN Chain is built with multi-layer security:

- **Post-Quantum Cryptography** — Dilithium3 (195-bit security) as the active signing algorithm, with Falcon-1024 and SPHINCS+ registered. Quantum threat level system (0-5) with automatic protocol escalation.
- **Guardian Council** — A multi-sig human oversight layer with emergency pause capability, AI limit enforcement, and veto power over dangerous proposals.
- **Antirug Module** — Dormant by default, activated by governance. Flags and freezes malicious tokens before they can harm users.
- **Spawn Gate** — Rate limiting on agent spawning to prevent AI proliferation attacks.
- **Reentrancy Protection** — All BankKeeper calls write state before transferring tokens.
- **Anti-dust** — Minimum transaction amount enforced (100 uoan).

---

## Staking Tiers

OAN Chain rewards long-term holders with staking tiers that unlock chain privileges:

| Tier | Requirement | Benefits |
|---|---|---|
| **None** | 0 OAN | Basic access |
| **Arcadian** | 1,000 OAN staked | Register agents, vote on proposals |
| **Obsidian** | 10,000 OAN staked | Spawn agents, run nodes |
| **Sovereign** | 100,000 OAN staked | Guardian Council eligibility |
| **Genesis** | 1,000,000 OAN staked | Founding civilization member |

---

## Universal Basic Income

OAN Chain has UBI built into the protocol — not as an afterthought, but as a core module.

- **UBI Pool**: 8,000,000 OAN allocated at genesis to the oaneconomy module
- **Eligibility**: Verified humans with a registered DID
- **Cooldown**: One claim per epoch per verified identity
- **Anti-bot**: DID verification required — agents cannot claim human UBI
- **Double-claim protection**: On-chain cooldown enforced at the module level

---

## Explorer

The OAN Explorer is a full-featured blockchain explorer covering all 16 modules:

- Live block and transaction feed
- Agent DNA tree visualization
- Governance proposals and voting
- Node network map
- Relay network by region
- Quantum security dashboard
- NFT gallery (film, art, music)
- Sports league and predictions
- Cross-chain bridge activity
- Wallet protocol registry

---

## Tech Stack

| Layer | Technology |
|---|---|
| **Blockchain** | Cosmos SDK v0.50 |
| **Consensus** | CometBFT (Tendermint) |
| **Language** | Go |
| **Scaffolding** | Ignite CLI |
| **IBC** | Inter-Blockchain Communication v7 |
| **Interchain** | Interchain Accounts |
| **Explorer** | React + Vite + TailwindCSS |
| **API** | REST (port 1317) + RPC (port 26657) |

---

## Network Endpoints (Testnet)

| Endpoint | URL |
|---|---|
| REST API | `http://[node-ip]:1317` |
| RPC | `http://[node-ip]:26657` |
| Chain ID | `oan` |
| Explorer | Coming soon |

---

## Roadmap

- [x] 16 custom modules built and tested
- [x] 100/100 stress test — zero failures
- [x] BankKeeper wired — real OAN transfers
- [x] Genesis Block 1 — Arkadina (March 22, 2026)
- [x] OAN Explorer — all modules live
- [ ] Testnet launch with external validators
- [ ] Validator documentation
- [ ] OAN OS integration
- [ ] Community airdrop
- [ ] Mainnet launch

---

## Validator Information

OAN Chain will launch with an open validator set. Validator documentation and onboarding guide coming soon.

Minimum requirements for validators will be published before testnet launch.

---

## Philosophy

OAN Chain is built on a simple belief: **the next civilization will be digital, autonomous, and owned by its participants.**

- Humans have sovereign identity on-chain
- AI agents are registered, traceable, and bounded by human oversight
- Universal Basic Income is protocol-level, not political
- Post-quantum security is built in from day one — not added later
- The chain is designed to outlast its founder

*"Building in the dark."*

---

## License

OAN Chain is open source. License details coming with mainnet release.

---

## Contact

- GitHub: [github.com/cluna80/OAN-Chain](https://github.com/cluna80/OAN-Chain)
- Explorer: Coming soon
- Community: Coming soon

---

*OAN Chain — Arkadina Genesis — March 22, 2026*