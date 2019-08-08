# Service Operator

We aim for this to be our generic service operator, with a hope for it to be a
provider agnostic one. The reason being, a desire to run the stack locally,
without the worry of providing live database connections to a local deployment.

## Adding new services

Say you have a desire to add a new service to the mix. We're using
[Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) as our generator for all the things.

To get started, [install Kubebuilder locally](https://book.kubebuilder.io/quick-start.html#installation) and run:

```sh
kubebuilder create api --group database --version v1beta1 --kind MySQL
...
kubebuilder create api --group queue --version v1beta1 --kind RabbitMQ
```

Where `--group` defines the general nature of the service and `--kind` defines
the engine we'd like to run.

This will generate some `api/`, `config/` and `controller/` files, which you
may want to look into. It usually comes with a nice set of comments, which will
explain to you, what needs to be done in each section to get to your desired
state.

Note, that because it's essentially a generator, for most changes to the
codebase, you may want to run `make generate`.

## PostgreSQL via AWS RDS

Given a configuration file of the form:

```yaml
---
apiVersion: database.gsp.k8s.io/v1beta1
kind: Postgres
metadata:
  name: postgres-sample
spec:
  aws:
    diskSizeGB: 150
    instanceType: db.m5.large
```

This service operator will create a Postgres database in AWS RDS, and generate
a secret only available to that namespace, containing all required data.
