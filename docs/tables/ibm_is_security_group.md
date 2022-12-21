# Table: ibm_is_security_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ | The date and time that the security group was created. | 
| crn | string | X | √ | The CRN for this security group. | 
| network_interfaces | json | X | √ | Array of references to network interfaces. | 
| akas | json | X | √ |  | 
| name | string | X | √ | The unique user-defined name for this security group. | 
| rules | json | X | √ | Array of rules for this security group. If no rules exist, all traffic will be denied. | 
| vpc | json | X | √ | The VPC this security group is a part of. | 
| account_id | string | X | √ | The account ID of this security group. | 
| tags | json | X | √ |  | 
| href | string | X | √ | The URL for this security group. | 
| region | string | X | √ | The region of this security group. | 
| title | string | X | √ |  | 
| id | string | X | √ | The unique identifier for this security group. | 
| resource_group | json | X | √ | The resource group for this security group. | 
| targets | json | X | √ | Array of references to targets. | 


