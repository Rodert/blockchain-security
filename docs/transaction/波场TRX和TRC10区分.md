TRON

![在这里插入图片描述](https://img-blog.csdnimg.cn/1d7b9f55f3c34772be587e5efce61a2a.png)



## 获取数据方法

- 远程连接 tron 节点方式： https://cn.developers.tron.network/docs/trongrid
- 扩展：trc10 - love币
- 调用tronscan浏览器判断交易什么代币
- 查询每个交易是什么类型  curl -X POST  https://api.trongrid.io/wallet/getblockbynum -d '{"num": 45014224}'  。
  详细参见：https://cn.developers.tron.network/docs/exchangewallet-integrate-with-the-tron-network
  - "type":"TransferAssetContract" trc10代币转账
  - "type":"transferContract" trx转账
  - "type":"TriggerSmartContract" 分为trc10和trx转账
- 搭建节点分为 fullnode 和 Soliditynode，查询 trc10 需要 Soliditynode ：https://cn.developers.tron.network/v4.0/docs/%E5%AE%8C%E6%95%B4%E8%8A%82%E7%82%B9




## 交易类型

```bash
### 区别不同交易的交易类型

"type": "TransferContract" trx、
"type": "TransferAssetContract" trc10、
"type": "TriggerSmartContract" 触发合约、
"type": "FreezeBalanceContract" 质押资产、冻结资产、
"type": "UnfreezeBalanceContract" 解锁资产、解冻资产、
"type": "VoteWitnessContract" 投票、
"type": "WithdrawBalanceContract" 领取投票收益、
"type": "AccountPermissionUpdateContract" 更新账户权限、
"type": "AccountCreateContract" 激活账户、
"type": "AccountUpdateContract" 修改账户名称、
"type": "ParticipateAssetIssueContract" 购买TRC10通证
"type": "AssetIssueContract" 发行TRC10通证
"type": "UnfreezeAssetContract" TRC10锁仓提取
"type": "UpdateAssetContract" 更新TRC10通证参数
"type": "CreateSmartContract" 创建智能合约
"type": "UpdateSettingContract" 更新合约参数
"type": "UpdateEnergyLimitContract" 更新合约能量限制
"type": "ClearABIContract" 清除合约ABI
"type": "ProposalCreateContract" 超级代表-创建提议
"type": "ProposalApproveContract" 赞成提议
"type": "ProposalDeleteContract" 撤销提议
"type": "WitnessCreateContract" 创建超级代表候选人
"type": "WitnessUpdateContract" 更新超级代表候选人信息
"type": "UpdateBrokerageContract" 更新超级代表佣金比例 
"type": "ExchangeTransactionContract" 执行Bancor交易
"type": "ExchangeCreateContract" 创建Bancor交易
"type": "ExchangeInjectContract" Bancor交易注资
"type": "ExchangeWithdrawContract" Bancor交易撤资

```


## 参考：

1. 获取所有 trc10 代币信息：https://api.shasta.trongrid.io/v1/assets ; https://cn.developers.tron.network/reference/list-all-assets-trc10-tokens-on-chain
2. 