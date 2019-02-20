# Cosmos SDK Tutorial `The nameservice app.`

## Application Goals

* 요구사항
  1. 사용자(user)가 `name`을 구매하고, 각 `name`에 값을 설정할 수 있음
  2. `name`의 소유자가 현재 가장 높은 입찰자(bidder)가 됨

블록체인 앱은 a replicated deterministic state machine임.  
(deterministic: 늘 동일한 결과가 나옴을 보장)  

개발자는 가장 먼저 state machine을 정의해야 함  
(state가 무엇인지, 시작 상태와 상태변화 메시지는 어떻게 생겼는가)  

Tendermint가 네트워크를 통한 복제를 처리해 줄 것임  

> Tendermint 란?  
> 블록체인의 3 layers 중 network layer와 consensus layer를 처리해 줌  
> (Tendermint is responsible for propagating and ordering transaction bytes.)  
> 
> Tendermint Core는 BFT알고리즘을 사용하여 합의를 이끌어냄  

modular framework인 Cosmos SDK를 이용해 state machine을 구성할 수 있음  
각 모듈은 각자의 message/transaction processoor로 이루어짐  

* `The nameservice app.`을 구성하는 3개의 모듈
  - auth: account, fee 정의
  - bank: token, token balance를 생성하고 처리
  - nameservice: 이 앱의 핵심 로직을 가지는 모듈

> 왜 validator set change를 다루는 모듈이 없는가?  
> Tendermint가 블록을 추가하는 합의에 이르기 위해 validator set에 의존하므로.  
> validator set change를 다루는 모듈이 없다면,   
> genesis file(`genesis.json`)에 정의된 validator set이 변화없이 동일하게 유지됨.  
> 이 앱에서는 위와 같이 처리함.  
> 만약 validator set을 변화하고 싶다면 SDK의 `staking module`을 사용할 수 있음.  


* State
  - SDK에서 모든 정보는 `multistore`에 저장되며, key/value 형태로 저장됨(KVStores in the Cosmos SDK)  
  - 이 예제에서는 `multistore`에 아래의 3개를 생성할 것임.  
    - nameStore: `name`과 `value`의 매핑을 저장
    - ownerStore: `name`과 `owner`의 매핑 저장
    - priceStore: `name`과 `price`의 매핑 저장  

* Messages
  - Message는 tx를 포함하며, 상태(state)를 변화시킴.  
  - 각 모듈은 각자의 message를 정의함.  
  - 이 예제에서는 아래 2개의 message를 구현할 것임.  
    - `MsgSetName`: name owner가 nameStore의 name에 대해 value를 설정 
    - `MsgBuyName`: account가 name을 구매하고 ownerStore에서 owner가 되는 것을 처리

  - tx가 Tendermint node에 도착할 때, ABCI를 통해 message가 전달되고 해석됨.  
    - 그 후 모듈로 라우팅되어 `Handler`에 정의된 로직에 의해 처리됨.  
    - 만약 state가 업데이트되면 `Handler`는 `Keeper`를 호출하여 업데이트를 처리함.  
