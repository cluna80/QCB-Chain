#!/bin/bash
set -e
cd ~/oan

echo "═══════════════════════════════════════════════════"
echo "  OAN → QCB Chain Rename"
echo "  QuarkCharmBit — Freedom at the Charm Level"
echo "═══════════════════════════════════════════════════"

echo "Step 1: Updating config.yml..."
cat > config.yml << 'EOF'
version: 1
validation: sovereign
accounts:
- name: alice
  coins:
  - 20000token
  - 200000000stake
  - 10000000000charmbits
- name: bob
  coins:
  - 10000token
  - 100000000stake
  - 10000000000charmbits
client:
  openapi:
    path: docs/static/openapi.yml
faucet:
  name: bob
  coins:
  - 5token
  - 100000stake
  - 1000000charmbits
validators:
- name: alice
  bonded: 100000000stake
- name: validator1
  bonded: 100000000stake
- name: validator2
  bonded: 200000000stake
- name: validator3
  bonded: 100000000stake
EOF
echo "  ✅ config.yml done"

echo "Step 2: go.mod..."
sed -i 's|^module oan$|module qcb|g' go.mod
echo "  ✅ go.mod done"

echo "Step 3: Address prefix..."
sed -i 's|AccountAddressPrefix = "oan"|AccountAddressPrefix = "qcb"|g' app/app.go
echo "  ✅ prefix done"

echo "Step 4: Go import paths..."
find . -name "*.go" -not -path "*/.git/*" | xargs sed -i 's|"oan/|"qcb/|g'
echo "  ✅ imports done"

echo "Step 5: Module names in Go files..."
find . -name "*.go" -not -path "*/.git/*" | xargs sed -i \
  's/oanagent/qcbagent/g; s/oanbridge/qcbbridge/g; s/oancomms/qcbcomms/g; s/oancompute/qcbcompute/g; s/oandao/qcbdao/g; s/oaneconomy/qcbeconomy/g; s/oanguardian/qcbguardian/g; s/oanidentity/qcbidentity/g; s/oanmarket/qcbmarket/g; s/oanmedia/qcbmedia/g; s/oannode/qcbnode/g; s/oanprotocol/qcbprotocol/g; s/oanqsec/qcbqsec/g; s/oanrelay/qcbrelay/g; s/oansports/qcbsports/g; s/oanwalletproto/qcbwalletproto/g'
echo "  ✅ module names done"

echo "Step 6: uoan → charmbits..."
find . -name "*.go" -not -path "*/.git/*" | xargs sed -i 's/"uoan"/"charmbits"/g'
find . -name "*.go" -not -path "*/.git/*" | xargs sed -i 's/uoan/charmbits/g'
echo "  ✅ denom done"

echo "Step 7: Renaming x/ directories..."
cd x/
for pair in "oanagent:qcbagent" "oanbridge:qcbbridge" "oancomms:qcbcomms" "oancompute:qcbcompute" "oandao:qcbdao" "oaneconomy:qcbeconomy" "oanguardian:qcbguardian" "oanidentity:qcbidentity" "oanmarket:qcbmarket" "oanmedia:qcbmedia" "oannode:qcbnode" "oanprotocol:qcbprotocol" "oanqsec:qcbqsec" "oanrelay:qcbrelay" "oansports:qcbsports" "oanwalletproto:qcbwalletproto"; do
  old="${pair%%:*}"; new="${pair##*:}"
  [ -d "$old" ] && mv "$old" "$new" && echo "  x/$old → x/$new"
done
cd ..
echo "  ✅ x/ dirs done"

echo "Step 8: Renaming api/ directories..."
[ -d "api/oan" ] && mv api/oan api/qcb && echo "  api/oan → api/qcb"
if [ -d "api/qcb" ]; then
  cd api/qcb/
  for pair in "oanagent:qcbagent" "oanbridge:qcbbridge" "oancomms:qcbcomms" "oancompute:qcbcompute" "oandao:qcbdao" "oaneconomy:qcbeconomy" "oanguardian:qcbguardian" "oanidentity:qcbidentity" "oanmarket:qcbmarket" "oanmedia:qcbmedia" "oannode:qcbnode" "oanprotocol:qcbprotocol" "oanqsec:qcbqsec" "oanrelay:qcbrelay" "oansports:qcbsports" "oanwalletproto:qcbwalletproto"; do
    old="${pair%%:*}"; new="${pair##*:}"
    [ -d "$old" ] && mv "$old" "$new"
  done
  cd ../..
fi
echo "  ✅ api/ dirs done"

echo "Step 9: Renaming proto/ directories..."
[ -d "proto/oan" ] && mv proto/oan proto/qcb && echo "  proto/oan → proto/qcb"
if [ -d "proto/qcb" ]; then
  cd proto/qcb/
  for pair in "oanagent:qcbagent" "oanbridge:qcbbridge" "oancomms:qcbcomms" "oancompute:qcbcompute" "oandao:qcbdao" "oaneconomy:qcbeconomy" "oanguardian:qcbguardian" "oanidentity:qcbidentity" "oanmarket:qcbmarket" "oanmedia:qcbmedia" "oannode:qcbnode" "oanprotocol:qcbprotocol" "oanqsec:qcbqsec" "oanrelay:qcbrelay" "oansports:qcbsports" "oanwalletproto:qcbwalletproto"; do
    old="${pair%%:*}"; new="${pair##*:}"
    [ -d "$old" ] && mv "$old" "$new"
  done
  cd ../..
fi
echo "  ✅ proto/ dirs done"

echo "Step 10: Updating proto files..."
find proto/ -name "*.proto" | xargs sed -i \
  's|package oan\.|package qcb.|g; s|option go_package = "oan/|option go_package = "qcb/|g; s/oanagent/qcbagent/g; s/oanbridge/qcbbridge/g; s/oancomms/qcbcomms/g; s/oancompute/qcbcompute/g; s/oandao/qcbdao/g; s/oaneconomy/qcbeconomy/g; s/oanguardian/qcbguardian/g; s/oanidentity/qcbidentity/g; s/oanmarket/qcbmarket/g; s/oanmedia/qcbmedia/g; s/oannode/qcbnode/g; s/oanprotocol/qcbprotocol/g; s/oanqsec/qcbqsec/g; s/oanrelay/qcbrelay/g; s/oansports/qcbsports/g; s/oanwalletproto/qcbwalletproto/g; s/charmbits/charmbits/g'
echo "  ✅ proto files done"

echo "Step 11: Updating cmd/..."
find cmd/ -name "*.go" | xargs sed -i 's/"oand"/"qcbd"/g' 2>/dev/null || true
echo "  ✅ cmd done"

echo ""
echo "Step 12: Building QCB Chain..."
ignite chain build 2>&1 | tail -40
echo ""
echo "═══════════════════════════════════════════════════"
echo "  RENAME COMPLETE — OAN is now QCB"
echo "═══════════════════════════════════════════════════"
