# kuubit-backend


API REST ROUTES

| Route      |      Method   |  Params                      |  Return                   | Content-Type     |
|:----------:|:-------------:|:----------------------------:|:-------------------------:|:----------------:|
| /          |  GET          | None                         | name, version, created_by | application/json |
| /login     |  POST         | email, password              | name, avatar, slug, token | application/json |
| /signup    |  POST         | name, email, slug, password  | name, avatar, slug, token | application/json |
