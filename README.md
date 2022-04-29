# eh-backend
## api specs
### Login Apis
#### ログイン [POST:/login]
<details><summary>Request</summary><div>
    + Body
        ```json
        {
            "user_name": "admin1",// required,max=64
            "password":"secret"// required
        }
        ```
</div></details>

<details><summary>Response 200 (application/json)</summary><div>
    + Body

        ```json
        {
            "session_token": "c9a882ce-b2f4-465a-821b-2871f9325fd1"
        }
        ```
</div></details>

<details><summary>Response 401 (application/json)</summary><div>

    + Body
        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>

### User Apis
#### ユーザーデータ取得 [GET:/users/${user_name}]
<details><summary>Request</summary><div>
    + Header
        ```json
        // required
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```
</div></details>

<details><summary>Response 200 (application/json)</summary><div>
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
</div></details>

<details><summary>Response 401 (application/json)</summary><div>
    + Body
        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>

<details><summary>Response 403 (application/json)</summary><div>
    + Body
        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>

<details><summary>Response 404 (application/json)</summary><div>
    + Body

        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>

#### ユーザーデータ登録 [POST:/users]
<details><summary>Request</summary><div>
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
</div></details>

<details><summary>Response 200 (application/json)</summary><div>
    + Body

        ```json
        {}
        ```
</div></details>
<details><summary>Response 401 (application/json)</summary><div>
    + Body
        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>
<details><summary>Response 403 (application/json)</summary><div>
    + Body
        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>
<details><summary>Response 409 (application/json)</summary><div>
    + Body

        ```json
        {
            "message": "Error message"
        }
        ```
</div></details>
