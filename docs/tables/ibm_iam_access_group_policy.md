# Table: ibm_iam_access_group_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ | The policy type. | 
| created_at | timestamp | X | √ | The time when the policy was created. | 
| created_by_id | string | X | √ | The iam ID of the entity that created the policy. | 
| last_modified_at | timestamp | X | √ | The timestamp when the policy was last modified. | 
| resources | json | X | √ | The resources associated with a policy. | 
| id | string | X | √ | The ID of the IAM user policy. | 
| group_id | string | X | √ | The ID of the IAM access group. | 
| description | string | X | √ | The description of the IAM access group. | 
| href | string | X | √ | The href link back to the policy. | 
| last_modified_by_id | string | X | √ | The iam ID of the entity that last modified the policy. | 
| subjects | json | X | √ | The subjects associated with a policy. | 
| roles | json | X | √ | A set of role cloud resource names (CRNs) granted by the policy. | 
| account_id | string | X | √ | ID of the account that this policy belongs to. | 


