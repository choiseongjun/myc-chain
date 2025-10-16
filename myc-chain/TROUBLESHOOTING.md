# MYC-Chain 문제 해결 가이드

## 🔧 "501 Not Implemented" 또는 REST API 접근 불가

### 원인
체인은 실행 중이지만 REST API 서버가 제대로 시작되지 않았거나 데이터베이스 손상

### 해결 방법

#### 방법 1: ignite chain serve로 재시작 (권장)

1. **모든 프로세스 중지**
```bash
pkill -f myc-chaind
pkill -f ignite
```

2. **데이터 초기화 및 재시작**
```bash
cd /mnt/c/studypj/myc/myc-chain
rm -rf ~/.myc-chain
ignite chain serve
```

`ignite chain serve`는 자동으로:
- ✅ 체인 초기화
- ✅ Alice, Bob 계정 생성
- ✅ REST API 서버 시작 (포트 1317)
- ✅ RPC 서버 시작 (포트 26657)
- ✅ 자동 리로드

#### 방법 2: 수동으로 API 서버 시작

1. **체인 초기화**
```bash
cd /mnt/c/studypj/myc/myc-chain
rm -rf ~/.myc-chain
ignite chain init
```

2. **노드 시작**
```bash
myc-chaind start &
```

3. **REST API 서버 별도 시작**
```bash
myc-chaind rest-server --enable-unsafe-cors &
```

하지만 **방법 1 (ignite chain serve)**을 강력 추천합니다!

---

## 🔍 API 테스트

REST API가 제대로 작동하는지 확인:

```bash
# 노드 정보 확인
curl http://localhost:1317/cosmos/base/tendermint/v1beta1/node_info

# 가맹점 목록 확인
curl http://localhost:1317/mycchain/payment/merchant

# 결제 목록 확인
curl http://localhost:1317/mycchain/payment/payment

# 정산 목록 확인
curl http://localhost:1317/mycchain/payment/settlement
```

성공 응답 예시:
```json
{
  "merchant": [],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

---

## 🚨 일반적인 문제들

### 1. "command not found: myc-chaind"

**해결:**
```bash
export PATH=$PATH:$HOME/go/bin
# 또는 전체 경로 사용
/home/choi/go/bin/myc-chaind version
```

### 2. "Error: go.mod not found"

**해결:**
```bash
cd /mnt/c/studypj/myc/myc-chain
# myc-chain 디렉토리 안에서 실행
```

### 3. "accepts 4 arg(s), received 3"

가맹점 생성 시 index 누락. 올바른 형식:
```bash
myc-chaind tx payment create-merchant [index] [name] [status] [registered-at]

# 예시
myc-chaind tx payment create-merchant starbucks "Starbucks" active 1697500000 --from alice --chain-id mycchain -y
```

### 4. CORS 에러 (프론트엔드에서)

**확인:**
`~/.myc-chain/config/app.toml` 파일에서:

```toml
[api]
enable = true
swagger = true
address = "tcp://0.0.0.0:1317"
```

체인 재시작 필요.

### 5. 포트가 이미 사용 중

**해결:**
```bash
# 사용 중인 프로세스 확인
lsof -i :1317
lsof -i :26657

# 프로세스 종료
kill -9 <PID>
```

---

## 📝 완전 재설정 (Clean Start)

모든 것을 처음부터 다시 시작:

```bash
# 1. 모든 프로세스 중지
pkill -f myc-chaind
pkill -f ignite

# 2. 데이터 삭제
rm -rf ~/.myc-chain

# 3. 재빌드
cd /mnt/c/studypj/myc/myc-chain
ignite chain build

# 4. 실행
ignite chain serve
```

---

## 🔄 정상 작동 체크리스트

체인이 제대로 실행되면:

- [ ] RPC 접근 가능: `curl http://localhost:26657/status`
- [ ] REST API 접근 가능: `curl http://localhost:1317/cosmos/base/tendermint/v1beta1/node_info`
- [ ] Payment 모듈 API: `curl http://localhost:1317/mycchain/payment/merchant`
- [ ] 블록 생성 중 (height 증가)
- [ ] Alice, Bob 계정 존재: `myc-chaind keys list`

---

## 💡 유용한 명령어

```bash
# 체인 상태 확인
myc-chaind status

# 계정 목록
myc-chaind keys list

# 잔액 확인
myc-chaind query bank balances [address]

# 로그 확인
tail -f ~/.myc-chain/myc-chaind.log

# 블록 높이 확인
curl -s http://localhost:26657/status | jq '.result.sync_info.latest_block_height'
```

---

## 📞 추가 도움이 필요하면

1. Ignite CLI 문서: https://docs.ignite.com
2. Cosmos SDK 문서: https://docs.cosmos.network
3. GitHub Issues: https://github.com/ignite/cli/issues
