Postgres commands

```
sudo -u postgres psql
create user fiber_api_user with encrypted password 'mypassword';
grant all privileges on database go_fiber_api to fiber_api_user;


```