#!/bin/bash

# MYC-Chain REST API 테스트 스크립트

API_BASE="http://localhost:1317/mycchain/payment"

echo "======================================"
echo "MYC-Chain REST API 테스트"
echo "======================================"
echo ""

echo "1️⃣  모든 가맹점 조회"
echo "GET ${API_BASE}/merchant"
echo ""
curl -s "${API_BASE}/merchant" | jq '.'
echo ""
echo ""

echo "2️⃣  특정 가맹점 조회 (starbucks)"
echo "GET ${API_BASE}/merchant/starbucks"
echo ""
curl -s "${API_BASE}/merchant/starbucks" | jq '.'
echo ""
echo ""

echo "3️⃣  모든 결제 내역 조회"
echo "GET ${API_BASE}/payment"
echo ""
curl -s "${API_BASE}/payment" | jq '.'
echo ""
echo ""

echo "4️⃣  특정 결제 조회 (ID: 0)"
echo "GET ${API_BASE}/payment/0"
echo ""
curl -s "${API_BASE}/payment/0" | jq '.'
echo ""
echo ""

echo "5️⃣  모든 정산 내역 조회"
echo "GET ${API_BASE}/settlement"
echo ""
curl -s "${API_BASE}/settlement" | jq '.'
echo ""
echo ""

echo "6️⃣  특정 정산 조회 (ID: 0)"
echo "GET ${API_BASE}/settlement/0"
echo ""
curl -s "${API_BASE}/settlement/0" | jq '.'
echo ""
echo ""

echo "======================================"
echo "테스트 완료!"
echo "======================================"
