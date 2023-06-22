# Table: jumpcloud_users

This table shows data for Jumpcloud Users.

The primary key for this table is **_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|_id (PK)|`utf8`|
|username|`utf8`|
|firstname|`utf8`|
|lastname|`utf8`|
|email|`utf8`|
|password|`utf8`|
|password_date|`utf8`|
|activated|`bool`|
|activation_key|`utf8`|
|expired_warned|`bool`|
|password_expired|`bool`|
|password_expiration_date|`timestamp[us, tz=UTC]`|
|pending_provisioning|`bool`|
|sudo|`bool`|
|unix_uid|`utf8`|
|unix_guid|`utf8`|
|enable_managed_uid|`bool`|
|enable_user_portal_multifactor|`bool`|
|totp_enabled|`bool`|
|attributes|`json`|
|tags|`list<item: utf8, nullable>`|
|externally_managed|`bool`|
|external_dn|`utf8`|
|external_source_type|`utf8`|