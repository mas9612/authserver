# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pkg/authserver/authserver.proto](#pkg/authserver/authserver.proto)
    - [CreateTokenRequest](#authserver.CreateTokenRequest)
    - [Token](#authserver.Token)
    - [ValidateTokenRequest](#authserver.ValidateTokenRequest)
    - [ValidateTokenResponse](#authserver.ValidateTokenResponse)
  
  
  
    - [Authserver](#authserver.Authserver)
  

- [Scalar Value Types](#scalar-value-types)



<a name="pkg/authserver/authserver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/authserver/authserver.proto
authserver.proto

Define Authserver Service and related messages.


<a name="authserver.CreateTokenRequest"></a>

### CreateTokenRequest
CreateTokenRequest represents the request of CreateToken.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | user is the username which want to be authenticate. |
| password | [string](#string) |  | password is the credential of given user. |
| orig_host | [string](#string) |  | orig_host is the hostname for which JWT token is valid. |






<a name="authserver.Token"></a>

### Token
Token represents the response of CreateToken.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token is the JWT token. |






<a name="authserver.ValidateTokenRequest"></a>

### ValidateTokenRequest
ValidateTokenRequest represents the request of ValidateToken.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token is the JWT token. |






<a name="authserver.ValidateTokenResponse"></a>

### ValidateTokenResponse
ValidateTokenResponse represents the response of ValidateToken.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  | valid represents whether given token is valid or not.

// err_code represents the error type if token is invalid. int32 err_code = 2; |





 

 

 


<a name="authserver.Authserver"></a>

### Authserver
Service to authenticate identity and issue JWT token.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateToken | [CreateTokenRequest](#authserver.CreateTokenRequest) | [Token](#authserver.Token) | CreateToken creates and returns new JWT token for requested identity. |
| ValidateToken | [ValidateTokenRequest](#authserver.ValidateTokenRequest) | [ValidateTokenResponse](#authserver.ValidateTokenResponse) | ValidateToken validates given token and returns its validity. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

