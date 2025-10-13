# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/chat/v1/chat_messages.proto](#api_chat_v1_chat_messages-proto)
    - [Chat](#api-chat-v1-Chat)
    - [CreateDirectChatRequest](#api-chat-v1-CreateDirectChatRequest)
    - [CreateDirectChatResponse](#api-chat-v1-CreateDirectChatResponse)
    - [GetChatRequest](#api-chat-v1-GetChatRequest)
    - [GetChatResponse](#api-chat-v1-GetChatResponse)
    - [ListChatMembersRequest](#api-chat-v1-ListChatMembersRequest)
    - [ListChatMembersResponse](#api-chat-v1-ListChatMembersResponse)
    - [ListMessagesRequest](#api-chat-v1-ListMessagesRequest)
    - [ListMessagesResponse](#api-chat-v1-ListMessagesResponse)
    - [ListUserChatsRequest](#api-chat-v1-ListUserChatsRequest)
    - [ListUserChatsResponse](#api-chat-v1-ListUserChatsResponse)
    - [Message](#api-chat-v1-Message)
    - [SendMessageRequest](#api-chat-v1-SendMessageRequest)
    - [SendMessageResponse](#api-chat-v1-SendMessageResponse)
    - [StreamMessagesRequest](#api-chat-v1-StreamMessagesRequest)
    - [StreamMessagesResponse](#api-chat-v1-StreamMessagesResponse)
  
- [api/chat/v1/chat_service.proto](#api_chat_v1_chat_service-proto)
    - [ChatService](#api-chat-v1-ChatService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_chat_v1_chat_messages-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/chat/v1/chat_messages.proto



<a name="api-chat-v1-Chat"></a>

### Chat
Chat - структура чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |
| participant_ids | [string](#string) | repeated | Id участников |






<a name="api-chat-v1-CreateDirectChatRequest"></a>

### CreateDirectChatRequest
CreateDirectChatRequest - запрос на создание личного чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| participant_id | [string](#string) |  | Id собеседника |






<a name="api-chat-v1-CreateDirectChatResponse"></a>

### CreateDirectChatResponse
CreateDirectChatResponse - ответ на создание чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id созданного чата |






<a name="api-chat-v1-GetChatRequest"></a>

### GetChatRequest
GetChatRequest - запрос получения чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |






<a name="api-chat-v1-GetChatResponse"></a>

### GetChatResponse
GetChatResponse - ответ получения чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat | [Chat](#api-chat-v1-Chat) |  | Данные чата |






<a name="api-chat-v1-ListChatMembersRequest"></a>

### ListChatMembersRequest
ListChatMembersRequest - запрос списка участников чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |






<a name="api-chat-v1-ListChatMembersResponse"></a>

### ListChatMembersResponse
ListChatMembersResponse - ответ списка участников чата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_ids | [string](#string) | repeated | Список Id участников |






<a name="api-chat-v1-ListMessagesRequest"></a>

### ListMessagesRequest
ListMessagesRequest - запрос списка сообщений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |
| limit | [uint32](#uint32) |  | Количество сообщений |
| last_message_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Время отправления последнего сообщения в списке (курсор) |






<a name="api-chat-v1-ListMessagesResponse"></a>

### ListMessagesResponse
ListMessagesResponse - ответ списка сообщений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#api-chat-v1-Message) | repeated | Список сообщений |
| last_message_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Следующее время отправления последнего сообщения в списке (курсор) |






<a name="api-chat-v1-ListUserChatsRequest"></a>

### ListUserChatsRequest
ListUserChatsRequest - запрос списка чатов пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Id пользователя |






<a name="api-chat-v1-ListUserChatsResponse"></a>

### ListUserChatsResponse
ListUserChatsResponse - ответ списка чатов пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chats | [Chat](#api-chat-v1-Chat) | repeated | Список чатов пользователя |






<a name="api-chat-v1-Message"></a>

### Message
Message - структура сообщения


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_id | [string](#string) |  | Id сообщения |
| chat_id | [string](#string) |  | Id чата |
| sender_id | [string](#string) |  | Id отправителя |
| text | [string](#string) |  | Текст сообщения |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Время создания сообщения |






<a name="api-chat-v1-SendMessageRequest"></a>

### SendMessageRequest
SendMessageRequest - запрос на отправку сообщения


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |
| text | [string](#string) |  | Текст сообщения |






<a name="api-chat-v1-SendMessageResponse"></a>

### SendMessageResponse
SendMessageResponse - ответ на отправку сообщения


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [Message](#api-chat-v1-Message) |  | Отправленное сообщение |






<a name="api-chat-v1-StreamMessagesRequest"></a>

### StreamMessagesRequest
StreamMessagesRequest - запрос на поток сообщений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chat_id | [string](#string) |  | Id чата |
| since_message_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Время отправления последнего сообщения |






<a name="api-chat-v1-StreamMessagesResponse"></a>

### StreamMessagesResponse
StreamMessagesResponse - ответ потока сообщений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [Message](#api-chat-v1-Message) |  | Новое сообщение |





 

 

 

 



<a name="api_chat_v1_chat_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/chat/v1/chat_service.proto


 

 

 


<a name="api-chat-v1-ChatService"></a>

### ChatService
ChatService - сервис управления чатами и отправки сообщений

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateDirectChat | [CreateDirectChatRequest](#api-chat-v1-CreateDirectChatRequest) | [CreateDirectChatResponse](#api-chat-v1-CreateDirectChatResponse) | Создать личный чат |
| GetChat | [GetChatRequest](#api-chat-v1-GetChatRequest) | [GetChatResponse](#api-chat-v1-GetChatResponse) | Получить информацию о чате |
| ListUserChats | [ListUserChatsRequest](#api-chat-v1-ListUserChatsRequest) | [ListUserChatsResponse](#api-chat-v1-ListUserChatsResponse) | Список чатов пользователя |
| ListChatMembers | [ListChatMembersRequest](#api-chat-v1-ListChatMembersRequest) | [ListChatMembersResponse](#api-chat-v1-ListChatMembersResponse) | Список участников чата |
| SendMessage | [SendMessageRequest](#api-chat-v1-SendMessageRequest) | [SendMessageResponse](#api-chat-v1-SendMessageResponse) | Отправить сообщение |
| ListMessages | [ListMessagesRequest](#api-chat-v1-ListMessagesRequest) | [ListMessagesResponse](#api-chat-v1-ListMessagesResponse) | История сообщений |
| StreamMessages | [StreamMessagesRequest](#api-chat-v1-StreamMessagesRequest) | [StreamMessagesResponse](#api-chat-v1-StreamMessagesResponse) stream | Стрим новых сообщений |

 



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

