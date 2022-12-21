# Table: ibm_is_subnet

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| available_ipv4_address_count | int | X | √ | The number of IPv4 addresses in this subnet that are not in-use, and have not been reserved by the user or the provider. | 
| network_acl | json | X | √ | The network ACL for this subnet. | 
| routing_table | json | X | √ | The routing table for this subnet. | 
| zone | json | X | √ | The zone this subnet resides in. | 
| account_id | string | X | √ | The account ID of this subnet. | 
| crn | string | X | √ | The CRN for this subnet. | 
| ip_version | string | X | √ | The IP version(s) supported by this subnet. | 
| vpc | json | X | √ | The VPC this subnet is a part of. | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ | The date and time that the subnet was created. | 
| ipv4_cidr_block | cidr | X | √ | The IPv4 range of the subnet, expressed in CIDR format. | 
| resource_group | json | X | √ | The resource group for this subnet. | 
| id | string | X | √ | The unique identifier for this subnet. | 
| name | string | X | √ | The unique user-defined name for this subnet. | 
| href | string | X | √ | The URL for this subnet. | 
| public_gateway | json | X | √ | The public gateway to handle internet bound traffic for this subnet. | 
| status | string | X | √ | The status of this subnet. | 
| total_ipv4_address_count | int | X | √ | The total number of IPv4 addresses in this subnet. | 
| region | string | X | √ | The region of this subnet. | 
| title | string | X | √ |  | 
| akas | json | X | √ |  | 


