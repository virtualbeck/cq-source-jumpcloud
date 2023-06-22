# Table: jumpcloud_systems

This table shows data for Jumpcloud Systems.

The primary key for this table is **_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|os|`utf8`|
|template_name|`utf8`|
|allow_ssh_root_login|`bool`|
|_id (PK)|`utf8`|
|last_contact|`utf8`|
|remote_ip|`utf8`|
|active|`bool`|
|ssh_root_enabled|`bool`|
|amazon_instance_id|`utf8`|
|ssh_pass_enabled|`bool`|
|version|`utf8`|
|agent_version|`utf8`|
|allow_public_key_authentication|`bool`|
|organization|`utf8`|
|created|`utf8`|
|arch|`utf8`|
|system_time_zone|`float64`|
|allow_ssh_password_authentication|`bool`|
|display_name|`utf8`|
|modify_ssh_d_config|`bool`|
|allow_multi_factor_authentication|`bool`|
|hostname|`utf8`|
|connection_history|`list<item: utf8, nullable>`|
|sshd_params|`json`|
|network_interfaces|`json`|
|tags|`list<item: utf8, nullable>`|