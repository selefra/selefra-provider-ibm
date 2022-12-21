# Table: ibm_is_vpc

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ | The date and time that the VPC was created. | 
| title | string | X | √ |  | 
| classic_access | bool | X | √ | Indicates whether this VPC is connected to Classic Infrastructure. | 
| crn | string | X | √ | The CRN for this VPC. | 
| status | string | X | √ | The status of this VPC. | 
| akas | json | X | √ |  | 
| default_network_acl | json | X | √ | The default network ACL to use for subnets created in this VPC. | 
| default_routing_table | json | X | √ | The default routing table to use for subnets created in this VPC. | 
| default_security_group | json | X | √ | The default security group to use for network interfaces created in this VPC. | 
| href | string | X | √ | The URL for this VPC. | 
| resource_group | json | X | √ | The resource group for this VPC. | 
| tags | json | X | √ |  | 
| id | string | X | √ | The unique identifier for this VPC. | 
| cse_source_ips | json | X | √ | Array of CSE source IP addresses for the VPC. The VPC will have one CSE source IP address per zone. | 
| account_id | string | X | √ | The account ID of this VPC. | 
| region | string | X | √ | The region of this VPC. | 
| name | string | X | √ | The unique user-defined name for this VPC. | 
| address_prefixes | json | X | √ | Array of all address pool prefixes for this VPC. | 


