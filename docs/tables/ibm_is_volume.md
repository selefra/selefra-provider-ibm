# Table: ibm_is_volume

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The region of this volume. | 
| name | string | X | √ | The user-defined name for this volume. | 
| encryption_key | string | X | √ | A reference to the root key used to wrap the data encryption key for the volume. This property will be present for volumes with an `encryption` type of `user_managed`. | 
| profile | json | X | √ | The profile for this volume. | 
| resource_group | json | X | √ | The resource group for this volume. | 
| zone | json | X | √ | The zone this volume resides in. | 
| akas | json | X | √ |  | 
| crn | string | X | √ | The CRN for this volume. | 
| capacity | int | X | √ | The capacity of the volume in gigabytes. | 
| iops | int | X | √ | The bandwidth for the volume. | 
| account_id | string | X | √ | The account ID of this volume. | 
| title | string | X | √ |  | 
| status_reasons | json | X | √ | The enumerated reason code values for this property will expand in the future. | 
| volume_attachments | json | X | √ | The collection of volume attachments attaching instances to the volume.. | 
| tags | json | X | √ |  | 
| id | string | X | √ | The unique identifier for this volume. | 
| status | string | X | √ | The status of the volume. | 
| created_at | timestamp | X | √ | The date and time that the volume was created. | 
| encryption | string | X | √ | The type of encryption used on the volume. | 
| href | string | X | √ | The URL for this volume. | 


