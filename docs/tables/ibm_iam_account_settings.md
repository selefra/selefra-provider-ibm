# Table: ibm_iam_account_settings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| session_expiration_in_seconds | string | X | √ | Defines the session expiration in seconds for the account. | 
| restrict_create_service_id | string | X | √ |  | 
| restrict_create_platform_api_key | string | X | √ | Indicates whether creating a platform API key is access controlled, or not. | 
| mfa | string | X | √ | Defines the MFA trait for the account. | 
| session_invalidation_in_seconds | string | X | √ | Defines the period of time in seconds in which a session will be invalidated due  to inactivity. | 
| history | json | X | √ | History of the Account Settings. | 
| account_id | string | X | √ | An unique ID of the account. | 
| allowed_ip_addresses | string | X | √ | The IP addresses and subnets from which IAM tokens can be created for the account. | 
| entity_tag | string | X | √ | Version of the account settings. | 


