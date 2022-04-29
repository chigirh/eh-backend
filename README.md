# eh-backend
## api specs
### Login Apis
#### ログイン [POST:/login]
+ Request
    + Body
        ```json
        {
            "user_name": "admin1",// required,max=64
            "password":"secret"// required
        }
        ```

+ Response 200 (application/json)

    + Body

        ```json
        {
            "user": {
                "id": "admin1",
                "first_name": "admin",
                "family_name": "developper",
                "password": "", // always empty
                "roles": [
                    "ADMIN"
                ]
            }
        }
        ```

+ Response 401 (application/json)

    + Body

        ```json
        {
            "message": "Error message"
        }
        ```

### User Apis
#### ユーザーデータ取得 [GET:/users/${user_name}]
+ Request
    + Header
        ```json
        // required
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```
+ Response 200 (application/json)

    + Body

        ```json
        {
            "user": {
                "id": "admin1",
                "first_name": "admin",
                "family_name": "developper",
                "password": "", // always empty
                "roles": [
                    "ADMIN"
                ]
            }
        }
        ```

+ Response 401 (application/json)
+ Response 403 (application/json)
+ Response 404 (application/json)

    + Body

        ```json
        {
            "message": "Error message"
        }
        ```

#### ユーザーデータ登録 [POST:/users]
+ Request
    + Header
        ```json
        // optional
        // can be executed without login
        x-master-key:xxxxxxxxxxxxxxxxxxxxxx
        // optional
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```
    + Body
        ```json
        {
            "user": {
                "id": "admin1", // required,max=64
                "first_name": "admin", // required,max=300
                "family_name": "developper", // required,max=300
                "password":"secret", // required
                "roles":["ADMIN"] // "ADMIN","CORP","GENE",min_size=1
            }
        }
        ```
+ Response 200 (application/json)

    + Body

        ```json
        {}
        ```

+ Response 401 (application/json)
+ Response 403 (application/json)
+ Response 409 (application/json)

    + Body

        ```json
        {
            "message": "Error message"
        }
        ```
