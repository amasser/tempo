## master / unreleased

* [CHANGE] Bloom filters are now sharded to reduce size and improve caching, as blocks grow. This is a **breaking change** and all data stored before this change will **not** be queryable. [192](https://github.com/grafana/tempo/pull/192)
* [ENHANCEMENT] CI checks for vendored dependencies using `make vendor-check`. Update CONTRIBUTING.md to reflect the same before checking in files in a PR. [#274](https://github.com/grafana/tempo/pull/274)