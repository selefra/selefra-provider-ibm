# Table: ibm_cos_bucket

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_date | timestamp | X | √ | The date when the bucket was created. | 
| acl | json | X | √ | The access control list (ACL) of a bucket. | 
| lifecycle_rules | json | X | √ | The lifecycle configuration information of the bucket. | 
| name | string | X | √ | Name of the bucket. | 
| sse_kp_enabled | bool | X | √ | Specifies whether the Bucket has Key Protect enabled. | 
| versioning_enabled | bool | X | √ | The versioning state of a bucket. | 
| versioning_mfa_delete | bool | X | √ | The MFA Delete status of the versioning state. | 
| public_access_block_configuration | json | X | √ | The public access block configuration information of the bucket. | 
| retention | json | X | √ | The retention configuration information of the bucket. | 
| website | json | X | √ | The lifecycle configuration information of the bucket. | 
| region | string | X | √ | The region of the bucket. | 
| sse_kp_customer_root_key_crn | string | X | √ | The root key used by Key Protect to encrypt this bucket. This value must be the full CRN of the root key. | 
| title | string | X | √ |  | 


