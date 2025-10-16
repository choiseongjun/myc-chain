# MYC-Chain 프론트엔드 연동 가이드

MYC-Chain은 REST API, gRPC, WebSocket을 통해 프론트엔드와 연동할 수 있습니다.

## 📡 API 엔드포인트

체인이 실행 중일 때 다음 포트들이 열립니다:

- **REST API**: `http://localhost:1317`
- **gRPC**: `localhost:9090`
- **RPC**: `http://localhost:26657`
- **WebSocket**: `ws://localhost:26657/websocket`

## 🔍 REST API 사용 예시

### 1. 가맹점 조회

#### 모든 가맹점 목록
```bash
curl http://localhost:1317/mycchain/payment/merchant
```

응답 예시:
```json
{
  "merchant": [
    {
      "index": "starbucks",
      "name": "Starbucks Coffee",
      "status": "active",
      "registeredAt": "1697500000"
    },
    {
      "index": "ediya",
      "name": "Ediya Coffee",
      "status": "active",
      "registeredAt": "1697500100"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

#### 특정 가맹점 조회
```bash
curl http://localhost:1317/mycchain/payment/merchant/starbucks
```

### 2. 결제 내역 조회

#### 모든 결제 목록
```bash
curl http://localhost:1317/mycchain/payment/payment
```

응답 예시:
```json
{
  "payment": [
    {
      "id": "0",
      "merchantId": "starbucks",
      "customerId": "customer001",
      "amount": "5000stake",
      "status": "completed",
      "createdAt": "1697600000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

#### 특정 결제 조회
```bash
curl http://localhost:1317/mycchain/payment/payment/0
```

### 3. 정산 내역 조회

#### 모든 정산 목록
```bash
curl http://localhost:1317/mycchain/payment/settlement
```

#### 특정 정산 조회
```bash
curl http://localhost:1317/mycchain/payment/settlement/0
```

### 4. 페이지네이션 (pagination)
```bash
# 페이지 크기 지정
curl "http://localhost:1317/mycchain/payment/payment?pagination.limit=10"

# 다음 페이지
curl "http://localhost:1317/mycchain/payment/payment?pagination.offset=10"
```

---

## 🌐 JavaScript/TypeScript 예시

### Fetch API 사용
```javascript
// 가맹점 목록 가져오기
async function getMerchants() {
  const response = await fetch('http://localhost:1317/mycchain/payment/merchant');
  const data = await response.json();
  return data.merchant;
}

// 특정 가맹점 조회
async function getMerchant(merchantId) {
  const response = await fetch(`http://localhost:1317/mycchain/payment/merchant/${merchantId}`);
  const data = await response.json();
  return data.merchant;
}

// 결제 내역 가져오기
async function getPayments() {
  const response = await fetch('http://localhost:1317/mycchain/payment/payment');
  const data = await response.json();
  return data.payment;
}

// 정산 내역 가져오기
async function getSettlements() {
  const response = await fetch('http://localhost:1317/mycchain/payment/settlement');
  const data = await response.json();
  return data.settlement;
}

// 사용 예시
getMerchants().then(merchants => {
  console.log('가맹점 목록:', merchants);
});
```

### Axios 사용
```javascript
import axios from 'axios';

const API_BASE = 'http://localhost:1317/mycchain/payment';

// 가맹점 API
export const merchantAPI = {
  getAll: () => axios.get(`${API_BASE}/merchant`),
  getOne: (id) => axios.get(`${API_BASE}/merchant/${id}`),
};

// 결제 API
export const paymentAPI = {
  getAll: () => axios.get(`${API_BASE}/payment`),
  getOne: (id) => axios.get(`${API_BASE}/payment/${id}`),
};

// 정산 API
export const settlementAPI = {
  getAll: () => axios.get(`${API_BASE}/settlement`),
  getOne: (id) => axios.get(`${API_BASE}/settlement/${id}`),
};
```

---

## ⚛️ React 컴포넌트 예시

```jsx
import React, { useState, useEffect } from 'react';

function MerchantList() {
  const [merchants, setMerchants] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchMerchants();
  }, []);

  const fetchMerchants = async () => {
    try {
      const response = await fetch('http://localhost:1317/mycchain/payment/merchant');
      const data = await response.json();
      setMerchants(data.merchant || []);
    } catch (error) {
      console.error('Error fetching merchants:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div>로딩 중...</div>;

  return (
    <div>
      <h2>가맹점 목록</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>이름</th>
            <th>상태</th>
            <th>등록일</th>
          </tr>
        </thead>
        <tbody>
          {merchants.map(merchant => (
            <tr key={merchant.index}>
              <td>{merchant.index}</td>
              <td>{merchant.name}</td>
              <td>{merchant.status}</td>
              <td>{new Date(merchant.registeredAt * 1000).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function PaymentList() {
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    fetchPayments();
  }, []);

  const fetchPayments = async () => {
    const response = await fetch('http://localhost:1317/mycchain/payment/payment');
    const data = await response.json();
    setPayments(data.payment || []);
  };

  return (
    <div>
      <h2>결제 내역</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>가맹점</th>
            <th>고객</th>
            <th>금액</th>
            <th>상태</th>
            <th>일시</th>
          </tr>
        </thead>
        <tbody>
          {payments.map(payment => (
            <tr key={payment.id}>
              <td>{payment.id}</td>
              <td>{payment.merchantId}</td>
              <td>{payment.customerId}</td>
              <td>{payment.amount}</td>
              <td>{payment.status}</td>
              <td>{new Date(payment.createdAt * 1000).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export { MerchantList, PaymentList };
```

---

## 🎨 Vue.js 컴포넌트 예시

```vue
<template>
  <div>
    <h2>가맹점 목록</h2>
    <div v-if="loading">로딩 중...</div>
    <table v-else>
      <thead>
        <tr>
          <th>ID</th>
          <th>이름</th>
          <th>상태</th>
          <th>등록일</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="merchant in merchants" :key="merchant.index">
          <td>{{ merchant.index }}</td>
          <td>{{ merchant.name }}</td>
          <td>{{ merchant.status }}</td>
          <td>{{ formatDate(merchant.registeredAt) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      merchants: [],
      loading: true
    }
  },
  mounted() {
    this.fetchMerchants()
  },
  methods: {
    async fetchMerchants() {
      try {
        const response = await fetch('http://localhost:1317/mycchain/payment/merchant')
        const data = await response.json()
        this.merchants = data.merchant || []
      } catch (error) {
        console.error('Error:', error)
      } finally {
        this.loading = false
      }
    },
    formatDate(timestamp) {
      return new Date(timestamp * 1000).toLocaleString()
    }
  }
}
</script>
```

---

## 💳 트랜잭션 전송 (결제/정산 생성)

프론트엔드에서 트랜잭션을 전송하려면 **CosmJS** 라이브러리를 사용합니다.

### 설치
```bash
npm install @cosmjs/stargate @cosmjs/proto-signing
```

### 예시 코드
```javascript
import { SigningStargateClient } from "@cosmjs/stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

async function createPayment() {
  // 1. 지갑 생성 (mnemonic 사용)
  const mnemonic = "your mnemonic words here...";
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "cosmos"
  });

  // 2. 계정 정보
  const [account] = await wallet.getAccounts();

  // 3. 클라이언트 연결
  const client = await SigningStargateClient.connectWithSigner(
    "http://localhost:26657",
    wallet
  );

  // 4. 트랜잭션 메시지 생성
  const msg = {
    typeUrl: "/mycchain.payment.MsgCreatePayment",
    value: {
      creator: account.address,
      id: "6",
      merchantId: "starbucks",
      customerId: "customer005",
      amount: "6000stake",
      status: "completed",
      createdAt: "1697700000"
    }
  };

  // 5. 트랜잭션 전송
  const fee = {
    amount: [{ denom: "stake", amount: "5000" }],
    gas: "200000"
  };

  const result = await client.signAndBroadcast(
    account.address,
    [msg],
    fee,
    "Payment created via frontend"
  );

  console.log("Transaction hash:", result.transactionHash);
  return result;
}
```

---

## 🔐 CORS 설정

프론트엔드가 다른 도메인에서 실행될 경우 CORS 설정이 필요합니다.

`~/.myc-chain/config/app.toml` 파일에서:

```toml
[api]
# Enable defines if the API server should be enabled.
enable = true

# Swagger defines if swagger documentation should automatically be registered.
swagger = true

# Address defines the API server to listen on.
address = "tcp://localhost:1317"

# CORS settings
[api.cors]
allowed-origins = ["*"]
allowed-methods = ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
allowed-headers = ["*"]
```

---

## 📊 실시간 데이터 업데이트 (WebSocket)

```javascript
const ws = new WebSocket('ws://localhost:26657/websocket');

ws.onopen = () => {
  // 새로운 블록 구독
  ws.send(JSON.stringify({
    jsonrpc: "2.0",
    method: "subscribe",
    id: 1,
    params: {
      query: "tm.event='NewBlock'"
    }
  }));
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('New block:', data);
  // 새 블록이 생성될 때마다 데이터 다시 가져오기
  fetchPayments();
};
```

---

## 🚀 빠른 시작 - 풀스택 예시

저장소에 간단한 Next.js 또는 React 예시 프로젝트를 만들어드릴까요?

1. **Next.js (추천)**: SSR + REST API 통합
2. **React + Vite**: SPA 빠른 개발
3. **Vue 3 + Vite**: Vue 생태계

어떤 것을 원하시나요?
