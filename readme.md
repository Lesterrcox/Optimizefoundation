# Optimizefoundation
Optimize is a layer 1 blockchain anyone can understand. It turns real-world things like gold, homes, government bonds into digital assets you can verify, track and trade. Think of it as a global, secure, tamper-proof spreadsheet that you can trust.  Why it matters? Trust is expensive; lawyers, middle men, paperwork, legal rules and delays. 

## Get started

### Using ignite cli
```
ignite chain serve -r -v
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Using local deployment script
```
cd scripts
./local-network
```

## Custom Modules

The blockchain includes three custom modules:

- **Validator Module**: Handles validator-related functionality and operations
- **OConsensus Module**: Manages consensus-related operations and parameters
- **Asset Module**: Handles asset management and operations

Each module is located in the `x` directory of the project.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/optimizeglobalcoin@latest! | sudo bash
```
`username/optimizeglobalcoin` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
