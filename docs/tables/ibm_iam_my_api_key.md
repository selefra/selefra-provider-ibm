# Table: ibm_iam_my_api_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The description of the API key. | 
| entity_tag | string | X | √ | Version of the API Key details object. | 
| account_id | string | X | √ | ID of the account that this API key authenticates for. | 
| modified_at | timestamp | X | √ | Specifies the date and time, the API key las modified. | 
| name | string | X | √ | Specifies the name of the API key. | 
| id | string | X | √ | Unique identifier of this API Key. | 
| iam_id | string | X | √ | The iam_id that this API key authenticates. | 
| created_at | timestamp | X | √ | Specifies the date and time, the API key is created. | 
| crn | string | X | √ | Cloud Resource Name of the API key. | 
| api_key | string | X | √ | The API key value. This property only contains the API key value for the following cases: create an API key, update a service ID API key that stores the API key value as retrievable, or get a service ID API key that stores the API key value as retrievable. | 
| history | json | X | √ | History of the API key. | 


