'use client';

import { useState, useEffect } from 'react';
import { merchantAPI, paymentAPI, settlementAPI, formatDate, formatAmount, getStatusColor } from '@/lib/api';
import type { Merchant, Payment, Settlement } from '@/lib/api';

export default function Home() {
  const [activeTab, setActiveTab] = useState<'merchants' | 'payments' | 'settlements'>('merchants');
  const [merchants, setMerchants] = useState<Merchant[]>([]);
  const [payments, setPayments] = useState<Payment[]>([]);
  const [settlements, setSettlements] = useState<Settlement[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showCreateModal, setShowCreateModal] = useState(false);

  useEffect(() => {
    loadData();
  }, [activeTab]);

  const loadData = async () => {
    setLoading(true);
    setError(null);
    try {
      if (activeTab === 'merchants') {
        const data = await merchantAPI.getAll();
        setMerchants(data);
      } else if (activeTab === 'payments') {
        const data = await paymentAPI.getAll();
        setPayments(data);
      } else if (activeTab === 'settlements') {
        const data = await settlementAPI.getAll();
        setSettlements(data);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : '데이터를 불러오는데 실패했습니다');
    } finally {
      setLoading(false);
    }
  };

  // 통계 계산
  const merchantStats = {
    total: merchants.length,
    active: merchants.filter(m => m.status === 'active').length,
  };

  const paymentStats = {
    total: payments.length,
    completed: payments.filter(p => p.status === 'completed').length,
    totalAmount: payments.reduce((sum, p) => sum + parseInt(p.amount.replace('stake', '')), 0),
  };

  const settlementStats = {
    total: settlements.length,
    completed: settlements.filter(s => s.status === 'completed').length,
    totalAmount: settlements.reduce((sum, s) => sum + parseInt(s.totalAmount.replace('stake', '')), 0),
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-600 via-purple-500 to-indigo-600">
      <div className="container mx-auto px-4 py-8">
        {/* Header */}
        <header className="bg-white rounded-lg shadow-lg p-8 mb-8">
          <h1 className="text-4xl font-bold text-purple-600 mb-2">
            🏪 MYC Chain 결제 관리 시스템
          </h1>
          <p className="text-gray-600">블록체인 기반 가맹점 결제 및 정산 플랫폼</p>
        </header>

        {/* Tabs */}
        <div className="flex gap-4 mb-6">
          <button
            onClick={() => setActiveTab('merchants')}
            className={`px-6 py-3 rounded-lg font-semibold transition-all ${
              activeTab === 'merchants'
                ? 'bg-white text-purple-600 shadow-lg'
                : 'bg-white/20 text-white hover:bg-white/30'
            }`}
          >
            가맹점 관리
          </button>
          <button
            onClick={() => setActiveTab('payments')}
            className={`px-6 py-3 rounded-lg font-semibold transition-all ${
              activeTab === 'payments'
                ? 'bg-white text-purple-600 shadow-lg'
                : 'bg-white/20 text-white hover:bg-white/30'
            }`}
          >
            결제 내역
          </button>
          <button
            onClick={() => setActiveTab('settlements')}
            className={`px-6 py-3 rounded-lg font-semibold transition-all ${
              activeTab === 'settlements'
                ? 'bg-white text-purple-600 shadow-lg'
                : 'bg-white/20 text-white hover:bg-white/30'
            }`}
          >
            정산 관리
          </button>
        </div>

        {/* Content */}
        <div className="bg-white rounded-lg shadow-lg p-8">
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-2xl font-bold text-gray-800">
              {activeTab === 'merchants' && '가맹점 목록'}
              {activeTab === 'payments' && '결제 내역'}
              {activeTab === 'settlements' && '정산 내역'}
            </h2>
            <div className="flex gap-3">
              {activeTab === 'merchants' && (
                <button
                  onClick={() => setShowCreateModal(true)}
                  className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                >
                  ➕ 가맹점 생성
                </button>
              )}
              <button
                onClick={loadData}
                className="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors"
              >
                🔄 새로고침
              </button>
            </div>
          </div>

          {/* Statistics Cards */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
            {activeTab === 'merchants' && (
              <>
                <StatCard label="전체 가맹점" value={merchantStats.total} />
                <StatCard label="활성 가맹점" value={merchantStats.active} />
              </>
            )}
            {activeTab === 'payments' && (
              <>
                <StatCard label="전체 결제" value={paymentStats.total} />
                <StatCard label="완료된 결제" value={paymentStats.completed} />
                <StatCard
                  label="총 결제액"
                  value={`${(paymentStats.totalAmount / 10000).toFixed(0)}만원`}
                />
              </>
            )}
            {activeTab === 'settlements' && (
              <>
                <StatCard label="전체 정산" value={settlementStats.total} />
                <StatCard label="완료된 정산" value={settlementStats.completed} />
                <StatCard
                  label="총 정산액"
                  value={`${(settlementStats.totalAmount / 10000).toFixed(0)}만원`}
                />
              </>
            )}
          </div>

          {/* Loading */}
          {loading && (
            <div className="text-center py-12">
              <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-purple-600"></div>
              <p className="mt-4 text-gray-600">데이터를 불러오는 중...</p>
            </div>
          )}

          {/* Error */}
          {error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {error}
            </div>
          )}

          {/* Tables */}
          {!loading && !error && (
            <>
              {activeTab === 'merchants' && <MerchantTable merchants={merchants} />}
              {activeTab === 'payments' && <PaymentTable payments={payments} />}
              {activeTab === 'settlements' && <SettlementTable settlements={settlements} />}
            </>
          )}
        </div>

        {/* Create Merchant Modal */}
        <CreateMerchantModal
          isOpen={showCreateModal}
          onClose={() => setShowCreateModal(false)}
          onCreated={loadData}
        />
      </div>
    </div>
  );
}

