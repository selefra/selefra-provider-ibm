# Table: ibm_cis_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The zone name. | 
| status | string | X | √ | The zone status. | 
| global_load_balancer | json | X | √ | The global load balancer of the zone. | 
| id | string | X | √ | The zone id. | 
| account_id | string | X | √ | An unique ID of the account. | 
| original_name_servers | json | X | √ | The original name servers of the zone. | 
| name_servers | json | X | √ | The name servers of the zone. | 
| original_dnshost | string | X | √ | The original DNS host of the zone. | 
| paused | bool | X | √ | Whether the zone is in paused state. | 
| title | string | X | √ |  | 
| akas | json | X | √ |  | 
| created_on | timestamp | X | √ | The date and time that the zone was created. | 
| modified_on | timestamp | X | √ | The date and time that the zone was updated. | 
| minimum_tls_version | string | X | √ | The tls version of the zone. | 
| original_registrar | string | X | √ | The original registrar of the zone. | 
| web_application_firewall | string | X | √ | The web application firewall state. | 
| dns_records | json | X | √ | DNS records for the domain. | 


