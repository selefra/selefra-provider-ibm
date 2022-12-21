# Table: ibm_certificate_manager_certificate

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| begins_on | timestamp | X | √ | The creation date of the certificate. | 
| id | string | X | √ | The ID of the certificate that is managed in certificate manager. | 
| certificate_manager_instance_id | string | X | √ | The CRN of the certificate manager service instance. | 
| algorithm | string | X | √ | The Algorithm of a certificate. | 
| auto_renew_enabled | bool | X | √ | The automatic renewal status of the certificate. | 
| order_policy_name | string | X | √ | The order policy name of the certificate. | 
| akas | json | X | √ |  | 
| description | string | X | √ | The description of the certificate. | 
| expires_on | timestamp | X | √ | The expiration date of the certificate. | 
| has_previous | bool | X | √ | Indicates whether a certificate has a previous version. | 
| imported | bool | X | √ | Indicates whether a certificate has imported or not. | 
| name | string | X | √ | The display name of the certificate. | 
| issuer | string | X | √ | The issuer of the certificate. | 
| title | string | X | √ |  | 
| issuance_info | json | X | √ | The issuance information of a certificate. | 
| key_algorithm | string | X | √ | An alphanumeric value identifying the account ID. | 
| account_id | string | X | √ | The account ID of this certificate. | 
| region | string | X | √ | The region of this certificate. | 
| status | string | X | √ | The status of a certificate. | 
| serial_number | string | X | √ | The serial number of a certificate. | 
| domains | json | X | √ | An array of valid domains for the issued certificate. The first domain is the primary domain, extra domains are secondary domains. | 
| rotate_keys | bool | X | √ | Rotate keys. | 