function StatCard({ label, value }: { label: string; value: string | number }) {
  return (
    <div className="bg-gradient-to-br from-purple-600 to-indigo-600 text-white p-6 rounded-lg shadow">
      <p className="text-sm opacity-90 mb-1">{label}</p>
      <p className="text-3xl font-bold">{value}</p>
    </div>
  );
}

function MerchantTable({ merchants }: { merchants: Merchant[] }) {
  return (
    <div className="overflow-x-auto">
      <table className="w-full">
        <thead>
          <tr className="bg-gray-50 border-b">
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">가맹점명</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">상태</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">등록일시</th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {merchants.map((merchant) => (
            <tr key={merchant.index} className="hover:bg-gray-50">
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{merchant.index}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-semibold">{merchant.name}</td>
              <td className="px-6 py-4 whitespace-nowrap">
                <span className={`px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full ${getStatusColor(merchant.status)}`}>
                  {merchant.status}
                </span>
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{formatDate(merchant.registered_at)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function PaymentTable({ payments }: { payments: Payment[] }) {
  return (
    <div className="overflow-x-auto">
      <table className="w-full">
        <thead>
          <tr className="bg-gray-50 border-b">
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">가맹점</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">고객</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">금액</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">상태</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">결제일시</th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {payments.map((payment) => (
            <tr key={payment.id} className="hover:bg-gray-50">
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{payment.id}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-semibold">{payment.merchantId}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{payment.customerId}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-bold">{formatAmount(payment.amount)}</td>
              <td className="px-6 py-4 whitespace-nowrap">
                <span className={`px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full ${getStatusColor(payment.status)}`}>
                  {payment.status}
                </span>
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{formatDate(payment.createdAt)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function SettlementTable({ settlements }: { settlements: Settlement[] }) {
  return (
    <div className="overflow-x-auto">
      <table className="w-full">
        <thead>
          <tr className="bg-gray-50 border-b">
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">가맹점</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">총 금액</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">정산일</th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">상태</th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {settlements.map((settlement) => (
            <tr key={settlement.id} className="hover:bg-gray-50">
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{settlement.id}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-semibold">{settlement.merchantId}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-bold">{formatAmount(settlement.totalAmount)}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{formatDate(settlement.settlementDate)}</td>
              <td className="px-6 py-4 whitespace-nowrap">
                <span className={`px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full ${getStatusColor(settlement.status)}`}>
                  {settlement.status}
                </span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function CreateMerchantModal({ isOpen, onClose, onCreated }: { isOpen: boolean; onClose: () => void; onCreated: () => void }) {
  const [formData, setFormData] = useState({
    index: '',
    name: '',
    status: 'active',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  const handleSubmit = async () => {
    if (!formData.index || !formData.name) {
      setError('모든 필드를 입력해주세요.');
      return;
    }

    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      const timestamp = Math.floor(Date.now() / 1000);

      // REST API를 통한 직접 트랜잭션 전송
      const { createMerchant } = await import('@/lib/tx-rest');

      const result = await createMerchant(
        formData.index,
        formData.name,
        formData.status,
        timestamp
      );

      if (result.code === 0) {
        setSuccess(true);
        setTimeout(() => {
          onCreated();
          onClose();
          setFormData({ index: '', name: '', status: 'active' });
          setSuccess(false);
        }, 2000);
      } else {
        setError(`트랜잭션 실패: ${result.rawLog || '알 수 없는 오류'}`);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : '트랜잭션 전송 중 오류가 발생했습니다.');
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 max-w-2xl w-full mx-4">
        <h3 className="text-2xl font-bold text-gray-800 mb-6">가맹점 생성</h3>

        <div className="space-y-4 mb-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              가맹점 ID (영문/숫자)
            </label>
            <input
              type="text"
              value={formData.index}
              onChange={(e) => setFormData({ ...formData, index: e.target.value })}
              placeholder="예: starbucks, cafe001"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 focus:border-transparent"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              가맹점명
            </label>
            <input
              type="text"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              placeholder="예: Starbucks Coffee"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 focus:border-transparent"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              상태
            </label>
            <select
              value={formData.status}
              onChange={(e) => setFormData({ ...formData, status: e.target.value })}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 focus:border-transparent"
            >
              <option value="active">활성</option>
              <option value="inactive">비활성</option>
            </select>
          </div>
        </div>

        {error && (
          <div className="mb-4 p-4 bg-red-100 border border-red-400 text-red-700 rounded-lg">
            {error}
          </div>
        )}

        {success && (
          <div className="mb-4 p-4 bg-green-100 border border-green-400 text-green-700 rounded-lg">
            ✅ 가맹점이 성공적으로 생성되었습니다!
          </div>
        )}

        <div className="flex justify-end gap-3">
          <button
            onClick={onClose}
            disabled={loading}
            className="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors disabled:opacity-50"
          >
            취소
          </button>
          <button
            onClick={handleSubmit}
            disabled={loading || !formData.index || !formData.name}
            className="px-6 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? (
              <span className="flex items-center gap-2">
                <span className="inline-block animate-spin rounded-full h-4 w-4 border-b-2 border-white"></span>
                전송 중...
              </span>
            ) : (
              '생성하기'
            )}
          </button>
        </div>
      </div>
    </div>
  );
}
