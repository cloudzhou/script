**浙江网商银行交易见证服务API**

**版本号：V**2.1.0

目录

[*1.* *系统概述* 3](#系统概述)

[*2.* *名词解释* 3](#名词解释)

[*3.* *产品简介* 4](#产品简介)

[*4.* *实施案例* 6](#实施案例)

[*5.* *总体架构* 7](#总体架构)

[*6.* *账户体系* 8](#账户体系)

[*7.* *接入指引* 9](#接入指引)

[*8.* *开发指引* 10](#开发指引)

[*8.1.* *指引说明* 10](#指引说明)

[*8.2.* *会员开户* 11](#会员开户)

[*8.3.* *会员绑卡* 13](#会员绑卡)

[*8.4.* *即时入账* 16](#即时入账)

[*8.5.* *担保入账* 18](#担保入账)

[*8.6.* *转账入账* 20](#转账入账)

[*8.7.* *充值记账* 21](#充值记账)

[*8.8.* *提现* 22](#提现)

[*8.9.* *冻结解冻* 24](#冻结解冻)

[*8.10.* *基础服务接口* 25](#基础服务接口)

[*9.* *接口规范* 25](#接口规范)

[*9.1.* *企业要求* 25](#企业要求)

[*9.2.* *接口说明* 25](#接口说明)

[*9.3.* *调用地址* 26](#调用地址)

[*9.4.* *调用示例* 26](#调用示例)

[*10.* *API列表* 28](#api列表)

[*10.1.* *会员API* 28](#会员api)

[*10.2.* *支付API* 28](#支付api)

[*10.3.* *基础API* 29](#基础api)

[*11.* *API说明* 29](#api说明)

[*11.1.* *公共请求参数* 29](#公共请求参数)

[*11.2.* *公共响应参数* 30](#公共响应参数)

[*11.3.* *\[10101\]mybank.tc.user.personal.register (开个人户)*
30](#mybank.tc.user.personal.register-开个人户)

[*11.4.* *\[10102\]mybank.tc.user.enterprise.register (开企业户)*
32](#mybank.tc.user.enterprise.register-开企业户)

[*11.5.* *\[10103\]mybank.tc.user.account.create (开户)*
34](#mybank.tc.user.account.create-开户)

[*11.6.* *\[10401\]mybank.tc.user.bankcard.bind(绑定银行卡)*
35](#mybank.tc.user.bankcard.bind绑定银行卡)

[*11.7.* *\[10402\]mybank.tc.user.bankcard.unbind(解绑银行卡)*
37](#mybank.tc.user.bankcard.unbind解绑银行卡)

[*11.8.* *\[10201\]mybank.tc.user.personal.info.modify (修改个人信息)*
38](#mybank.tc.user.personal.info.modify-修改个人信息)

[*11.9.* *\[10202\]mybank.tc.user.enterprise.info.modify (修改企业信息)*
39](#mybank.tc.user.enterprise.info.modify-修改企业信息)

[*11.10.* *\[10301\]mybank.tc.user.personal.info.query (查询个人信息)*
41](#mybank.tc.user.personal.info.query-查询个人信息)

[*11.11.* *\[10302\]mybank.tc.user.enterprise.info.query(查询企业信息)*
42](#mybank.tc.user.enterprise.info.query查询企业信息)

[*11.12.* *\[10303\]mybank.tc.user.account.balance(查余额)*
44](#mybank.tc.user.account.balance查余额)

[*11.13.* *\[10403\]mybank.tc.user.bankcard.query(查询银行卡列表)*
45](#mybank.tc.user.bankcard.query查询银行卡列表)

[*11.14.* *\[10404\]mybank.tc.user.info.img.upload (文件上传)*
47](#mybank.tc.user.info.img.upload-文件上传)

[*11.15.* *\[20101\]mybank.tc.trade.pay.instant (即时交易入账)*
48](#mybank.tc.trade.pay.instant-即时交易入账)

[*11.16.* *\[20102\]mybank.tc.trade.pay.ensure (担保交易入账)*
54](#mybank.tc.trade.pay.ensure-担保交易入账)

[*11.17.* *\[20103\]mybank.tc.trade.settle(担保交易结算入账)*
57](#mybank.tc.trade.settle担保交易结算入账)

[*11.18.* *\[20105\]mybank.tc.trade.deposit(充值记账)*
58](#mybank.tc.trade.deposit充值记账)

[*11.19.* *\[20108\]mybank.tc.trade.refund (退款入账)*
60](#mybank.tc.trade.refund-退款入账)

[*11.20.* *\[20106\]mybank.tc.trade.paytocard(单笔提现)*
63](#mybank.tc.trade.paytocard单笔提现)

[*11.21.* *\[20109\]mybank.tc.trade.refundticket (退票)*
65](#mybank.tc.trade.refundticket-退票)

[*11.22.* *\[20104\]mybank.tc.trade.transfer(转账入账)*
66](#mybank.tc.trade.transfer转账入账)

[*11.23.* *\[20201\]mybank.tc.trade.query (交易流水查询)*
68](#mybank.tc.trade.query-交易流水查询)

[*11.24.* *\[20202\]mybank.tc.trade.info.query (交易详情查询)*
69](#mybank.tc.trade.info.query-交易详情查询)

[*11.25.* *\[20301\]mybank.tc.trade.funds.freeze (冻结资金)*
71](#mybank.tc.trade.funds.freeze-冻结资金)

[*11.26.* *\[20302\]mybank.tc.trade.funds.unfreeze (解冻资金)*
72](#mybank.tc.trade.funds.unfreeze-解冻资金)

[*11.27.* *\[20401\]mybank.tc.trade.carryover(资金结转)*
73](#mybank.tc.trade.carryover资金结转)

[*11.28.* *\[30101\]mybank.tc.user.area.query (省、市查询)*
75](#mybank.tc.user.area.query-省市查询)

[*11.29.* *\[30102\]mybank.tc.user.area.bank.query (银行列表查询)*
76](#mybank.tc.user.area.bank.query-银行列表查询)

[*11.30.* *\[30103\]mybank.tc.user.cardbin.query (卡BIN查询)*
78](#mybank.tc.user.cardbin.query-卡bin查询)

[*12.* *交易对账* 79](#交易对账)

[*12.1.* *对账说明* 79](#对账说明)

[*12.2.* *错账处理* 79](#错账处理)

[*13.* *商户异步通知* 80](#商户异步通知)

[*13.1.* *通知说明* 80](#通知说明)

[*13.2.* *RSA签名机制* 81](#rsa签名机制)

[*13.3.* *异步通知接口* 82](#异步通知接口)

[*14.* *附录* 83](#附录)

[*14.1.* *交易状态* 83](#交易状态)

[*14.2.* *公共错误码* 85](#公共错误码)

[*14.3.* *证件类型* 87](#证件类型)

[*14.4.* *银行总行列表* 88](#银行总行列表)

[*14.5.* *银行机构编码列表* 93](#银行机构编码列表)

[*14.6.* *第三方支付机构列表* 94](#第三方支付机构列表)

[*14.7.* *折扣支付方式* 94](#折扣支付方式)

[*14.8.* *商户通知类型* 95](#商户通知类型)

[*14.9.* *账户类型* 95](#账户类型)

[*14.10.* *支付方式* 95](#支付方式)

[*14.11.* *业务产品编码* 96](#业务产品编码)

系统概述
========

网商银行推出交易见证平台旨在为电商交易平台、交易所平台等平台型的企业提供一站式的金融服务解决方案，解决从资金存管、资金收付、资金增值等一系列的金融服务。

与网商银行交易见证平台的对接，可为平台带来：

-   提升平台资金管理的内控水平，提升平台的公信力，增强平台获客能力

-   一站式的综合金融服务服务解决电商平台所需的主要金融服务需求

-   为电商平台提供便捷的、低成本的支付和收款、清分等服务，大大提高平台交易资金的使用安全程度与处理效率s

-   通过内嵌的资金增值服务，提高平台账户和用户账户内留存资金的收益

目前，网商银行的交易见证平台可为客户提供交易资金监管和交易资金代付两大功能。

**交易资金监管**：银行为平台开立资金存管汇总账户，为每位平台上的会员建立资金存管明细账户，该明细账户应与交易会员指定的同名银行结算账户建立出入金转账对应关系；在此前提下，银行发挥账户管理、资金结算等方面的金融服务优势及第三方存管的作用，为电商平台和会员提供出入金、交收清算和对账服务。

**交易资金代付**：可一站式实现平台资金对个人银行卡、对公账户等所有类型账户的实时代付需求，目前代发业务手续费全免。

名词解释
========

  术语                   定义
  ---------------------- --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  合作方业务平台         接入网商交易见证系统的客户平台，例如商城。
  合作方业务平台订单号   合作方业务平台提交的交易号，对于合作方全局唯一。
  充值                   通过网银或其它渠道给自己账户充值。交易见证平台提供的“充值记账”接口仅用作同步合作方业务平台的充值信息。
  提现                   将账户资金转入自己的银行借记卡。
  转账                   资金在会员账户间转移的支付动作称为转账。交易见证平台提供的“转账入账”接口仅用作同步合作方业务平台的转账信息。
  即时到账交易           买家下单并付款成功后金额直接入卖家账户，区别于“担保交易”。交易见证平台提供的“即时交易入账”接口仅用作同步合作方业务平台的即时到账交易信息。
  担保交易               区别于“即时到账交易”，买家付款成功后，金额入担保账户，需要买家确认到货后，发起“确认收货”，金额才转入卖家账户。交易见证平台提供的“担保交易入账”接口仅用作同步合作方业务平台的担保交易信息。
  确认收货               买家付款成功，在收到货物后，将担保金额结算给卖家。交易见证平台提供的“担保交易结算入账”接口仅用作同步合作方业务平台的确认收货交易信息。
  退款                   买家付款完成或者交易完成后，由于买家或者卖家的原因需要退款时，将支付款退还给买家。交易见证平台提供的“退款入账”接口仅用作同步合作方业务平台的退款信息。
  交易对账               交易对账是指以银行流水为基础，通过金额、银行订单号、银行等要素的匹配，与入账流水一一勾兑的过程，交易对账的目的主要关注交易成功。
  资金结转               交易对账完成后，对于勾对一致的流水，将其对应的余额，在系统内核销对应待清算资金。

产品简介
========

-   会员账户管理

> 交易见证平台与客户平台会员系统账户一一对应，客户可以通过交易见证平台提供的接口对会员账户进行新增、修改、查询等操作。另外可对会员账户余额进行冻结、解冻操作。

-   银行卡管理

对平台客户的会员银行卡进行管理，可以对银行卡进行新增、删除、查询等操作，同一个用户可以绑定多张银行卡。另外交易见证平台会对提现银行卡进行校验，确认提现银行卡与提现账户为同一用户所有。

-   交易数据同步

客户平台交易成功后，将交易同步给交易见证平台，交易见证平台支持充值、支付、转账、退款等交易信息同步，同步成功后，交易见证平台异步通知客户系统。

-   交易资金代付

适用于客户平台下会员账户实时提现到银行卡操作，交易见证平台提供单笔提现，提现成功后，结果异步通知给客户。

-   对账

> 交易见证平台提供交易账单与账户余额账单下载，客户根据交易账单与业务系统对账，确保两边交易数据一致。通过余额账单对账与业务系统会员账户余额对账，确保两边账户数据一致。

-   理财能力

交易见证平台接入的余利宝是蚂蚁金服旗下网商银行代销的面向小微企业的货币基金产品，可以作为企业的流动性资金管理、增值的工具。具有超高流动性，低风险等特色。在交易见证平台上，合作方业务平台通过控台页面，直接使用账户中的余额申购余利宝。从而实现资金收益最大化。当前理财产品申购、赎回均不收费。

特点：高流动性：单日赎回500万以内实时到账。低门槛：一分钱起购。全天候：7×24小时全天候交易。低风险：对接低风险的货币基金

**会员交易见证功能示意图**

![](media/image1.png){width="6.966666666666667in"
height="3.1831266404199474in"}

**典型交易流程图**

**（会员A直接通过入款渠道订单支付给B见证账户）**

![](media/image2.emf){width="6.947916666666667in"
height="5.496508092738408in"}

**典型交易流程图**

**（会员A先通过入款渠道给自己充值，再用见证账户余额订单支付给B见证账户）**

![](media/image3.emf){width="6.885416666666667in"
height="6.298082895888014in"}

实施案例
========

**案例1：**

某二手车平台，成功对接交易见证平台，完成交易资金监管与交易资金代付。该平台主要业务包括：会员开户、会员银行卡管理、会员账户充值、二手车买卖、二手车拍卖保证金缴纳、购买平台增值服务、会员提现。

总体架构
========

> ![](media/image4.emf){width="5.805555555555555in"
> height="2.796527777777778in"}

账户体系
========

![](media/image5.emf){width="6.483333333333333in"
height="1.2076804461942257in"}

![](media/image6.emf){width="6.277387357830271in" height="4.525in"}

-   银行结算账户：

<!-- -->

-   **网商监管账户：**平台在网商银行开立的资金监管账户，用于汇总核算平台所有会员及平台自有的交易资金。其下面包含着会员虚拟户、平台自有虚拟户、内部户虚拟户以及银存虚拟户。实际资金则存放在网商监管账户里，而子账户的余额则是记录这些钱的分别归属，并不是实际资金。网商监管账户余额与见证虚拟账户间余额满足如下两个等式规则：

1.  网商监管账户余额=所有会员虚拟户余额+所有内部户虚拟户余额+平台自有虚拟户余额。

2.  网商监管账户余额=“银存-网商银行存款虚拟户”余额。

-   **平台自有结算账户：**平台在网商或其他银行开立的实体结算账户。平台可将归属于平台自己的资金提现到该账户（目前只能在PC-WEB端操作）。

-   **会员绑定的银行结算账户：**会员在网商或其他银行开立的实体结算账户。会员可将归属于会员自己的资金提现到该账户（目前只能通过提现接口操作）

<!-- -->

-   见证虚拟账户：

<!-- -->

-   **银存虚拟户：**目前有银存-网商银行存款账户（见证虚拟账户），该账户与网商银行监管账户（银行结算账户）一一对应，基于账实相符的原则，这两个账户间余额需保持一致。

-   **会员虚拟户：**平台通过接口调用的方式为会员在网商银行见证系统开立的会员分账账户。会员开户时候，系统默认开立BASIC-基本户（会员虚拟基本户），会员虚拟户的余额可用于平台体系内交易支付（会员作为在实际交易背景中的买家，将资金支付给卖家），也可用于提现到会员绑定的实体结算账户。

-   **平台自有虚拟户：**归属于平台自己的虚拟子账户，分为自有-基本户、自有-收入户和自有-费用户3个分账账户。

    -   **自有-基本户：**平台自有资金归集分账账户，该账户默认绑定平台上线提交的平台自有结算账户。平台可通过自主汇划的方式，将资金从自有结算账户转入到平台网商监管账户，并在PC-WEB上通过登账交易调增“自有-基本户”余额。

    -   **自有-收入户：**核算平台在交易过程中收取的手续费。即交易过程中，若平台收取交易手续费，则手续费收入记入该分账账户。平台若希望提取该部分收益，可先在PC-WEB端将“自有-收入户”的余额结转到“自有-基本户”，再从“自有-基本户”提现到绑定的平台自有结算账户。

    -   **自有-营销户：**在折扣支付场景下，需平台垫付营销费用。即交易过程中，若涉及到折扣，则折扣金额从该分账账户扣减。
        平台可以在PC-WEB端从“自有-基本户”转入到“自有-营销户”。

-   **内部虚拟户：**业务资金过度账户。目前主要有担保过渡户和待清算过渡户两个。

    -   **担保过渡户：**用户在担保支付场景下。比如：买家支付订单款项后，资金不马上给到卖家，资金先进入到担保过渡户，后续买家确认收到货物后，再通过担保结算入账将款项支付给卖家。该账户余额只能通过担保入账、担保结算入账、担保退款入账三个交易场景变更，平台无法直接通过登账或余额支付等方式变更。

    -   **待清算过渡户：**即时入账、担保入账、充值记账三个入账场景下，
        若非余额支付，即通过入款渠道进行支付，平台先调用入款渠道接口完成支付，再调用网商见证接口进行入账。一般入款渠道资金T+1结算（需结算到网商监管账户）。因此，见证系统在调增会员余额同时，记待清算过渡户，表示该笔资金在途等待实际清算。分录为：

> 借：待清算过渡户
>
> 贷：会员基本户（或担保过渡户）
>
> 实际资金T+1到账后，系统自动结转银存。
>
> 借：银存
>
> 贷：已清算过渡户
>
> 平台可通过资金结转交易通知网商见证系统进行资金清算。分录为：
>
> 借：已清算过渡户
>
> 贷：待清算过渡户

接入指引
========

  步骤   事项                 说明                                                                                                                                                                                                                                                                                               相关文档
  ------ -------------------- -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------
  1.     商务接洽，协议签订   与网商签订协议，完成银行开户。该项任务是后续对接交易见证平台的前提条件。客户在网商银行完成银行账户开户，签订对接交易见证平台协议，申请天威诚信证书。                                                                                                                                               
  1.     需求分析与提交       该项任务是为帮助交易见证平台充分了解客户需求，评估对接交易见证平台可行性及功能点。客户根据自身业务场景，整理出会员开户、会员银行卡管理、充值记账、即时入账、担保方式入账、会员提现等业务场景，通过需求申请表详细描述需求场景，交易见证平台进行评估，确认所有场景均可支持，可进行下一步联调工作。   
  1.     联调对接             交易见证平台为客户开通联调测试账号，根据客户业务需求，指导客户联调测试，完成上线方案准备。客户拿到交易见证平台提供的测试账号，根据自己的业务场景对接联调测试，完成联调确认所有场景都已跑通后，制定上线迁移方案。                                                                                   
  1.     上线投产             客户完成联调后，确认上线方案。交易见证平台给客户开通相关生产环境账号，配置生产环境相关配置信息。                                                                                                                                                                                                   

> 补充说明：
>
> 客户若涉及存量业务，需将实际资金转由网商银行账户存管。
>
> 上线时将实际提现渠道切换为网商银行，客户上线生产环境后，交易见证平台配合客户确认是否上线成功，确认上线成功后，系统对接完成。

![](media/image10.emf){width="5.7965277777777775in"
height="0.6111111111111112in"}

1.  开发指引
    ========

    1.  指引说明
        --------

交易接口：即时到账、担保交易、结算、充值、退款、单笔体现、转账，系统根据outer\_trade\_no（合作方业务平台订单号）做判重处理。

-   若系统返回明确失败或成功，平台不需重试。 若平台重试，
    则不管原交易处理成功还是失败，均返回“DUPLICATE\_REQUEST\_N-合作方业务平台订单号重复”。

{"charset":"UTF-8","error\_code":"DUPLICATE\_REQUEST\_NO","error\_message":"合作方业务平台订单号重复","is\_success":"F"}

-   若接口调用超时或系统返回SYSTEM\_ERROR， 平台若重试，
    则按如下2种情况处理：

<!-- -->

-   第一次请求见证系统若未落地，则继续处理， 并返回处理结果。

-   第一次已落地， 则返回“DUPLICATE\_REQUEST\_NO
    -合作方业务平台订单号重复”，平台可调用“20202-交易详情查询”接口同步见证系统处理状态。

    1.  会员开户
        --------

        1.  ### 场景介绍

会员在平台注册的时候，平台需在网商见证系统开立分账账户。该分账账户用以核算归属于会员的资金份额。账户余额可用于向绑定卡提现等业务场景。

![](media/image11.png){width="6.791666666666667in"
height="4.395833333333333in"}

见证系统在首次开户时默认开通会员基本户。若平台业务场景需要在会员项下开立其他类型的分账账户。可追加开立其他分账账户。如：可在基本户外再开立保证金账户，分账核算会员的保证金。

![](media/image12.png){width="2.4791666666666665in" height="1.15625in"}

### 涉及接口

\[10101\]mybank.tc.user.personal.register (开个人户)

\[10102\]mybank.tc.user.enterprise.register (开企业户)

\[10103\]mybank.tc.user.account.create (开户)

\[10201\]mybank.tc.user.personal.info.modify (修改个人信息)

\[10202\]mybank.tc.user.enterprise.info.modify (修改企业信息)

\[10301\]mybank.tc.user.personal.info.query (查询个人信息)

\[10302\]mybank.tc.user.enterprise.info.query(查询企业信息)

\[10303\]mybank.tc.user.account.balance(查余额)

![](media/image13.png){width="5.59375in" height="4.71875in"}

### 实现参考

#### 个人开户

![](media/image14.png){width="3.1006003937007875in"
height="3.0555555555555554in"}

1.  会员绑卡
    --------

    1.  ### 场景介绍

会员开户后，
平台可给该会员绑定银行卡。会员提现只能提现到已绑定的银行号。这里的银行卡泛指银行实体账号，含个人存折、对公账号等非银联卡性质的银行账号。

![](media/image15.png){width="5.483333333333333in"
height="2.945790682414698in"}

-   绑定的提现账号支持对公或对私。

-   银行端可根据配置决定会员开户时候证件名（real\_name）是否需与绑定账号的户名（account\_name）一致。

-   一个会员可绑定多个提现账号，但户名必须一致。

-   绑卡时联行号（branch\_no）必输。branch\_no获取流程和规则如下：

<!-- -->

-   步骤1：会员选择绑定对公账号还是个人账号

-   步骤2：若会员选择对公：则提供相关栏位供会员录入“省、市”，然后选择对应支行（见证提供相关接口查询支行联行号和名称），再录入其他信息（账号、户名、手机号等信息）后提交绑卡。

-   步骤3：若会员选择对私：则在会员录入卡号后可调用见证提供的卡bin校验接口获取对应总行联行号和名称，
    若卡bin校验不通过。则提供相关栏位供会员录入”省、市”，然后选择对应支行（见证提供相关接口查询查询支行联行号和名称），再录入其他信息（户名、手机号等信息）后提交绑卡。

![](media/image16.png){width="7.268055555555556in"
height="4.8069444444444445in"}

### 涉及接口

\[10401\]mybank.tc.user.bankcard.bind(绑定银行卡)

[]{#_Toc491271908 .anchor}\[30104\]mybank.tc.user.cardbin.query
(卡BIN核验)

[]{#_Toc491271906 .anchor}\[30101\]mybank.tc.user.area.query
(省、市查询)

[]{#_Toc491271907 .anchor}\[30102\]mybank.tc.user.area.bank.query
(银行列表查询)

\[10402\]mybank.tc.user.bankcard.unbind(解绑银行卡)

### 实现参考

#### 对公绑卡

![](media/image17.png){width="3.3125in" height="5.020833333333333in"}

#### 对私绑卡

![](media/image18.png){width="7.268055555555556in"
height="7.598611111111111in"}

1.  即时入账
    --------

    1.  ### 场景介绍

> [[]{#OLE_LINK38 .anchor}]{#OLE_LINK37
> .anchor}合作方业务平台的买家在业务平台选择商品下单，可选择如下几种支付方式进行交易。
>
> 余额支付：使用买家子账户余额支付。
>
> 快捷支付：平台自接入款渠道的快捷支付。
>
> 网银支付：平台自接入款渠道的网银网关支付。
>
> POS支付：平台自接POS收单机构支付。
>
> 除余额支付外，快捷支付、网银支付、POS支付均支持匿名支付，即买家无见证账户，直接支付给卖家见证账户。
>
> ![](media/image19.png){width="7.268055555555556in"
> height="1.6555555555555554in"}
>
> 若会员选择使用见证账户余额进行支付，则平台需通过即时入账接口将买家见证账户资金结算给卖家见证账户，支付方式余额支付；
>
> 若会员通过其他入款渠道完成付款，则平台需在入款渠道支付成功后，通过即时入账接口直接将支付款项结算给卖家见证账户。支付方式可以为：快捷支付、网银支付或POS支付。此方式资金不留存买家见证账户，直接结算给卖家见证账户。
>
> 若选择折扣支付，则折扣金额从平台的营销费用账户中扣除。

### 涉及接口

\[20101\]mybank.tc.trade.pay.instant (即时交易入账)

\[20108\]mybank.tc.trade.refund (退款入账)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[10303\]mybank.tc.user.account.balance(查余额)

\[20401\]mybank.tc.trade.carryover(资金结转)

![](media/image20.png){width="5.010416666666667in"
height="3.3568963254593176in"}

![](media/image21.png){width="5.083179133858268in"
height="5.943258967629046in"}

1.  担保入账
    --------

    1.  ### 场景介绍

担保交易区别于即时入账，即时入账交易在交易成功后，卖家实时收入资金。而担保交易，则分为担保入账账、担保结算入账和担保退款入账3个环节。应用于担保交易场景，比如：买家支付订单款项后，资金不马上给到卖家，由平台做担保支付，资金先进入到平台担保账号，后续买家确认收到货物后，再由平台将款项支付给卖家。

担保交易资金流示意：

![](media/image22.png){width="6.141666666666667in"
height="1.4416666666666667in"}

a.  担保交易入账：交易成功后，扣减买家资金后，
    > 资金收入平台担保账号，担保账号资金平台无法通过接口直接转账。

b.  担保结算入账：可通过该交易将订单存入担保账号资金结算给给订单的卖家。暂只支持全额结算，结算若涉及分润，可上送任润参数。

c.  担保退款入账：可通过该交易将订单存入担保账号资金退换给买家。
    > 系统支持原路退回，原入账时支付方式为余额支付，则退回买家见证账号，若涉及手续费或分润，则原手续费和分润资金也原路退回
    > 。原入账支付方式为非余额支付（网银支付、快捷支付、POS支付），见证将资金退回到入款渠道对应的银存账号，
    > 若非余额字符，则实际资金退回需平台自行对接的入款渠道处理。

    1.  ### 涉及接口

\[20102\]mybank.tc.trade.pay.ensure (担保交易入账)

\[20103\]mybank.tc.trade.settle(担保交易结算入账)

\[20108\]mybank.tc.trade.refund (退款入账)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[10303\]mybank.tc.user.account.balance(查余额)

\[20401\]mybank.tc.trade.carryover(资金结转)

![](media/image23.png){width="3.9951826334208222in"
height="3.658333333333333in"}

![](media/image24.png){width="4.364969378827647in"
height="5.509743000874891in"}

1.  转账入账
    --------

    1.  ### 场景介绍

若会员账户除了基本户还有其他类型的账户，
平台可通过转账入账的方式实现同一个会员的不同账户之间划转。
比如：会员基本户余额缴纳保证金。

![](media/image25.png){width="3.875in" height="2.1979166666666665in"}

### 涉及接口

\[20104\]mybank.tc.trade.transfer(转账入账)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[10303\]mybank.tc.user.account.balance(查余额)

![](media/image26.png){width="5.165017497812773in"
height="2.2666666666666666in"}

1.  充值记账
    --------

    1.  ### 场景介绍

适用于合作方业务平台接入第三方支付渠道（非网商入金渠道），用户充值，当充值交易在合作方业务平台完成后，调用该接口通知交易见证平台同步充值交易。

![](media/image27.png){width="5.791666666666667in"
height="1.6458650481189852in"}

### 涉及接口

\[20105\]mybank.tc.trade.deposit(充值记账)

\[20108\]mybank.tc.trade.refund (退款入账)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[10303\]mybank.tc.user.account.balance(查余额)

\[20401\]mybank.tc.trade.carryover(资金结转)

![](media/image28.png){width="4.954426946631671in" height="5.51875in"}

1.  提现
    ----

    1.  ### 场景介绍

会员见证账户余额可通过提现的方式将实际资金转入绑定的实体银行账号。

![](media/image15.png){width="5.483333333333333in"
height="2.945790682414698in"}

1.  会员未绑定银行账号，则不允许发起提现。

2.  会员提现不保证实时到账，
    网商银行准实时发出，具体到账时间以收款银行实际入账处理为准。

3.  对公账户大于5万提现走人民银行大额支付通道。只支持在工作日09:00-17:00之间发起（有可能根据人行通知调整）。

4.  若平台绑定提现账户时，有做三四要素鉴权验证，则提现成功率会很高，反之则无法避免“卡号不存在”，“账户户名不符”等非网商银行原因导致的提现失败。

5.  提现为异步处理。提现结果以银行异步通知为准。

6.  当提现转出成功（20106异步通知处理状态为成功），但由于收款银行方面处理失败时，导致资金退回网商银行，需由平台发起退票操作。退票只支持全额退票，收取的手续费也会全额退回。平台可调用20109接口将退票信息提交见证系统。

    1.  ### 涉及接口

\[20106\]mybank.tc.trade.paytocard(单笔提现)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[10303\]mybank.tc.user.account.balance(查余额)

\[20109\]mybank.tc.trade.refundticket (退票)

![](media/image29.png){width="6.135416666666667in"
height="5.135416666666667in"}

1.  冻结解冻
    --------

    1.  ### 场景介绍

为保障交易能正常进行，在交易处理前，调用此接口冻结相关会员账户的资金。也可以用于避免交易出现风险时，冻结会员账户中的资金。平台可调用10303查询冻结余额，
冻结余额=账面余额-可用余额。

### 涉及接口

\[20104\]mybank.tc.trade.transfer(转账入账)

\[20201\]mybank.tc.trade.query (交易流水查询)

\[20202\]mybank.tc.trade.info.query (交易详情查询)

\[10303\]mybank.tc.user.account.balance(查余额)

\[20301\]mybank.tc.trade.funds.freeze (冻结资金)

\[20302\]mybank.tc.trade.funds.unfreeze (解冻资金)

![](media/image30.png){width="4.907327209098862in" height="3.3in"}

1.  基础服务接口
    ------------

    1.  ### 场景介绍

会员银行卡绑定时可以查询省、市、银行列表；还可以查询卡bin信息。

### 涉及接口

> \[30101\] mybank.tc.user.area.query (省、市、区查询)
>
> \[30102\] mybank.tc.user.area.bank.query (银行列表查询)
>
> \[30103\] mybank.tc.user.cardbin.query (卡bin查询)

1.  接口规范
    ========

    1.  企业要求
        --------

-   开通网商银行企业账户和企业网银；

-   与网商银行签订银企直连协议；

-   申请天威诚信网商银行企业证书。

    1.  接口说明
        --------

1、合作方业务平台和交易见证平台之间通过https协议来进行通信，接口以URL的形式提供以post的请求方式处理，接口说明中描述了post的请求参数。

文档的接口包括两种：服务接口和通知接口。服务接口由支付系统提供，供商户调用；通知接口由商户提供，供支付系统调用。虽然通知接口由商户提供，但是仍由支付系统制定接口规范。服务接口包括接口说明中的所有接口信息；通知接口目前仅包括交易状态和退款状态的通知等。

2、接口参数中金额单位统一为”元”。

调用地址
--------

SIT环境调用地址：http://test.tc.mybank.cn/gop/gateway.do

调用示例
--------

**import** java.io.BufferedReader;

**import** java.io.IOException;

**import** java.io.InputStream;

**import** java.io.InputStreamReader;

**import** java.util.ArrayList;

**import** java.util.HashMap;

**import** java.util.List;

**import** java.util.Map;

**import** org.apache.http.HttpResponse;

**import** org.apache.http.NameValuePair;

**import** org.apache.http.client.HttpClient;

**import** org.apache.http.client.entity.UrlEncodedFormEntity;

**import** org.apache.http.client.methods.HttpPost;

**import** org.apache.http.impl.client.HttpClientBuilder;

**import** org.apache.http.message.BasicNameValuePair;

**public** **class** GatewayClientSample {

**private** **static** HttpClient *httpClient* =
HttpClientBuilder.*create*().build();

**public** **static** **void** main(String\[\] args) {

String gatewayUrl = "http://test.tc.mybank.cn/gop/gateway.do";

/\*

\* 请求参数设置详见各接口

\*\*/

*doHttpClientPost*(gatewayUrl, params);

}

**public** **static** String doHttpClientPost(String url, Map&lt;String,
String&gt; params) {

HttpResponse response;

**try** {

HttpPost post = **new** HttpPost(url);

List&lt;NameValuePair&gt; ps = *buildPostParams*(params);

post.setEntity(**new** UrlEncodedFormEntity(ps, "UTF-8"));

response = *httpClient*.execute(post);

**if** (response.getStatusLine().getStatusCode() == 200) {//
网关调用成功

String rst = *inputStreamToStr*(response.getEntity().getContent(),
"UTF-8");

System.***out***.println("=======================================");

System.***out***.println(String.*format*("httpClient Post调用结果：%s",
rst));

System.***out***.println("=======================================");

**return** rst;

}

} **catch** (Exception e) {

System.***out***.println("=======================================");

System.***out***.println(String.*format*("httpClient Post 请求失败：{}",
e));

System.***out***.println("=======================================");

}

**return** **null**;

}

**private** **static** List&lt;NameValuePair&gt;
buildPostParams(Map&lt;String, String&gt; params) {

**if** (params == **null** || params.size() == 0)

**return** **null**;

List&lt;NameValuePair&gt; results = **new**
ArrayList&lt;NameValuePair&gt;();

**for** (Map.Entry&lt;String, String&gt; entry : params.entrySet()) {

String key = entry.getKey();

String value = entry.getValue();

results.add(**new** BasicNameValuePair(key, value));

}

**return** results;

}

**private** **static** String inputStreamToStr(InputStream is, String
charset) **throws** IOException {

BufferedReader in = **new** BufferedReader(**new** InputStreamReader(is,
"ISO-8859-1"));

StringBuffer buffer = **new** StringBuffer();

String line = "";

**while** ((line = in.readLine()) != **null**) {

buffer.append(line);

}

**return** **new** String(buffer.toString().getBytes("ISO-8859-1"),
charset);

}

}

注：依赖包

&lt;dependency&gt;

&lt;groupId&gt;org.apache.httpcomponents&lt;/groupId&gt;

&lt;artifactId&gt;*httpclient*&lt;/artifactId&gt;

&lt;version&gt;4.5.2&lt;/version&gt;

&lt;/dependency&gt;

[]{#_Toc435460930 .anchor}

1.  API列表
    =======

    1.  会员API
        -------

  **功能**       **编号**   **接口名**       **API列表**
  -------------- ---------- ---------------- ---------------------------------------
  开户           10101      开个人户         mybank.tc.user.personal.register
                 10102      开企业户         mybank.tc.user.enterprise.register
                 10103      开户             mybank.tc.user.account.create
  修改会员信息   10201      修改个人信息     mybank.tc.user.personal.info.modify
                 10202      修改企业信息     mybank.tc.user.enterprise.info.modify
  查询           10301      查询个人信息     mybank.tc.user.personal.info.query
                 10302      查询企业信息     mybank.tc.user.enterprise.info.query
                 10303      查询余额         mybank.tc.user.account.balance
  绑定银行卡     10401      绑定银行卡       mybank.tc.user.bankcard.bind
                 10402      解绑银行卡       mybank.tc.user.bankcard.unbind
                 10403      查询银行卡列表   mybank.tc.user.bankcard.query
                 10404      上传文件         mybank.tc.user.info.img.upload

支付API
-------

  **功能**        **编号**   **接口名**     **API列表**
  --------------- ---------- -------------- --------------------------------
  支付            20101      即时交易       mybank.tc.trade.pay.instant
                  20102      担保交易       mybank.tc.trade.pay.ensure
                  20103      担保交易结算   mybank.tc.trade.settle
                  20105      充值           mybank.tc.trade.deposit
                  20106      单笔提现       mybank.tc.trade.paytocard
                  20108      退款入账       mybank.tc.trade.refund
                  20109      退票           mybank.tc.trade.refundticket
  交易查询        20201      交易流水查询   mybank.tc.trade.query
                  20202      交易详情查询   mybank.tc.trade.info.query
  资金冻结/解冻   20301      冻结资金       mybank.tc.trade.funds.freeze
                  20302      解冻资金       mybank.tc.trade.funds.unfreeze
  结转            20401      资金结转       mybank.tc.trade.carryover

基础API
-------

  **功能**           **编号**   **接口名**       **API列表**
  ------------------ ---------- ---------------- --------------------------------------
                     30101      省、市查询       mybank.tc.user.area.query
  绑卡基础信息查询   30102      银行列表查询     mybank.tc.user.area.bank.query
                     30103      卡bin查询        mybank.tc.user.cardbin.query
  合作方余额查询     30104      合作方余额查询   mybank.tc.user.partner.balance.query

1.  API说明
    =======

    1.  公共请求参数
        ------------

请求基本参数是下面所有请求接口都要加上的。

  -----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数**      **类型（长度）**   **是否必填**   **描述**                                                                                                     **样例**
  ------------- ------------------ -------------- ------------------------------------------------------------------------------------------------------------ ----------------------------------
  version       Number(5)          不可空         接口版本，设置为2.0                                                                                          2.0

  partner\_id   String(32)         不可空         交易见证平台分配给合作方业务平台签约唯一ID                                                                   208800115994003

  charset       String(10)         不可空         合作方业务平台使用的编码格式，如utf-8、gbk、gb2312等                                                         GBK

  sign          String(256)        不可空         商户生成签名字符串所使用的签名算法类型，目前只支持天威证书签名（使用天威诚信网商银行企业证书对数据签名）。   TWSIGN
                                                                                                                                                               
                                                  除了sign、sign\_type两个字段，以及空和空白字段外，其它字段都会参与算签。                                     e8qdwl9caset5zugii2r7q0k8ikopxor

  sign\_type    String(10)         不可空         签名方式只支持TWSIGN                                                                                         TWSIGN

  service       String(64)         不可空         API接口名称                                                                                                  mybank.tc.user.personal.register

  memo          String(1000)       可空           备注信息                                                                                                     
  -----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

公共响应参数
------------

部分接口调用接口成功时，会同步返回。

注意：这里的成功只是调用接口成功，不代表业务的成功，业务是否成功，依赖异步通知消息。同步返回时，需要的基本参数：

  参数             类型（长度）   是否必填   描述                                                             样例
  ---------------- -------------- ---------- ---------------------------------------------------------------- -------------------------
  is\_success      String(1)      不可空     成功标识。表示接口调用是否成功，并不表明业务处理结果。           T
  charset          String(10)     不可空     参数编码字符集。业务平台使用的编码格式，如utf-8、gbk、gb2312等   GBK
  error\_code      String(30)     可空       返回错误码参见附录                                               PARTNER\_ID\_NOT\_EXIST
  error\_message   String(200)    可空       返回错误原因参见附录                                             合作方Id不存在
  memo             String(1000)   可空       说明信息                                                         

1.  \[10101\]mybank.tc.user.personal.register (开个人户) 
    -----------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口创建个人会员。

1.  默认开通个人会员的BASIC-基本户（账户类型account\_type参考附件枚举）。

2.  uid为合作方业务平台自定义，字母或数字，不能重复，不建议使用手机号作为uid。重复提交的开户请求根据uid
    作幂等返回成功。接口若调用超时，平台可重新使用相同uid重复提交。

3.  接口返回的member\_id，暂保留使用。

4.  接口调用超时，平台也可调用10301接口查询银行处理结果并记录返回的member\_id。

5.  个人会员数量没有限制。

    1.  ### 业务请求参数

  参数                类型（长度）   是否必填   描述                                     样例
  ------------------- -------------- ---------- ---------------------------------------- -----------------------------------------------------
  uid                 String(32)     不可空     合作方业务平台的用户ID                   10010020304
  mobile              String(11)     可空       合作方业务平台用户手机号                 13813252598
  email               String(30)     可空       邮箱号。                                 [*example@example.com*](mailto:example@example.com)
  real\_name          String(50)     不可空     真实姓名                                 许三多
  member\_name        String(50)     不可空     会员名称。用户昵称(平台个人会员登录名)   angel
  certificate\_type   String(20)     不可空     证件类型（见附录）。目前只支持身份证。   ID\_CARD
  certificate\_no     String(50)     不可空     作为会员实名认证通过后的证件号           210213198602165207
  is\_verify          String(1)      可空       是否认证                                 Y:是 N:否
  is\_active          String(1)      可空       预留字段，可不填                         格式T,F

### 业务响应参数

  参数               类型（长度）   是否必填   描述     样例
  ------------------ -------------- ---------- -------- ---------------
  member\_id         String(12)     不可空     会员号   1000000012345
  sub\_account\_no   String(10)     不可空     子账户   4992739872

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("member\_name", "59a2x");

params.put("mobile", "13674401802");

params.put("partner\_id", "200000900001");

params.put("real\_name", "3uu8r");

params.put("service", "mybank.tc.user.personal.register");

params.put("sign", "gt8D96yo3KT…");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","member\_id":"100000900005","sub\_account\_no":"4992739872"}

异常：{"charset":"UTF-8","error\_code":"IDENTITY\_EXIST\_ERROR","error\_message":"会员标识已经存在","is\_success":"F"}

1.  \[10102\]mybank.tc.user.enterprise.register (开企业户)
    ------------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口创建企业会员。

1.  默认开通企业会员的BASIC-基本户（账户类型account\_type参考附录账户类型枚举）。

2.  uid为合作方业务平台自定义，字母或数字，不能重复，不建议使用手机号作为uid。重复提交的开户请求根据uid
    作幂等返回成功。接口若调用超时，平台可重新使用相同uid重复提交。

3.  接口返回的member\_id，暂保留使用。

4.  接口调用超时，平台也可调用10301接口查询银行处理结果并记录返回的member\_id。

5.  个人会员数量没有限制。

    1.  ### 业务请求参数

  参数                               类型（长度）   是否必填   描述                                                                样例
  ---------------------------------- -------------- ---------- ------------------------------------------------------------------- -------------------------------
  uid                                String(32)     不可空     合作方业务平台用户ID                                                10010020304
  enterprise\_name                   String(200)    不可空     企业名称                                                            
  member\_name                       String(50)     可空       企业简称                                                            
  legal\_person                      String(50)     可空       企业法人姓名                                                        
  legal\_person\_certificate\_type   String(20)     可空       法人证件类型                                                        见附录[*证件类型*](#证件类型)
  legal\_person\_certificate\_no     String(50)     可空       法人证件号                                                          
  legal\_person\_phone               String(11)     可空       法人手机号码                                                        13812345600
  website                            String(128)    可空       企业网址                                                            
  address                            String(200)    可空       企业地址                                                            
  license\_no                        String(15)     可空       执照号,若营业执照号和统一社会信用代码都填写，则取统一社会信用代码   若做企业实名认证则不可空
  license\_address                   String(200)    可空       营业执照所在地                                                      
  license\_expire\_date              String(8)      可空       执照过期日（营业期限）yyyymmdd                                      20100131
  business\_scope                    String(1024)   可空       营业范围                                                            
  telephone                          String(20)     可空       联系电话                                                            
  organization\_no                   String(128)    可空       组织机构代码                                                        
  unified\_social\_credit\_code      String(18)     可空       统一社会信用代码                                                    若做企业实名认证则不可空
  summary                            String(1024)   可空       企业简介                                                            
  open\_account\_license             String(32)     可空       开户许可证                                                          
  is\_verify                         String(1)      可空       是否认证                                                            Y:是 N:否
  login\_name                        String(50)     可空       预留字段，可不填                                                    
  is\_active                         String(1)      可空       预留字段，可不填                                                    T

### 业务响应参数

  参数               类型（长度）   是否必填   描述     样例
  ------------------ -------------- ---------- -------- --------------
  member\_id         String(12)     不可空     会员号   200000001234
  sub\_account\_no   String(10)     不可空     子账户   0000053587

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("address", "3pfszesj11");

params.put("business\_scope", "8zb041y842");

params.put("enterprise\_name", "31foq0aqrsbbrig0eg9w");

params.put("legal\_person", "1ac91");

params.put("legal\_person\_certificate\_no", "1089793429234934940");

params.put("legal\_person\_certificate\_type", "ID\_CARD");

params.put("legal\_person\_phone", "13146047508");

params.put("license\_address", "9k1z95nj5n");

params.put("license\_expire\_date", "20150118");

params.put("license\_no", "u65l9550l6");

params.put("member\_name", "y98v7");

params.put("open\_account\_license", "5385195217");

params.put("organization\_no", "7583062346");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.enterprise.register ");

params.put("sign", "udiLvxzuR1....");

params.put("sign\_type", "RSA");

params.put("summary", "2l920t97");

params.put("telephone", "14868838");

params.put("uid", "30cmjp7p8y");

params.put("unified\_social\_credit\_code", "40k2c993qg");

params.put("version", "2.0");

params.put("website", "23638xf1b5");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","member\_id":"200000900007","sub\_account\_no":"0000053587"}

异常：{"charset":"UTF-8","error\_code":"IDENTITY\_EXIST\_ERROR","error\_message":"会员标识已经存在","is\_success":"F"}

1.  \[10103\]mybank.tc.user.account.create (开户)
    ---------------------------------------------

    1.  ### 接口说明

> 平台调用该接口创建会员基本户之外的其他账户。

1.  同一个账户类型只能有一个账户。

2.  uid为合作方业务平台自定义，字母或数字，不能重复，不建议使用手机号作为uid。重复提交的开户请求根据uid
    作幂等返回成功。接口若调用超时，平台可重新使用相同uid重复提交。

3.  接口返回的account\_id，暂保留使用。

4.  接口调用超时，平台也可调用103031口查询银行处理结果。

5.  会员账户类型不能以P开头。

    1.  ### 业务请求参数

  参数                类型（长度）   是否必填   描述                                 样例
  ------------------- -------------- ---------- ------------------------------------ -------------
  uid                 String(32)     不可空     合作方业务平台用户ID                 10010020304
  account\_type       String(20)     不可空     账户类型                             DEPOSIT
  alias               String(10)     不可空     账户别名                             保证金户
  account\_identity   String(32)     可空       账户标识，用以区分同类型的多个账户   

### 业务响应参数

  参数           类型（长度）   是否必填   描述           样例
  -------------- -------------- ---------- -------------- -----------------------------
  account\_id    String(32)     不可空     会员账户号     200100100210000369876500001
  create\_time   String(19)     不可空     账户创建时间   2015-07-09 17:46:42

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("account\_type", "BASIC");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.account.create");

params.put("sign", "2tr4o1nbz79nj....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","account\_id":"20010010011100000900005000002","create\_time":"2017-06-28
21:00:40","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"SYSTEM\_ERROR","error\_message":"系统内部错误","is\_success":"F","memo":"参数错误\[账户类型非法\]"}

1.  \[10401\]mybank.tc.user.bankcard.bind(绑定银行卡)
    -------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口给会员绑定银行卡。

1.  这里的银行卡泛指银行实体账号，含个人存折、对公账号等非银联卡性质的银行账号。

2.  同一个会员可绑定多张卡。但银行开户名须保持一致，否则接口返回失败。

3.  接口并不拒绝绑定贷记卡或准贷记卡，平台根据自己的业务场景选用，若无特殊场景要求支持贷记卡，一般建议在平台端控制不提交贷记卡进行绑定。

4.  调用该接口前可调用银行提供的联行号查询相关接口获取bank\_name、bank\_branch、branch\_no信息。

5.  平台也可以不调用银行提供的联行号查询接口，而使用本地数据自行匹配，平台若使用本地数据，需确保数据的正确性和完整性。

6.  平台需要记录接口返回的bank\_id。后续其他接口， 如：提现等，
    统一使用bank\_id做绑卡标识。

7.  绑卡成功后，若需要修改绑卡信息，可通过解绑，然后再绑定的方式实现。

8.  绑定银行卡，不能是银行二、三类账户。必须是在银行临柜开立的结算账户。

    1.  ### 业务请求参数

  ----------------------------------------------------------------------------------------------------------------------------
  参数                类型（长度）   是否必填   描述                                           样例
  ------------------- -------------- ---------- ---------------------------------------------- -------------------------------
  uid                 String(32)     不可空     合作方业务平台用户ID                           10010020304

  bank\_name          String(50)     不可空     银行全称                                       中国工商银行

  bank\_branch        String(255)    可空       支行名称                                       中国工商银行南京南山支行

  branch\_no          String(12)     不可空     联行号获取规请参考开发[*指引8.3*](#会员绑卡)   313703018022

  bank\_account\_no   String(32)     不可空     银行账号/卡号，明文。                          6225885425698745

  account\_name       String(50)     不可空     银行开户名，明文。                             测试用户

  card\_type          String(10)     不可空     卡类型                                         DC
                                                                                               
                                                DC 借记                                        
                                                                                               
                                                CC 贷记（信用卡）                              
                                                                                               
                                                PB 存折                                        
                                                                                               
                                                OC 其他                                        

  card\_attribute     String(10)     不可空     卡属性：                                       C
                                                                                               
                                                C 对私                                         
                                                                                               
                                                B 对公                                         

  verify\_type        String(1)      可空       验证类型                                       3:三要素验证
                                                                                               
                                                                                               4:四要素验证

  certificate\_type   String(20)     可空       证件类型                                       见附录[*证件类型*](#证件类型)

  certificate\_no     String(50)     可空       证件号                                         

  province            String(128)    可空       省份                                           江苏省

  city                String(128)    可空       城市                                           南京市

  reserved\_mobile    String(11)     可空       银行预留手机号                                 13788888888

  pay\_attribute      String(10)     可空       支付属性：                                     NORMAL
                                                                                               
                                                NORMAL 普通(默认)                              
  ----------------------------------------------------------------------------------------------------------------------------

### 业务响应参数

  参数       类型（长度）   是否必填   描述                                     样例
  ---------- -------------- ---------- ---------------------------------------- ------
  bank\_id   String(32)     不可空     银行卡id，银行卡在交易见证平台的绑卡id   

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("account\_name", "测试用户");

params.put("bank\_account\_no", "6225885425698745");

params.put("bank\_branch", "浦东支行");

params.put("bank\_code", "CMB");

params.put("bank\_name", "招商银行");

params.put("branch\_no", "123212524541");

params.put("card\_attribute", "C");

params.put("card\_type", "DC");

params.put("city", "上海市");

params.put("partner\_id", "200000900001");

params.put("pay\_attribute", "normal");

params.put("province", "上海市");

params.put("reserved\_mobile", "13866668888");

params.put("service", "mybank.tc.user.bankcard.bind");

params.put("sign", "olVuorUy6Tbne....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");;

### 响应示例

正常：{"charset":"UTF-8","bank\_id":"2526","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"FIXED\_SIZE\_ERROR\\":\[\\"支行号应为12位长度不合法\\"\]}"}

1.  \[10402\]mybank.tc.user.bankcard.unbind(解绑银行卡)
    ---------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口给会员解绑已绑定的银行卡。

1.  这里的银行卡泛指银行实体账号，含个人存折、对公账号等非银联卡性质的银行账号。

2.  解绑前，平台可调用10403接口查询已绑定的银行卡。

3.  解绑后，系统删除该银行卡与会员的绑定关系。提现接口不支持已解绑的银行卡提现。

    1.  ### 业务请求参数

  参数       类型（长度）   是否必填   描述                                     样例
  ---------- -------------- ---------- ---------------------------------------- --------------
  uid        String(32)     不可空     合作方业务平台用户ID                     100210054875
  bank\_id   String(32)     不可空     银行卡id，银行卡在交易见证平台的绑卡id   10023

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("bank\_id", "2526");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.bankcard.unbind");

params.put("sign", "mANIS1PpqbKU....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"REQUIRED\_FIELD\_NOT\_EXIST\\":\[\\"银行卡ID不能为空\\"\]}"}

1.  \[10201\]mybank.tc.user.personal.info.modify (修改个人信息)
    -----------------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口修改会员相关信息。

1.  除uid外，其他要素均可修改。

    1.  ### 业务请求参数

  ---------------------------------------------------------------------------------------------------------------------------------------------------------------
  参数                类型（长度）   是否必填   描述                                                        样例
  ------------------- -------------- ---------- ----------------------------------------------------------- -----------------------------------------------------
  uid                 String(32)     不可空     合作方业务平台用户ID                                        1001200002460019

  real\_name          String(50)     可空       真实姓名。                                                  
                                                                                                            
                                                当可空时，不修改原有信息                                    

  member\_name        String(50)     可空       用户昵称(平台个人会员登录名) 。                             
                                                                                                            
                                                当可空时，不修改原有信息                                    

  certificate\_type   String(20)     可空       证件类型，暂只支持身份证。见附录*[证件类型](#证件类型)。*   

  certificate\_no     String(50)     可空       作为会员实名认证通过后的身份证号。                          
                                                                                                            
                                                当可空时，不修改原有信息                                    

  mobile              String(11)     可空       会员手机号                                                  13788888888
                                                                                                            
                                                当可空时，不修改原有信息                                    

  email               String(30)     可空       邮箱号。当手机号为空时，邮箱不能为空                        [*example@example.com*](mailto:example@example.com)
  ---------------------------------------------------------------------------------------------------------------------------------------------------------------

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("member\_name", "es2f0");

params.put("partner\_id", "200000900001");

params.put("real\_name", "es2f0");

params.put("service", " mybank.tc.user.personal.info.modify");

params.put("sign", "xhsUZZPbdfZOv9....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","member\_id":"100000900005"}

异常：{"charset":"UTF-8","error\_code":"MEMBER\_NOT\_EXIST","error\_message":"会员不存在","is\_success":"F"}

1.  \[10202\]mybank.tc.user.enterprise.info.modify (修改企业信息)
    -------------------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口修改会员相关信息。

1.  除uid外，其他要素均可修改。

    1.  ### 业务请求参数

  参数                               类型（长度）   是否必填   描述                             样例
  ---------------------------------- -------------- ---------- -------------------------------- -------------------------------
  uid                                String(32)     不可空     合作方业务平台用户ID             200000491600
  enterprise\_name                   String(200)    可空       企业名称                         
  member\_name                       String(50)     可空       会员名称                         
  legal\_person                      String(50)     可空       企业法人                         
  legal\_person\_phone               String(11)     可空       法人手机号码                     13788888888
  legal\_person\_certificate\_type   String(20)     可空       法人证件类型                     见附录[*证件类型*](#证件类型)
  legal\_person\_certificate\_no     String(50)     可空       法人证件号                       
  website                            String(128)    可空       企业网址                         
  address                            String(200)    可空       企业地址                         
  license\_no                        String(128)    可空       执照号                           
  license\_address                   String(200)    可空       营业执照所在地                   
  license\_expire\_date              String(8)      可空       执照过期日（营业期限）yyyymmdd   20180101
  business\_scope                    String(1024)   可空       营业范围                         
  telephone                          String(20)     可空       联系电话                         
  organization\_no                   String(128)    可空       组织机构代码                     
  open\_account\_license             String(32)     可空       开户许可证                       
  unified\_scoial\_credit\_code      String(128)    可空       统一社会信用代码                 
  Summary                            String(1024)   可空       企业简介                         

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("address", "9t2h6v4c70");

params.put("business\_scope", "78pqtd27k7");

params.put("enterprise\_name", "o114h00mk423ve9tu800");

params.put("legal\_person\_phone", "18880034817");

params.put("license\_address", "09ifa25s5t");

params.put("license\_expire\_date", "20150116");

params.put("license\_no", "80x154ax57");

params.put("member\_name", "938nf");

params.put("organization\_no", "3198565729");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.enterprise.info.modify");

params.put("sign", "Dap6GRPdz8NnCB....");

params.put("sign\_type", "RSA");

params.put("summary", "2h6sfqrh");

params.put("telephone", "54288002");

params.put("uid", "56p7usf46c");

params.put("version", "2.0");

params.put("website", "0x0yxh9p6m");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","member\_id":"200000900003"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"FIELD\_TYPE\_ERROR\\":\[\\"法人证件\\"\]}"}

1.  \[10301\]mybank.tc.user.personal.info.query (查询个人信息)
    ----------------------------------------------------------

    1.  ### 接口说明

> 平台调用该接口查询个人会员信息。

### 业务请求参数

  参数   类型（长度）   是否必填   描述                   样例
  ------ -------------- ---------- ---------------------- --------------
  Uid    String(32)     不可空     合作方业务平台用户ID   100000795683

### 业务响应参数

  参数                类型（长度）   描述                                      样例
  ------------------- -------------- ----------------------------------------- -----------------------
  Uid                 String(32)     合作方业务平台用户ID                      100000795683
  Mobile              String(11)     手机号                                    13906889493
  Email               String(30)     邮箱号                                    <example@example.com>
  real\_name          String(50)     真实姓名                                  
  member\_name        String(50)     用户昵称(平台个人会员登录名)              
  certificate\_type   String(20)     证件类型，见附录[*证件类型*](#证件类型)   
  certificate\_no     String(50)     作为会员实名认证通过后的身份证号          
  member\_id          String(32)     会员Id                                    

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.personal.info.query");

params.put("sign", "A3vhlJSakixiB1JG....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应实例

正常：{"charset":"UTF-8","is\_success":"T","member\_id":"100000900005","member\_name":"es2f0","mobile":"13674401802","real\_name":"es2f0","uid":"02196aur2e"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_SIGN","error\_message":"验签未通过","is\_success":"F"}

1.  \[10302\]mybank.tc.user.enterprise.info.query(查询企业信息)
    -----------------------------------------------------------

    1.  ### 接口说明

        平台调用该接口查询个人会员信息。

    2.  ### 业务请求参数

  参数   类型（长度）   是否必填   描述                   样例
  ------ -------------- ---------- ---------------------- --------------
  Uid    String(32)     不可空     合作方业务平台用户ID   200000491600

### 业务响应参数

  参数                               类型（长度）   描述                                          样例
  ---------------------------------- -------------- --------------------------------------------- --------------
  Uid                                String(32)     合作方业务平台用户ID                          200001940062
  enterprise\_name                   String(200)    企业名称                                      
  member\_name                       String(50)     会员名称                                      
  legal\_person                      String(50)     企业法人                                      
  legal\_person\_phone               String(11)     法人手机号码                                  13812340945
  legal\_person\_certificate\_type   String(20)     法人证件类型，见附录[*证件类型*](#证件类型)   
  legal\_person\_certificate\_no     String(50)     法人证件号                                    
  website                            String(128)    企业网址                                      
  address                            String(200)    企业地址                                      
  license\_no                        String(128)    执照号                                        
  license\_address                   String(200)    营业执照所在地                                
  license\_expire\_date              String(8)      执照过期日（营业期限）yyyymmdd                20100131
  business\_scope                    String(1024)   营业范围                                      
  telephone                          String(20)     联系电话                                      
  organization\_no                   String(128)    组织机构代码                                  
  open\_account\_license             String(32)     开户许可证                                    
  unified\_scoial\_credit\_code      String(128)    统一社会信用代码                              
  summary                            String(1024)   企业简介                                      

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.enterprise.info.query");

params.put("sign", "DdYg027HqBTOH....");

params.put("sign\_type", "RSA");

params.put("uid", "30cmjp7p8y");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","address":"3pfszesj11","business\_scope":"8zb041y842","enterprise\_name":"31foq0aqrsbbrig0eg9w","is\_success":"T","legal\_person":"1ac91","legal\_person\_certificate\_no":"1089793429234934940","legal\_person\_certificate\_type":"ID\_CARD","legal\_person\_phone":"13146047508","license\_address":"9k1z95nj5n","license\_expire\_date":1421510400000,"license\_no":"u65l9550l6","member\_id":"200000900007","member\_name":"y98v7","open\_account\_license":"5385195217","organization\_no":"7583062346","summary":"2l920t97","telephone":"14868838","uid":"30cmjp7p8y","unified\_scoial\_credit\_code":"40k2c993qg","website":"23638xf1b5"}

异常：{"charset":"UTF-8","error\_code":"MEMBER\_NOT\_EXIST","error\_message":"会员不存在","is\_success":"F"}

1.  \[10303\]mybank.tc.user.account.balance(查余额)
    -----------------------------------------------

    1.  ### 接口说明

平台可调用该接口根据会员余额

1.  根据uid和account\_type查询账户余额。若指定account\_type，则获取获取单个账户余额；反之查询改会员项下所有账户余额。

    1.  ### 业务请求参数

  参数                类型（长度）   是否必填   描述                                       样例
  ------------------- -------------- ---------- ------------------------------------------ ------------
  Uid                 String(32)     不可空     合作方业务平台用户ID                       U500000001
  account\_type       String(20)     可空       账户类型，若为空，则默认查询会员所有账户   BASIC
  account\_identity   String(32)     可空       预留字段，可不填(账户标识)                 

### 业务响应参数

  参数                  类型（长度）   是否必填   描述                                               样例
  --------------------- -------------- ---------- -------------------------------------------------- -----------------------------
  Uid                   String(32)     不可空     合作方业务平台用户ID                               U500000001
  account\_list                                   账户列表                                           
  └account\_id          String(25)     不可空     会员账户号                                         200100100210000369876500001
  └account\_type        String(20)     不可空     账户类型                                           BASIC
  └balance              Number         不可空     账面余额。单位为：RMB Yuan。精确到小数点后两位。   10.00
  └available\_balance   Number         不可空     可用余额。单位为：RMB Yuan。精确到小数点后两位。   8.00
  └sub\_account\_no     String(10)     可空       子账户号                                           4992739872

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
*String*&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.account.balance");

params.put("sign", "19VpdIS4....");

params.put("sign\_type", "RSA");

params.put("uid", "56p7usf46c");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","account\_list":\[{"account\_id":"20080010011200000900003000001","account\_type":"BASIC","available\_balance":"0.00","balance":"0.00","sub\_account\_no":"0000000042"}\],"is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"MEMBER\_NOT\_EXIST","error\_message":"会员不存在","is\_success":"F"}

1.  \[10403\]mybank.tc.user.bankcard.query(查询银行卡列表)
    ------------------------------------------------------

    1.  ### 接口说明

        平台可调用该接口查询会员绑定的银行卡。

<!-- -->

1.  这里的银行卡泛指银行实体账号，含个人存折、对公账号等非银联卡性质的银行账号。

2.  若查询无记录，接口返回失败。

    1.  使用场景

        合作方业务平台，调用该接口可以查询交易见证平台会员绑定的银行卡列表。

    <!-- -->

    1.  ### 业务请求参数

  **参数**   **类型（长度）**   **是否必填**   **描述**                            **样例**
  ---------- ------------------ -------------- ----------------------------------- --------------
  uid        String(32)         不可空         合作方业务平台用户ID (字母或数字)   100210054875

### 业务响应参数

  -------------------------------------------------------------------------------------------------------------------------------
  **参数**             **类型（长度）**   **是否必填**   **描述**                             **样例**
  -------------------- ------------------ -------------- ------------------------------------ -----------------------------------
  bank\_list                                             有效的银行卡列表                     

  └bank\_id            String(10)         不可空         银行卡id                             100056

  └bank\_code          String(10)         不可空         银行编号 [*见附录*](#银行总行列表)   ICBC,CCB

  └bank\_name          String(50)         不可空         银行全称 [*见附录*](#银行总行列表)   中国工商银行

  └bank\_account\_no   String(32)         不可空         银行卡号。明文                       

  └account\_name       String(50)         不可空         开户名。明文                         

  └province            String(128)        不可空         省份                                 上海市，江苏省

  └city                String(128)        不可空         城市                                 上海市，南京市

  └branch\_no          String(12)         不可空         支行号（联行号）                     

  └bank\_branch        String(50)         不可空         分支行名称                           

  └reserved\_mobile    String(11)         可空           预留手机号                           13788888888

  └pay\_attribute      String(10)         可空           卡支付属                             NORMAL

  └card\_attribute     String(10)         可空           卡属性                               C 对私
                                                                                              
                                                                                              B 对公

  └card\_type          String(10)         可空           卡类型                               卡类型
                                                                                              
                                                                                              DC 借记
                                                                                              
                                                                                              CC 贷记（信用卡）
                                                                                              
                                                                                              PB 存折
                                                                                              
                                                                                              OC 其他

  └certificate\_no     String(50)         可空           证件号                               

  └certificate\_type   String(20)         可空           证件类型                             见附录[*证件类型*](#证件类型)

  └is\_verified        String(20)         可空           验证状态                             UNVERIFIED:未认证 VERIFIED:已认证
  -------------------------------------------------------------------------------------------------------------------------------

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.bankcard.query");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("uid", "02196aur2e");

params.put("version", "2.0");

### 响应示例

正常：{"charset":"UTF-8","bank\_list":\[{"account\_name":"张三","bank\_account\_no":"6225880134287138","bank\_branch":"中国农业银行股份有限公司南京高淳新区支行","bank\_code":"ABC","bank\_id":"2535","bank\_name":"中国农业银行股份有限公司","branch\_no":"103301012525","card\_attribute":"C","card\_type":"DC","certificate\_no":"430681198508229015","certificate\_type":"ID\_CARD","city":"南京市","is\_verified":"VERIFIED","pay\_attribute":"NORMAL","province":"江苏省","reserved\_mobile":"13810610842"},{"account\_name":"张三","bank\_account\_no":"6225880134287143","bank\_branch":"中国农业银行股份有限公司南京高淳新区支行","bank\_code":"ABC","bank\_id":"2633","bank\_name":"中国农业银行股份有限公司11","branch\_no":"103301012525","card\_attribute":"C","card\_type":"OC","city":"南京市","is\_verified":"UNVERIFIED","pay\_attribute":"NORMAL","province":"江苏省","reserved\_mobile":"13810610842"}\],"is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"REQUIRED\_FIELD\_NOT\_EXIST\\":\[\\"合作方业务平台用户ID不能为空\\"\]}"}

1.  \[10404\]mybank.tc.user.info.img.upload (文件上传)
    --------------------------------------------------

    1.  ### 接口说明

        平台可调用该接直接上传压缩文件。

<!-- -->

1.  文件上传只支持压缩包格式zip，rar文件。

2.  文件上传需关联业务订单号。

    1.  ### 使用场景

        合作方业务平台在开户时调用该接口可以上传会员开户相关资料压缩包；交易请求时，调用接口上传交易凭证文件压缩包

    2.  ### 业务请求参数

  ---------------------------------------------------------------------------------------------------------------------------------------------------------
  1.  **参数**        **类型（长度）**   **是否必填**   **描述**                                                             **样例**
                                                                                                                             
  ------------------- ------------------ -------------- -------------------------------------------------------------------- ------------------------------
  File                String             不可空         上传的文件内容,需要对文件数组进行base64位加密(文件不超过2M)。        支持的文件类型： zip,rar
                                                                                                                             
                                                        支持上传压缩包                                                       

  file\_name          String(20)         不可空         文件名称（要求带文件后缀），文件命名规则：                           OPEN\_ACCOUNT-01.rar
                                                                                                                             
                                                        文件用途-业务订单号.rar(ip)                                          

  Purpose             String(50)         不可空         文件用途,目前提供如下4个选择：FILE\_PURPOSE\_CETIFICATION-身份认证   FILE\_PURPOSE\_OPEN\_ACCOUNT
                                                                                                                             
                                                        FILE\_PURPOSE\_OPEN\_ACCOUNT-开户                                    
                                                                                                                             
                                                        FILE\_PURPOSE\_ENTRY-登账                                            
                                                                                                                             
                                                        FILE\_PURPOSE\_TRADE\_VOUCHER-交易凭证                               

  rela\_voucher\_no   String(32)         不可空         关联业务订单号,商户业务系统的订单号。                                310000001212450
  ---------------------------------------------------------------------------------------------------------------------------------------------------------

### 业务返回参数

  **参数**    **类型（长度）**   **是否必填**   **描述**   **样例**
  ----------- ------------------ -------------- ---------- -----------------------------------------------------------------------
  file\_url   String(500)        不可空         文件路径   http://baseintra.vfinance.cn/ufs/download/public/20170704/test001.rar

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.user.info.img.upload");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("file", "UmFyIRoHAM. . .");

params.put("file\_name", "test001.rar");

params.put("purpose", "FILE\_PURPOSE\_CETIFICATION");

params.put("rela\_voucher\_no", "232323232323");

### 响应示例

正常：{"charset":"UTF-8","file\_url":"http://baseintra.vfinance.cn/ufs/download/public/20170704/test001.rar","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"FIELD\_TYPE\_ERROR\\":\[\\"文件名称格式\\"\]}"}

1.  \[20101\]mybank.tc.trade.pay.instant (即时交易入账)
    ---------------------------------------------------

    1.  ### 接口说明

        平台可调用该接口实现订单交易，合作方业务平台的买家在业务平台选择商品下单，付款成功后款项直接结算给卖家账户，业务平台交易成功后，将成功的交易同步给交易见证平台。

<!-- -->

1.  即时交易入账是交易见证平台用于同步客户业务平台即时到账类交易的入账接口，发起即时到账交易请求后，会同步返回接口请求结果。即时到账交易的处理结果以后台通知或者查询接口结果为准。交易见证系统会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“交易处理结果通知”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证系统会重试若干次，以避免掉单现象。

2.  平台专属入款渠道编码：平台在对接申请表里需填写入款渠道。申请表提交后，银行小二会分配平台专属入款渠道编码，编码规则：入款渠道编码+5位序号。如：入款渠道为支付宝ALIPAY，网商分配反馈的编码则可以是ALIPAY00012，具体编码以小二反馈信息为准。平台在调用即时入账接口的时候，
    入款渠道需填写专属入款渠道代码。

3.  手续费：即时入账可以收取手续费，
    手续费可以只收买家也可以只收卖家，也可以既收买家也收卖家。

    a.  若收取买家手续费,则交易金额=订单金额，买家实际支出金额=交易金额+收取买家手续费金额.
        假设订单金额10元,收取买家1元手续费,则交易金额需填写10元，买家手续费填写1元，交易成功后,平台手续费账号收入1元,卖家收入10，买家实际支出11元。

    b.  若收取卖家手续费，则订单金额=交易金额，卖家收入金额=交易金额-手续费金额。设订单金额10元,
        收取卖家手续费1元,则交易成功后,卖家收入9元,平台手续费账号收入1元，买家实际支出10元。

    c.  若同时收取买家和卖家各1元手续费. 则综合上述计算规则.
        假设订单金额10元,收取买家和卖家各1元手续费. 则交易金额为10元.
        卖家实际收入9元,平台手续费账号收入2元，买家实际支出11元。

4.  分润:即时入账支持分润控制.即卖家收入金额需根据规则进行分润.假设卖家A预计收入10元,
    接口上送分润规则,需分给B和C各1元。 则交易成功后，A实际收入8元，
    B和C各收入1元。

5.  折扣:系统支持买家以一定折扣支付金额。折扣方式暂只支持补贴，即由平台补贴折扣金额。假设交易金额10元，折扣为2元。则卖家实际收入10元，买家只需支出8元，其中2元由平台营销费用账户支出。

6.  订单相关信息：即时入账交易，平台需上送交易订单信息，即明确该笔交易的交易背景和用途。

    1.  ### 业务请求参数

  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数 **                **类型（长度） **   **是否必填**   **描述**                                                                                                                                                                                                                                         **样例 **
  ------------------------ ------------------- -------------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  outer\_trade\_no         String(32)          不可空         合作方业务平台订单号                                                                                                                                                                                                                             2017032216593367523352596551
                                                                                                                                                                                                                                                                                                               
                                                              合作方提交的交易号，对于合作方全局唯一                                                                                                                                                                                                           

  outer\_inst\_order\_no   String(32)          可空           合作方平台自接入款渠道发起的支付，上送支付渠道的支付订单号。若为pay\_method余额支付，则该栏位可空。该订单号编码规则银行不做要求，以合作方平台与入款渠道约定规则为准。                                                                            20101018001052601340050

  white\_channel\_code     String(32)          可空           平台专属入款渠道编码，pay\_method非余额支付方式时必填。该栏位的可选列表由网商银行小二根据平台递交的申请表分配并反馈。编码规则：入款渠道编码+5位序号。如：入款渠道为ALIPAY，网商分配反馈的编码则可以是ALIPAY00012，具体编码以小二反馈信息为准。   ALIPAY00012

  buyer\_id                String(32)          不可空         买家在业务平台的ID（UID）。匿名支付场景下，此栏位填写anonymous                                                                                                                                                                                   U500000001

  buyer\_ip                String(32)          可空           用户在业务平台下单的时候的ip地址，公网IP，不是内网IP                                                                                                                                                                                             202.114.12.45
                                                                                                                                                                                                                                                                                                               
                                                              用于风控校验                                                                                                                                                                                                                                     

  pay\_method              String(1000)        不可空         支付方式，格式为Json，具体说明见下方接口参数补充说明。                                                                                                                                                                                           {"pay\_method":"QPAY","amount":"0.3","memo":"ALIPAY,C,DC, N6228480210599399511","extension":""}

                           String(1000)        可空           折扣支付方式，无折扣则该栏位无需填写。 格式为JsonList，具体说明见下方接口参数补充说明。                                                                                                                                                          {discount\_type: "allowancePay",instCode: "DISCOUNTPAY",amountInfo: "25.00"}\]

  fee\_info                String(100)         可空           手续费信息，无手续费则该栏位无需填写。说明：                                                                                                                                                                                                     {"sellerFee":"0.5","buyerFee":"0.3"}
                                                                                                                                                                                                                                                                                                               
                                                              如果只收买家手续费则卖家不填，如：                                                                                                                                                                                                               
                                                                                                                                                                                                                                                                                                               
                                                              {"buyerFee":"0.3"}                                                                                                                                                                                                                               
                                                                                                                                                                                                                                                                                                               
                                                              sellerFee:卖家手续费                                                                                                                                                                                                                             
                                                                                                                                                                                                                                                                                                               
                                                              buyerFee：买家手续费                                                                                                                                                                                                                             

  subject                  String(256)         不可空         商品的标题/交易标题/订单标题/订单关键字等。                                                                                                                                                                                                      油漆
                                                                                                                                                                                                                                                                                                               
                                                              该参数最长为128个汉字。                                                                                                                                                                                                                          

  price                    Number(10)          不可空         商品单价。单位为：RMB Yuan。取值范围为\[0.01，1000000.00\]，精确到小数点后两位。                                                                                                                                                                 115.00

  quantity                 Number(5)           不可空         商品的数量。                                                                                                                                                                                                                                     2

  total\_amount            Number(12)          不可空         交易金额=（商品单价×商品数量）。卖家实际扣款和卖家实际收款金额计算规则请参考接口参数补充说明。                                                                                                                                                   230.00

  royalty\_parameters      String(300)         可空           交易金额分润账号集，接口参数补充说明                                                                                                                                                                                                             \[{"uid":"10012547853","account\_type":"BASIC","account\_identity":"","amount":"10.00","royalty\_rule":""},{"uid":"10012685423","account\_type":"BASIC","account\_identity ":"","amount":"10.00","royalty\_rule":""}\]

  seller\_id               String(32)          不可空         卖家在业务平台的用户ID（UID）                                                                                                                                                                                                                    10004563451

  account\_type            String(20)          不可空         卖家账户类型                                                                                                                                                                                                                                     BASIC

  body                     String(200)         可空           商品描述。对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加。                                                                                                                                                                     白色墙面漆 5KG

  show\_url                String(200)         可空           商品展示URL。收银台页面上，商品展示的超链接。                                                                                                                                                                                                    http://www.custom.com/?product-9.html

  notify\_url              String(200)         不可空         服务器异步通知页面路径。该接口同步只受理请求，交易见证平台订单成功后会主动通知商户业务平台里指定的页面http路径。对应异步通知的“**交易状态变更通知**”。                                                                                           http://mas.test.custom.net/atinterface/receive\_notify.htm

  shop\_name               String(128)         可空           具体交易发生的店铺名称                                                                                                                                                                                                                           xxx官方旗舰店

  product\_code            String(5)           可空           预留字段，可不填。                                                                                                                                                                                                                               

  operator\_id             String(32)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  return\_url              String(1000)        可空           预留字段，可不填。                                                                                                                                                                                                                               

  go\_cashier              String(1)           可空           预留字段，可不填                                                                                                                                                                                                                                 

  is\_web\_access          String(1)           可空           预留字段，可不填。                                                                                                                                                                                                                               

  account\_identity        String(32)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  expiration\_time         String(10)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  submit\_time             String(14)          可空           预留字段，可不填。                                                                                                                                                                                                                               
  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

【接口参数补充说明】

1.  **pay\_method格式为json：**pay\_method(支付方式)，amount(金额)，memo(备注)，extension(扩展信息，保留使用，默认为空）。

> ****pay\_method(支付方式)****可以为如下选择：

-   **BALANCE**-余额支付

-   **POS**-POS支付

-   **QPAY**-快捷

-   **ONLINE\_BANK**-网银

> ****memo(备注**)**格式为以逗号分隔的字符串：银行代码,对公/对私：B/C,借记卡/贷记卡：DC=借记卡支付,CC=贷记卡支付)

-   当支付方式为网银、POS方式时，memo必填。样例：ALIPAY,C,DC

-   当支付方式为快捷时memo必填, 需要在后面增加B+bankid或N+
    bank\_account\_no

> 样例：ALIPAY,C,DC,N6228480210599399511或ALIPAY,C,DC,B10599399511

-   当支付方式为余额时，memo为空默认基本户支付，若使用其它账户支付则填写为对应的账户类型，例如：营销费用户(PFEE)支付，memo设置为PFEE

> **pay\_method样例**

-   以余额支付为例：

> {"pay\_method":"BALANCE","amount":"0.3","memo":"","extension":""}

-   以工商银行-POS为例：

> {"pay\_method":"POS","amount":"0.3","memo":"ICBC,C,DC","extension":""}

-   以支付宝-网银网关为例：

> {"pay\_method":"ONLINE\_BANK","amount":"0.3","memo":"ALIPAY,C,DC","extension":""}

-   以支付宝-快捷为例：

> {"pay\_method":"QPAY","amount":"0.3","memo":"ALIPAY,C,DC,
> N6228480210599399511","extension":""}
>
> 或
>
> {"pay\_method":"QPAY","amount":"0.3","memo":"ALIPAY,C,DC,
> B10599399511","extension":""}

1.  **discount\_pay\_method字段**为jsonList格式：discount\_type(折扣类型),
    instCode(折扣来源机构编码),amountInfo(折扣金额)。

> ****discount\_type(折扣类型)****：暂只支持补贴方式，固定填写allowancePay。注意：不能填写ALLOWANCEPAY，大小写须与allowancePay一致。补贴方式默认是平台进行补贴，实际资金在交易时扣减**平台营销费用子账户**。
>
> ****instCode(折扣来源渠道编码)****：固定填写DISCOUNTPAY。若填错，则报错：“不支持的折扣支付方式”。
>
> ****amountInfo(折扣金额)****：RMB
> Yuan。取值范围为\[0.01，100000000.00\]，精确到小数点后两位。

**discount\_pay\_method样例：**

> 平台补贴25元：
>
> \[{discount\_type: "allowancePay",instCode: "DISCOUNTPAY",amountInfo:
> "25.00"}\]

1.  **royalty\_parameters字段**为jsonList格式：uid（外部用户号)(不可空)，account\_type(账号类型)(不可空)，account\_identity(账户标识)(保留使用，固定填空)，amount(金额)(不可空)，royalty\_rule(分润规则)(
    保留使用，固定填空)

样例：

\[{"uid":"10012547853","account\_type":"BASIC","account\_identity":"","amount":"10.00","royalty\_rule":""},{"uid":"10012685423","account\_type":"BASIC","account\_identity
":"","amount":"10.00","royalty\_rule":""}\]

支付时，则默认买家是分润资金付款方；退款时则卖家是分润资金付款方；

同一笔交易的分润处理，必须遵循原则：分润资金付款方要先收到钱才能再付给其他被分润者，收到的钱一定要大于等于付出的钱，即先入后出，入款要大于等于出款；

1.  接口金额计算规则：

> 如果一笔交易100元，则total\_amount(交易金额)为100元，如果对买家按1%收取收费则手续费信息中手续费金额为1元，如果有折扣20元则折扣支付金额为20元，最终买家付款81元，卖家收到100元。若同时收取卖家手续费金额1元，则卖家收到99元。若还存在分润支出10元，则卖家收到89元。

### 业务响应参数

[*只有公共响应参数 *](#公共响应参数)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", " mybank.tc.trade.pay.instant");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("account\_type", "PBASIC");

params.put("body", "爱丽舍-三厢 1.6 MT");

params.put("buyer\_id", "p80aw8b06r");

params.put("buyer\_ip", "202.114.12.45");

params.put("buyer\_mobile", "202.114.12.45");

params.put("discount\_pay\_method",
"\[{\\"discount\_type\\":\\"bonusPay\\",\\"instCode\\":\\"HONGBAOPAY\\",\\"amountInfo\\":\\"1\^0.20|2\^0.30\\"}\]");

params.put("expiration\_time", "7d");

params.put("fee\_info",
"{\\"sellerFee\\":\\"0.5\\",\\"buyerFee\\":\\"0.5\\"}");

params.put("outer\_inst\_order\_no", "242332434345278");

params.put("outer\_trade\_no", "2017070416434026864850763232");

params.put("pay\_method",
"{\\"pay\_method\\":\\"online\_bank\\",\\"amount\\":\\"0.5\\",\\"memo\\":\\"ALIPAY,C,DC\\",\\"extension\\":\\"\\"}");

params.put("price", "0.50");

params.put("quantity", "2");

params.put("return\_url",
"http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("seller\_id", "1234656@qq.com");

params.put("shop\_name", "上海大众4S店");

params.put("show\_url", "http://www.test.com/?product-9.html");

params.put("subject", "东风雪铁龙");

params.put("submit\_time", "20140526090530");

params.put("total\_amount", "0.50");

params.put("white\_channel\_code", "ALIPAY10916");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"处理成功"}

异常：{"charset":"UTF-8","error\_code":"TRADE\_PAY\_MATCH\_ERROR","error\_message":"交易与支付金额不匹配","is\_success":"F"}

1.  \[20102\]mybank.tc.trade.pay.ensure (担保交易入账)
    --------------------------------------------------

    1.  ### 接口说明

<!-- -->

1.  担保交易入账是交易见证平台用于同步合作方业务平台担保类交易的入账接口，发起担保交易请求后，会同步返回接口请求结果。担保交易的处理结果以后台通知或者查询接口结果为准。交易见证系统会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“交易处理结果通知”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证系统会重试若干次，以避免掉单现象。

    a.  担保交易资金流示意：

> ![](media/image22.png){width="6.141666666666667in"
> height="1.4416666666666667in"}

a.  担保交易入账：交易成功后，扣减买家资金后，
    资金收入平台担保账号，担保账号资金平台无法通过接口直接转账。

b.  担保结算入账：可通过该交易将订单存入担保账号资金结算给给订单的卖家。暂只支持全额结算，结算若涉及分润，可上送分润参数。

c.  担保退款入账：可通过该交易将订单存入担保账号资金退换给买家。
    系统支持原路退回，原入账时支付方式为余额支付，则退回买家见证账号，若涉及手续费或分润，则原手续费和分润资金也原路退回
    。原入账支付方式为非余额支付（网银支付、快捷支付、POS支付），见证将资金退回到入款渠道对应的银存账号，
    若非余额字符，则实际资金退回需平台自行对接的入款渠道处理。

<!-- -->

1.  分润规则：参考即时入账场景。

2.  手续费规则：参考即时入账场景。

3.  折扣规则：参考即时入账场景。

    1.  ### 业务请求参数

  ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数 **                                                             **类型（长度） **   **是否必填**   **描述**                                                                                                                                                                                                                                         **样例 **
  --------------------------------------------------------------------- ------------------- -------------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ -------------------------------------------------------------------------------------------------------- --
  outer\_trade\_no                                                      String(32)          不可空         合作方业务平台订单号                                                                                                                                                                                                                             2017032216593367523352596552

  outer\_inst\_order\_no                                                String(32)          可空           合作方平台自接入款渠道发起的支付，上送支付渠道的支付订单号。若为pay\_method余额支付，则该栏位可空。该订单号编码规则银行不做要求，以合作方平台与入款渠道约定规则为准。                                                                            20101018001052601340050

  white\_channel\_code                                                  String(32)          可空           平台专属入款渠道编码，pay\_method非余额支付方式时必填。该栏位的可选列表由网商银行小二根据平台递交的申请表分配并反馈。编码规则：入款渠道编码+5位序号。如：入款渠道为ALIPAY，网商分配反馈的编码则可以是ALIPAY00012，具体编码以小二反馈信息为准。   ALIPAY00012

  buyer\_id                                                             String(32)          不可空         买家在业务平台的ID（UID）。匿名支付场景下，此栏位填写anonymous                                                                                                                                                                                   U500000001

  buyer\_ip                                                             String(32)          可空           用户在业务平台下单的时候的ip地址，公网IP，不是内网IP                                                                                                                                                                                             202.114.12.45
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           用于风控校验                                                                                                                                                                                                                                     

  pay\_method                                                           String(1000)        不可空         支付方式，格式为Json，具体说明见下方接口参数补充说明。                                                                                                                                                                                           {"pay\_method":"ALIPAY00012","amount":"0.3","memo":"ALIPAY,C,DC, N6228480210599399511","extension":""}

  [[]{#OLE_LINK16 .anchor}]{#OLE_LINK15 .anchor}discount\_pay\_method   String(1000)        可空           折扣支付方式，无折扣则该栏位无需填写。 格式为JsonList，具体说明见下方接口参数补充说明。                                                                                                                                                          {discount\_type: "allowancePay",instCode: "DISCOUNTPAY",amountInfo: "25.00"}\]

  fee\_info                                                             String(100)         可空           手续费信息，无手续费则该栏位无需填写。说明：                                                                                                                                                                                                     {"sellerFee":"0.5","buyerFee":"0.3"}
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           如果只收买家手续费则卖家不填，如：                                                                                                                                                                                                               
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           {"buyerFee":"0.3"}                                                                                                                                                                                                                               
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           sellerFee:卖家手续费                                                                                                                                                                                                                             
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           buyerFee：买家手续费                                                                                                                                                                                                                             

  subject                                                               String(256)         不可空         商品的标题/交易标题/订单标题/订单关键字等。                                                                                                                                                                                                      油漆
                                                                                                                                                                                                                                                                                                                                                            
                                                                                                           该参数最长为128个汉字。                                                                                                                                                                                                                          

  price                                                                 Number(10)          不可空         商品单价。单位为：RMB Yuan。取值范围为\[0.01，1000000.00\]，精确到小数点后两位。                                                                                                                                                                 115.00

  quantity                                                              Number(5)           不可空         商品的数量。                                                                                                                                                                                                                                     2

  total\_amount                                                         Number(12)          不可空         交易金额=（商品单价×商品数量）                                                                                                                                                                                                                   230.00

  ensure\_amount                                                        Number(12)          不可空         担保金额，和交易金额相等                                                                                                                                                                                                                         230.00

  seller\_id                                                            String(32)          不可空         卖家在业务平台的用户ID（UID）                                                                                                                                                                                                                    10004563451

  account\_type                                                         String(20)          不可空         卖家账户类型                                                                                                                                                                                                                                     BASIC

  body                                                                  String(200)         可空           商品描述。对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加。                                                                                                                                                                     白色墙面漆 5KG

  show\_url                                                             String(200)         可空           商品展示URL。收银台页面上，商品展示的超链接。                                                                                                                                                                                                    http://www.custom.com/?product-9.html

  notify\_url                                                           String(200)         不可空         服务器异步通知页面路径。交易见证平台订单成功后会主动通知商户业务平台里指定的页面http路径。对应异步通知的“交易状态变更通知”                                                                                                                       http://mas.test.custom.net/atinterface/receive\_notify.htm

  shop\_name                                                            String(128)         可空           具体交易发生的卖方名称                                                                                                                                                                                                                           xxx官方旗舰店

  product\_code                                                         String(5)           可空           预留字段，可不填。                                                                                                                                                                                                                               

  operator\_id                                                          String(32)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  return\_url                                                           String(1000)        可空           预留字段，可不填。                                                                                                                                                                                                                               

  go\_cashier                                                           String(1)           可空           预留字段，可不填。                                                                                                                                                                                                                               

  is\_web\_access                                                       String(1)           可空           预留字段，可不填。                                                                                                                                                                                                                               

  account\_identity                                                     String(32)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  expiration\_time                                                      String(10)          可空           预留字段，可不填。                                                                                                                                                                                                                               

  submit\_time                                                          String(14)          可空           预留字段，可不填。                                                                                                                                                                                                                               
  ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

【参数说明】

> 相关参数格式参考即时入账交易说明。

### 业务响应参数

[*只有公共响应参数 *](#公共响应参数)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.pay.ensure");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("account\_type", "PBASIC");

params.put("body", "爱丽舍-三厢 1.6 MT");

params.put("buyer\_id", "p80aw8b06r");

params.put("buyer\_ip", "202.114.12.45");

params.put("buyer\_mobile", "202.114.12.45");

params.put("discount\_pay\_method",
"\[{\\"discount\_type\\":\\"bonusPay\\",\\"instCode\\":\\"HONGBAOPAY\\",\\"amountInfo\\":\\"1\^0.20|2\^0.30\\"}\]");

params.put("expiration\_time", "7d");

params.put("fee\_info", " {"sellerFee":"0.5","buyerFee":"0.5"}");

params.put("outer\_inst\_order\_no", "242332434345278");

params.put("outer\_trade\_no", "2017070416434026864850763232");

params.put("pay\_method",
"{\\"pay\_method\\":\\"online\_bank\\",\\"amount\\":\\"0.5\\",\\"memo\\":\\"ALIPAY,C,DC\\",\\"extension\\":\\"\\"}");

params.put("price", "0.50");

params.put("quantity", "2");

params.put("return\_url",
"http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("seller\_id", "1234656@qq.com");

params.put("shop\_name", "上海大众4S店");

params.put("show\_url", "http://www.test.com/?product-9.html");

params.put("subject", "东风雪铁龙");

params.put("submit\_time", "20140526090530");

params.put("total\_amount", "1");

params.put("ensure\_amount", "1");

params.put("white\_channel\_code", "ALIPAY10916");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"处理成功"}

异常：{"charset":"UTF-8","error\_code":"DUPLICATE\_REQUEST\_NO","error\_message":"合作方业务平台订单号重复","is\_success":"F"}

1.  \[20103\]mybank.tc.trade.settle(担保交易结算入账)
    -------------------------------------------------

    1.  ### 接口说明

        本接口与担保交易入账接口联合使用。对应买家确认收货场景，客户确认收货后，款项从平台担保户结算给卖家。

<!-- -->

1.  结算入账需关联原担保入账业务订单号。

2.  暂只支持全额结算

3.  手续费：参考即时入账说明

4.  分润：参考即时入账说明

    1.  ### 业务请求参数

  -------------------------------------------------------------------------------------------------------------------------------------------
  **参数**              **类型（长度）**   **是否必填**   **描述**                                                     **样例**
  --------------------- ------------------ -------------- ------------------------------------------------------------ ----------------------
  outer\_trade\_no      String(32)         不可空         合作方业务平台订单号（原担保入账交易合作方业务平台订单号）   2013112405052323

  royalty\_parameters   String(300)        可空           交易金额分润账号集。                                         格式见参数说明

  fee\_info             String(100)        可空           手续费信息，json格式                                         {"sellerFee":"0.5"}
                                                                                                                       
                                                                                                                       说明：
                                                                                                                       
                                                                                                                       sellerFee:卖家手续费

  operator\_id          String(32)         可空           预留字段，可不填                                             

  is\_web\_access       String(1)          可空           预留字段，可不填                                             
  -------------------------------------------------------------------------------------------------------------------------------------------

【参数说明】

royalty\_parameters参数参考即时入账交易说明：

fee\_info参数参考即时入账交易说明：

### 业务响应参数

[*只有公共响应参数 *](#公共响应参数)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.settle");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("fee\_info", "
{\\"sellerFee\\":\\"0.5\\",\\"buyerFee\\":\\"0.5\\"}");

params.put("outer\_trade\_no", "2017070416434026864850763232");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"结算成功"}

异常：{"charset":"UTF-8","error\_code":"DUPLICATE\_REQUEST\_NO","error\_message":"合作方业务平台订单号重复","is\_success":"F"}

1.  \[20105\]mybank.tc.trade.deposit(充值记账)
    ------------------------------------------

    1.  ### 功能说明

        适用于合作方业务平台接入第三方支付渠道（非网商入金渠道），用户充值，当充值交易在合作方业务平台完成后，调用该接口通知交易见证平台同步充值交易。

<!-- -->

1.  发起充值记账请求后，会同步返回接口请求结果。的理结果以后台通知或者查询接口结果为准。交易见证系统会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“交易处理结果通知”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证系统会重试若干次，以避免掉单现象。

2.  交易成功后，实时增加会员可用余额。

3.  网商渠道入金不需通过该接口同步。

    1.  ### 使用场景

    2.  ### 业务请求参数 

  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  参数                     类型（长度）    是否必填   描述                                                                                                                                                                                                            样例
  ------------------------ --------------- ---------- --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- --------------------------------------------------------------------------------------------------------
  outer\_trade\_no         String(32)      不可空     合作方业务平台订单号                                                                                                                                                                                            10000043551252390

  uid                      String(32)      不可空     合作方业务平台用户ID（UID）                                                                                                                                                                                     U500000001

  outer\_inst\_order\_no   String(32)      不可空     合作方平台自接入款渠道发起的支付，上送支付渠道的支付订单号。该订单号编码规则银行不做要求，以合作方平台与入款渠道约定规则为准。                                                                                  20101018001052601340050

  white\_channel\_code     String(32)      不可空     平台专属入款渠道编码，该栏位的可选列表由网商银行小二根据平台递交的申请表分配并反馈。编码规则：入款渠道编码+5位序号。如：入款渠道为ALIPAY，网商分配反馈的编码则可以是ALIPAY00012，具体编码以小二反馈信息为准。   ALIPAY00012

  account\_type            String(20)      不可空     充值账户类型                                                                                                                                                                                                    BASIC

  amount                   Number(12)      不可空     充值金额                                                                                                                                                                                                        100.00

  notify\_url              String（256）   不可空     异步回调地址                                                                                                                                                                                                    http://mas.test.custom.net/atinterface/receive\_notify.htm

  ip                       String(32)      可空       用户在业务平台下单的时候的ip地址，公网IP，不是内网IP                                                                                                                                                            202.114.12.45
                                                                                                                                                                                                                                                                      
                                                      用于风控校验                                                                                                                                                                                                    

  pay\_method              String(1000)    不可空     支付方式，不能为余额支付方式。格式参考即时入账pay\_method                                                                                                                                                       {"pay\_method":"ALIPAY00012","amount":"0.3","memo":"ALIPAY,C,DC, N6228480210599399511","extension":""}

  fee\_info                String(100)     可空       手续费信息，json格式。buyerFee即收取会员手续费。会员实际收入金额=充值金额，平台手续费账户收入手续费金额，入款渠道需支付金额=充值金额+收取会员手续费金额。                                                       {"buyerFee":"1.00"}

  product\_code            String(5)       可空       预留字段，可不填                                                                                                                                                                                                

  operator\_id             String(32)      可空       预留字段，可不填                                                                                                                                                                                                

  account\_identity        String(32)      可空       预留字段，可不填（账户标识）                                                                                                                                                                                    

  go\_cashier              String(1)       可空       预留字段，可不填                                                                                                                                                                                                

  is\_web\_access          String(1)       可空       预留字段，可不填                                                                                                                                                                                                
  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
*String*&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", " mybank.tc.trade.deposit");

params.put("sign", "19VpdIS4....");

params.put("sign\_type", "RSA");

params.put("uid", "56p7usf46c");

params.put("version", "2.0");

params.put("account\_type", " PBASIC");

params.put("amount", " 0.50");

params.put("fee\_info", "{\\"buyerFee\\":\\"0.5\\"}");

params.put("go\_cashier", " Y");

params.put("ip", " 202.114.12.45");

params.put("is\_web\_access", " Y");

params.put("notify\_url", "
http://func121admin.vfinance.cn/gop-test/mag/asynNotify.htm");

params.put("operator\_id", " 10005454");

params.put("outer\_inst\_order\_no", " 23232323245874545434");

params.put("outer\_trade\_no", " 2017070416153879277603086084");

params.put("pay\_method", "
{\\"pay\_method\\":\\"online\_bank\\",\\"amount\\":\\"0.50\\",\\"memo\\":\\"ALIPAY,C,DC\\",\\"extension\\":\\"\\"}");

params.put("return\_url", "
http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("white\_channel\_code", " ALIPAY10916");

params.put("memo", "{\\"real\_name\\":\\"张三\\"}");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"处理成功"}

异常：{"charset":"UTF-8","error\_code":"PAY\_METHOD\_ERROR","error\_message":"支付方式错误","is\_success":"F","memo":"支付方式不支持"}

\[20108\]mybank.tc.trade.refund (退款入账) 
-------------------------------------------

[[[[[[]{#_Toc481701749 .anchor}]{#_Toc481701748 .anchor}]{#_Toc481701747
.anchor}]{#_Toc481701746 .anchor}]{#_Toc481701745
.anchor}]{#_Toc481701744 .anchor}

1.  ### 功能说明

    该接口支持即时入账退款和担保入账退款。

<!-- -->

1.  发起退款请求后，会同步返回接口请求结果。退款交易的处理结果以后台通知或者查询接口结果为准。交易见证系统会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“交易处理结果通知”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证系统会重试若干次，以避免掉单现象。

2.  见证体系内资金原路退回。
    若原交易支付方式非余额支付，则平台也需保证实际资金通过原入款渠道原路退回，若入款渠道退款机制为轧差退款，则完成见证资金退回处理即可。

3.  退款需关联原即时入行或担保入账业务订单号。

4.  退款支持单笔交易分多次退款，多次退款需要提交原始的合作方业务平台唯一订单号和设置不同的合作方业务平台订单号。一笔退款失败后重新提交，要采用原来的退款单号。总退款金额不能超过用户实际支付金额。

5.  退款支持退还手续费、分润、折扣及交易本金的退还。

    1.  ### 业务请求参数

  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  参数                     类型（长度）   是否必填   描述                                                                                                                                                                                                                                       样例
  ------------------------ -------------- ---------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  outer\_trade\_no         String(32)     不可空     合作方业务平台订单号                                                                                                                                                                                                                       30000043551252390

  orig\_outer\_trade\_no   String(32)     不可空     原始的合作方业务平台唯一订单号。同交易中的一致。                                                                                                                                                                                           2017032216593367523352596552

  outer\_inst\_order\_no   String(32)     可空       合作方平台自接入款渠道发起的支付，上送支付渠道的支付订单号。若为pay\_method余额支付，则该栏位可空。该订单号编码规则银行不做要求，以合作方平台与入款渠道约定规则为准。                                                                      20101018001052601340050

  refund\_amount           Number(12)     不可空     退款金额。支持部分退款，退款金额不能大于交易金额。                                                                                                                                                                                         50.00
                                                                                                                                                                                                                                                                                                
                                                     单位为：RMB Yuan，精确到小数点后两位。                                                                                                                                                                                                     
                                                                                                                                                                                                                                                                                                
                                                     缺省值为0元。                                                                                                                                                                                                                              

  royalty\_parameters      String(1000)   可空       交易金额退分润账号集。                                                                                                                                                                                                                     \[{"uid":"10012547853","account\_type":"BASIC","account\_identity":"","amount":"10.00","royalty\_rule":""},{"uid":"10012685423","account\_type":"BASIC","account\_identity ":"","amount":"10.00","royalty\_rule":""}\]
                                                                                                                                                                                                                                                                                                
                                                     账号必须在原账号集中。分润账号退款金额不可大于原分润金额。                                                                                                                                                                                 
                                                                                                                                                                                                                                                                                                
                                                     参数规则说明请参考即时入账参数说明。                                                                                                                                                                                                       

  notify\_url              String(1000)   不可空     异步回调平台的url                                                                                                                                                                                                                          http://www.xxx.com/callback.htm
                                                                                                                                                                                                                                                                                                
                                                     对应异步通知的“交易处理结果通知”                                                                                                                                                                                                           

  refund\_fee\_info        String(100)    可空       手续费信息，json格式                                                                                                                                                                                                                       {"sellerRefundFee":"0.5","buyerRefundFee":"0.3"}
                                                                                                                                                                                                                                                                                                
                                                                                                                                                                                                                                                                                                说明：
                                                                                                                                                                                                                                                                                                
                                                                                                                                                                                                                                                                                                sellerRefundFee:退卖家手续费
                                                                                                                                                                                                                                                                                                
                                                                                                                                                                                                                                                                                                buyerRefundFee：退买家手续费

  refund\_info             String()       可空       没有折扣支付退款时，退款信息refund\_info可空，有折扣支付退款时，退款信息必填，但如果折扣不退款填写方式{"buyerRefund":"1.5"}，如果交易不退款{" discountRefund ":"0.5"}如果交易和折扣都退款{"buyerRefund":"1.5", " discountRefund ":"0.5"}   {"buyerRefund":"1.5", " discountRefund ":"0.5"}
                                                                                                                                                                                                                                                                                                
                                                                                                                                                                                                                                                                                                buyerRefund：买家支付退款金额
                                                                                                                                                                                                                                                                                                
                                                                                                                                                                                                                                                                                                discountRefund：折扣支付退款金额

  refund\_ensure\_amount   Number(12)     可空       预留字段，可不填                                                                                                                                                                                                                           

  operator\_id             String(32)     可空       预留字段，可不填                                                                                                                                                                                                                           

  is\_web\_access          String(1)      可空       预留字段，可不填                                                                                                                                                                                                                           
  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### 业务响应参数

[*只有公共响应参数 *](#公共响应参数)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.refund");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("outer\_trade\_no", " 2017070416473019902397952974");

params.put("return\_url","http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("orig\_outer\_trade\_no", "2017070417004018132311422378");

params.put("is\_web\_access", "Y");

params.put("notify\_url", "
http://func121admin.vfinance.cn/gop-test/mag/asynNotify.htm");

params.put("orig\_outer\_trade\_no", " 2017070417263432363163866103");

params.put("refund\_amount", "2.0");

params.put("refund\_fee\_info",
"{\\"sellerRefundFee\\":\\"0.5\\",\\"buyerRefundFee\\":\\"0.5\\"}");

params.put("refund\_info",
"{\\"buyerRefund\\":\\"1.5\\",\\"discountRefund\\":\\"0.5\\"}");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"DUPLICATE\_REQUEST\_NO","error\_message":"合作方业务平台订单号重复","is\_success":"F"}

1.  \[20106\]mybank.tc.trade.paytocard(单笔提现) 
    ---------------------------------------------

    1.  ### 接口说明

        平台可调用20106接口发起提现。将会员或合作方自有账户资金余额提现到指定银行卡。

<!-- -->

1.  发起单笔提现交易请求后，会同步返回接口请求结果。单笔提现交易的处理结果以后台通知或者查询接口结果为准。交易见证平台会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“提现状态变更通知（单笔）”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证平台会重试若干次，以避免出现掉单现象。

2.  当提现成功，但由于银行方面处理失败时，需发起退票操作。退票只支持全额退票，收取的手续费也会全额退回。平台可调用20109接口将退票信息提交见证系统。

3.  会员未绑定银行账号，则不允许发起提现。

4.  会员提现不保证实时到账，
    网商银行准实时发出，具体到账时间以收款银行实际入账处理为准。

5.  对公账户大于5万提现走人民银行大额支付通道。只支持在工作日09:00-17:00之间发起（有可能根据人行通知调整）。

6.  若平台绑定提现账户时，有做三四要素鉴权验证，则提现成功率会很高，反之则无法避免“卡号不存在”，“账户户名不符”等非网商银行原因导致的提现失败。

7.  提现为异步处理。提现结果以银行异步通知为准。

8.  若平台超过一定时间还未收到异步通知，
    则可以调用20202接口查询银行端交易状态。若查询返回trade\_status为提现成功或提现失败，平台可以此为准做后续处理，若返回提现处理中，则平台需继续查询，直到返回提现成功或提现失败。

9.  若异步通知结果为明确成功或失败，平台可根据通知结果做后续处理，
    若为系统异常，则不可重试，需联系网商确认结果后，再进行对应处理。

    1.  ### 业务请求参数

+-------------+-------------+-------------+-------------+-------------+
| **参数**    | **类型（长度）** | **是否必填** | **描述** | **样例** |
+=============+=============+=============+=============+=============+
| outer\_trad | String(32)  | 不可空      | 合作方业务平台订单号* | 20101018001 |
| e\_no       |             |             | *\          | 05260134005 |
|             |             |             | **合作方提交的交易号 | 2  |
|             |             |             | ，对于合作方全局唯一 |    |
+-------------+-------------+-------------+-------------+-------------+
| uid         | String(32)  | 不可空      | 合作方业务平台用户ID | U500000001 |
+-------------+-------------+-------------+-------------+-------------+
| outer\_inst | String(32)  | 不可空      | 合作方对接出款渠道使用 | 20101018001 |
| \_order\_no |             |             | 的提现订单号。若出款渠 | 05260134005 |
|             |             |             | 道是网商银行，则此处填 | 2 |
|             |             |             | 写与outer\_tr |           |
|             |             |             | ade\_no保持一致 |         |
|             |             |             | 。          |             |
+-------------+-------------+-------------+-------------+-------------+
| white\_chan | String(32)  | 不可空      | 平台专属出款渠道编码， | MYBANK00012 |
| nel\_code   |             |             | 该栏位的可选列表由网商 |  |
|             |             |             | 银行小二根据平台递交的 |  |
|             |             |             | 申请表分配并反馈。编码 |  |
|             |             |             | 规则：出款渠道编码+5 |    |
|             |             |             | 位序号。如：出款渠道为 |  |
|             |             |             | 网商，网商分配反馈的编 |  |
|             |             |             | 码则可以是MYBANK |        |
|             |             |             | 00012，具体编码以 |       |
|             |             |             | 小二反馈信息为准。 |      |
+-------------+-------------+-------------+-------------+-------------+
| account\_ty | String(20)  | 不可空      | 账户类型,会员提现，暂 | BASIC |
| pe          |             |             | 只支持BASIC |             |
+-------------+-------------+-------------+-------------+-------------+
| bank\_id    | String(20)  | 不可空      | 会员在交易见证平台绑定 |  |
|             |             |             | 的银行卡id  |             |
|             |             |             | ，即10401接口返回 |       |
|             |             |             | 的bank\_id。 |            |
+-------------+-------------+-------------+-------------+-------------+
| amount      | Number(12)  | 不可空      | 提现金额。金额必须不大 | 15000.00 |
|             |             |             | 于账户可用余额 |          |
+-------------+-------------+-------------+-------------+-------------+
| notify\_url | String(200) | 不可空      | 服务器异步通知页面路径 | http://biz/ |
|             |             |             | 。          | atinterface |
|             |             |             |             | /receive\_n |
|             |             |             | 对应异步通知的“提现状 | otify.htm |
|             |             |             | 态变更通知（单笔）” |     |
+-------------+-------------+-------------+-------------+-------------+
| fee\_info   | String(100) | 可空        | 手续费信息，json格 | {"buyerFee" |
|             |             |             | 式。buyerFee即 | :"0.3"}  |
|             |             |             |             |             |
|             |             |             |             |             |
|             |             |             | 收取会员提现手续费。 |    |
+-------------+-------------+-------------+-------------+-------------+
| is\_web\_ac | String(1)   | 可空        | 预留字段，可不填 |        |
| cess        |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| account\_id | String(32)  | 可空        | 预留字段，可不填（账户 |  |
| entity      |             |             | 标识）      |             |
+-------------+-------------+-------------+-------------+-------------+
| product\_co | String(5)   | 可空        | 预留字段，可不填 |        |
| de          |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| bank\_accou | String(32)  | 可空        | 预留字段，可不填（银行 |  |
| nt\_no      |             |             | 卡号）      |             |
+-------------+-------------+-------------+-------------+-------------+
| pay\_attrib | String(10)  | 可空        | 预留字段，可不填（卡支 | NORMAL普通卡 |
| ute         |             |             | 付属性）    |             |
|             |             |             |             |             |
|             |             |             |             | QPAY快捷    |
+-------------+-------------+-------------+-------------+-------------+

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.paytocard");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("account\_type", " BASIC");

params.put("amount", " 0.01");

params.put("return\_url",
"http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("bank\_account\_no", "6214830215878947");

params.put("bank\_id", "2651");

params.put("is\_web\_access", "Y");

params.put("notify\_url","http%3A%2F%2Ffunc121admin.vfinance.cn%2Fgop-test%2Fmag%2FasynNotify.htm");

params.put("outer\_inst\_order\_no", "adadsdasd312311");

params.put("outer\_trade\_no", "2017070418180427553316851673");

params.put("pay\_attribute", "NORMAL");

params.put("product\_code", "20201");

params.put("white\_channel\_code", "WXPAY00001");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"CARD\_INFO\_MATCH\_ERROR","error\_message":"银行卡信息有误","is\_success":"F"}

1.   \[20109\]mybank.tc.trade.refundticket (退票) 
    ----------------------------------------------

    1.  ### 接口说明

        若提现因某些原因，导致收款行将资金退回，平台通过企业网银等渠道了解到退票信息后，可调用该接口将原提现交易状态修改为退票，同时资金退还原提现支出账户。

<!-- -->

1.  需关联原合作方业务平台订单号

2.  只支持全额退回

3.  提现收取的手续费也会全额退回。

    1.  ### 业务请求参数

  ---------------------------------------------------------------------------------------------------------------------------------------
  参数                           类型（长度）   是否必填   描述                                                 样例
  ------------------------------ -------------- ---------- ---------------------------------------------------- -------------------------
  request\_no                    String(32)     不可空     业务平台请求号                                       20101018001052601340053

  orig\_outer\_inst\_order\_no   String(32)     不可空     1.  原提现交易合作方对接出款渠道使用的提现订单号。   20101018001052601340052
                                                                                                                
  ---------------------------------------------------------------------------------------------------------------------------------------

### 业务响应参数

[*只有公共响应参数 *](#公共响应参数)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.refundticket");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("orig\_outer\_inst\_order\_no", "adadsdasd3123");

params.put("return\_url", "
http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("request\_no", "2017070418033851902195792718");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"E0001","error\_message":"原提现机构订单已退票","is\_success":"F"}

1.  \[20104\]mybank.tc.trade.transfer(转账入账)
    -------------------------------------------

    1.  ### 接口说明

        平台可调用该接口实现一个会员不同账户之间的转账。

<!-- -->

1.  发起转账入账交易请求后，会同步返回接口请求结果。处理结果以后台通知或者查询接口结果为准。交易见证系统会主动通知合作方业务平台的异步回调地址，通过异步通知接口中“交易处理结果通知”返回交易处理结果和交易完成时间，需要合作方业务平台保存信息，用于与交易见证平台进行交易对账。如果异步通知失败，交易见证系统会重试若干次，以避免掉单现象。

2.  不支持不同会员间账户转账。

    1.  ### 业务请求参数

+-------------+-------------+-------------+-------------+-------------+
| **参数**    | **类型（长度）** | **是否必填** | **描述** | **样例** |
+=============+=============+=============+=============+=============+
| outer\_trad | String(32)  | 不可空      | 合作方业务平台订单号* | 30000043551 |
| e\_no       |             |             | *\          | 252390      |
|             |             |             | **合作方提交的交易号 |    |
|             |             |             | ，对于合作方全局唯一 |    |
+-------------+-------------+-------------+-------------+-------------+
| fundin\_uid | String(32)  | 不可空      | 入款用户Id,金额增加 | U50000001 |
|             |             |             | 方的用户ID（UID） |       |
|             |             |             | 或会员ID（内部会员I |     |
|             |             |             | D）         |             |
+-------------+-------------+-------------+-------------+-------------+
| fundin\_acc | String(20)  | 不可空      | 入款账户类型。金额增加 | BASIC |
| ount\_type  |             |             | 的账户的账户类型 |        |
+-------------+-------------+-------------+-------------+-------------+
| fundout\_ui | String(32)  | 不可空      | 用户Id,金额减少方的 | U50000001 |
| d           |             |             | 用户ID（UID），或 |       |
|             |             |             | 会员ID（内部会员ID |      |
|             |             |             | ）          |             |
+-------------+-------------+-------------+-------------+-------------+
| fundout\_ac | String(20)  | 不可空      | 账户类型,金额减少的账 | BASIC |
| count\_type |             |             | 户的账户类型 |            |
+-------------+-------------+-------------+-------------+-------------+
| transfer\_a | Number(10)  | 不可空      | 转账金额。必须大于0 | 210.00 |
| mount       |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| notify\_url | String(200) | 可空        | 服务器异步通知页面路径 | http://biz/ |
|             |             |             | 。支付平台服务器主动通 | atinterface |
|             |             |             | 知业务平台里指定的页面 | /receive\_n |
|             |             |             | http路径。  | otify.htm   |
|             |             |             |             |             |
|             |             |             | 对应异步通知的“交易处 |   |
|             |             |             | 理结果通知” |             |
+-------------+-------------+-------------+-------------+-------------+
| is\_web\_ac | String(1)   | 可空        | 预留字段，可不填 |        |
| cess        |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| transfer\_t | String(20)  | 可空        | 预留字段，可不填 |        |
| ype         |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| fundout\_ac | String(32)  | 可空        | 预留字段，可不填（金额 |  |
| count\_iden |             |             | 减少的账户的账户标识） |  |
| tity        |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| fundin\_acc | String(32)  | 可空        | 预留字段，可不填（金额 |  |
| ount\_ident |             |             | 增加的账户的账户标识） |  |
| ity         |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+
| product\_co | String(5)   | 可空        | 预留字段，可不填 |        |
| de          |             |             |             |             |
+-------------+-------------+-------------+-------------+-------------+

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.transfer");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("outer\_trade\_no", " 2017070416473019902397952974");

params.put("return\_url","http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("fundin\_account\_type", "BASIC");

params.put("fundin\_uid", "p80aw8b06r");

params.put("fundout\_account\_type", "PBASIC");

params.put("fundout\_uid", "1234656@qq.com");

params.put("is\_web\_access", "Y");

params.put("notify\_url",
"http://func121admin.vfinance.cn/gop-test/mag/asynNotify.htm");

params.put("transfer\_amount", "0.01");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"转账-付款成功,
场次执行结果：处理成功,处理成功"}

异常：{"charset":"UTF-8","error\_code":"SYSTEM\_ERROR","error\_message":"F0201余额不足20010010011100001520002000001","is\_success":"F"}

1.  \[20201\]mybank.tc.trade.query (交易流水查询)
    ---------------------------------------------

    1.  ### 接口说明

        该接口提供交易订单的查询，平台方可以通过该接口主动查询订单状态，完成下一步的业务逻辑。需要调用查询接口的情况：
        当平台方业务平台、网络、服务器等出现异常情况，平台方业务系统最终未接收到异步通知； 调用接口后，返回系统错误或未知交易状态情况。接口支持查询一段时间（时间指的是订单创建时间）内所有订单状态。

    2.  ### 业务请求参数

  **参数**          **类型（长度）**   **是否必填**   **描述**                                                                         **样例**
  ----------------- ------------------ -------------- -------------------------------------------------------------------------------- ----------------
  start\_time       String(14)         不可空         开始时间。14位,格式为：年\[4位\]月\[2位\]日\[2位\]时\[2位\]分\[2位\]秒\[2位\]    20140101020101
  end\_time         String(14)         不可空         结束时间。14位,格式为：年\[4位\]月\[2位\]日\[2位\]时\[2位\]分\[2位\]秒\[2位\],   20140617020101
  is\_web\_access   String(1)          可空           预留字段，可不填                                                                 
  product\_code     String(20)         可空           预留字段，可不填                                                                 

### 业务响应参数

  **参数**                 **类型（长度）**   **是否必填**   **描述**                                                                             **样例**
  ------------------------ ------------------ -------------- ------------------------------------------------------------------------------------ --------------------------
  outer\_trade\_no         String(32)         不可空         合作方业务平台唯一订单号                                                             订单编号：20131105154925
  inner\_trade\_no         String(32)         不可空         交易见证平台交易订单号。                                                             2013110514543
  gmt\_close               String(14)         可空           交易完成的时间.14位,格式为：年\[4位\]月\[2位\]日\[2位\]时\[2位\]分\[2位\]秒\[2位\]   20131117020101
  trade\_process\_status   String(50)         不可空         交易处理状态.                                                                        [*见附录*](#交易状态)

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.query");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("end\_time", "20170704235959");

params.put("return\_url", "
http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("is\_web\_access", "Y");

params.put("product\_code", "231123");

params.put("start\_time", "20170704000000");

### 响应示例

正常：{"charset":"UTF-8","flows":\[{"gmt\_close":"2017-07-04
18:16:00","inner\_trade\_no":"101149916335913620172","outer\_trade\_no":"2017070418155865208950377691","trade\_process\_status":"TRADE\_FAILED"},{"gmt\_close":"2017-07-04
18:16:43","inner\_trade\_no":"101149916340144720173","outer\_trade\_no":"2017070418164023251594523327","trade\_process\_status":"TRADE\_FINISHED"}\],"is\_success":"T"}

异常：{"charset":"UTF-8","error\_code":"BUSINESS\_VOUCHER\_INEXISTENCE","error\_message":"查询无此交易记录","is\_success":"F"}

1.  \[20202\]mybank.tc.trade.info.query (交易详情查询) 
    ---------------------------------------------------

    1.  ### 接口说明

        可以查询单笔交易详情，包括交易状态、交易处理时间等信息。

    2.  ### 业务请求参数 

  **参数 **            **类型（长度） **   **是否必填 **   **描述 **              **样例 **
  -------------------- ------------------- --------------- ---------------------- ---------------------------
  > outer\_trade\_no   String(32)          不可空          合作方业务平台订单号   > 20101018001052601340052

### 业务响应参数 

  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数 **                  **类型（长度） **   **是否必填 **   **描述 **                                                                     **样例 **
  -------------------------- ------------------- --------------- ----------------------------------------------------------------------------- -------------------------
  > outer\_trade\_no         > String(32)        不可空          合作方业务平台订单号                                                          20101018001052601340052

  > inner\_trade\_no         > String(32)        不可空          交易见证系统的交易凭证号                                                      2013110514543

  > outer\_inst\_order\_no   > String(32)        不可空          平台专属出款渠道                                                              2013110514543

  > product\_code            > String(5)         可空            业务产品码 见附录说明                                                         

  > subject                  > String(256)       可空            商品的标题/交易标题/订单标题/订                                               油漆
                                                                                                                                               
                                                                 单关键字等。                                                                  

  > trade\_amount            > Number            不可空          交易金额，精确到小数点后两位                                                  656.00

  > discount\_amount         > Number            可空            折扣金额，精确到小数点后两位                                                  1.00

  > buyer\_fee               > Number            可空            买家手续费，精确到小数点后两位                                                1.00

  > seller\_fee              > Number            可空            卖家手续费，精确到小数点后两位                                                1.00

  > fail\_reason             > String            可空            失败原因                                                                      

  > trade\_time              > String(14)        不可空          交易创建时间.数字串，一共 14 位格式为：年\[4 位\]月 \[2 位\]日\[2 位\]时\[2   20131117020101
                                                                                                                                               
                                                                 位\]分\[2 位\]秒\[2 位\]                                                      

  > trade\_status            > String(50)        不可空          交易状态.                                                                     [*见附录*](#交易状态)

  > memo                     > String(50)        可空            备注                                                                          
  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.info.query");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("outer\_trade\_no", " 2017070416473019902397952974");

params.put("return\_url", "
http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

### 响应示例

正常：{"buyer\_fee":"0.20","charset":"UTF-8","discount\_amount":"0.30","inner\_trade\_no":"101150036961709340774","outer\_inst\_order\_no":"instant\_0718\_01","is\_success":"T","outer\_trade\_no":"2017071817201372868556467173","seller\_fee":"0.10","subject":"东风雪铁龙","trade\_amount":"1.00","trade\_status":"TRADE\_FINISHED","trade\_time":"2017-07-18
17:20:33"}

异常：{"charset":"UTF-8","error\_code":"BUSINESS\_VOUCHER\_INEXISTENCE","error\_message":"查询无此交易记录","is\_success":"F"}

1.   \[20301\]mybank.tc.trade.funds.freeze (冻结资金)
    -------------------------------------------------

    1.  ### 功能说明

        可调用该接口将会员账户下的资金操作为冻结状态。

<!-- -->

1.  可调用10303接口查询余额。冻结余额=账面余额-可用余额。

2.  冻结支持目前只支持对基本户进行冻结。

3.  可调用20302接口解冻。

    1.  ### 业务请求参数

  **参数**            **类型（长度）**   **是否必填**   **描述**                         **样例**
  ------------------- ------------------ -------------- -------------------------------- -------------------
  outer\_trade\_no    String(32)         不可空         合作方业务平台订单号             30000043551252390
  uid                 String(32)         不可空         合作方业务平台用户ID（UID）      U500000001
  account\_type       Sting(20)          不可空         账户类型，暂只支持BASIC-基本户   BASIC
  amount              Number(12)         不可空         冻结金额.金额不大于账户余额      1000.00
  frozen\_desc        String(20)         可空           冻结描述                         业务冻结
  account\_identity   String(32)         可空           预留字段，可不填                 
  fund\_prop\_type    String(2)          可空           预留字段，可不填                 
  is\_web\_access     String(1)          可空           预留字段，可不填                 

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.funds.freeze");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("outer\_trade\_no", " 2017070416473019902397952974");

params.put("return\_url","http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("account\_type", "BASIC");

params.put("amount", "0.03");

params.put("frozen\_desc", "冻结任性");

params.put("fund\_prop\_type", "DR");

params.put("is\_web\_access", "Y");

params.put("uid", "p80aw8b06r");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：
{"charset":"UTF-8","error\_code":"USER\_ACCOUNT\_NOT\_EXIST","error\_message":"p80aw8b06r账号不存在","is\_success":"F"}

1.  \[20302\]mybank.tc.trade.funds.unfreeze (解冻资金)
    --------------------------------------------------

    1.  ### 功能说明

        将指定会员账户下冻结的资金解冻，恢复为正常状态。

<!-- -->

1.  可调用10303接口查询余额。冻结余额=账面余额-可用余额。

2.  解冻金额不大于冻结金额。

    1.  ### 业务请求参数

  ---------------------------------------------------------------------------------------------------------------------------------------
  **参数**                 **类型（长度）**   **是否必填**   **描述**                                                 **样例**
  ------------------------ ------------------ -------------- -------------------------------------------------------- -------------------
  outer\_trade\_no         String(32)         不可空         合作方业务平台订单号**\                                  30000043551252398
                                                             **合作方提交的交易号，对于合作方全局唯一                 

  orig\_outer\_trade\_no   String(32)         不可空         原始的业务平台唯一订单号。同冻结资金的订单号中的一致。   30000043551252390

  amount                   Number(12)         不可空         解冻金额，金额不大于冻结金额                             1000.00

  is\_web\_access          String(1)          可空           预留字段，可不填                                         
  ---------------------------------------------------------------------------------------------------------------------------------------

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.funds.unfreeze");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("outer\_trade\_no", " 2017070416473019902397952974");

params.put("return\_url","http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("orig\_outer\_trade\_no", "2017070417004018132311422378");

params.put("amount", "0.03");

params.put("is\_web\_access", "Y");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T"}

异常：
{"charset":"UTF-8","error\_code":"SYSTEM\_ERROR","error\_message":"F0501解冻金额超限:payer-100001520002-20010010011100001520002000001,0.02&lt;0.03","is\_success":"F"}

[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[]{#_Toc481842435
.anchor}]{#_Toc481841192 .anchor}]{#_Toc481839955
.anchor}]{#_Toc481838718 .anchor}]{#_Toc481837481
.anchor}]{#_Toc481709442 .anchor}]{#_Toc481706970
.anchor}]{#_Toc481704498 .anchor}]{#_Toc481705734
.anchor}]{#_Toc481708206 .anchor}]{#_Toc481702030
.anchor}]{#_Toc481698438 .anchor}]{#_Toc481609941
.anchor}]{#_Toc481697207 .anchor}]{#_Toc481691979
.anchor}]{#_Toc481842352 .anchor}]{#_Toc481841109
.anchor}]{#_Toc481839872 .anchor}]{#_Toc481838635
.anchor}]{#_Toc481837398 .anchor}]{#_Toc481705651
.anchor}]{#_Toc481706887 .anchor}]{#_Toc481708123
.anchor}]{#_Toc481704415 .anchor}]{#_Toc481698355
.anchor}]{#_Toc481709359 .anchor}]{#_Toc481609858
.anchor}]{#_Toc481701947 .anchor}]{#_Toc481691896
.anchor}]{#_Toc481697124 .anchor}]{#_Toc481842351
.anchor}]{#_Toc481841108 .anchor}]{#_Toc481839871
.anchor}]{#_Toc481838634 .anchor}]{#_Toc481837397
.anchor}]{#_Toc481691895 .anchor}]{#_Toc481709358
.anchor}]{#_Toc481698354 .anchor}]{#_Toc481609857
.anchor}]{#_Toc481708122 .anchor}]{#_Toc481706886
.anchor}]{#_Toc481701946 .anchor}]{#_Toc481704414
.anchor}]{#_Toc481705650 .anchor}]{#_Toc481697123
.anchor}]{#_Toc481842338 .anchor}]{#_Toc481841095
.anchor}]{#_Toc481839858 .anchor}]{#_Toc481838621
.anchor}]{#_Toc481837384 .anchor}]{#_Toc481698341
.anchor}]{#_Toc481706873 .anchor}]{#_Toc481697110
.anchor}]{#_Toc481705637 .anchor}]{#_Toc481609844
.anchor}]{#_Toc481709345 .anchor}]{#_Toc481708109
.anchor}]{#_Toc481691882 .anchor}]{#_Toc481701933
.anchor}]{#_Toc481704401 .anchor}]{#_Toc481842337
.anchor}]{#_Toc481841094 .anchor}]{#_Toc481839857
.anchor}]{#_Toc481838620 .anchor}]{#_Toc481837383
.anchor}]{#_Toc481691881 .anchor}]{#_Toc481704400
.anchor}]{#_Toc481701932 .anchor}]{#_Toc481706872
.anchor}]{#_Toc481698340 .anchor}]{#_Toc481709344
.anchor}]{#_Toc481705636 .anchor}]{#_Toc481697109
.anchor}]{#_Toc481609843 .anchor}]{#_Toc481708108
.anchor}]{#_Toc481842336 .anchor}]{#_Toc481841093
.anchor}]{#_Toc481839856 .anchor}]{#_Toc481838619
.anchor}]{#_Toc481837382 .anchor}]{#_Toc481701931
.anchor}]{#_Toc481706871 .anchor}]{#_Toc481705635
.anchor}]{#_Toc481691880 .anchor}]{#_Toc481697108
.anchor}]{#_Toc481704399 .anchor}]{#_Toc481709343
.anchor}]{#_Toc481708107 .anchor}]{#_Toc481609842
.anchor}]{#_Toc481698339 .anchor}]{#_Toc481842335
.anchor}]{#_Toc481841092 .anchor}]{#_Toc481839855
.anchor}]{#_Toc481838618 .anchor}]{#_Toc481837381
.anchor}]{#_Toc481709342 .anchor}]{#_Toc481589630
.anchor}]{#_Toc481608791 .anchor}]{#_Toc481704398
.anchor}]{#_Toc481698338 .anchor}]{#_Toc481697107
.anchor}]{#_Toc481708106 .anchor}]{#_Toc481705634
.anchor}]{#_Toc481504755 .anchor}]{#_Toc481691879
.anchor}]{#_Toc481706870 .anchor}]{#_Toc481701930
.anchor}]{#_Toc481609841 .anchor}]{#_Toc481509219
.anchor}]{#_Toc481842334 .anchor}]{#_Toc481841091
.anchor}]{#_Toc481839854 .anchor}]{#_Toc481838617
.anchor}]{#_Toc481837380 .anchor}]{#_Toc481708105
.anchor}]{#_Toc481701929 .anchor}]{#_Toc481709341
.anchor}]{#_Toc481697106 .anchor}]{#_Toc481691878
.anchor}]{#_Toc481698337 .anchor}]{#_Toc481608790
.anchor}]{#_Toc481509218 .anchor}]{#_Toc481704397
.anchor}]{#_Toc481705633 .anchor}]{#_Toc481609840
.anchor}]{#_Toc481706869 .anchor}]{#_Toc481589629
.anchor}]{#_Toc481504754 .anchor}]{#_Toc481842333
.anchor}]{#_Toc481841090 .anchor}]{#_Toc481839853
.anchor}]{#_Toc481838616 .anchor}]{#_Toc481837379
.anchor}]{#_Toc481705632 .anchor}]{#_Toc481697105
.anchor}]{#_Toc481509217 .anchor}]{#_Toc481691877
.anchor}]{#_Toc481504753 .anchor}]{#_Toc481608789
.anchor}]{#_Toc481589628 .anchor}]{#_Toc481701928
.anchor}]{#_Toc481709340 .anchor}]{#_Toc481698336
.anchor}]{#_Toc481704396 .anchor}]{#_Toc481609839
.anchor}]{#_Toc481708104 .anchor}]{#_Toc481706868
.anchor}]{#_Toc481842254 .anchor}]{#_Toc481841011
.anchor}]{#_Toc481839774 .anchor}]{#_Toc481838537
.anchor}]{#_Toc481837300 .anchor}]{#_Toc481701849
.anchor}]{#_Toc481609760 .anchor}]{#_Toc481589549
.anchor}]{#_Toc481504674 .anchor}]{#_Toc481705553
.anchor}]{#_Toc481608710 .anchor}]{#_Toc481704317
.anchor}]{#_Toc481691798 .anchor}]{#_Toc481706789
.anchor}]{#_Toc481509138 .anchor}]{#_Toc481698257
.anchor}]{#_Toc481697026 .anchor}]{#_Toc481709261
.anchor}]{#_Toc481708025 .anchor}]{#_Toc481842253
.anchor}]{#_Toc481841010 .anchor}]{#_Toc481839773
.anchor}]{#_Toc481838536 .anchor}]{#_Toc481837299
.anchor}]{#_Toc481704316 .anchor}]{#_Toc481504673
.anchor}]{#_Toc481697025 .anchor}]{#_Toc481708024
.anchor}]{#_Toc481509137 .anchor}]{#_Toc481608709
.anchor}]{#_Toc481589548 .anchor}]{#_Toc481709260
.anchor}]{#_Toc481701848 .anchor}]{#_Toc481705552
.anchor}]{#_Toc481706788 .anchor}]{#_Toc481698256
.anchor}]{#_Toc481609759 .anchor}]{#_Toc481691797
.anchor}]{#_Toc481842252 .anchor}]{#_Toc481841009
.anchor}]{#_Toc481839772 .anchor}]{#_Toc481838535
.anchor}]{#_Toc481837298 .anchor}]{#_Toc481697024
.anchor}]{#_Toc481709259 .anchor}]{#_Toc481589547
.anchor}]{#_Toc481691796 .anchor}]{#_Toc481609758
.anchor}]{#_Toc481608708 .anchor}]{#_Toc481504672
.anchor}]{#_Toc481509136 .anchor}]{#_Toc481705551
.anchor}]{#_Toc481708023 .anchor}]{#_Toc481701847
.anchor}]{#_Toc481698255 .anchor}]{#_Toc481704315
.anchor}]{#_Toc481706787 .anchor}

1.  \[20401\]mybank.tc.trade.carryover(资金结转)
    --------------------------------------------

    1.  ### 功能说明

        交易订单完成后，可通过该接口，将结转订单总金额归并到银存账户。

<!-- -->

1.  对于入款交易，结转后，对于入款交易，结转后，系统自动核销入款渠道待清算。

2.  合作方将已同步到见证平台需要结转的交易生成文件，上传到交易见证平台sftp目录，然后调用该接口向交易见证平台发起结转指令后，平台才处理结转请求。

3.  平台需保证文件名唯一

    1.  ### 业务请求参数 

  **参数**           **类型（长度）**   > **是否必填**   > **描述**                            **样例**
  ------------------ ------------------ ---------------- ------------------------------------- ---------------------------------------------------------
  > request\_no      String(32)         不可空           业务请求号                            > 20131105154925
  > file\_path       String(500)        不可空           文件所在SFTP服务器路径                /226610000051657147614/20170308/
  > file\_name       String(150)        不可空           文件名称,格式说明请参考下方参数说明   h2h\_carryover\_226610000051657147614\_201706280001.csv
  > funds\_channel   String(64)         不可空           结转的资金渠道                        > MYBANK1010
  > total\_count     Number(10)         不可空           总笔数。对应提现csv文件的总笔数       > 3
  > total\_amount    Number(15)         不可空           总金额。对应提现csv文件的总金额       > 5.94
  > memo             String(64)         可空             备注                                  

【参数说明】

> 1\. 文件为csv格式，file\_name文件名称规范命名：h2h\_carryover\_\${ pid
> }\_\${ request\_no }.csv
>
> pid为商户平台ID
>
> request\_no为业务请求流水号。流水号建议以YYYYMMDD开头。
>
> 文件参考：
>
> 文件栏位定义如下
>
> outerInstOrderNo：合作方平台自接入/出款渠道订单号。用以关联原即时入账、担保入账、充值入账、退款、提现等场景下使用的入/出款渠道订单号。
>
> amount：支付金额。
>
> bizType：业务类型。I 入款，O 提现，B 退款
>
> tranDate:实际出金/入金日期。格式YYYY-MM-DD。

1.  ### 请求示例

    示例公共代码详见 [*接口规范-调用示例*](#调用示例)

Map&lt;String, String&gt; params = **new** HashMap&lt;String,
String&gt;();

params.put("charset", "UTF-8");

params.put("partner\_id", "200000900001");

params.put("service", "mybank.tc.trade.carryover");

params.put("sign", "A3vhlJSakixi....");

params.put("sign\_type", "RSA");

params.put("version", "2.0");

params.put("file\_name", "
h2h\_batchPay\_226610000051657147614\_20170704173510990696577876358618.xls");
params.put("return\_url",
"http://func121admin.vfinance.cn/gop-test/mag/syncNotify.htm");

params.put("file\_path", "/226610000051657147614/20170704");

params.put("funds\_channel", "WXPAY00001");

params.put("operator\_id", "10005454");

params.put("request\_no","20170704173510990696577876358618");

params.put("total\_amount", "0.01");

params.put("total\_count", "1");

### 响应示例

正常：{"charset":"UTF-8","is\_success":"T","memo":"结转申请成功"}

异常：{"charset":"UTF-8","error\_code":"E0001","error\_message":"重复的业务请求号:20170704183115556780518691564350","is\_success":"F"}

1.  \[30101\]mybank.tc.user.area.query (省、市查询)
    -----------------------------------------------

    1.  ### 功能说明

省、市、区信息查询

### 使用场景

合作方会员在绑卡时，填写绑卡信息时使用。

### 业务请求参数

  参数                     类型（长度）   是否必填   描述                                     样例
  ------------------------ -------------- ---------- ---------------------------------------- ------
  parent\_district\_code   String(32)     不可空     父节点。若查询省、直辖市，则填写ROOT。   ROOT

**参数说明：**

查询省市区域parent\_district\_code 传的是父节点的区域代码
  查询结果返回区域的信息。例如：查询省、直辖市其值为ROOT，查询市其值为市所属的省的区域代码；查询区其值为该区所属的市的区域代码。

**使用步骤：**

1.  parent\_district\_code 的值设为ROOT，查询出所有的省、直辖市

2.  parent\_district\_code的值设为1中查询出的district\_code(地区编号)，查询省下的市或直辖市

3.  parent\_district\_code的值设为2中查询出的district\_code(地区编号)，查询市或直辖市下的区

    1.  ### 业务返回参数

  参数                      类型（长度）   是否必填   描述                     样例
  ------------------------- -------------- ---------- ------------------------ --------
  district\_list                                      地区列表                 
  └district\_name           String(32)     不可空     地区名称（省、市、区）   安徽省
  └district\_code           String(10)     不可空     地区编号                 340000
  └post\_code               String(64)     可空       邮编（省、直辖市为空）   
  └parent\_diatrict\_code   String(32)     不可空     父节点                   ROOT
  └memo                     String(256)    可空       备注                     

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#_调用实例)

Map&lt;String, String&gt; params = new HashMap&lt;String, String&gt;();

> params.put("charset", "UTF-8");
>
> params.put("partner\_id", "200004620102");
>
> params.put("service", " mybank.tc.user.area.query");
>
> params.put("version", "2.0");
>
> params.put("sign", "MIIGGAYJKoZIhvcNAQ…");
>
> params.put("sign\_type", "RSA");
>
> params.put("parent\_district\_code", "ROOT");
>
> ;

### 返回示例

正常：{“charset”:”UTF-8”,”district\_list”:\[{“district\_name”:”安徽省”,”
district\_code”:”340000”,”post\_code”:”0254521”,”parent\_code”:”ROOT”}\],”is\_success”:”T”，”memo”:””,”partner\_id”:”188888888888”}

异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_SIGN","error\_message":"验签未通过","is\_success":"F"}

1.   \[30102\]mybank.tc.user.area.bank.query (银行列表查询) 
    --------------------------------------------------------

    1.  ### 功能说明

银行列表联行号信息查询

### 使用场景

合作方会员在绑卡时，填写绑卡信息时使用。

### 业务请求参数

  参数                 类型（长度）   是否必填   描述                                        样例
  -------------------- -------------- ---------- ------------------------------------------- --------
  parent\_branch\_no   String(20)     不可空     父联行号，[*参考附录14.4*](#银行总行列表)   ROOT
  area\_code           String(20)     不可空     区域编码                                    341800
  key\_words           String(255)    可空       查询关键字                                  宣州

**参数说明：**

parent\_branch\_code：查询区域的所有银行其值为ROOT，若查询具体银行其值为具体的银行的branch\_no.

area\_code：只能查市下的分支行信息，省级的不可以 所以area\_code
要传市的区域码。

key\_words：如果传入关键字会根据关键字过滤查询结果，过滤规则为
返回的银行名称中包含此关键字的银行列表

### 业务返回参数

  参数               类型（长度）   是否必填   描述       样例
  ------------------ -------------- ---------- ---------- ----------------------------------
  area\_bank\_list                             银行列表   
  └bank\_name        String(32)     不可空     银行名称   安徽宣州湖商村镇银行股份有限公司
  └branch\_no        String(32)     不可空     联行号     320377100015
  └memo              String(256)    可空       备注       

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#_调用实例)

> Map&lt;String, String&gt; params = new HashMap&lt;String,
> String&gt;();
>
> params.put("charset", "UTF-8");
>
> params.put("partner\_id", "200004620102");
>
> params.put("service", "mybank.tc.user.area.bank.query");
>
> params.put("version", "2.0");
>
> params.put("sign", "MIIGGAYJKoZIhvcNAQ…");
>
> params.put("sign\_type", "RSA");
>
> params.put("parent\_branch\_no", "ROOT");
>
> params.put("key\_words ", "");
>
> params.put("area\_code", "340000");

### 返回示例

> 正常：{“charset”:”UTF-8”,”area\_bank\_list”:\[{“bank\_name”:”安徽宣州湖商村镇银行股份有限公司”,”branch\_code”:”320377100015”
> }\],”is\_success”:”T”，”memo”:””,”partner\_id”:”188888888888”}
>
> 异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_SIGN","error\_message":"验签未通过","is\_success":"F"}

1.   \[30103\]mybank.tc.user.cardbin.query (卡BIN查询)
    --------------------------------------------------

    1.  ### 功能说明

        本接口提供银行卡bin查询

    2.  ### 使用场景

        适用于合作方业务平台在会员绑卡时对绑定的银行卡进行卡bin校验。

    3.  ### 业务请求参数 

  参数                类型（长度）   是否必填   描述       样例
  ------------------- -------------- ---------- ---------- ------------------
  bank\_account\_no   String(32)     不可空     银行卡号   8888888453427229

### 业务返回参数

  参数              类型（长度）   是否必填   描述           样例
  ----------------- -------------- ---------- -------------- -----------------------------------
  card\_bin\_info   String(32)     不可空     卡bin信息      
  └card\_bin        String(32)     不可空     卡bin信息      888888
  └card\_type       String(32)     不可空     卡类型         DC:借记卡，CC:贷记卡,SCC:准贷记卡
  └branch\_no       String(32)     不可空     机构联行号     323331000001
  └bank\_name       String(128)    不可空     所属银行名称   浙江网商银行股份有限公司

### 请求示例

示例公共代码详见 [*接口规范-调用示例*](#_调用实例)

> Map&lt;String, String&gt; params = new HashMap&lt;String,
> String&gt;();
>
> params.put("charset", "UTF-8");
>
> params.put("partner\_id", "200004620102");
>
> params.put("service", "mybank.tc.user.area.bank.query");
>
> params.put("version", "2.0");
>
> params.put("sign", "MIIGGAYJKoZIhvcNAQ…");
>
> params.put("sign\_type", "RSA");
>
> params.put("bank\_account\_no ", "8888888453427229");

### 返回示例

> 正常：
> {"charset":"UTF-8","is\_success":"T","card\_bin\_info":{"bank\_name":"招商银行","branch\_no":"308584000013","card\_bin":"622588","card\_type":"DC"}}
>
> 异常：{"charset":"UTF-8","error\_code":"ILLEGAL\_SIGN","error\_message":"验签未通过","is\_success":"F"}

1.  \[30104\] mybank.tc.user.partner.balance.query (合作方账户余额查询)
    -------------------------------------------------------------------

    1.  ### 功能说明

合作方账户余额查询。

### 使用场景

合作方可根据自己的业务调用此接口查询网商账户余额。

### 业务请求参数

无

### 业务返回参数

  --------------------------------------------------------------------------------------------
  参数                 类型（长度）   是否必填   描述                                   样例
  -------------------- -------------- ---------- -------------------------------------- ------
  card\_no             String(32)     不可空     账号                                   

  account\_name        String(32)     不可空     账户名                                 

  available\_balance   Number         不可空     可用余额                               
                                                                                        
                                                 单位为：RMB 元。精确到小数点后两位。   

  actual\_balance      Number         不可空     实际余额                               
                                                                                        
                                                 单位为：RMB 元。精确到小数点后两位。   

  freeze\_amount       Number         可空       冻结金额                               
                                                                                        
                                                 单位为：RMB 元。精确到小数点后两位。   
  --------------------------------------------------------------------------------------------

### 请求示例

示例公共代码详见
[*接口规范-调用示例*](file:///D:\E\svn全集\网商银行_mybank_new\04-设计文档\0401-对外接口文档\新建文件夹\浙江网商银行交易见证服务API-余额查询.docx#_调用实例)

> Map&lt;String, String&gt; params = new HashMap&lt;String,
> String&gt;();
>
> params.put("charset", "UTF-8");
>
> params.put("partner\_id", "200004620102");
>
> params.put("service", " mybank.tc.user.partner.balance.query");
>
> params.put("version", "2.0");
>
> params.put("sign", "MIIGGAYJKoZIhvcNAQ…");
>
> params.put("sign\_type", "RSA");

### 返回示例

> 正常：
>
> {"charset":"UTF-8","is\_success":"T"," card\_no":"888888\*\*\*365","
> account\_name
> ":"xxx网商银行账户","available\_balance":"10.00","actual\_balance":"10.00","freeze\_amount":"0.00"}
>
> 异常：
> {"charset":"UTF-8","error\_code":"ILLEGAL\_ARGUMENT","error\_message":"参数校验未通过","is\_success":"F","memo":"{\\"REQUIRED\_FIELD\_NOT\_EXIST\\":\[\\"合作方平台ID不能为空\\"\]}"}

1.  交易对账
    ========

    1.  对账说明
        --------

所谓交易对账是商户交易流水的对账，对账流程如下：

1.  商户同步交易流水到交易见证平台，交易见证平台提供同步的交易流水下载功能，商户登录PC-WEB下载交易流水。

2.  商户T+1日下载T日的交易流水对账文件。

3.  交易对账以每笔流水的 渠道+业务类型+机构订单号+金额为匹配条件。

4.  商户对账完成后发送结转请求到见证系统。

    1.  错账处理
        --------

<!-- -->

1.  商户异步通知
    ============

    1.  通知说明
        --------

        1.  ### 约定规范

<!-- -->

1.  通知采用POST方式发送

2.  签名原文使用“charset”字符集编码进行签名；

3.  商户正常处理完通知后必须输出“success”（不包含引号且小写）；

4.  签名原文未经URL Encoding；

5.  如果商户处理通知没有输出“success”，则服务器会不断重发通知，重试时间间隔为：2m，10m，10m，1h，2h，6h，15h（共通知8次后不再通知）；

    1.  ### 通知条件

  **通知类型**   **通知状态**            **说明**
  -------------- ----------------------- ------------
  交易状态同步   PAY\_SUCCESS            买家已付款
                 TRADE\_SUCCESS          交易成功
                 TRADE\_FINISHED         交易结束
                 TRADE\_FAILED           交易失败
                 TRADE\_CLOSED           交易关闭
  退款状态同步   REFUND\_FINISH          退款成功
                 REFUND\_FAIL            退款失败
                 REFUND\_PROCESS         退款中
  支付状态同步   PAY\_SUCCESS            支付成功
                 PAY\_FAIL               支付失败
  提现状态同步   WITHDRAWAL\_SUBMITTED   已提交银行
                 TRADE\_FINISHED         提现成功
                 TRADE\_FAILED           提现失败
                 TRADE\_FAILED           已提交
                 RETURN\_TICKET          已退票
  转账状态同步   TRADE\_FINISHED         转账成功
                 TRADE\_FAILED           转账失败

1.  RSA签名机制
    -----------

    1.  ### 基本说明

<!-- -->

1.  请求的所有参数，需要根据参数名=参数值的格式，按首字符字典顺序（ascii值大小）排序，若遇到相同首字符，则判断第二个字符，以此类推，待签名字符串需要以“参数名1=参数值1&参数名2=参数值2&….&参数名N=参数值N”的规则进行拼接。

2.  在对请求的参数做签名时，这些参数必须来源于请求参数列表，并且除去列表中的参数sign、sign\_type。

3.  在对请求的参数做签名时，对于请求参数列表中那些可空的参数，如果选择使用它们，那么这些参数的参数值必须不能为空或空值。

4.  签名时将字符转化成字节流时指定的字符集与charset保持一致；如果传递了charset参数，这个参数也应该包含在待签名数据中。

5.  待签名数据应该是参数原始值而不是url
    encoding之后的值，例如：调用某接口需要对请求参数email进行数字签名，那么待签名数据应该是email=test@msn.com，而不是email=test%40msn.com。

    1.  ### 签名字符串

<!-- -->

1.  需要参与签名的参数
    在请求参数列表中，除去sign、sign\_type两个参数外，其他需要使用到的参数皆是要签名的参数。（个别接口中参数sign\_type也需要参与签名。）在通知返回参数列表中，除去sign、sign\_type两个参数外，凡是通知返回回来的参数皆是要签名的参数。

2.  生成待签名字符串

> 对于如下的参数数组：
>
> string\[\] parameters={
>
> "service=create\_partner\_trade\_by\_buyer",
>
> "partner\_id=2088002007018916",
>
> "charset=gbk",
>
> "return\_url=http://www.test.com/qijian/return\_url.asp",
>
> "out\_trade\_no=709651609727679",
>
> "subject=nokia n8",
>
> "price=3003",
>
> "quantity=1",
>
> "seller\_email=zhoubo\_seller@qjtest.com"
>
> };
>
> 对数组里的每一个值从a到z的顺序排序，若遇到相同首字母，则看第二个字母，以此类推。
>
> 排序完成之后，再把所有数组值以“&”字符连接起来，如：
>
> charset=gbk&out\_trade\_no=709651609727679&partner\_id=2088002007018916&payment\_type=1
>
> &price=3003&quantity=1&return\_url=http://www.test.com/qijia/return\_url.asp
>
> &seller\_email=zhoubo\_seller@test.com&service=create\_partner\_trade\_by\_buyer&subject=nokia
> n8
>
> 这串字符串便是待签名字符串

1.  没有值的参数无需传递，也无需包含到待签名数据中；签名时将字符转化成字节流时指定的字符集与charset保持一致；如果传递了charset参数，这个参数也应该包含在待签名数据中；

> 根据HTTP协议要求，传递参数的值中如果存在特殊字符（如：&、@等），那么该值需要做URL
> Encoding，这样请求接收方才能接收到正确的参数值。这种情况下，待签名数据应该是原生值而不是encoding之后的值。例如：调用某接口需要对请求参数email进行数字签名，那么待签名数据应该是email=test@msn.com，而不是email=test%40msn.com。

1.  异步通知接口
    ------------

    1.  ### 公共响应参数

  -----------------------------------------------------------------------------------------------------------------------------------------
  **参数**       **参数名称**     **类型长度**   **参数说明**                            **是否必填**   **样例**
  -------------- ---------------- -------------- --------------------------------------- -------------- -----------------------------------
  notify\_id     通知ID           String(32)     通知的唯一标识                          不可空         2bf30f9cc5c646f5acbdde31d91947df

  notify\_type   通知类型         String         交易通知此字段为：trade\_status\_sync   不可空         trade\_status\_sync(交易状态同步)

  notify\_time   通知时间         String         通知的发送时间，格式：                  不可空         20131101102030
                                                                                                        
                                                 yyyyMMddHHmmss                                         

  charset        参数字符集编码   String         和商户请求中的字符集编码一样            不可空         UTF-8

  sign           签名             String         见[*签名机制*](#rsa签名机制)            不可空         

  sign\_type     签名方式         String         RSA，必须大写                           不可空         RSA

  version        版本号           String         接口版本号                              不可空         2.0

  productCode    产品编码         String         预留字段                                可空           
  -----------------------------------------------------------------------------------------------------------------------------------------

### 交易处理结果通知

> 说明:除退票、退款外的，其他交易类型包括充值、转账、及时到账交易、担保交易使用该异步通知，

当一笔交易的交易状态是交易结束（TRADE\_FINISHED）时，才表示这笔交易成功。

### 通知结果

  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数**                                  **参数名称**         **类型长度**   **参数说明**                                                    **是否必填**   **样例**
  ----------------------------------------- -------------------- -------------- --------------------------------------------------------------- -------------- ------------------------------
  outer\_trade\_no                          商户网站唯一订单号   String(32)     合作方业务平台唯一订单号                                        不可空         2015070609402382115027336039

  inner\_trade\_no                          支付系统交易号       String(64)     交易凭证号                                                      不可空         101143651423149320100

  gmt\_close                                交易处理完成时间     String         交易处理完成时间，交易状态为：TRADE\_FINISHED时才有值，格式：   **不**可空     20131101102030
                                                                                                                                                               
                                                                                yyyyMMddHHmmss                                                                 
                                                                                                                                                               
                                                                                **商户平台需要落地该时间，作为与交易见证平台的对账时间**                       

  []{#_Hlk439080084 .anchor}trade\_status   交易状态             String(20)     见附录                                                          不可空         TRADE\_FINISHED

  error\_code                               错误代码             String(10)                                                                     可空           

  error\_msg                                错误信息             String(500)                                                                    可空           
  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### 提现通知结果(单笔)

  -------------------------------------------------------------------------------------------------------------------------------------------------------
  **参数**             **参数名称**             **类型长度**   **参数说明**                                         **是否必填**   **样例**
  -------------------- ------------------------ -------------- ---------------------------------------------------- -------------- ----------------------
  outer\_trade\_no     商户网站提现唯一订单号   String(32)     标识商户提现订单的唯一凭证号，失败后再发起需要更换   不可空         29388034012938822409

  inner\_trade\_no     交易见证平台交易号       String(32)     提现交易订单号                                       可空           

  withdrawal\_amount   提现金额                 Number         提现金额，单位为RMB-Yuan，精确到小数点后两位         可空           

  withdrawal\_status   提现状态                 String         参见*[提现状态](#提现状态明细)(明细)*                不可空         

  error\_code          错误代码                 String         失败时指出具体原因                                   可空           

  error\_msg           错误信息                 String         提现失败原因                                         可空           

  gmt\_close           交易处理完成时间         String         格式：                                               不可空         
                                                                                                                                   
                                                               yyyyMMddHHmmss                                                      

  uid                  用户ID                   String         合作方业务平台用户ID                                 不可空         

  account\_type        账户类型                 String(5)      账户类型                                             不可空         BASIC

  account\_identity    账户标识                 String(32)     预留字段，可不填                                     可空           
  -------------------------------------------------------------------------------------------------------------------------------------------------------

1.  附录
    ====

    1.  交易状态
        --------

        1.  ### 担保交易状态

  枚举名称           枚举说明
  ------------------ ------------------------
  WAIT\_BUYER\_PAY   交易创建，等待买家付款
  PAY\_SUCCESS       买家已付款
  TRADE\_SUCCESS     交易成功
  TRADE\_FINISHED    交易结束
  TRADE\_FAILED      交易失败
  TRADE\_CLOSED      交易关闭

### 即时到账交易状态

  枚举名称           枚举说明
  ------------------ ------------------------
  WAIT\_BUYER\_PAY   交易创建，等待买家付款
  TRADE\_SUCCESS     交易成功
  TRADE\_FINISHED    交易结束
  TRADE\_FAILED      交易失败
  TRADE\_CLOSED      交易关闭

### 退款状态

  枚举名称          枚举说明
  ----------------- ----------
  REFUND\_FINISH    退款成功
  REFUND\_FAIL      退款失败
  REFUND\_PROCESS   退款中

### 支付状态

  **枚举名称**   **枚举说明**
  -------------- --------------
  PAY\_SUCCESS   支付成功
  PAY\_FAIL      支付失败

### 提现状态(明细)

  **枚举名称**      **枚举说明**
  ----------------- --------------
  TRADE\_FINISHED   提现成功
  TRADE\_FAILED     提现失败
  TRADE\_SUCCESS    已提交

[]{#_提现状态(批量) .anchor}

### 退票

  **枚举名称**     **枚举说明**
  ---------------- --------------
  REFUND\_TICKET   退票成功

### 转账状态

  **枚举名称**      **枚举说明**
  ----------------- --------------
  TRADE\_FINISHED   转账成功
  TRADE\_FAILED     转账失败

1.  公共错误码
    ----------

    1.  ### 系统错误码

  **错误代码（error\_code）**       **含义**
  --------------------------------- ----------------------------
  SYSTEM\_ERROR                     系统内部错误
  SESSION\_TIMEOUT                  session超时
  ILLEGAL\_ACCESS\_SWITCH\_SYSTEM   商户不允许访问该类型的接口

### 业务错误码

  **错误代码（error\_code）**           **含义**
  ------------------------------------- ---------------------------------------------------------
  REQUIRED\_FIELD\_NOT\_EXIST           必填字段未填
  FIELD\_LENGTH\_EXCEEDS\_LIMIT         字段长度超过限制
  FIELD\_TYPE\_ERROR                    字段类型错误
  PARTNER\_ID\_NOT\_EXIST               合作方Id不存在
  TRADE\_DATA\_MATCH\_ERROR             交易信息有误
  TRADE\_AMOUNT\_MATCH\_ERROR           交易内金额不匹配
  TRADE\_PAY\_MATCH\_ERROR              交易与支付金额不匹配
  TRADE\_NO\_MATCH\_ERROR               交易号信息有误
  ILLEGAL\_REQUEST                      风控未通过
  SELLER\_NOT\_EXIST                    交易卖家不存在
  TRADE\_BUYER\_NOT\_MATCH              输入的买家与交易买家不匹配
  TRADE\_SELLER\_NOT\_MATCH             输入的卖家与交易卖家不匹配
  TOTAL\_FEE\_LESSEQUAL\_ZERO           交易总额小于等于0
  TOTAL\_FEE\_GREATER\_THAN\_MAX        担保交易单笔总金额不得超过1000000(100万)
  EXTERFACE\_INVOKE\_CONTEXT\_EXPIRED   接口调用上下文过期
  ILLEGAL\_SIGN\_TYPE                   签名类型不正确
  ILLEGAL\_SIGN                         验签未通过
  ILLEGAL\_ARGUMENT                     参数校验未通过
  ILLEGAL\_SERVICE                      服务接口不存在
  ILLEGAL\_ID\_TYPE                     ID类型不存在
  USER\_ACCOUNT\_NOT\_EXIST             用户账号不存在
  MEMBER\_ID\_NOT\_EXIST                用户MemberId不存在
  MEMBER\_OR\_ACCOUNT\_QUERY\_FAIL      会员/账户询失败(可能是memberID不存在也可能是账户不存在)
  MOBILE\_NOT\_EXIST                    用户手机号不存在
  ILLEGAL\_BUYER\_INFO                  买家内部Id，外部Id或手机号不匹配
  ILLEGAL\_SELLER\_INFO                 卖家内部Id，外部Id或手机号不匹配
  ILLEGAL\_ROYALTY\_PARAMETERS          分润账号集错误
  ILLEGAL\_SUBSCRIPTION\_ORDER\_NO      订金下订单号错误
  ILLEGAL\_SUBSCRIPTION                 订金金额错误
  ILLEGAL\_REFUND\_AMOUNT               退款金额信息错误
  PAY\_METHOD\_ERROR                    支付方式错误
  ILLEGAL\_PAY\_METHOD                  支付方式未授权
  DUPLICATE\_REQUEST\_NO                重复的请求号
  ILLEGAL\_OUTER\_TRADE\_NO             交易订单号不存在
  ILLEGAL\_DATE\_FORMAT                 日期格式错误
  ILLEGAL\_AMOUNT\_FORMAT               金额格式错误
  OPERATOR\_ID\_NOT\_EXIST              操作员Id不存在
  ILLEGAL\_ENSURE\_AMOUNT               担保金额信息错误
  ILLEGAL\_TIME\_INTERVAL               时间区间错误
  ILLEGAL\_PAY\_ERROR                   []{#OLE_LINK11 .anchor}支付方式错误
  ILLEGAL\_ORDER\_NO\_FORMAT            订单号格式错误
  ORIGINAL\_VOUCHER\_INEXISTENCE        原始凭证号不存在
  PARSE\_DISCOUNT\_PAY\_METHOD\_ERROR   折扣支付信息有误
  NOT\_SUPPORT\_DISCOUNT\_PAY\_METHOD   不支持的折扣方式
  ACCOUNT\_NAME\_DECRYPT\_ERROR         银行卡账户名解密失败
  CARD\_NO\_DECRYPT\_ERROR              银行卡号解密失败
  CHARSET\_ERROR                        编码类型错误
  ACCOUNT\_TYPE\_NOT\_SUPPORTED         账户类型不支持
  TRADE\_TYPE\_NOT\_SUPPORTED           交易类型不支持
  ILLEGAL\_FUND\_TYPE                   资金类型不支持
  TRANSFER\_TYPE\_NOT\_SUPPORTED        转账类型不支持
  CARD\_ATTRIBUTE\_NOT\_SUPPORTED       卡属性不支持，只支持 B对公，C对私
  CARD\_TYPE\_NOT\_SUPPORTED            卡类型不支持，只支持 D借记，C贷记
  CHECK\_PAY\_CHANNEL\_FAIL             网银支付渠道校验失败
  PARTNER\_ID\_ERROR                    商户信息错误
  CERT\_ERROR                           签名证书错误
  CERT\_EXPIRED                         签名证书过期
  FIELD\_LENGTH\_EXCEEDS\_ERROR         字段长度不合法
  FIELD\_CANNOT\_BE\_ZERO               字段值不允许为0
  FIELD\_CHECK\_ERROR                   参数类型校验错误
  CERT\_DOWNLOAD\_FAIL                  证书下载失败
  BANK\_CARD\_NOT\_MEMBERID             信息不匹配
  WHITE\_CARD\_IS\_EXIST                白名单卡已存在
  WHITE\_CARD\_IS\_NOT\_EXIST           白名单卡不存在
  BUSINESS\_VOUCHER\_INEXISTENCE        查询无此交易记录
  UPLOAD\_IMAGE\_COUNT\_LIMIT           上传图片数量超出限制
  FILE\_NAME\_ERROR                     文件名格式错误
  AMOUNT\_NOT\_NULL                     支付方式不为空时，充值金额不能为空
  ILLEGAL\_INVALIDDATE\_FORMAT          过期时间格式错误
  UFS\_ERROR                            UFS调用失败
  CERT\_ADDREPEAT\_FAIL                 证书重复添加
  ILLEGAL\_BANKCARD                     明细中部分银行卡验证不通过
  CARRYOVER\_DATA\_MATCH\_ERROR         结转信息有误
  UID\_NOT\_EXIST                       用户uId不存在
  ILLEGAL\_ACCOUNT\_TYPE                账户信息不正确
  UPLOAD\_FILE\_ERROR                   上传文件错误
  PAY\_ATTRIBUTE\_ERROR                 银行卡支付属性错误
  ILLEGAL\_ARGUMENT                     参数校验未通过
  ILLEGAL\_ARGUMENT\_MEMBER             会员校验未通过
  ILLEGAL\_ARGUMENT\_ACCOUNT            账户校验未通过
  DUPLICATE\_REQUEST\_NO                重复的请求号
  ILLEGAL\_ARGUMENT\_BANK\_CARD         银行卡校验未通过
  BANK\_WITHDRAWAL\_FAIL                银行出款失败
  BATCH\_ILLEGAL\_ARGUMENT              批次校验不通过
  BATCH\_DETAIL\_ILLEGAL\_ARGUMENT      批次明细校验不通过

证件类型 
---------

  枚举名称                     枚举值     枚举说明
  ---------------------------- ---------- ------------------------
  对私                                    
  ID\_CARD                     ID\_CARD   身份证
  TEMP\_ID\_CARD               TIC        临时身份证
  RESIDENCE\_BOOKLET           RB         户口簿
  PASSPORT                     PP         护照
  ARMY\_OFFICER\_CARD          AOC        军官证
  ARMY\_CIVILIAN\_CADRE        ACC        军人文职干部证
  SOLDIER\_CARD                SC         士兵证
  HK\_MC\_PASS                 HMP        港澳居民来往内地通行证
  TW\_PASS                     TWP        台湾同胞来往内地通行证
  TW\_RETURN                   TWR        台湾回乡证
  FOREIGNER\_RESIDENCE         FR         外国人居留证
  FOREIGN\_PASSPORT            FPP        外国护照
  POLICE\_OFFICER\_CARD        POC        警官证
  ARMED\_POLICE\_CARD          APC        武警证
  OTHER                        OTHER      其他证件
  对公：                                  
  ORGANIZATION\_NO             ORGN       组织机构代码证号
  BUSINESS\_LICENSE\_NO        BULN       营业执照号码
  CERT\_OF\_REGIST             COR        登记证书
  NATIONAL\_TAX\_NO            NTN        国税登记证号码
  LOCAL\_TAX\_NO               LTN        地税登记证号码
  PERMIT\_OPEN\_ACCOUNT        POA        开户许可证
  BUSINESS\_UNIT\_NO           BUN        事业单位编号
  COMPANY\_OTHER               CO         其他证件
  FINANCIAL\_LICENSE\_NUMBER   FLN        金融许可证编号

银行总行列表
------------

  **银行行号**   **银行机构名称**
  -------------- -----------------------------------------------
  313493080539   洛阳银行股份有限公司
  402332010004   宁波鄞州农村商业银行股份有限公司(鄞州银行)
  313146000019   廊坊银行股份有限公司
  314302066666   无锡农村商业银行股份有限公司（不对外）
  313228060009   营口沿海银行股份有限公司
  313655091983   自贡银行股份有限公司
  313821050016   甘肃银行股份有限公司
  313881000002   乌鲁木齐银行清算中心
  313673093259   四川天府银行股份有限公司
  402301099998   江苏省农村信用社联合社
  313662000015   遂宁银行
  313658000014   长城华西银行股份有限公司
  313641099995   海南银行股份有限公司
  323290000016   上海华瑞银行
  531290088881   花旗银行(中国)有限公司
  314304099999   江苏江南农村商业银行股份有限公司(不对外)
  313234001089   朝阳银行股份有限公司
  313657092617   泸州市商业银行
  323653010015   重庆富民银行股份有限公司
  313501080608   焦作中旅银行股份有限公司
  313665092924   乐山市商业银行
  313463000993   泰安银行股份有限公司
  313741095715   云南红塔银行股份有限公司
  717110000010   中德住房储蓄银行
  323331000001   浙江网商银行股份有限公司
  314110000011   天津滨海农村商业银行股份有限公司
  402602000018   东莞农村商业银行股份有限公司
  313221030008   盛京银行清算中心
  313138000019   张家口银行股份有限公司
  321667090019   重庆三峡银行股份有限公司
  402701002999   贵州省农村信用社联合社
  313551070008   华融湘江银行股份有限公司
  313175000011   晋中银行
  313736000019   曲靖市商业银行
  313397075189   泉州银行股份有限公司
  402791000010   陕西省农村信用社联合社资金清算中心
  313193057846   乌海银行股份有限公司
  313421087506   江西银行股份有限公司
  402581090008   广东省农村信用社联合社
  781393010011   厦门国际银行股份有限公司
  787290000019   富邦华一银行有限公司
  314651000000   成都农商银行
  313338009688   金华银行股份有限公司
  402421099990   江西省农村信用社联合社
  313134000011   保定银行股份有限公司
  313332090019   宁波通商银行股份有限公司
  325290000012   上海银行股份有限公司
  313148053964   衡水银行股份有限公司
  402521090019   武汉农村商业银行股份有限公司
  323584000888   深圳前海微众银行股份有限公司
  313455000018   东营银行股份有限公司
  313495081900   平顶山银行股份有限公司
  317110010019   天津农村商业银行股份有限公司
  402451000010   山东省农村信用社联合社
  313521006000   湖北银行股份有限公司
  402651020006   四川省农村信用社联合社
  313491099996   中原银行股份有限公司
  313651099999   成都银行
  502290000006   东亚银行（中国）有限公司
  314302200018   江苏江阴农村商业银行股份有限公司
  313168000003   晋城银行
  313226009000   丹东银行清算中心
  402491000026   河南省农村信用社联合社资金清算中心(不转汇)
  402221010013   辽宁省农村信用社联合社运营管理部
  313231000013   辽阳银行
  313424076706   九江银行股份有限公司
  314641000014   海口联合农村商业银行股份有限公司
  313454000016   枣庄银行股份有限公司
  313791000015   西安银行股份有限公司
  313312300018   江苏长江商业银行
  314305106644   太仓农村商业银行
  402584009991   深圳农村商业银行股份有限公司
  313671000017   宜宾市商业银行
  313872097457   石嘴山银行股份有限公司
  313701099012   贵州银行股份有限公司
  402100000018   北京农村商业银行股份有限公司
  313586000006   广东华兴银行股份有限公司
  313451000019   齐鲁银行
  402361018886   安徽省农村信用联社资金清算中心
  313223007007   鞍山市商业银行
  313192000013   包商银行股份有限公司
  313100000013   北京银行
  318110000014   渤海银行股份有限公司
  313143005157   沧州银行
  313141052422   承德银行股份有限公司
  313222080002   大连银行
  313468000015   德州银行股份有限公司
  313602088017   东莞银行股份有限公司
  313205057830   鄂尔多斯银行股份有限公司
  314588000016   佛山顺德农村商业银行股份有限公司
  313391080007   福建海峡银行股份有限公司
  402391000068   福建省农村信用社联合社
  313229000008   阜新银行结算中心
  313731010015   富滇银行股份有限公司运营管理部
  313428076517   赣州银行股份有限公司
  313591001001   广东南粤银行股份有限公司
  306581000003   广发银行股份有限公司
  313611001018   广西北部湾银行
  402611099974   广西壮族自治区农村信用社联合社
  314581000011   广州农村商业银行股份有限公司
  313581003284   广州银行
  313701098010   贵阳市商业银行
  313617000018   桂林银行股份有限公司
  313261000018   哈尔滨银行结算中心
  402641000014   海南省农村信用社联合社资金清算中心
  313127000013   邯郸市商业银行股份有限公司
  597100000014   韩亚银行（中国）有限公司
  313521000011   汉口银行资金清算中心
  313331000014   杭州银行股份有限公司
  313121006888   河北银行股份有限公司
  315456000105   恒丰银行
  313227600018   葫芦岛银行股份有限公司
  402521000032   湖北省农村信用社联合社结算中心
  313336071575   湖州银行股份有限公司
  304100040000   华夏银行股份有限公司总行
  319361000013   徽商银行股份有限公司
  402241000015   吉林省农村信用社联合社（不办理转汇业务）
  313241066661   吉林银行
  313461000012   济宁银行股份有限公司
  313335081005   嘉兴银行股份有限公司清算中心(不对外办理业务）
  314305506621   江苏常熟农村商业银行股份有限公司清算中心
  313301099999   江苏银行股份有限公司
  301290000007   交通银行
  313227000012   锦州银行
  313161000017   晋商银行股份有限公司
  313882000012   昆仑银行股份有限公司
  314305206650   昆山农村商业银行
  313463400019   莱商银行
  313821001016   兰州银行股份有限公司
  313473070018   临商银行股份有限公司
  313614000012   柳州银行股份有限公司清算中心
  313261099913   龙江银行股份有限公司
  313659000016   绵阳市商业银行
  313301008887   南京银行股份有限公司
  313191000011   内蒙古银行
  313332082914   宁波银行股份有限公司
  402871099996   宁夏黄河农村商业银行股份有限公司
  313871000007   宁夏银行总行清算中心
  313656000019   攀枝花市商业银行
  307584007998   平安银行（原深圳发展银行）
  313453001017   齐商银行
  596110000013   企业银行（中国）有限公司
  313452060150   青岛银行
  313851000018   青海银行股份有限公司营业部
  313473200011   日照银行股份有限公司
  313393080005   厦门银行股份有限公司
  322290000011   上海农村商业银行
  310290000013   上海浦东发展银行
  313433076801   上饶银行
  313337009004   绍兴银行股份有限公司营业部
  313305066661   苏州银行股份有限公司
  313345001665   台州银行股份有限公司
  313110000017   天津银行股份有限公司
  313465000010   威海市商业银行
  313458000013   潍坊银行
  313333007331   温州银行股份有限公司
  314305400015   吴江农村商业银行清算中心
  595100000007   新韩银行（中国）有限公司
  313131000016   邢台银行股份有限公司
  309391000011   兴业银行总行
  313456000108   烟台银行股份有限公司
  313228000276   营口银行股份有限公司资金清算中心
  593100000020   友利银行(中国)有限公司
  402731057238   云南省农村信用社联合社
  314305670002   张家港农村商业银行
  313791030003   长安银行股份有限公司
  313551088886   长沙银行股份有限公司
  308584000013   招商银行股份有限公司
  313338707013   浙江稠州商业银行
  313345400010   浙江民泰商业银行
  402331000007   浙江省农村信用社联合社
  313345010019   浙江泰隆商业银行
  316331000018   浙商银行
  313491000232   郑州银行
  102100099996   中国工商银行
  303100000006   中国光大银行
  105100000017   中国建设银行股份有限公司总行
  305100000013   中国民生银行
  103100000026   中国农业银行股份有限公司
  104100000004   中国银行总行
  403100000004   中国邮政储蓄银行有限责任公司
  302100011000   中信银行股份有限公司
  314653000011   重庆农村商业银行股份有限公司
  313653000013   重庆银行
  313585000990   珠海华润银行股份有限公司清算中心

银行机构编码列表
----------------

  **银行编码**   **银行名称**
  -------------- ----------------------
  CMB            招商银行
  ICBC           中国工商银行
  CCB            中国建设银行
  SPDB           上海浦东发展银行
  ABC            中国农业银行
  SDB            深圳发展银行
  CIB            兴业银行
  CEB            中国光大银行
  CMBC           中国民生银行
  CITIC          中信银行
  GDB            广东发展银行
  SZPAB          平安银行
  BOC            中国银行
  COMM           中国交通银行
  PSBC           中国邮政储蓄银行
  BOS            上海银行
  CRB            华润银行
  SDEB           顺德农商
  HXB            华夏银行
  BCCB           北京银行
  CBHB           渤海银行
  EGBANK         恒丰银行
  CZB            浙商银行
  HCCB           杭州银行
  BSB            包商银行
  CQCB           重庆银行
  BOCD           成都银行
  HBCB           哈尔滨银行
  GDNYB          南粤银行
  JSBC           江苏银行
  NBCB           宁波银行
  NJCB           南京银行
  ZZCB           郑州银行
  TCCB           天津银行
  BJRCB          北京农商行
  CCQTGB         重庆三峡银行
  CITYBANK       城市商业银行
  COUNTYBANK     村镇银行
  CSCB           长沙银行
  CYB            集友银行
  CZCB           浙江稠州商业银行
  GNXS           广州市农信社
  GZCB           广州市商业银行
  HKBCHINA       汉口银行
  HKBEA          东亚银行
  HNNXS          湖南农信社
  HSBANK         徽商银行
  RCB            农村商业银行
  RCC            农村信用合作社
  SCB            渣打银行
  SDE            顺德农信社
  SHRCB          上海农村商业银行
  SNXS           深圳农村商业银行
  SXJS           晋城市商业银行
  UCC            城市信用合作社
  URCB           农村合作银行
  WZCB           温州市商业银行
  YDXH           尧都信用合作联社
  ZHNX           珠海市农村信用合作社
  SRCB           上海农商
  MYBANK         网商银行
  UPOP           银联支付

第三方支付机构列表
------------------

  **银行编码**   **银行名称**
  -------------- --------------
  WXPAY          微信支付
  WEIHUIPAY      微汇快捷
  TENPAY         财付通
  99BILL         快钱渠道
  SINAPAY        新浪支付
  LIANLIANPAY    连连支付
  UMPAY          联动优势
  ALIPAY         支付宝

折扣支付方式
------------

该字段为jsonList格式

  ---------------------------------------------------------------------------------------------------------------------------------------
  **字段名称**     **说明**                                                                           **范式**
  ---------------- ---------------------------------------------------------------------------------- -----------------------------------
  discount\_type   折扣类型，取值为：                                                                 Json样式：
                                                                                                      
                   补贴(allowancePay)                                                                 \[{discount\_type:”allowancePay”,
                                                                                                      
                                                                                                      instCode:”DISCOUNTPAY”,
                                                                                                      
                                                                                                      amountInfo:”25.00”}\]

  instCode         机构编码(折扣来源平台)                                                             
                                                                                                      
                   DISCOUNTPAY                                                                        

  amountInfo       根据不同折扣类型取值                                                               
                                                                                                      
                   补贴：补贴金额                                                                     
                                                                                                      
                   其中金额单位为：RMB Yuan。取值范围为\[0.01，100000000.00\]，精确到小数点后两位。   
  ---------------------------------------------------------------------------------------------------------------------------------------

商户通知类型
------------

  **通知类型编码**                    **通知类型说明**
  ----------------------------------- ------------------
  trade\_status\_sync                 交易状态同步
  refund\_status\_sync                退款状态同步
  transfer\_status\_sync              转账状态同步
  deposit\_status\_sync               充值状态同步
  withdrawal\_status\_sync            提现状态同步
  batched\_withdrawal\_status\_sync   批量提现状态同步
  pay\_status\_sync                   支付状态同步

账户类型
--------

  **账户类型**   **账户类型说明**
  -------------- ------------------
  BASIC          基本户
  DEPOSIT        保证金户
  INCOME         收入户
  FINACNE        理财户
  CREDIT         贷记户
  COUPON         红包户
  CUSTOM         自定义账户
  FEE            费用户

> []{#_文件用途 .anchor}

支付方式
--------

  **枚举名称**   **枚举说明**
  -------------- --------------
  BALANCE        余额
  QPAY           快捷
  ONLINE\_BANK   网银
  POS            POS支付

业务产品编码
------------

  **业务类型**   **产品编码**
  -------------- --------------
  充值           10101
  即时到账       20201
  提现           10210
  代付           10240
  会员转账       10310
  担保交易       20202
  冻结解冻       40101
  登帐           30201


