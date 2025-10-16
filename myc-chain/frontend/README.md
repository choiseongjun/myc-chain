# MYC-Chain 프론트엔드 대시보드

블록체인 기반 결제 관리 시스템의 웹 대시보드입니다.

## 🚀 빠른 시작

### 1. 체인 실행
```bash
cd /mnt/c/studypj/myc/myc-chain
ignite chain serve
```

### 2. 웹 브라우저에서 열기
`index.html` 파일을 브라우저에서 직접 열기

또는 간단한 HTTP 서버 사용:
```bash
cd frontend
python3 -m http.server 8000
# 또는
npx serve .
```

그리고 http://localhost:8000 접속

## 📊 기능

### ✅ 가맹점 관리
- 등록된 모든 가맹점 조회
- 가맹점 상태 확인 (active/pending)
- 실시간 통계

### 💳 결제 내역
- 모든 결제 트랜잭션 조회
- 결제 금액 및 상태 확인
- 총 결제액 통계

### 💰 정산 관리
- 가맹점별 정산 내역
- 정산 상태 추적 (completed/pending/processing)
- 총 정산액 표시

## 🔧 커스터마이징

### API 엔드포인트 변경
`index.html` 파일에서 `API_BASE` 상수 수정:

```javascript
const API_BASE = 'http://localhost:1317/mycchain/payment';
```

### 스타일 변경
`<style>` 태그 내의 CSS를 수정하여 디자인 커스터마이징

## 📱 반응형 디자인
모바일, 태블릿, 데스크톱 모두 지원합니다.

## 🔄 자동 새로고침
각 탭에서 "🔄 새로고침" 버튼을 클릭하여 최신 데이터를 가져올 수 있습니다.

## ⚠️ CORS 이슈 해결

만약 CORS 오류가 발생한다면, 체인의 설정을 수정하세요:

`~/.myc-chain/config/app.toml`:
```toml
[api]
enable = true
swagger = true
address = "tcp://0.0.0.0:1317"
```

그리고 체인을 재시작하세요.
