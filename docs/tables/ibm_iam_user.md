# Table: ibm_iam_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | An alphanumeric value identifying the user profile. | 
| realm | string | X | √ | The realm of the user. The value is either IBMid or SL. | 
| first_name | string | X | √ | The first name of the user. | 
| last_name | string | X | √ | The last name of the user. | 
| state | string | X | √ | The state of the user. Possible values are PROCESSING, PENDING, ACTIVE, DISABLED_CLASSIC_INFRASTRUCTURE, and VPN_ONLY. | 
| phonenumber | string | X | √ | The phone number of the user. | 
| settings | json | X | √ | User settings. | 
| iam_id | string | X | √ | An alphanumeric value identifying the user's IAM ID. | 
| user_id | string | X | √ | The user ID used for login. | 
| email | string | X | √ | The email of the user. | 
| alt_phonenumber | string | X | √ | The alternative phone number of the user. | 
| photo | string | X | √ | A link to a photo of the user. | 
| account_id | string | X | √ | An alphanumeric value identifying the account ID. | 


