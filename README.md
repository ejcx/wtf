# What the fuzz!
Simplified, pre-packaged, artisinal, hand-crafted fuzzing with AFL.

# How does this work.
Run doesn't actually run anything yet.
```
  docker build .
  docker run -it <image_id> /bin/bash
  make -f Makefile.fuzz jq
```
