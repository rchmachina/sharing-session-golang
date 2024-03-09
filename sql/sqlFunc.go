package sql

var (
	CreateTable = `
        
        CREATE TABLE users (
            user_id CHAR(36) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
            user_name VARCHAR(255) NOT NULL unique,
            address VARCHAR(255),
            roles CHAR(25) NOT NULL,
            hashed_password CHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );`
	UpdateUserFunction = `
        CREATE FUNCTION users_update(data JSON)
        RETURNS JSON
        DETERMINISTIC
        BEGIN
            DECLARE new_password VARCHAR(255);
            DECLARE id_user VARCHAR(255);
            DECLARE new_user_name VARCHAR(255);
            DECLARE new_address VARCHAR(255);
            DECLARE user_exists INT;

            SET new_password = JSON_UNQUOTE(JSON_EXTRACT(data, '$.newPassword'));
            SET id_user = JSON_UNQUOTE(JSON_EXTRACT(data, '$.userId'));
            SET new_user_name = JSON_UNQUOTE(JSON_EXTRACT(data, '$.userName'));
            SET new_address = JSON_UNQUOTE(JSON_EXTRACT(data, '$.newAddress')); -- Corrected syntax and variable name

            SELECT COUNT(*) INTO user_exists FROM users WHERE user_id = id_user;
            
            IF user_exists = 1 THEN 
                UPDATE users
                SET 
                    hashed_password = CASE 
                        WHEN new_password IS NOT NULL AND LENGTH(new_password) != 0 THEN new_password 
                        ELSE hashed_password 
                        END,
                    address = CASE 
                                WHEN new_address IS NOT NULL AND LENGTH(new_address) != 0 THEN new_address 
                                ELSE address  
                            END,
                    user_name = CASE 
                                WHEN new_user_name IS NOT NULL AND LENGTH(new_user_name) != 0 THEN new_user_name 
                                ELSE user_name
                            end,
                    updated_at = NOW()
                WHERE 
                    user_id = id_user;
                RETURN JSON_OBJECT('status', 'success', 'message', CONCAT(id_user," is updated")); -- Corrected variable name
            ELSE 
                RETURN JSON_OBJECT('status', 'failed', 'message', 'User not found');
            END IF;
        END;
        `
	DeleteUserFunction = `CREATE FUNCTION users_delete(data JSON)
        RETURNS JSON
        DETERMINISTIC
        BEGIN
            DECLARE uid VARCHAR(255);
            DECLARE user_exists INT;
            SET uid = JSON_EXTRACT(data, '$.userId');
            
            SELECT COUNT(*) INTO user_exists FROM users WHERE user_id = uid;
            
            IF user_exists > 0 THEN
                DELETE FROM users WHERE user_id = uid;
                RETURN JSON_OBJECT('status', 'success', 'message', CONCAT('Deleted user id ', uid));
            ELSE
                RETURN '{"status": "failed", "message": "User not found"}';
            END IF;
        END;`
	IsUserExistFunction = `CREATE FUNCTION users_login(data JSON)
        RETURNS JSON
        deterministic
        BEGIN
            DECLARE user_data JSON;
            DECLARE username VARCHAR(50);
            SET username = JSON_UNQUOTE(JSON_EXTRACT(data, '$.userName'));
            
            SELECT JSON_OBJECT(
                'userName', user_name,
                'hashedPassword', hashed_password,
                'roles', roles,
                'UserId', user_id,
                'address', address
            ) INTO user_data
            FROM (
                SELECT user_id, user_name, address, roles, hashed_password
                FROM users
                WHERE user_name = username
                limit 1
            ) AS filtered_users;
        
            RETURN COALESCE(user_data, JSON_OBJECT());
        END;
        `
	CreateUserFunction = `CREATE FUNCTION users_create_user(data JSON)
    RETURNS JSON deterministic 
    BEGIN
        DECLARE username VARCHAR(255);
        DECLARE hashed_password VARCHAR(255);
        DECLARE roles VARCHAR(255);
        DECLARE address VARCHAR(255);
    
        SET username = JSON_UNQUOTE(json_extract(data, '$.userName'));
        SET hashed_password = JSON_UNQUOTE(json_extract(data, '$.password'));
        SET roles = JSON_UNQUOTE(json_extract(data, '$.roles'));
           SET address = JSON_UNQUOTE(json_extract(data, '$.address'));
    
        INSERT INTO users (user_name, hashed_password, roles, created_at,address)
        VALUES (username, hashed_password , roles, NOW(),address);
    
        RETURN '{"status": "success"}';
    END;`
	ReadAllUsersFunction = `CREATE FUNCTION users_get_all(data JSON) RETURNS json 
    DETERMINISTIC 
    BEGIN 
    DECLARE
        search_query VARCHAR(255);
    DECLARE
        page INT;
    DECLARE
        page_size INT;
    DECLARE
        offset_val INT;
    DECLARE
        get_user_data JSON;
    DECLARE
        total_records INT;
        -- Extract parameters from JSON
        SET
            search_query = JSON_UNQUOTE(
                JSON_EXTRACT(data, '$.searchQuery')
            );
        SET page = JSON_UNQUOTE(JSON_EXTRACT(data, '$.page'));
        SET
            page_size = JSON_UNQUOTE(
                JSON_EXTRACT(data, '$.pageSize')
            );
        SET offset_val = (page - 1) * page_size;
        SELECT COUNT(*) AS total_records FROM users INTO total_records;
        if page = 0 and page_size = 0 THEN 
    
            SELECT JSON_OBJECT(
                    'getUserData', JSON_ARRAYAGG(
                        JSON_OBJECT(
                            'userId', user_id, 'userName', user_name, 'Address', Address, 'roles', roles, 'createdAt', created_at, 'updatedAt', updated_at
                        )
                    ), 'totalRecords', total_records
                ) INTO get_user_data
            FROM (
                    SELECT
                        user_id, user_name, Address, roles, created_at, updated_at
                    FROM users
                    WHERE
                        user_name LIKE CONCAT('%', search_query,'%')
                    ORDER BY user_name
                ) AS filtered_users;
      
        else
        SELECT JSON_OBJECT(
                'getUserData', JSON_ARRAYAGG(
                    JSON_OBJECT(
                        'userId', user_id, 'userName', user_name, 'Address', Address, 'roles', roles, 'createdAt', created_at, 'updatedAt', updated_at
                    )
                ), 'totalRecords', total_records, 'page', page
            ) INTO get_user_data
        FROM (
                SELECT
                    user_id, user_name, Address, roles, created_at, updated_at
                FROM users
                WHERE
                    user_name LIKE CONCAT('%', search_query,'%')
                ORDER BY user_name
                LIMIT page_size
                OFFSET
                    offset_val
            ) AS filtered_users;
        end if;
        RETURN get_user_data;
    END;`
)
