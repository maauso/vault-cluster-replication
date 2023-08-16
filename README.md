# (WIP) vault-cluster-replication

## Test it locally

### Prerequisites

- [Tilt](https://tilt.dev/)
- [Docker](https://www.docker.com/)
- [kind](https://kind.sigs.k8s.io/)

### Setup

The test environment is based on Tilt.
The tilt fins can be found in the [tilt](./tilt) directory.

#### Run Tilt:

Run the following command in the terminal to start Tilt:

```bash
tilt up
```

Tilt will orchestrate the creation of two Vault clusters named `vault-1` and `vault-2` within a Kubernetes cluster (
using KinD).

Throughout the process, there are two specific manual actions that require your attention. These actions involve
unsealing the Vault clusters and creating an `appRole` for the application's interaction.

Here's a step-by-step breakdown of the process:

1. Tilt will initiate the deployment of the `vault-1` cluster. At this point, your manual intervention is needed. You
   should perform the following steps:
   - Unseal the `vault-1` cluster.
   - Create an `appRole` tailored for the application.
   - Execute the `vault-1-operator-init.sh` script via the Tilt UI to set everything in motion.

2. Following the successful deployment of `vault-1`, Tilt will proceed to set up the `vault-2` cluster. Similarly, this
   phase requires your input:
   - Unseal the `vault-2` cluster.
   - Establish the corresponding `appRole` configuration for the application.
   - Initiate the `vault-2-operator-init.sh` script through the Tilt UI.

By following these steps, you'll ensure the proper unsealing of both Vault clusters and the creation of
application-specific `appRole` configurations. Tilt streamlines the deployment process, while your manual involvement
guarantees the appropriate setup of each cluster and the seamless integration of the application.

#### Access Vault UI:

Tilt will set up port forwarding for you, so you can access the Vault UI in your browser:

For vault-1, visit http://localhost:8200
For vault-2, visit http://localhost:8300

As a result of this setup, 2 new files will be created in the [tilt](./tilt) directory:

- `vault-1_unseal_keys.json`
- `vault-2_unseal_keys.json`

They will contain the unseal and root tokens for each Vault cluster. You can use these tokens to access the Vault UIs.




