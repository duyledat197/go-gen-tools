# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [enum.proto](#enum-proto)
    - [ProjectStatus](#enum_pb-ProjectStatus)
    - [UserRole](#enum_pb-UserRole)
  
- [example.proto](#example-proto)
    - [HelloReply](#example-HelloReply)
    - [HelloRequest](#example-HelloRequest)
  
    - [Greeter](#example-Greeter)
  
- [hub.proto](#hub-proto)
    - [CreateHubRequest](#hub-CreateHubRequest)
    - [CreateHubResponse](#hub-CreateHubResponse)
    - [GetHubByIDRequest](#hub-GetHubByIDRequest)
    - [GetHubByIDResponse](#hub-GetHubByIDResponse)
    - [GetListHubRequest](#hub-GetListHubRequest)
    - [GetListHubResponse](#hub-GetListHubResponse)
    - [Hub](#hub-Hub)
    - [UpdateHubRequest](#hub-UpdateHubRequest)
    - [UpdateHubResponse](#hub-UpdateHubResponse)
  
    - [HubService](#hub-HubService)
  
- [search.proto](#search-proto)
    - [SearchTeamHubRequest](#search-SearchTeamHubRequest)
    - [SearchTeamHubResponse](#search-SearchTeamHubResponse)
  
    - [SearchService](#search-SearchService)
  
- [team.proto](#team-proto)
    - [CreateTeamRequest](#team-CreateTeamRequest)
    - [CreateTeamResponse](#team-CreateTeamResponse)
    - [GetListTeamRequest](#team-GetListTeamRequest)
    - [GetListTeamResponse](#team-GetListTeamResponse)
    - [GetTeamByIDRequest](#team-GetTeamByIDRequest)
    - [GetTeamByIDResponse](#team-GetTeamByIDResponse)
    - [Team](#team-Team)
    - [UpdateTeamRequest](#team-UpdateTeamRequest)
    - [UpdateTeamResponse](#team-UpdateTeamResponse)
  
    - [TeamService](#team-TeamService)
  
- [user.proto](#user-proto)
    - [CreateUserRequest](#user-CreateUserRequest)
    - [CreateUserResponse](#user-CreateUserResponse)
    - [GetListUserRequest](#user-GetListUserRequest)
    - [GetListUserResponse](#user-GetListUserResponse)
    - [GetUserByIDRequest](#user-GetUserByIDRequest)
    - [GetUserByIDResponse](#user-GetUserByIDResponse)
    - [UpdateUserRequest](#user-UpdateUserRequest)
    - [UpdateUserResponse](#user-UpdateUserResponse)
    - [User](#user-User)
  
    - [UserService](#user-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="enum-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## enum.proto


 


<a name="enum_pb-ProjectStatus"></a>

### ProjectStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ProjectStatus_NONE | 0 |  |
| ProjectStatus_IN_PROGRESS | 1 |  |
| ProjectStatus_PAUSE | 2 |  |
| ProjectStatus_CLOSE | 3 |  |
| ProjectStatus_DRAFT | 4 |  |



<a name="enum_pb-UserRole"></a>

### UserRole


| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_UNKNOWN | 0 |  |
| SUPER_ADMIN | 1 |  |
| ADMIN | 2 |  |
| SELLER | 3 |  |
| MANAGER | 4 |  |


 

 

 



<a name="example-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## example.proto



<a name="example-HelloReply"></a>

### HelloReply
The response message containing the greetings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="example-HelloRequest"></a>

### HelloRequest
The request message containing the user&#39;s name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |





 

 

 


<a name="example-Greeter"></a>

### Greeter
The greeting service definition.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SayHello | [HelloRequest](#example-HelloRequest) | [HelloReply](#example-HelloReply) | Sends a greeting |

 



<a name="hub-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hub.proto



<a name="hub-CreateHubRequest"></a>

### CreateHubRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hub | [Hub](#hub-Hub) |  |  |






<a name="hub-CreateHubResponse"></a>

### CreateHubResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="hub-GetHubByIDRequest"></a>

### GetHubByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hubID | [string](#string) |  |  |






<a name="hub-GetHubByIDResponse"></a>

### GetHubByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Hub](#hub-Hub) |  |  |






<a name="hub-GetListHubRequest"></a>

### GetListHubRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="hub-GetListHubResponse"></a>

### GetListHubResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Hub](#hub-Hub) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="hub-Hub"></a>

### Hub



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| location_id | [string](#string) |  |  |






<a name="hub-UpdateHubRequest"></a>

### UpdateHubRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hub | [Hub](#hub-Hub) |  |  |






<a name="hub-UpdateHubResponse"></a>

### UpdateHubResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="hub-HubService"></a>

### HubService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetHubByID | [GetHubByIDRequest](#hub-GetHubByIDRequest) | [GetHubByIDResponse](#hub-GetHubByIDResponse) |  |
| CreateHub | [CreateHubRequest](#hub-CreateHubRequest) | [CreateHubResponse](#hub-CreateHubResponse) |  |
| GetList | [GetListHubRequest](#hub-GetListHubRequest) | [GetListHubResponse](#hub-GetListHubResponse) |  |
| UpdateHub | [UpdateHubRequest](#hub-UpdateHubRequest) | [UpdateHubResponse](#hub-UpdateHubResponse) |  |

 



<a name="search-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## search.proto



<a name="search-SearchTeamHubRequest"></a>

### SearchTeamHubRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| q | [string](#string) |  |  |






<a name="search-SearchTeamHubResponse"></a>

### SearchTeamHubResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [team.Team](#team-Team) | repeated |  |
| hubs | [hub.Hub](#hub-Hub) | repeated |  |





 

 

 


<a name="search-SearchService"></a>

### SearchService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTeamHub | [SearchTeamHubRequest](#search-SearchTeamHubRequest) | [SearchTeamHubResponse](#search-SearchTeamHubResponse) |  |

 



<a name="team-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## team.proto



<a name="team-CreateTeamRequest"></a>

### CreateTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [Team](#team-Team) |  |  |






<a name="team-CreateTeamResponse"></a>

### CreateTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="team-GetListTeamRequest"></a>

### GetListTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="team-GetListTeamResponse"></a>

### GetListTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Team](#team-Team) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="team-GetTeamByIDRequest"></a>

### GetTeamByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teamID | [string](#string) |  |  |






<a name="team-GetTeamByIDResponse"></a>

### GetTeamByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Team](#team-Team) |  |  |






<a name="team-Team"></a>

### Team



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| type | [string](#string) |  |  |
| location_id | [string](#string) |  |  |
| hub_id | [string](#string) |  |  |






<a name="team-UpdateTeamRequest"></a>

### UpdateTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [Team](#team-Team) |  |  |






<a name="team-UpdateTeamResponse"></a>

### UpdateTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="team-TeamService"></a>

### TeamService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTeamByID | [GetTeamByIDRequest](#team-GetTeamByIDRequest) | [GetTeamByIDResponse](#team-GetTeamByIDResponse) |  |
| CreateTeam | [CreateTeamRequest](#team-CreateTeamRequest) | [CreateTeamResponse](#team-CreateTeamResponse) |  |
| GetList | [GetListTeamRequest](#team-GetListTeamRequest) | [GetListTeamResponse](#team-GetListTeamResponse) |  |
| UpdateTeam | [UpdateTeamRequest](#team-UpdateTeamRequest) | [UpdateTeamResponse](#team-UpdateTeamResponse) |  |

 



<a name="user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |






<a name="user-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="user-GetListUserRequest"></a>

### GetListUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="user-GetListUserResponse"></a>

### GetListUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [User](#user-User) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="user-GetUserByIDRequest"></a>

### GetUserByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [string](#string) |  |  |






<a name="user-GetUserByIDResponse"></a>

### GetUserByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [User](#user-User) |  |  |






<a name="user-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |






<a name="user-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="user-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| type | [string](#string) |  |  |
| team_id | [string](#string) |  |  |





 

 

 


<a name="user-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserByID | [GetUserByIDRequest](#user-GetUserByIDRequest) | [GetUserByIDResponse](#user-GetUserByIDResponse) |  |
| CreateUser | [CreateUserRequest](#user-CreateUserRequest) | [CreateUserResponse](#user-CreateUserResponse) |  |
| GetList | [GetListUserRequest](#user-GetListUserRequest) | [GetListUserResponse](#user-GetListUserResponse) |  |
| UpdateUser | [UpdateUserRequest](#user-UpdateUserRequest) | [UpdateUserResponse](#user-UpdateUserResponse) |  |

 



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

