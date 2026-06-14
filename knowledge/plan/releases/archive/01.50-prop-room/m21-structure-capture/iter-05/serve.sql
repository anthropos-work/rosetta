-- M21 iter-05: register content collections + replicate prod's public-read access chain.
INSERT INTO directus.directus_policies (id, name, icon, description, enforce_tfa, admin_access, app_access)
VALUES ('abf8a154-5b1c-4a46-ac9c-7300570f4f17','$t:public_label','public','$t:public_description',false,false,false)
ON CONFLICT (id) DO NOTHING;

INSERT INTO directus.directus_access (id, role, "user", policy, sort)
VALUES ('91063d51-4951-41ea-8817-d1fb331f45bc',NULL,NULL,'abf8a154-5b1c-4a46-ac9c-7300570f4f17',1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO directus.directus_collections (collection, hidden, singleton, accountability, collapse, versioning)
VALUES ('simulations',false,false,'all','open',false),
       ('skill_paths',false,false,'all','open',false),
       ('roles',false,false,'all','open',false),
       ('sequences',false,false,'all','open',false),
       ('sequences_roles',false,false,'all','open',false)
ON CONFLICT (collection) DO NOTHING;

INSERT INTO directus.directus_permissions (collection, action, permissions, validation, presets, fields, policy)
VALUES ('simulations','read','{"_and":[{"status":{"_eq":"published"}}]}'::json,'{}'::json,NULL,'*','abf8a154-5b1c-4a46-ac9c-7300570f4f17'),
       ('skill_paths','read','{"_and":[{"status":{"_eq":"published"}}]}'::json,'{}'::json,NULL,'*','abf8a154-5b1c-4a46-ac9c-7300570f4f17'),
       ('roles','read','{}'::json,'{}'::json,NULL,'*','abf8a154-5b1c-4a46-ac9c-7300570f4f17'),
       ('sequences','read','{}'::json,'{}'::json,NULL,'*','abf8a154-5b1c-4a46-ac9c-7300570f4f17'),
       ('sequences_roles','read','{}'::json,'{}'::json,NULL,'*','abf8a154-5b1c-4a46-ac9c-7300570f4f17');
