// Cosmos SDK 트랜잭션 처리
import { SigningStargateClient, StargateClient, defaultRegistryTypes } from '@cosmjs/stargate';
import { DirectSecp256k1HdWallet, Registry } from '@cosmjs/proto-signing';
import { stringToPath } from '@cosmjs/crypto';

const RPC_ENDPOINT = process.env.NEXT_PUBLIC_RPC_ENDPOINT || 'http://localhost:26657';
const CHAIN_ID = 'mycchain';

// Alice의 니모닉 (초기화 시 생성됨)
const ALICE_MNEMONIC = 'combine any whip scissors mean endless improve yellow company interest donate renew foot gloom crawl account finish trophy salt bless rather image happy cake';

// 커스텀 메시지 타입 정의
const mycchainTypes = [
  ['/mycchain.payment.MsgCreateMerchant', {
    encode: (message: any) => {
      // 간단한 인코딩 (실제로는 protobuf 사용)
      return new Uint8Array(0);
    },
    decode: (input: Uint8Array) => {
      return {};
    },
  }],
  ['/mycchain.payment.MsgCreatePayment', {
    encode: (message: any) => {
      return new Uint8Array(0);
    },
    decode: (input: Uint8Array) => {
      return {};
    },
  }],
  ['/mycchain.payment.MsgCreateSettlement', {
    encode: (message: any) => {
      return new Uint8Array(0);
    },
    decode: (input: Uint8Array) => {
      return {};
    },
  }],
];

// 지갑 생성 및 클라이언트 연결
export async function getSigningClient() {
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(ALICE_MNEMONIC, {
    prefix: 'cosmos',
  });

  const [firstAccount] = await wallet.getAccounts();

  // 커스텀 레지스트리 생성
  const registry = new Registry([...defaultRegistryTypes, ...mycchainTypes as any]);

  const client = await SigningStargateClient.connectWithSigner(
    RPC_ENDPOINT,
    wallet,
    {
      registry,
      gasPrice: {
        amount: '0.025',
        denom: 'stake',
      },
    }
  );

  return { client, address: firstAccount.address };
}

// REST API를 통한 트랜잭션 전송 (더 간단한 방법)
export async function createMerchant(
  index: string,
  name: string,
  status: string,
  registeredAt: number
) {
  const { client, address } = await getSigningClient();

  // Amino 메시지 형식 사용 (CosmJS가 자동으로 인코딩)
  const aminoMsg = {
    type: 'mycchain/CreateMerchant',
    value: {
      creator: address,
      index,
      name,
      status,
      registered_at: registeredAt.toString(),
    },
  };

  // 트랜잭션 전송
  const fee = {
    amount: [{ denom: 'stake', amount: '5000' }],
    gas: '200000',
  };

  // signAndBroadcast 사용 (일반 객체로 전달)
  const result = await client.signAndBroadcast(
    address,
    [aminoMsg as any],
    fee,
    ''
  );

  return result;
}

// 결제 생성 트랜잭션
export async function createPayment(
  merchantId: string,
  customerId: string,
  amount: string,
  status: string,
  createdAt: number
) {
  const { client, address } = await getSigningClient();

  const msg = {
    typeUrl: '/mycchain.payment.MsgCreatePayment',
    value: {
      creator: address,
      merchantId,
      customerId,
      amount,
      status,
      createdAt: createdAt.toString(),
    },
  };

  const fee = {
    amount: [{ denom: 'stake', amount: '5000' }],
    gas: '200000',
  };

  const result = await client.signAndBroadcast(address, [msg], fee, '');

  return result;
}

// 정산 생성 트랜잭션
export async function createSettlement(
  merchantId: string,
  totalAmount: string,
  settlementDate: number,
  status: string
) {
  const { client, address } = await getSigningClient();

  const msg = {
    typeUrl: '/mycchain.payment.MsgCreateSettlement',
    value: {
      creator: address,
      merchantId,
      totalAmount,
      settlementDate: settlementDate.toString(),
      status,
    },
  };

  const fee = {
    amount: [{ denom: 'stake', amount: '5000' }],
    gas: '200000',
  };

  const result = await client.signAndBroadcast(address, [msg], fee, '');

  return result;
}

// 잔액 조회
export async function getBalance(address: string) {
  const client = await StargateClient.connect(RPC_ENDPOINT);
  const balance = await client.getAllBalances(address);
  return balance;
}
