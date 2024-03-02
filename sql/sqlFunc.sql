CREATE TABLE users (
    user_id CHAR(36) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    roles CHAR(25) NOT NULL,
    hashed_password CHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);




CREATE FUNCTION users_create_user(data JSON)
RETURNS JSON deterministic 
BEGIN
    DECLARE username VARCHAR(255);
    DECLARE hashed_password VARCHAR(255);
    DECLARE roles VARCHAR(255);
    DECLARE address VARCHAR(255);

    SET username = json_extract(data, '$.userName');
    SET hashed_password = json_extract(data, '$.password');
    SET roles = json_extract(data, '$.roles');
   	SET address = json_extract(data, '$.address');

    INSERT INTO users (user_name, hashed_password, roles, created_at,address)
    VALUES (username, hashed_password , roles, NOW(),address);

    RETURN '{"status": "success"}';
END;

select users_create_user('{"userName":"testing user","password":"124142","roles":"admin"}');
select users_create_user('{"userName":"testing user1","password":"124142","roles":"admin"}');
select users_create_user('{"userName":"testing user2","password":"124142","roles":"admin"}');
select users_create_user('{"userName":"testing user3","password":"124142","roles":"admin"}');
select users_create_user('{"userName":"testing user3","password":"124142","roles":"admin"}');
select users_create_user('{"userName":"testing user4","password":"124142","roles":"admin"}');







CREATE FUNCTION users_get_all(data JSON)
RETURNS JSON
DETERMINISTIC
BEGIN
    DECLARE search_query VARCHAR(255);
    DECLARE page INT;
    DECLARE page_size INT;
    DECLARE offset_val INT;
    DECLARE searchonly JSON;
    DECLARE total_records INT;

    SET search_query = JSON_EXTRACT(data, '$.searchQuery');
    SET page = IFNULL (JSON_EXTRACT(data, '$.page'), 1);
    SET page_size = IFNULL(JSON_EXTRACT(data, '$.pageSize'), 10);
    SET offset_val = (page - 1) * page_size;

    SET searchonly = JSON_OBJECT('users', JSON_ARRAY());

    SELECT COUNT(*) INTO total_records
    FROM users
    WHERE user_name LIKE CONCAT('%', CASE WHEN LENGTH(search_query) > 0 THEN search_query ELSE '' END, '%');

    SELECT JSON_OBJECT(
        'users', JSON_ARRAYAGG(
            JSON_OBJECT(
                'userId', user_id,
                'userName', user_name,
                'Address', Address,
                'roles', roles,
                'createdAt', created_at,
                'updatedAt', updated_at
            )
        ),
        'total_records', total_records,
        'page', page
    ) INTO searchonly
    FROM (
        SELECT user_id, user_name, Address, roles, created_at, updated_at
        FROM users
        WHERE user_name LIKE CONCAT('%', CASE WHEN LENGTH(search_query) > 0 THEN search_query ELSE '' END, '%')
        ORDER BY user_name
        LIMIT page_size OFFSET offset_val
    ) AS filtered_users;

    RETURN searchonly;
END;



CREATE FUNCTION users_delete(data JSON)
	RETURNS JSON
	DETERMINISTIC
	BEGIN
		DECLARE uid varchar(255);
    	SET uid = JSON_EXTRACT(data, '$.userId');
		DELETE FROM users WHERE user_id = uid;
	RETURN '{"status": "success"}';
END;
	

CREATE FUNCTION users_update(params JSON)
    RETURNS JSON
    DETERMINISTIC
BEGIN
    DECLARE new_password VARCHAR(255);
    DECLARE uid VARCHAR(255);
    DECLARE user_name VARCHAR(255);
	DECLARE new_address VARCHAR(255);
    SET new_password = JSON_EXTRACT(params, '$.newPassword');
    SET uid = JSON_EXTRACT(params, '$.userId');
    SET user_name = JSON_EXTRACT(params, '$.userName');
	SET new_address = JSON_EXTRACT(params, '$newAdress');

    UPDATE users
    SET 
        user_name = username,
        hashed_password = CASE 
                             WHEN new_password IS NOT NULL AND LENGTH(new_password) != 0 THEN new_password 
                             ELSE hashed_password -- Keep the existing password if new_password is not provided or empty
                         END,
        Address = CASE 
                  WHEN new_address IS NOT NULL AND LENGTH(new_address) != 0 THEN new_address 
                  ELSE address  
                  END
    WHERE 
        user_id = uid;


    RETURN JSON_OBJECT('status', 'User updated successfully');
END;
