# Table: ibm_iam_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| actions | json | X | √ | The actions of the role. | 
| id | string | X | √ | The role ID. | 
| crn | string | X | √ | The Cloud Resource Name (CRN) that uniquely identifies your cloud resources. | 
| created_by_id | string | X | √ | The IAM ID of the entity that created the role. | 
| last_modified_at | timestamp | X | √ | The timestamp when the role was last modified. | 
| account_id | string | X | √ | An alphanumeric value identifying the account ID. | 
| display_name | string | X | √ | The display name of the role that is shown in the console. | 
| description | string | X | √ | The description of the role. | 
| created_at | timestamp | X | √ | The timestamp when the role was created. | 
| last_modified_by_id | string | X | √ | The IAM ID of the entity that last modified the policy. | 
| name | string | X | √ | The name of the role that is used in the CRN. | 
| service_name | string | X | √ | The service name. | 


