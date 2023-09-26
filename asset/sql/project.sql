SELECT * FROM projects WHERE id='c254bc39-f4f9-4c4b-a38a-10679d0054d7';

SELECT group_id FROM project_groups WHERE project_id='c254bc39-f4f9-4c4b-a38a-10679d0054d7';

SELECT key, value, version, created_at, updated_at  FROM secrets
WHERE project_id='c254bc39-f4f9-4c4b-a38a-10679d0054d7';

SELECT group_id,user_id FROM group_users
WHERE group_id='a18eab29-b2e2-4d6f-999e-5bf10c483ee9';

SELECT r.name FROM group_roles gr
INNER JOIN roles r ON r.name=gr.role
WHERE gr.group_id='a18eab29-b2e2-4d6f-999e-5bf10c483ee9';

SELECT r.name FROM project_groups pg
INNER JOIN group_users gu ON gu.group_id=pg.group_id
INNER JOIN group_roles gr ON gr.group_id=pg.group_id
INNER JOIN roles r ON r.name=gr.role
WHERE gu.user_id='d11ca42c-63b3-4e81-bc61-a99067ecedc6' AND pg.group_id='eca46247-e9f7-49a4-8a7e-e5bda8bcfb34';

SELECT user_id FROM group_users WHERE group_id=$1