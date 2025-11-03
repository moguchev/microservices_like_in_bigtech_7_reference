# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/users/v1/users_messages.proto](#api_users_v1_users_messages-proto)
    - [CreateProfileRequest](#api-users-v1-CreateProfileRequest)
    - [CreateProfileResponse](#api-users-v1-CreateProfileResponse)
    - [GetProfileByIDRequest](#api-users-v1-GetProfileByIDRequest)
    - [GetProfileByIDResponse](#api-users-v1-GetProfileByIDResponse)
    - [GetProfileByNicknameRequest](#api-users-v1-GetProfileByNicknameRequest)
    - [GetProfileByNicknameResponse](#api-users-v1-GetProfileByNicknameResponse)
    - [SearchByNicknameRequest](#api-users-v1-SearchByNicknameRequest)
    - [SearchByNicknameResponse](#api-users-v1-SearchByNicknameResponse)
    - [UpdateProfileRequest](#api-users-v1-UpdateProfileRequest)
    - [UpdateProfileRequest.UpdateProfileFields](#api-users-v1-UpdateProfileRequest-UpdateProfileFields)
    - [UpdateProfileResponse](#api-users-v1-UpdateProfileResponse)
    - [UserProfile](#api-users-v1-UserProfile)
  
- [api/users/v1/users_service.proto](#api_users_v1_users_service-proto)
    - [UserService](#api-users-v1-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_users_v1_users_messages-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/users/v1/users_messages.proto



<a name="api-users-v1-CreateProfileRequest"></a>

### CreateProfileRequest
CreateProfileRequest - запрос на создание профиля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя |
| nickname | [string](#string) |  | Никнейм |
| bio | [string](#string) | optional | Биография |
| avatar_url | [string](#string) | optional | URL аватара |
| email | [string](#string) |  | Почта |






<a name="api-users-v1-CreateProfileResponse"></a>

### CreateProfileResponse
CreateProfileResponse - ответ на создание профиля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-users-v1-UserProfile) |  | Созданный профиль |






<a name="api-users-v1-GetProfileByIDRequest"></a>

### GetProfileByIDRequest
GetProfileByIDRequest - запрос на получение профиля по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id пользователя |






<a name="api-users-v1-GetProfileByIDResponse"></a>

### GetProfileByIDResponse
GetProfileByIDResponse - ответ на получение профиля по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-users-v1-UserProfile) |  | Профиль пользователя |






<a name="api-users-v1-GetProfileByNicknameRequest"></a>

### GetProfileByNicknameRequest
GetProfileByNicknameRequest - запрос на получение профиля по никнейму


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  | Никнейм |






<a name="api-users-v1-GetProfileByNicknameResponse"></a>

### GetProfileByNicknameResponse
GetProfileByNicknameResponse - ответ на получение профиля по никнейму


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-users-v1-UserProfile) |  | Профиль пользователя |






<a name="api-users-v1-SearchByNicknameRequest"></a>

### SearchByNicknameRequest
SearchByNicknameRequest - запрос поиска пользователей по никнейму


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  | Поисковый запрос |
| limit | [uint32](#uint32) |  | Лимит количества результатов |






<a name="api-users-v1-SearchByNicknameResponse"></a>

### SearchByNicknameResponse
SearchByNicknameResponse - ответ поиска пользователей по никнейму


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| results | [UserProfile](#api-users-v1-UserProfile) | repeated | Список найденных профилей |






<a name="api-users-v1-UpdateProfileRequest"></a>

### UpdateProfileRequest
UpdateProfileRequest - запрос на обновление профиля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя |
| update_fields | [UpdateProfileRequest.UpdateProfileFields](#api-users-v1-UpdateProfileRequest-UpdateProfileFields) |  | Обновляемые поля |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | FieldMask

См. https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/patch_feature/#patch-feature |






<a name="api-users-v1-UpdateProfileRequest-UpdateProfileFields"></a>

### UpdateProfileRequest.UpdateProfileFields
Обновляемые поля профиля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) | optional | Никнейм |
| bio | [string](#string) | optional | Биография |
| avatar_url | [string](#string) | optional | URL аватара |






<a name="api-users-v1-UpdateProfileResponse"></a>

### UpdateProfileResponse
UpdateProfileResponse - ответ на обновление профиля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-users-v1-UserProfile) |  | Обновлённый профиль |






<a name="api-users-v1-UserProfile"></a>

### UserProfile
UserProfile - структура профиля пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя |
| nickname | [string](#string) |  | Никнейм |
| bio | [string](#string) |  | Биография |
| avatar_url | [string](#string) |  | URL аватара |





 

 

 

 



<a name="api_users_v1_users_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/users/v1/users_service.proto


 

 

 


<a name="api-users-v1-UserService"></a>

### UserService
UserService - управление профилями пользователей

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProfile | [CreateProfileRequest](#api-users-v1-CreateProfileRequest) | [CreateProfileResponse](#api-users-v1-CreateProfileResponse) | Создать профиль |
| UpdateProfile | [UpdateProfileRequest](#api-users-v1-UpdateProfileRequest) | [UpdateProfileResponse](#api-users-v1-UpdateProfileResponse) | Обновить профиль |
| GetProfileByID | [GetProfileByIDRequest](#api-users-v1-GetProfileByIDRequest) | [GetProfileByIDResponse](#api-users-v1-GetProfileByIDResponse) | Получить профиль по ID |
| GetProfileByNickname | [GetProfileByNicknameRequest](#api-users-v1-GetProfileByNicknameRequest) | [GetProfileByNicknameResponse](#api-users-v1-GetProfileByNicknameResponse) | Получить профиль по никнейму |
| SearchByNickname | [SearchByNicknameRequest](#api-users-v1-SearchByNicknameRequest) | [SearchByNicknameResponse](#api-users-v1-SearchByNicknameResponse) | Поиск пользователей по никнейму |

 



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

