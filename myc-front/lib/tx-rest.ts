// REST API를 통한 트랜잭션 전송 (SigningStargateClient 사용)
import { SigningStargateClient, defaultRegistryTypes } from '@cosmjs/stargate';
import { DirectSecp256k1HdWallet, Registry, GeneratedType } from '@cosmjs/proto-signing';
import { Writer } from 'protobufjs/minimal';

const RPC_ENDPOINT = process.env.NEXT_PUBLIC_RPC_ENDPOINT || 'http://localhost:26657';
const CHAIN_ID = 'mycchain';

// Alice의 니모닉
const ALICE_MNEMONIC = 'combine any whip scissors mean endless improve yellow company interest donate renew foot gloom crawl account finish trophy salt bless rather image happy cake';

// protobuf 인코더
const MsgCreateMerchant = {
  encode(message: any, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index);
    }
    if (message.name !== '') {
      writer.uint32(26).string(message.name);
    }
    if (message.status !== '') {
      writer.uint32(34).string(message.status);
    }
    if (message.registeredAt !== 0) {
      writer.uint32(40).int64(message.registeredAt);
    }
    return writer;
  },
  decode(input: any) {
    return {};
  },
  fromJSON(object: any) {
    return object;
  },
  toJSON(message: any) {
    return message;
  },
  fromPartial(object: any) {
    return object;
  },
};

// 가맹점 생성
export async function createMerchant(
  index: string,
  name: string,
  status: string,
  registeredAt: number
) {
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(ALICE_MNEMONIC, {
    prefix: 'cosmos',
  });

  const [account] = await wallet.getAccounts();

  // 커스텀 레지스트리 생성
  const registry = new Registry([
    ...defaultRegistryTypes,
    ['/mycchain.payment.v1.MsgCreateMerchant', MsgCreateMerchant as GeneratedType],
  ]);

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

  // 메시지 생성
  const msg = {
    typeUrl: '/mycchain.payment.v1.MsgCreateMerchant',
    value: {
      creator: account.address,
      index,
      name,
      status,
      registeredAt: registeredAt, // int64로 전달
    },
  };

  // Fee 설정
  const fee = {
    amount: [{ denom: 'stake', amount: '5000' }],
    gas: '200000',
  };

  // 트랜잭션 전송
  const result = await client.signAndBroadcast(
    account.address,
    [msg],
    fee,
    ''
  );

  return {
    code: result.code,
    txhash: result.transactionHash,
    rawLog: result.rawLog,
  };
}
