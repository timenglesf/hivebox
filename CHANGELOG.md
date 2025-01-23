# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.2.9](https://github.com/timenglesf/hivebox/compare/v0.2.8...v0.2.9) (2025-01-23)


### Features

* **ci:** refactor CI workflows for modularity ([48be3cd](https://github.com/timenglesf/hivebox/commit/48be3cd76f32045864c7b9aa72a886df04fab10a))


### Bug Fixes

* rename secret in reusable workflow ([b911ad9](https://github.com/timenglesf/hivebox/commit/b911ad979737924f4805e316a596b68d3b5eafd2))

### [0.2.8](https://github.com/timenglesf/hivebox/compare/v0.2.7...v0.2.8) (2025-01-19)

### [0.2.7](https://github.com/timenglesf/hivebox/compare/v0.2.6...v0.2.7) (2025-01-19)

### [0.2.6](https://github.com/timenglesf/hivebox/compare/v0.2.5...v0.2.6) (2025-01-19)


### Features

* add CD workflow for container build and push ([8513efb](https://github.com/timenglesf/hivebox/commit/8513efb120a4de43083c2bc2536d3c25525221cd))
* add Kubernetes configs and update timeouts ([811dfa6](https://github.com/timenglesf/hivebox/commit/811dfa68a05d1d4bcc99aa2580dd121a9b5b9366))

### [0.2.5](https://github.com/timenglesf/hivebox/compare/v0.2.4...v0.2.5) (2025-01-18)


### Features

* **api:** add temperature status and improve tests ([4c08d00](https://github.com/timenglesf/hivebox/commit/4c08d00fa8f551fccf504ffca798a5a0301287c0))


### Bug Fixes

* pin checkout actions by sha ([9d953c1](https://github.com/timenglesf/hivebox/commit/9d953c1d9ac9811a0180f49fbec13c5fd8506f94))
* update permissions and correct typo in CI workflow ([85bfadb](https://github.com/timenglesf/hivebox/commit/85bfadbae5e0fe01b4429f296d7e9d313de72846))

### [0.2.4](https://github.com/timenglesf/hivebox/compare/v0.2.3...v0.2.4) (2025-01-16)


### Bug Fixes

* add permissions for packages to ci.yaml ([5518e80](https://github.com/timenglesf/hivebox/commit/5518e80d8449588f894845ad524622728f7e5c3e))

### [0.2.3](https://github.com/timenglesf/hivebox/compare/v0.2.2...v0.2.3) (2025-01-16)

### [0.2.2](https://github.com/timenglesf/hivebox/compare/v0.1.2...v0.2.2) (2025-01-16)


### Features

* **ci:** add multi-arch container build and push ([df4ddfb](https://github.com/timenglesf/hivebox/commit/df4ddfb12d668727d2e4754927cd0fddc019c590))

### [0.2.1](https://github.com/timenglesf/hivebox/compare/v0.1.2...v0.2.1) (2025-01-16)

### Bug Fixes

* **ci:** typos in release.yaml ([6dd45ad](https://github.com/timenglesf/hivebox/commit/6dd45ad2dae0abf4802d2a4067a678dd677829c8))

## [0.2.0](https://github.com/timenglesf/hivebox/compare/v0.1.0...v0.2.0) (2025-01-16)

### Features

* **api:** add temperature endpoint ([d3aefa8](https://github.com/timenglesf/hivebox/commit/d3aefa87307fa26e497e965e8096fcf9ed16d500))
* **ci:** add Go style and containerization jobs ([df5e332](https://github.com/timenglesf/hivebox/commit/df5e332dbe3064dcccad8e9f403ecf81c715eb91))
* **ci:** add multi-platform build support ([48492bb](https://github.com/timenglesf/hivebox/commit/48492bbdba73b0225ee7480670756cf541cb0135))
* **ci:** add release workflow and update Dockerfile ([5b758cb](https://github.com/timenglesf/hivebox/commit/5b758cbd376decbacd550b10d857bb65107b814c))
* **ci:** enhance containerization jobs ([e5e90a7](https://github.com/timenglesf/hivebox/commit/e5e90a7ad7acde7d77aeac136a5d8540373129c8))
* **sensebox:** add dummy data and service implementation ([8167554](https://github.com/timenglesf/hivebox/commit/816755410c4c887c90080bf4ef8cdc80d0172700))

### Bug Fixes

* add apk update command ([dbbb832](https://github.com/timenglesf/hivebox/commit/dbbb8325efb21d803452bd3ee8c0199889655599))
* change git version ([aace425](https://github.com/timenglesf/hivebox/commit/aace425b86f13c77e108e377e9bada45ad9f2e6b))
* **ci:** update url in wget command ([60a02dc](https://github.com/timenglesf/hivebox/commit/60a02dc515956b700cdea7d00ba8bb18ce7c869e))
* Docker repository var names ([83d0d0a](https://github.com/timenglesf/hivebox/commit/83d0d0aefde33ab12419470da6bd75504ff9dbf4))
* **Dockerfile:** correct COPY command destination path ([90243c3](https://github.com/timenglesf/hivebox/commit/90243c35067601cfd5f584aa5f609359df68fd32))
* **Dockerfile:** correct git version and add build platform ([f40b3c7](https://github.com/timenglesf/hivebox/commit/f40b3c74245771e8fd1881533aae707116bd5bcb))
* **Dockerfile:** pin versions and remove --platform flag ([619e8d0](https://github.com/timenglesf/hivebox/commit/619e8d0d634fa641c5925183493d19fdad502dd2))
* env name for DockerHub login ([89f8b6e](https://github.com/timenglesf/hivebox/commit/89f8b6e45116443c1e9ba57f517ded15a0548308))
* fix incorrect conditional in getSesnseBoxIds ([c559a65](https://github.com/timenglesf/hivebox/commit/c559a65d26b896aefda641ee46724935eaf9c5f6))
* linting error in Dockerfile ([519c9e0](https://github.com/timenglesf/hivebox/commit/519c9e05b6a858fc7b7d79389e64b1d84fe35617))
* overwriting default sensebox ids ([e28c5bd](https://github.com/timenglesf/hivebox/commit/e28c5bd360cb2113d41a666243c3b3c5e53c1a99))
* unhandled error ([b6cdd7e](https://github.com/timenglesf/hivebox/commit/b6cdd7ee90639b954ea1958a588f278675e53e22))

### [0.2.1](https://github.com/timenglesf/hivebox/compare/v0.2.0...v0.2.1) (2025-01-16)

### Bug Fixes

* **ci:** typos in release.yaml ([6dd45ad](https://github.com/timenglesf/hivebox/commit/6dd45ad2dae0abf4802d2a4067a678dd677829c8))

## [0.2.0](https://github.com/timenglesf/hivebox/compare/v0.1.0...v0.2.0) (2025-01-16)

### Features

* **api:** add temperature endpoint ([d3aefa8](https://github.com/timenglesf/hivebox/commit/d3aefa87307fa26e497e965e8096fcf9ed16d500))
* **ci:** add Go style and containerization jobs ([df5e332](https://github.com/timenglesf/hivebox/commit/df5e332dbe3064dcccad8e9f403ecf81c715eb91))
* **ci:** add multi-platform build support ([48492bb](https://github.com/timenglesf/hivebox/commit/48492bbdba73b0225ee7480670756cf541cb0135))
* **ci:** add release workflow and update Dockerfile ([5b758cb](https://github.com/timenglesf/hivebox/commit/5b758cbd376decbacd550b10d857bb65107b814c))
* **ci:** enhance containerization jobs ([e5e90a7](https://github.com/timenglesf/hivebox/commit/e5e90a7ad7acde7d77aeac136a5d8540373129c8))
* **sensebox:** add dummy data and service implementation ([8167554](https://github.com/timenglesf/hivebox/commit/816755410c4c887c90080bf4ef8cdc80d0172700))

### Bug Fixes

* add apk update command ([dbbb832](https://github.com/timenglesf/hivebox/commit/dbbb8325efb21d803452bd3ee8c0199889655599))
* change git version ([aace425](https://github.com/timenglesf/hivebox/commit/aace425b86f13c77e108e377e9bada45ad9f2e6b))
* **ci:** update url in wget command ([60a02dc](https://github.com/timenglesf/hivebox/commit/60a02dc515956b700cdea7d00ba8bb18ce7c869e))
* Docker repository var names ([83d0d0a](https://github.com/timenglesf/hivebox/commit/83d0d0aefde33ab12419470da6bd75504ff9dbf4))
* **Dockerfile:** correct COPY command destination path ([90243c3](https://github.com/timenglesf/hivebox/commit/90243c35067601cfd5f584aa5f609359df68fd32))
* **Dockerfile:** correct git version and add build platform ([f40b3c7](https://github.com/timenglesf/hivebox/commit/f40b3c74245771e8fd1881533aae707116bd5bcb))
* **Dockerfile:** pin versions and remove --platform flag ([619e8d0](https://github.com/timenglesf/hivebox/commit/619e8d0d634fa641c5925183493d19fdad502dd2))
* env name for DockerHub login ([89f8b6e](https://github.com/timenglesf/hivebox/commit/89f8b6e45116443c1e9ba57f517ded15a0548308))
* fix incorrect conditional in getSesnseBoxIds ([c559a65](https://github.com/timenglesf/hivebox/commit/c559a65d26b896aefda641ee46724935eaf9c5f6))
* linting error in Dockerfile ([519c9e0](https://github.com/timenglesf/hivebox/commit/519c9e05b6a858fc7b7d79389e64b1d84fe35617))
* overwriting default sensebox ids ([e28c5bd](https://github.com/timenglesf/hivebox/commit/e28c5bd360cb2113d41a666243c3b3c5e53c1a99))
* unhandled error ([b6cdd7e](https://github.com/timenglesf/hivebox/commit/b6cdd7ee90639b954ea1958a588f278675e53e22))

## [0.1.0](https://github.com/timenglesf/hivebox/compare/v0.0.1...v0.1.0) (2025-01-12)

### Features

* add CI workflow and improve error handling ([faea8a5](https://github.com/timenglesf/hivebox/commit/faea8a50b652522e40ccf52be6292c9f639729ec))
* add configuration and version handling ([439aaf3](https://github.com/timenglesf/hivebox/commit/439aaf3fd3ecc3df3ad59d743cc9f9247e95454f))
* add version endpoint and test utilities ([7ce6fd4](https://github.com/timenglesf/hivebox/commit/7ce6fd44231dd4b327282a32f9c51682ba2eb513))
* **api:** add version endpoint and server setup ([3329259](https://github.com/timenglesf/hivebox/commit/3329259aa9a7bf5231fa7608feb8c98bb67b2c72))

### Bug Fixes

* add missing error handling to writeJSON ([6a5fb9b](https://github.com/timenglesf/hivebox/commit/6a5fb9ba8e069850ebee3595ebe3cf435c0e184b))

### 0.0.1 (2025-01-10)

### Features

* add Dockerfile 00ae572
* print version fe12abb
