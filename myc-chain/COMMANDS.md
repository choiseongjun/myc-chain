# MYC-Chain 테스트 명령어 모음

## 📋 가맹점 생성 (Merchant)

```bash
# 1. 스타벅스 가맹점 등록
myc-chaind tx payment create-merchant starbucks "Starbucks Coffee" active 1697500000 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 2. 이디야 가맹점 등록
myc-chaind tx payment create-merchant ediya "Ediya Coffee" active 1697500100 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 3. 투썸플레이스 가맹점 등록
myc-chaind tx payment create-merchant twosome "A Twosome Place" pending 1697500200 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 4. 메가커피 가맹점 등록
myc-chaind tx payment create-merchant mega "Mega Coffee" active 1697500300 --from bob --chain-id mycchain --node http://localhost:26657 -y
```

형식: `create-merchant [index] [name] [status] [registered-at]`
- index: 가맹점 고유 ID (예: starbucks, ediya)
- name: 가맹점 이름
- status: 상태 (active, pending, inactive)
- registered-at: 등록 시간 (Unix timestamp)

---

## 💳 결제 생성 (Payment)

```bash
# 1. 스타벅스 결제 #1 (5000원)
myc-chaind tx payment create-payment 0 starbucks customer001 5000stake completed 1697600000 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 2. 스타벅스 결제 #2 (12000원)
myc-chaind tx payment create-payment 1 starbucks customer002 12000stake completed 1697600100 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 3. 이디야 결제 #1 (3500원)
myc-chaind tx payment create-payment 2 ediya customer001 3500stake completed 1697600200 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 4. 이디야 결제 #2 (4200원, pending)
myc-chaind tx payment create-payment 3 ediya customer003 4200stake pending 1697600300 --from alice --chain-id mycchain --node http://localhost:26657 -y

# 5. 투썸플레이스 결제 (8500원)
myc-chaind tx payment create-payment 4 twosome customer002 8500stake completed 1697600400 --from bob --chain-id mycchain --node http://localhost:26657 -y

# 6. 메가커피 결제 (2800원)
myc-chaind tx payment create-payment 5 mega customer004 2800stake completed 1697600500 --from bob --chain-id mycchain --node http://localhost:26657 -y
```

형식: `create-payment [id] [merchant-id] [customer-id] [amount] [status] [created-at]`
- id: 결제 고유 번호 (0, 1, 2, ...)
- merchant-id: 가맹점 ID
- customer-id: 고객 ID
- amount: 금액 (단위: stake)
- status: 상태 (completed, pending, failed)
- created-at: 생성 시간 (Unix timestamp)

---

## 💰 정산 생성 (Settlement)

```bash
# 1. 스타벅스 일일 정산 (17000원)
myc-chaind tx payment create-settlement 0 starbucks 17000stake 1697702400 completed --from alice --chain-id mycchain --node http://localhost:26657 -y

# 2. 이디야 일일 정산 (7700원)
myc-chaind tx payment create-settlement 1 ediya 7700stake 1697702400 pending --from alice --chain-id mycchain --node http://localhost:26657 -y

# 3. 투썸플레이스 일일 정산 (8500원)
myc-chaind tx payment create-settlement 2 twosome 8500stake 1697702400 processing --from bob --chain-id mycchain --node http://localhost:26657 -y

# 4. 메가커피 일일 정산 (2800원)
myc-chaind tx payment create-settlement 3 mega 2800stake 1697702400 completed --from bob --chain-id mycchain --node http://localhost:26657 -y
```

형식: `create-settlement [id] [merchant-id] [total-amount] [settlement-date] [status]`
- id: 정산 고유 번호 (0, 1, 2, ...)
- merchant-id: 가맹점 ID
- total-amount: 총 금액 (단위: stake)
- settlement-date: 정산일 (Unix timestamp)
- status: 상태 (completed, pending, processing, failed)

---

## 🔍 조회 명령어 (Query)

### 가맹점 조회
```bash
# 모든 가맹점 목록
myc-chaind query payment list-merchant --node http://localhost:26657

# 특정 가맹점 조회
myc-chaind query payment show-merchant starbucks --node http://localhost:26657
myc-chaind query payment show-merchant ediya --node http://localhost:26657
```

### 결제 조회
```bash
# 모든 결제 목록
myc-chaind query payment list-payment --node http://localhost:26657

# 특정 결제 조회 (ID로)
myc-chaind query payment show-payment 0 --node http://localhost:26657
myc-chaind query payment show-payment 1 --node http://localhost:26657
```

### 정산 조회
```bash
# 모든 정산 목록
myc-chaind query payment list-settlement --node http://localhost:26657

# 특정 정산 조회 (ID로)
myc-chaind query payment show-settlement 0 --node http://localhost:26657
myc-chaind query payment show-settlement 1 --node http://localhost:26657
```

---

## ⚙️ 환경 설정

### PATH 추가 (영구적)
```bash
# ~/.bashrc 또는 ~/.zshrc에 추가
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### 또는 임시로 PATH 추가
```bash
export PATH=$PATH:$HOME/go/bin
```

### 전체 경로 사용 (PATH 추가 없이)
모든 명령어에서 `myc-chaind` 대신 `/home/choi/go/bin/myc-chaind` 사용

---

## 📝 유용한 팁

### 체인 시작
```bash
# Ignite로 시작 (개발 모드)
ignite chain serve

# 또는 직접 시작
myc-chaind start
```

### 계정 확인
```bash
# Alice 주소 확인
myc-chaind keys show alice --keyring-backend test

# Bob 주소 확인
myc-chaind keys show bob --keyring-backend test
```

### 잔액 확인
```bash
myc-chaind query bank balances $(myc-chaind keys show alice -a --keyring-backend test) --node http://localhost:26657
```
