# Table: ibm_iam_access_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The ID of the IAM access group. | 
| last_modified_at | timestamp | X | √ | Specifies the date and time, the group las modified. | 
| href | string | X | √ | An url to the given group resource. | 
| name | string | X | √ | Name of the access group. | 
| is_federated | bool | X | √ | This is set to true if rules exist for the group. | 
| created_at | timestamp | X | √ | The timestamp the group was created at. | 
| description | string | X | √ | The description of the IAM access group. | 
| created_by_id | string | X | √ | The iam_id of the entity that created the group. | 
| account_id | string | X | √ | ID of the account that this group belongs to. | 


