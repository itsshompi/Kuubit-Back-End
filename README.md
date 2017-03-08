# kuubit-backend


##### Kuubit API Routes

| Route      |      Method   |  Params                      |  Return                   | Content-Type     |
|:-----------|:-------------:|:-----------------------------|:--------------------------|:-----------------|
| /          |  GET          | None                         | name, description, version, webiste, url, repository, created_by, github, contact | application/json |
| /auth/signin     |  POST         | email, password              | name, avatar, slug, token | application/json |
| /auth/signup    |  POST         | name, email, slug, password  | name, avatar, slug, token | application/json |
