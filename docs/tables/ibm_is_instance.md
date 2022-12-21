# Table: ibm_is_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| floating_ips | json | X | √ | Floating IPs allow inbound and outbound traffic from the Internet to an instance | 
| profile | json | X | √ | The profile for this virtual server instance. | 
| vpc | json | X | √ | The VPC this virtual server instance resides in. | 
| zone | json | X | √ | The zone this virtual server instance resides in. | 
| title | string | X | √ |  | 
| crn | string | X | √ | The CRN for this virtual server instance. | 
| status | string | X | √ | The status of the virtual server instance. | 
| created_at | timestamp | X | √ | The date and time that the virtual server instance was created. | 
| akas | json | X | √ |  | 
| name | string | X | √ | The user-defined name for this virtual server instance (and default system hostname). | 
| gpu | json | X | √ | The virtual server instance GPU configuration. | 
| vcpu | json | X | √ | The virtual server instance VCPU configuration. | 
| network_interfaces | json | X | √ | A collection of the virtual server instance's network interfaces, including the primary network interface. | 
| primary_network_interface | json | X | √ | Specifies the primary network interface. | 
| resource_group | json | X | √ | The resource group for this instance. | 
| volume_attachments | json | X | √ | A collection of the virtual server instance's volume attachments, including the boot volume attachment. | 
| bandwidth | int | X | √ | The total bandwidth (in megabits per second) shared across the virtual server instance's network interfaces. | 
| disks | json | X | √ | A collection of the instance's disks. | 
| image | json | X | √ | The image the virtual server instance was provisioned from. | 
| boot_volume_attachment | json | X | √ | Specifies the boot volume attachment. | 
| account_id | string | X | √ | The account ID of this instance. | 
| region | string | X | √ | The region of this instance. | 
| tags | json | X | √ |  | 
| id | string | X | √ | The unique identifier for this virtual server instance. | 
| href | string | X | √ | The URL for this virtual server instance. | 
| memory | int | X | √ | The amount of memory, truncated to whole gibibytes. | 


