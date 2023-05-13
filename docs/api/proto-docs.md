<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [elysium/elysium.proto](#elysium/elysium.proto)
    - [Params](#elysium.Params)
    - [TokenMapping](#elysium.TokenMapping)
    - [TokenMappingChangeProposal](#elysium.TokenMappingChangeProposal)
  
- [elysium/genesis.proto](#elysium/genesis.proto)
    - [GenesisState](#elysium.GenesisState)
  
- [elysium/query.proto](#elysium/query.proto)
    - [ContractByDenomRequest](#elysium.ContractByDenomRequest)
    - [ContractByDenomResponse](#elysium.ContractByDenomResponse)
    - [DenomByContractRequest](#elysium.DenomByContractRequest)
    - [DenomByContractResponse](#elysium.DenomByContractResponse)
    - [ReplayBlockRequest](#elysium.ReplayBlockRequest)
    - [ReplayBlockResponse](#elysium.ReplayBlockResponse)
  
    - [Query](#elysium.Query)
  
- [elysium/tx.proto](#elysium/tx.proto)
    - [MsgConvertVouchers](#elysium.MsgConvertVouchers)
    - [MsgConvertVouchersResponse](#elysium.MsgConvertVouchersResponse)
    - [MsgTransferTokens](#elysium.MsgTransferTokens)
    - [MsgTransferTokensResponse](#elysium.MsgTransferTokensResponse)
    - [MsgUpdateTokenMapping](#elysium.MsgUpdateTokenMapping)
    - [MsgUpdateTokenMappingResponse](#elysium.MsgUpdateTokenMappingResponse)
  
    - [Msg](#elysium.Msg)
  
- [icactl/v1/params.proto](#icactl/v1/params.proto)
    - [Params](#elysium.icactl.v1.Params)
  
- [icactl/v1/genesis.proto](#icactl/v1/genesis.proto)
    - [GenesisState](#elysium.icactl.v1.GenesisState)
  
- [icactl/v1/query.proto](#icactl/v1/query.proto)
    - [QueryInterchainAccountAddressRequest](#elysium.icactl.v1.QueryInterchainAccountAddressRequest)
    - [QueryInterchainAccountAddressResponse](#elysium.icactl.v1.QueryInterchainAccountAddressResponse)
    - [QueryParamsRequest](#elysium.icactl.v1.QueryParamsRequest)
    - [QueryParamsResponse](#elysium.icactl.v1.QueryParamsResponse)
  
    - [Query](#elysium.icactl.v1.Query)
  
- [icactl/v1/tx.proto](#icactl/v1/tx.proto)
    - [MsgRegisterAccount](#elysium.icactl.v1.MsgRegisterAccount)
    - [MsgRegisterAccountResponse](#elysium.icactl.v1.MsgRegisterAccountResponse)
    - [MsgSubmitTx](#elysium.icactl.v1.MsgSubmitTx)
    - [MsgSubmitTxResponse](#elysium.icactl.v1.MsgSubmitTxResponse)
  
    - [Msg](#elysium.icactl.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="elysium/elysium.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## elysium/elysium.proto



<a name="elysium.Params"></a>

### Params
Params defines the parameters for the elysium module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ibc_ely_denom` | [string](#string) |  |  |
| `ibc_timeout` | [uint64](#uint64) |  |  |
| `elysium_admin` | [string](#string) |  | the admin address who can update token mapping |
| `enable_auto_deployment` | [bool](#bool) |  |  |






<a name="elysium.TokenMapping"></a>

### TokenMapping
TokenMapping defines a mapping between native denom and contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `contract` | [string](#string) |  |  |






<a name="elysium.TokenMappingChangeProposal"></a>

### TokenMappingChangeProposal
TokenMappingChangeProposal defines a proposal to change one token mapping.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `contract` | [string](#string) |  |  |
| `symbol` | [string](#string) |  | only when updating elysium (source) tokens |
| `decimal` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="elysium/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## elysium/genesis.proto



<a name="elysium.GenesisState"></a>

### GenesisState
GenesisState defines the elysium module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#elysium.Params) |  | params defines all the paramaters of the module. |
| `external_contracts` | [TokenMapping](#elysium.TokenMapping) | repeated |  |
| `auto_contracts` | [TokenMapping](#elysium.TokenMapping) | repeated | this line is used by starport scaffolding # genesis/proto/state this line is used by starport scaffolding # ibc/genesis/proto |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="elysium/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## elysium/query.proto



<a name="elysium.ContractByDenomRequest"></a>

### ContractByDenomRequest
ContractByDenomRequest is the request type of ContractByDenom call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="elysium.ContractByDenomResponse"></a>

### ContractByDenomResponse
ContractByDenomRequest is the response type of ContractByDenom call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  |  |
| `auto_contract` | [string](#string) |  |  |






<a name="elysium.DenomByContractRequest"></a>

### DenomByContractRequest
DenomByContractRequest is the request type of DenomByContract call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  |  |






<a name="elysium.DenomByContractResponse"></a>

### DenomByContractResponse
DenomByContractResponse is the response type of DenomByContract call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="elysium.ReplayBlockRequest"></a>

### ReplayBlockRequest
ReplayBlockRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msgs` | [ethermint.evm.v1.MsgEthereumTx](#ethermint.evm.v1.MsgEthereumTx) | repeated | the eth messages in the block |
| `block_number` | [int64](#int64) |  |  |
| `block_hash` | [string](#string) |  |  |
| `block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="elysium.ReplayBlockResponse"></a>

### ReplayBlockResponse
ReplayBlockResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `responses` | [ethermint.evm.v1.MsgEthereumTxResponse](#ethermint.evm.v1.MsgEthereumTxResponse) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="elysium.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ContractByDenom` | [ContractByDenomRequest](#elysium.ContractByDenomRequest) | [ContractByDenomResponse](#elysium.ContractByDenomResponse) | ContractByDenom queries contract addresses by native denom | GET|/elysium/v1/contract_by_denom/{denom}|
| `DenomByContract` | [DenomByContractRequest](#elysium.DenomByContractRequest) | [DenomByContractResponse](#elysium.DenomByContractResponse) | DenomByContract queries native denom by contract address | GET|/elysium/v1/denom_by_contract/{contract}|
| `ReplayBlock` | [ReplayBlockRequest](#elysium.ReplayBlockRequest) | [ReplayBlockResponse](#elysium.ReplayBlockResponse) | ReplayBlock replay the eth messages in the block to recover the results of false-failed txs. | |

 <!-- end services -->



<a name="elysium/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## elysium/tx.proto



<a name="elysium.MsgConvertVouchers"></a>

### MsgConvertVouchers
MsgConvertVouchers represents a message to convert ibc voucher coins to elysium evm coins.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="elysium.MsgConvertVouchersResponse"></a>

### MsgConvertVouchersResponse
MsgConvertVouchersResponse defines the ConvertVouchers response type.






<a name="elysium.MsgTransferTokens"></a>

### MsgTransferTokens
MsgTransferTokens represents a message to transfer elysium evm coins through ibc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="elysium.MsgTransferTokensResponse"></a>

### MsgTransferTokensResponse
MsgTransferTokensResponse defines the TransferTokens response type.






<a name="elysium.MsgUpdateTokenMapping"></a>

### MsgUpdateTokenMapping
MsgUpdateTokenMapping defines the request type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `contract` | [string](#string) |  |  |
| `symbol` | [string](#string) |  | only when updating elysium (source) tokens |
| `decimal` | [uint32](#uint32) |  |  |






<a name="elysium.MsgUpdateTokenMappingResponse"></a>

### MsgUpdateTokenMappingResponse
MsgUpdateTokenMappingResponse defines the response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="elysium.Msg"></a>

### Msg
Msg defines the Elysium Msg service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConvertVouchers` | [MsgConvertVouchers](#elysium.MsgConvertVouchers) | [MsgConvertVouchersResponse](#elysium.MsgConvertVouchersResponse) | ConvertVouchers defines a method for converting ibc voucher to elysium evm coins. | |
| `TransferTokens` | [MsgTransferTokens](#elysium.MsgTransferTokens) | [MsgTransferTokensResponse](#elysium.MsgTransferTokensResponse) | TransferTokens defines a method to transfer elysium evm coins to another chain through IBC | |
| `UpdateTokenMapping` | [MsgUpdateTokenMapping](#elysium.MsgUpdateTokenMapping) | [MsgUpdateTokenMappingResponse](#elysium.MsgUpdateTokenMappingResponse) | UpdateTokenMapping defines a method to update token mapping | |

 <!-- end services -->



<a name="icactl/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## icactl/v1/params.proto



<a name="elysium.icactl.v1.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minTimeoutDuration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | minTimeoutDuration defines the minimum value of packet timeout when submitting transactions to host chain on behalf of interchain account |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="icactl/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## icactl/v1/genesis.proto



<a name="elysium.icactl.v1.GenesisState"></a>

### GenesisState
GenesisState defines the icactl module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#elysium.icactl.v1.Params) |  | params defines the genesis parameters |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="icactl/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## icactl/v1/query.proto



<a name="elysium.icactl.v1.QueryInterchainAccountAddressRequest"></a>

### QueryInterchainAccountAddressRequest
QueryInterchainAccountAddressRequest defines the request for the InterchainAccountAddress query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `connectionId` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="elysium.icactl.v1.QueryInterchainAccountAddressResponse"></a>

### QueryInterchainAccountAddressResponse
QueryInterchainAccountAddressResponse defines the response for the InterchainAccountAddress query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interchainAccountAddress` | [string](#string) |  |  |






<a name="elysium.icactl.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="elysium.icactl.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#elysium.icactl.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="elysium.icactl.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#elysium.icactl.v1.QueryParamsRequest) | [QueryParamsResponse](#elysium.icactl.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/elysium/icactl/v1/params|
| `InterchainAccountAddress` | [QueryInterchainAccountAddressRequest](#elysium.icactl.v1.QueryInterchainAccountAddressRequest) | [QueryInterchainAccountAddressResponse](#elysium.icactl.v1.QueryInterchainAccountAddressResponse) | InterchainAccountAddress queries the interchain account address for given `connectionId` and `owner` | GET|/elysium/icactl/v1/interchain_account_address/{connectionId}/{owner}|

 <!-- end services -->



<a name="icactl/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## icactl/v1/tx.proto



<a name="elysium.icactl.v1.MsgRegisterAccount"></a>

### MsgRegisterAccount
MsgRegisterAccount defines the request message for MsgRegisterAccount


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner represents the owner of the interchain account |
| `connectionId` | [string](#string) |  | connectionId represents the IBC `connectionId` of the host chain |






<a name="elysium.icactl.v1.MsgRegisterAccountResponse"></a>

### MsgRegisterAccountResponse
MsgRegisterAccountResponse defines the response message for MsgRegisterAccount






<a name="elysium.icactl.v1.MsgSubmitTx"></a>

### MsgSubmitTx
MsgSubmitTx defines the request message for MsgSubmitTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner represents the owner of the interchain account |
| `connectionId` | [string](#string) |  | connectionId represents the IBC `connectionId` of the host chain |
| `msgs` | [google.protobuf.Any](#google.protobuf.Any) | repeated | msgs represents the transactions to be submitted to the host chain |
| `timeoutDuration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | timeoutDuration represents the timeout duration for the IBC packet from last block |






<a name="elysium.icactl.v1.MsgSubmitTxResponse"></a>

### MsgSubmitTxResponse
MsgSubmitTxResponse defines the response message for MsgSubmitTx





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="elysium.icactl.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RegisterAccount` | [MsgRegisterAccount](#elysium.icactl.v1.MsgRegisterAccount) | [MsgRegisterAccountResponse](#elysium.icactl.v1.MsgRegisterAccountResponse) | RegisterAccount registers an interchain account on host chain with given `connectionId` | |
| `SubmitTx` | [MsgSubmitTx](#elysium.icactl.v1.MsgSubmitTx) | [MsgSubmitTxResponse](#elysium.icactl.v1.MsgSubmitTxResponse) | SubmitTx submits a transaction to the host chain on behalf of interchain account | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

