# rancher-wrangler

## Documentation

To view generated docs for rancher-wrangler, run the below command and point your brower to http://127.0.0.1:6060/pkg/

    godoc -http=:6060 2>/dev/null &

### Buffalo

http://gobuffalo.io/docs/getting-started

### Pop/Soda

http://gobuffalo.io/docs/db

## Database Configuration

 	development:
 		dialect: postgres
 		database: rancher-wrangler_development
 		user: <username>
 		password: <password>
 		host: 127.0.0.1
 		pool: 5

 	test:
 		dialect: postgres
 		database: rancher-wrangler_test
 		user: <username>
 		password: <password>
 		host: 127.0.0.1

 	production:
 		dialect: postgres
 		database: rancher-wrangler_production
 		user: <username>
 		password: <password>
 		host: 127.0.0.1
 		pool: 25

 ### Running Migrations

    buffalo soda migrate

 ## Run Tests

    buffalo test

 ## Run in dev

    buffalo dev

[Powered by Buffalo](http://gobuffalo.io)

