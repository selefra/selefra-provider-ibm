# Table: ibm_account

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ | The type of the account. | 
| state | string | X | √ | The current state of the account. | 
| currency_code | string | X | √ | Specifies the currency type. | 
| owner_guid | string | X | √ | An unique Id of the account owner. | 
| owner_unique_id | string | X | √ | An unique identifier of the account owner. | 
| organizations | json | X | √ | A list of organizations the account is associated. | 
| members | json | X | √ | A list of members associated with this account. | 
| name | string | X | √ | Specifies the name of the account. | 
| guid | string | X | √ | An unique ID of the account. | 
| country_code | string | X | √ | Specifies the country code. | 
| customer_id | string | X | √ | The customer ID of the account. | 
| owner_user_id | string | X | √ | The owner user ID used for login. | 


