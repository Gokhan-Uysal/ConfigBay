SELECT * FROM projects WHERE id='c254bc39-f4f9-4c4b-a38a-10679d0054d7';

SELECT key, value, version, created_at, updated_at  FROM secrets
WHERE project_id='c254bc39-f4f9-4c4b-a38a-10679d0054d7';

SELECT group_id,user_id FROM group_users
WHERE group_id='a18eab29-b2e2-4d6f-999e-5bf10c483ee9';