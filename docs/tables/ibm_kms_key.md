# Table: ibm_kms_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_date | timestamp | X | √ | The timestamp when the key material was created. | 
| description | string | X | √ | A text field used to provide a more detailed description of the key. | 
| payload | string | X | √ | Specifies the key payload. | 
| state | string | X | √ | The key state based on NIST SP 800-57. States are integers and correspond to the Pre-activation = 0, Active = 1, Suspended = 2, Deactivated = 3, and Destroyed = 5 values. | 
| extractable | bool | X | √ | Indicates whether the key material can leave the service, or not. | 
| key_ring_id | string | X | √ | An ID that identifies the key ring. | 
| aliases | json | X | √ | A list of key aliases. | 
| crn | string | X | √ | The Cloud Resource Name (CRN) that uniquely identifies your cloud resources. | 
| imported | bool | X | √ | Indicates whether the key was originally imported or generated in Key Protect. | 
| encrypted_nonce | string | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ | An unique identifier of the key. | 
| dual_auth_delete | json | X | √ | Metadata that indicates the status of a dual authorization policy on the key. | 
| algorithm_type | string | X | √ | Specifies the key algorithm. | 
| type | string | X | √ | Specifies the MIME type that represents the key resource. | 
| deleted | bool | X | √ | Indicates whether the key has been deleted, or not. | 
| expiration | timestamp | X | √ | The date the key material will expire. | 
| last_update_date | timestamp | X | √ | The date when the key metadata was last modified. | 
| akas | json | X | √ |  | 
| name | string | X | √ | A human-readable name assigned to your key for convenience. | 
| deletion_date | timestamp | X | √ | The date the key material was destroyed. | 
| key_version | json | X | √ | Properties associated with a specific key version. | 
| rotation_policy | json | X | √ | Key rotation policy. | 
| account_id | string | X | √ | The account ID of this key. | 
| region | string | X | √ | The region of this key. | 
| title | string | X | √ |  | 
| instance_id | string | X | √ | The key protect instance GUID. | 
| deleted_by | string | X | √ | The unique identifier for the resource that deleted the key. | 
| encryption_algorithm | string | X | √ |  | 
| last_rotate_date | timestamp | X | √ | The date when the key was last rotated. | 
| created_by | string | X | √ | The unique identifier for the resource that created the key. | 


