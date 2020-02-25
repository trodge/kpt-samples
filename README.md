# kpt Packages

kpt packages are consumed using the following command:

```
kpt pkg get https://github.com/trodge/kpt-samples.git/pubsub pubsub
```

# kpt Configuration

List kpt setters

`kpt cfg list-setters pubsub`

Set field values

`kpt cfg set pubsub retain true`

# create-setters.sh

This script will create setters for kpt. Its parameter should be the path to a
yaml file in a directory which also contains a file named "fields" which
contains, as tab seperated values, the names, descriptions, and types of all
fields for which setters should be created.
