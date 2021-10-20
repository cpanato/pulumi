### Improvements

- [cli] Add `--json` flag to `up`, `destroy` and `refresh`.

  Passing the `--json` flag to `up`, `destroy` and `refresh` will stream JSON events from the engine to stdout. 
  For `preview`, the existing functionality of outputting a JSON object at the end of preview is maintained. 
  However, the streaming output can be extended to `preview` by using the `PULUMI_ENABLE_STREAMING_JSON_PREVIEW` environment variable.

  [#8275](https://github.com/pulumi/pulumi/pull/8275)

### Bug Fixes
