Postgres commands

```
sudo -u postgres psql
create user fiber_api_user with encrypted password 'mypassword';
grant all privileges on database go_fiber_api to fiber_api_user;


```

# Steps:
## To create migration
- Install golang migrate library
- `migrate create -ext sql -dir internal/db/migration -seq init_schema` 
## To handle dirty migration
- `update schema_migrations set dirty=false where version=5;`