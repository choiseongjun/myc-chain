// MYC-Chain API 클라이언트

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:1317/myc-chain/payment/v1';

export interface Merchant {
  index: string;
  name: string;
  status: string;
  registered_at: string;
  creator?: string;
}

export interface Payment {
  id: string;
  merchantId: string;
  customerId: string;
  amount: string;
  status: string;
  createdAt: string;
}

export interface Settlement {
  id: string;
  merchantId: string;
  totalAmount: string;
  settlementDate: string;
  status: string;
}

export interface ApiResponse<T> {
  [key: string]: T[];
  pagination?: {
    next_key: string | null;
    total: string;
  };
}

// 트랜잭션 브로드캐스트를 위한 타입
export interface TxResponse {
  code: number;
  txhash: string;
  raw_log: string;
}

// 가맹점 API
export const merchantAPI = {
  async getAll(): Promise<Merchant[]> {
    const response = await fetch(`${API_BASE}/merchant`);
    if (!response.ok) throw new Error('Failed to fetch merchants');
    const data: ApiResponse<Merchant> = await response.json();
    return data.merchant || [];
  },

  async getOne(id: string): Promise<Merchant> {
    const response = await fetch(`${API_BASE}/merchant/${id}`);
    if (!response.ok) throw new Error(`Failed to fetch merchant ${id}`);
    const data = await response.json();
    return data.merchant;
  },

  async create(index: string, name: string, status: string, registeredAt: number): Promise<TxResponse> {
    // Cosmos SDK 트랜잭션 생성
    const tx = {
      body: {
        messages: [{
          "@type": "/mycchain.payment.MsgCreateMerchant",
          creator: "cosmos1qzpt5xt06qr70mqsqh64lqua5ahxtwp25lw87p", // alice 주소
          index,
          name,
          status,
          registeredAt: registeredAt.toString(),
        }],
        memo: "",
      },
      auth_info: {
        signer_infos: [],
        fee: {
          amount: [{ denom: "stake", amount: "200" }],
          gas_limit: "200000",
        },
      },
      signatures: [],
    };

    // 실제 환경에서는 Keplr나 다른 지갑을 통해 서명해야 하지만,
    // 여기서는 백엔드 CLI를 사용하도록 에러 메시지 반환
    throw new Error('트랜잭션 생성은 CLI를 통해 수행해야 합니다. 가맹점 생성 폼은 CLI 명령어를 생성합니다.');
  },
};

// 결제 API
export const paymentAPI = {
  async getAll(): Promise<Payment[]> {
    const response = await fetch(`${API_BASE}/payment`);
    if (!response.ok) throw new Error('Failed to fetch payments');
    const data: ApiResponse<Payment> = await response.json();
    return data.payment || [];
  },

  async getOne(id: string): Promise<Payment> {
    const response = await fetch(`${API_BASE}/payment/${id}`);
    if (!response.ok) throw new Error(`Failed to fetch payment ${id}`);
    const data = await response.json();
    return data.payment;
  },
};

// 정산 API
export const settlementAPI = {
  async getAll(): Promise<Settlement[]> {
    const response = await fetch(`${API_BASE}/settlement`);
    if (!response.ok) throw new Error('Failed to fetch settlements');
    const data: ApiResponse<Settlement> = await response.json();
    return data.settlement || [];
  },

  async getOne(id: string): Promise<Settlement> {
    const response = await fetch(`${API_BASE}/settlement/${id}`);
    if (!response.ok) throw new Error(`Failed to fetch settlement ${id}`);
    const data = await response.json();
    return data.settlement;
  },
};

// 유틸리티 함수
export const formatDate = (timestamp: string): string => {
  const date = new Date(parseInt(timestamp) * 1000);
  return date.toLocaleString('ko-KR');
};

export const formatAmount = (amount: string): string => {
  const num = parseInt(amount.replace('stake', ''));
  return num.toLocaleString() + '원';
};

export const getStatusColor = (status: string): string => {
  switch (status.toLowerCase()) {
    case 'active':
    case 'completed':
      return 'text-green-700 bg-green-100';
    case 'pending':
      return 'text-yellow-700 bg-yellow-100';
    case 'processing':
      return 'text-blue-700 bg-blue-100';
    case 'failed':
    case 'inactive':
      return 'text-red-700 bg-red-100';
    default:
      return 'text-gray-700 bg-gray-100';
  }
};
