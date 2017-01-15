# rancher-wrangler a.k.a "The Wrangler"

## Documentation
Wrangler is a Central Web Interface to see detailed Information about multiple Rancher Management Servers.
 This is not a Web application for Rancher Server Federation (e.g Kubernetes Federation). In it's present state,
 this tool is very bare bones, Hosts added are then queried every so often for data via the API, such as the Server Version, Docker version, etc.

 ## Getting Started
 For dependency management, we're using [Glide](https://github.com/Masterminds/glide).

 * UNIX/Linux

 ```
 curl https://glide.sh/get | sh
 ```

 * MacOSX

 ```
 brew install glide
 ```

 * Ubuntu

 ```
 sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
 sudo apt-get install glide
 ```

 Versions are managed in the `glide.yaml` file and a subsequent `glide.lock` file is created.

 Once Glide is installed we need to make sure that we have the PostgreSQL image downloaded, have the container built, schema loaded, and dependencies installed. To accomplish this, run the below command.

 ### Running Migrations

    buffalo db migrate

 ## Run Tests

    buffalo test

 ## Run in dev

    buffalo dev

 ## Releases will be pushed to Docker Hub as Images

    bscott/rancher-wrangler

[Powered by Buffalo](http://gobuffalo.io)

