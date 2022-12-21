# Table: ibm_is_network_acl

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The unique identifier for this network ACL | 
| vpc | json | X | √ | he VPC this network ACL is a part of. | 
| account_id | string | X | √ | The account ID of this subnet. | 
| akas | json | X | √ |  | 
| tags | json | X | √ |  | 
| href | string | X | √ | The URL for this network ACL. | 
| resource_group | json | X | √ | The resource group for this network ACL. | 
| subnets | json | X | √ | The subnets to which this network ACL is attached. | 
| title | string | X | √ |  | 
| rules | json | X | √ | The ordered rules for this network ACL. If no rules exist, all traffic will be denied. | 
| region | string | X | √ | The region of this subnet. | 
| name | string | X | √ | The user-defined name for this network ACL. | 
| crn | string | X | √ | The CRN for this network ACL. | 
| created_at | timestamp | X | √ | The date and time that the network ACL was created. | 


