# Authorization with OPA

This folder holds the policy for the gocook service.
To adjust the policy, simply modify the policies.rego file and trigger a rebuild of the .tar.gz.

This can be done via the follwing command:

`opa build src/authorization -o src/authorization/bundle.tar.gz`

The bundle-server will provide the .tar.gz file for the OPA container.