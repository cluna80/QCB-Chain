#!/bin/bash
echo "Patching QCB genesis..."

ALICE=$(qcbd keys show alice --address 2>/dev/null)
BOB=$(qcbd keys show bob --address 2>/dev/null)
FOUNDER="qcb1kqyhad9acgt2tdqypt36zdve3c9024e2jmnseg"

echo "Alice: $ALICE"
echo "Bob:   $BOB"

python3 << PYEOF
import json, sys

ALICE   = "$ALICE"
BOB     = "$BOB"
FOUNDER = "$FOUNDER"

with open('/home/cluna80/.qcb/config/genesis.json') as f:
    g = json.load(f)

app = g['app_state']
g['genesis_name'] = 'Arkadina'

TOTAL     = 40_000_000_000_000
COMMUNITY = 12_000_000_000_000
ECOSYSTEM =  4_000_000_000_000
FOUNDER_A =  8_000_000_000_000

app['bank']['supply'] = [{'denom': 'charmbits', 'amount': str(TOTAL)}]
app['bank']['balances'] = [
    {'address': ALICE,   'coins': [{'denom': 'charmbits', 'amount': str(COMMUNITY)}]},
    {'address': BOB,     'coins': [{'denom': 'charmbits', 'amount': str(ECOSYSTEM)}]},
    {'address': FOUNDER, 'coins': [{'denom': 'charmbits', 'amount': str(FOUNDER_A)}]},
]
app['staking']['params']['bond_denom'] = 'charmbits'
app['staking']['params']['min_commission_rate'] = '0.050000000000000000'
app['mint']['params']['mint_denom'] = 'charmbits'
app['mint']['minter']['inflation'] = '0.050000000000000000'
if 'crisis' in app:
    app['crisis']['constant_fee'] = {'denom': 'charmbits', 'amount': '1000000000'}
if 'gov' in app and 'params' in app['gov']:
    app['gov']['params']['min_deposit'] = [{'denom': 'charmbits', 'amount': '10000000000'}]

with open('/home/cluna80/.qcb/config/genesis.json', 'w') as f:
    json.dump(g, f, indent=2)

print(f"✅ Genesis patched: Alice={ALICE[:20]}... gets 12M QCB")
PYEOF
