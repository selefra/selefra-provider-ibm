# Table: ibm_resource_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_default | bool | X | √ | Indicates whether this resource group is default of the account or not. | 
| title | string | X | √ |  | 
| quota_id | string | X | √ | An alpha-numeric value identifying the quota ID associated with the resource group. | 
| quota_url | string | X | √ | The URL to access the quota details that associated with the resource group. | 
| akas | json | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ | An alpha-numeric value identifying the resource group. | 
| name | string | X | √ | The human-readable name of the resource group. | 
| crn | string | X | √ | The full CRN (cloud resource name) associated with the resource group. | 
| state | string | X | √ | The state of the resource group. | 
| updated_at | timestamp | X | √ | The date when the resource group was last updated. | 
| resource_linkages | string | X | √ | An array of the resources that linked to the resource group. | 
| created_at | timestamp | X | √ | The date when the resource group was initially created. | 
| payment_methods_url | string | X | √ | The URL to access the payment methods details that associated with the resource group. | 
| teams_url | string | X | √ | The URL to access the team details that associated with the resource group. | 
| account_id | string | X | √ | An alpha-numeric value identifying the account ID. | 


