# kpt Packages

kpt packages are consumed using the following commands:

`export SRC_REPO=git@github.com:trodge/kpt-samples.git
kpt pkg get $SRC_REPO/pubsub-subscription pubsub-subscription`

# kpt Configuration

List kpt setters

`kpt cfg list-setters pubsub-subscription`

Set field values

`kpt cfg set pubsub-subscription/ retain true`
