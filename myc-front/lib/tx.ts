// 간단한 트랜잭션 전송 (REST API 직접 사용)
import { DirectSecp256k1HdWallet } from '@cosmjs/proto-signing';
import { toBase64 } from '@cosmjs/encoding';

const RPC_ENDPOINT = process.env.NEXT_PUBLIC_RPC_ENDPOINT || 'http://localhost:26657';
const CHAIN_ID = 'mycchain';

// Alice의 니모닉
const ALICE_MNEMONIC = 'combine any whip scissors mean endless improve yellow company interest donate renew foot gloom crawl account finish trophy salt bless rather image happy cake';

// RPC를 통해 트랜잭션 브로드캐스트
export async function broadcastTx(txBytes: Uint8Array) {
  const response = await fetch(`${RPC_ENDPOINT}/broadcast_tx_sync`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      jsonrpc: '2.0',
      id: 1,
      method: 'broadcast_tx_sync',
      params: {
        tx: toBase64(txBytes),
      },
    }),
  });

  const data = await response.json();
  return data.result;
}

// 계정 정보 조회
async function getAccount(address: string) {
  const response = await fetch(
    `http://localhost:1317/cosmos/auth/v1beta1/accounts/${address}`
  );
  const data = await response.json();
  return data.account;
}

// 가맹점 생성 (간단한 버전 - CLI 프록시 사용)
export async function createMerchant(
  index: string,
  name: string,
  status: string,
  registeredAt: number
) {
  // 백엔드 API를 통해 CLI 명령 실행
  const response = await fetch('/api/tx/create-merchant', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      index,
      name,
      status,
      registeredAt,
    }),
  });

  if (!response.ok) {
    throw new Error('Failed to create merchant');
  }

  return await response.json();
}
