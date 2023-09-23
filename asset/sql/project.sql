SELECT * FROM projects WHERE id='c00e7304-2dfa-4cfc-b0e6-16d4a359dff4';

SELECT * FROM groups WHERE project_id='c00e7304-2dfa-4cfc-b0e6-16d4a359dff4';

SELECT * FROM secrets
WHERE project_id='c00e7304-2dfa-4cfc-b0e6-16d4a359dff4';

SELECT user_id FROM group_users
WHERE group_id='a80e209b-8f35-4ab9-ae75-8423e97820ec';

SELECT r.name FROM group_roles gr
                       INNER JOIN roles r ON r.name=gr.role
WHERE gr.group_id='a80e209b-8f35-4ab9-ae75-8423e97820ec';