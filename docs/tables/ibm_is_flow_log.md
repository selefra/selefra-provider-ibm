# Table: ibm_is_flow_log

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ | The unique user-defined name for this flow log collector. | 
| active | bool | X | √ | Indicates whether this collector is active. | 
| resource_group | json | X | √ | The resource group for this flow log collector. | 
| vpc | json | X | √ | The VPC this flow log collector is associated with. | 
| region | string | X | √ | The region of this flow log collector. | 
| id | string | X | √ | The unique identifier for this flow log collector. | 
| href | string | X | √ | The URL for this flow log collector. | 
| crn | string | X | √ | The CRN for this flow log collector. | 
| auto_delete | bool | X | √ | If set to `true`, this flow log collector will be automatically deleted when the target is deleted. | 
| akas | json | X | √ |  | 
| lifecycle_state | string | X | √ | The lifecycle state of the flow log collector. | 
| created_at | timestamp | X | √ | The date and time that the flow log collector was created. | 
| storage_bucket | json | X | √ | The Cloud Object Storage bucket where the collected flows are logged. | 
| target | json | X | √ | The target this collector is collecting flow logs for. | 
| account_id | string | X | √ | The account ID of this flow log collector. | 


