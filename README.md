# What the fuzz!
Simplified, pre-packaged, artisinal, hand-crafted, dockerized
fuzzing and crash collection with AFL.

Spin up your own fuzz farm quickly and easily. Write a test,
orchestrate the container in the cloud, and collect crashes
on Amazon S3.

# How does it work?
### Building a fuzzing image
Build and instrument programs for fuzzing with AFL in docker.

To build a docker image that will fuzz a specific program when it
is started, run a docker build, and specify a fuzztgt as a build arg.

```
docker build --build-arg fuzztgt=jq  .
docker run ...
```

The current fuzz targets have working harnasses.
 - jq
 - bc

### Collecting Crashes.
In order to collect crashes you need to specify your S3
details in `crashbot/crashbot.go`. You don't need to know
Golang to do this. You'll see a revoked old pair of S3
credentials are currently there to guide you.

