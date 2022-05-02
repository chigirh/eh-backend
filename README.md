# eh-backend
## build
```bash
docker-compose build
docker-compose up -d
```
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
### Auth Apis
#### パスワード変更 [POST:/auth]
+ Request
    + Body
        ```json
        {
            "before":"secret", // required
            "after":"secret2" // required
        }
        ```

+ Response 200 (application/json)

    + Body

        ```json
        {}
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

### Schedules Apis
#### スケジュール登録 [POST:/schedules]
+ Request
    + Header
        ```json
        // required
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```

    + Body

        ```json
            {
                "schedules": [
                    {
                        "date":"2022-05-06", // uuuu-MM-dd
                        "periods":[1,2,3,4,5,6,7,8,9,10,11,12,14,15] // min=1,max=48
                    }
                ]
            }
        ```
+ Response 200 (application/json)

    + Body

        ```json
        {}
        ```

+ Response 403 (application/json)
    + Body

        ```json
        {
            "message": "Error message"
        }
        ```

#### スケジュール集計 [POST:/schedules/aggregate]
+ Request
    + Header
        ```json
        // optional
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```
    + Query parameter
        ```json
        /schedules/aggregate?from=2022-05-01&to=2022-05-10
        ```
+ Response 200 (application/json)

    + Body

        ```json
        {
            "aggregates": [
                {
                    "date": "2022-05-01",
                    "periods": [
                        {
                            "period": 1,
                            "count": 10
                        },
                        {
                            "period": 2,
                            "count": 11
                        }
                    ]
                },
                {
                    "date": "2022-05-02",
                    "periods": [
                        {
                            "period": 12,
                            "count": 12
                        },
                        {
                            "period": 13,
                            "count": 12
                        }
                    ]
                }
            ]
        }
        ```
#### 詳細取得 [POST:/schedules/details]
+ Request
    + Header
        ```json
        // optional
        // issue with login api
        x-session-token:e261c5e5-02ad-49da-a90d-19a6c6eecb75
        ```
    + Query parameter
        ```json
        /details?date=2022-05-01&period=1
        ```
+ Response 200 (application/json)

    + Body

        ```json
        {
            "aggregates": [
                {
                    "date": "2022-05-01",
                    "periods": [
                        {
                            "period": 1,
                            "count": 10
                        },
                        {
                            "period": 2,
                            "count": 11
                        }
                    ]
                },
                {
                    "date": "2022-05-02",
                    "periods": [
                        {
                            "period": 12,
                            "count": 12
                        },
                        {
                            "period": 13,
                            "count": 12
                        }
                    ]
                }
            ]
        }
        ```

+ Response 403 (application/json)

    + Body

        ```json
        {
            "message": "Error message"
        }
        ```
