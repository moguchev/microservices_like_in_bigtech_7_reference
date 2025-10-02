# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/social/v1/social_messages.proto](#api_social_v1_social_messages-proto)
    - [AcceptFriendRequestRequest](#api-social-v1-AcceptFriendRequestRequest)
    - [AcceptFriendRequestResponse](#api-social-v1-AcceptFriendRequestResponse)
    - [DeclineFriendRequestRequest](#api-social-v1-DeclineFriendRequestRequest)
    - [DeclineFriendRequestResponse](#api-social-v1-DeclineFriendRequestResponse)
    - [FriendRequest](#api-social-v1-FriendRequest)
    - [ListFriendsRequest](#api-social-v1-ListFriendsRequest)
    - [ListFriendsResponse](#api-social-v1-ListFriendsResponse)
    - [ListRequestsRequest](#api-social-v1-ListRequestsRequest)
    - [ListRequestsResponse](#api-social-v1-ListRequestsResponse)
    - [RemoveFriendRequest](#api-social-v1-RemoveFriendRequest)
    - [RemoveFriendResponse](#api-social-v1-RemoveFriendResponse)
    - [SendFriendRequestRequest](#api-social-v1-SendFriendRequestRequest)
    - [SendFriendRequestResponse](#api-social-v1-SendFriendRequestResponse)
  
    - [FriendRequestStatus](#api-social-v1-FriendRequestStatus)
  
- [api/social/v1/social_service.proto](#api_social_v1_social_service-proto)
    - [SocialService](#api-social-v1-SocialService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_social_v1_social_messages-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/social/v1/social_messages.proto



<a name="api-social-v1-AcceptFriendRequestRequest"></a>

### AcceptFriendRequestRequest
AcceptFriendRequestRequest - запрос на принятие заявки в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Id заявки |






<a name="api-social-v1-AcceptFriendRequestResponse"></a>

### AcceptFriendRequestResponse
AcceptFriendRequestResponse - ответ на принятие заявки


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request | [FriendRequest](#api-social-v1-FriendRequest) |  | Принятая заявка |






<a name="api-social-v1-DeclineFriendRequestRequest"></a>

### DeclineFriendRequestRequest
DeclineFriendRequestRequest - запрос на отклонение заявки в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Id заявки |






<a name="api-social-v1-DeclineFriendRequestResponse"></a>

### DeclineFriendRequestResponse
DeclineFriendRequestResponse - ответ на отклонение заявки


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request | [FriendRequest](#api-social-v1-FriendRequest) |  | Отклонённая заявка |






<a name="api-social-v1-FriendRequest"></a>

### FriendRequest
FriendRequest - структура заявки в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Id заявки |
| status | [FriendRequestStatus](#api-social-v1-FriendRequestStatus) |  | Статус заявки |






<a name="api-social-v1-ListFriendsRequest"></a>

### ListFriendsRequest
ListFriendsRequest - запрос списка друзей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя |
| limit | [int32](#int32) |  | Лимит количества друзей |
| cursor | [string](#string) |  | Курсор для постраничной навигации |






<a name="api-social-v1-ListFriendsResponse"></a>

### ListFriendsResponse
ListFriendsResponse - ответ списка друзей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| friend_user_ids | [string](#string) | repeated | Id друзей |
| next_cursor | [string](#string) |  | Следующий курсор |






<a name="api-social-v1-ListRequestsRequest"></a>

### ListRequestsRequest
ListRequestsRequest - запрос списка заявок в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя, для которого получаем заявки |






<a name="api-social-v1-ListRequestsResponse"></a>

### ListRequestsResponse
ListRequestsResponse - ответ списка заявок в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requests | [FriendRequest](#api-social-v1-FriendRequest) | repeated | Список заявок |






<a name="api-social-v1-RemoveFriendRequest"></a>

### RemoveFriendRequest
RemoveFriendRequest - запрос на удаление из друзей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id друга, которого нужно удалить |






<a name="api-social-v1-RemoveFriendResponse"></a>

### RemoveFriendResponse
RemoveFriendResponse - ответ на удаление из друзей






<a name="api-social-v1-SendFriendRequestRequest"></a>

### SendFriendRequestRequest
SendFriendRequestRequest - запрос на отправку заявки в друзья


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя, которому отправляется заявка |






<a name="api-social-v1-SendFriendRequestResponse"></a>

### SendFriendRequestResponse
SendFriendRequestResponse - ответ SendFriendRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request | [FriendRequest](#api-social-v1-FriendRequest) |  | Заявка в друзья |





 


<a name="api-social-v1-FriendRequestStatus"></a>

### FriendRequestStatus
FriendRequestStatus - статус заявки в друзья

| Name | Number | Description |
| ---- | ------ | ----------- |
| FRIEND_REQUEST_STATUS_UNSPECIFIED | 0 | Не указан |
| FRIEND_REQUEST_STATUS_PENDING | 1 | Ожидает ответа |
| FRIEND_REQUEST_STATUS_ACCEPTED | 2 | Принята |
| FRIEND_REQUEST_STATUS_DECLINED | 3 | Отклонена |


 

 

 



<a name="api_social_v1_social_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/social/v1/social_service.proto


 

 

 


<a name="api-social-v1-SocialService"></a>

### SocialService
SocialService - сервис управления друзьями и заявками

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendFriendRequest | [SendFriendRequestRequest](#api-social-v1-SendFriendRequestRequest) | [SendFriendRequestResponse](#api-social-v1-SendFriendRequestResponse) | Отправить заявку в друзья |
| ListRequests | [ListRequestsRequest](#api-social-v1-ListRequestsRequest) | [ListRequestsResponse](#api-social-v1-ListRequestsResponse) | Список входящих заявок |
| AcceptFriendRequest | [AcceptFriendRequestRequest](#api-social-v1-AcceptFriendRequestRequest) | [AcceptFriendRequestResponse](#api-social-v1-AcceptFriendRequestResponse) | Принять заявку |
| DeclineFriendRequest | [DeclineFriendRequestRequest](#api-social-v1-DeclineFriendRequestRequest) | [DeclineFriendRequestResponse](#api-social-v1-DeclineFriendRequestResponse) | Отклонить заявку |
| RemoveFriend | [RemoveFriendRequest](#api-social-v1-RemoveFriendRequest) | [RemoveFriendResponse](#api-social-v1-RemoveFriendResponse) | Удалить друга |
| ListFriends | [ListFriendsRequest](#api-social-v1-ListFriendsRequest) | [ListFriendsResponse](#api-social-v1-ListFriendsResponse) | Список друзей |

 



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

