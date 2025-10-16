# MYC-Front 빠른 시작 가이드

## 1️⃣ 블록체인 실행

```bash
cd /mnt/c/studypj/myc/myc-chain
ignite chain serve
```

## 2️⃣ 프론트엔드 실행

```bash
cd /mnt/c/studypj/myc/myc-front
npm run dev
```

브라우저에서 http://localhost:3000 접속

## 3️⃣ 테스트 데이터 입력

### 가맹점 생성
```bash
myc-chaind tx payment create-merchant starbucks "Starbucks Coffee" active 1697500000 --from alice --chain-id mycchain --node http://localhost:26657 -y

myc-chaind tx payment create-merchant ediya "Ediya Coffee" active 1697500100 --from alice --chain-id mycchain --node http://localhost:26657 -y
```

### 결제 생성  
```bash
myc-chaind tx payment create-payment 0 starbucks customer001 5000stake completed 1697600000 --from alice --chain-id mycchain --node http://localhost:26657 -y

myc-chaind tx payment create-payment 1 starbucks customer002 12000stake completed 1697600100 --from alice --chain-id mycchain --node http://localhost:26657 -y
```

### 정산 생성
```bash
myc-chaind tx payment create-settlement 0 starbucks 17000stake 1697702400 completed --from alice --chain-id mycchain --node http://localhost:26657 -y
```

## 4️⃣ 대시보드에서 확인

http://localhost:3000 에서:
- 가맹점 관리 탭
- 결제 내역 탭  
- 정산 관리 탭

모든 데이터를 확인할 수 있습니다!
