# kpt Packages

kpt packages are consumed using the following commands:

```
export SRC_REPO=git@github.com:trodge/kpt-samples.git
kpt pkg get $SRC_REPO/pubsub-subscription pubsub-subscription
```

# kpt Configuration

List kpt setters

`kpt cfg list-setters pubsub-subscription`

Set field values

`kpt cfg set pubsub-subscription/ retain true`

# create-setters.sh

This script will create setters for kpt. Its parameter should be the path to a
yaml file in a directory which also contains a file named "fields" which
contains, as tab seperated values, the names, descriptions, and types of all
fields for which setters should be created.
