# Changelog

## [0.2.0](https://github.com/unfunco/anthropic-sdk-go/compare/v0.1.0...v0.2.0) (2024-07-03)


### New features

* Add a constructor for the Transport ([#18](https://github.com/unfunco/anthropic-sdk-go/issues/18)) ([7707c5f](https://github.com/unfunco/anthropic-sdk-go/commit/7707c5f905c7207487656fe0b468ae4c849020d0))
* Add convenience methods for pointer creation ([#24](https://github.com/unfunco/anthropic-sdk-go/issues/24)) ([90623c7](https://github.com/unfunco/anthropic-sdk-go/commit/90623c7beca3a808aae0fa60a5a90dcf0ddc1dc4))
* Add support for the Claude 3.5 Sonnet model ([#27](https://github.com/unfunco/anthropic-sdk-go/issues/27)) ([d0f5ca9](https://github.com/unfunco/anthropic-sdk-go/commit/d0f5ca9e9a964ba3ba8ca7e62adea048b9cef14c))


### Bug fixes

* Fix an ineffectual assignment to err ([#11](https://github.com/unfunco/anthropic-sdk-go/issues/11)) ([fdf2a64](https://github.com/unfunco/anthropic-sdk-go/commit/fdf2a64a04274d1b5aafd7a643462aca4cd8bfe0))
* Fix name of the claude-3-haiku-20240307 const ([#13](https://github.com/unfunco/anthropic-sdk-go/issues/13)) ([430676d](https://github.com/unfunco/anthropic-sdk-go/commit/430676de0a1c718f76dcf948dc6d09cc7e1171bd))
* Fix the format of the user-agent header ([#23](https://github.com/unfunco/anthropic-sdk-go/issues/23)) ([de05838](https://github.com/unfunco/anthropic-sdk-go/commit/de058382074e1e2084975d2274580c92687484cf))
* Return an error if the context is nil ([#22](https://github.com/unfunco/anthropic-sdk-go/issues/22)) ([c284ec6](https://github.com/unfunco/anthropic-sdk-go/commit/c284ec6d43950fa4401d8711d4fec402206496a7))


### Miscellaneous

* Add workflow badges to the README ([#15](https://github.com/unfunco/anthropic-sdk-go/issues/15)) ([059718d](https://github.com/unfunco/anthropic-sdk-go/commit/059718d9b917acf5bf67d1acdad10d5f04fd7ff2))
* Automatically add labels to pull requests ([#17](https://github.com/unfunco/anthropic-sdk-go/issues/17)) ([5c1335c](https://github.com/unfunco/anthropic-sdk-go/commit/5c1335ca51aeb7bdb4d0dc7c73ed72114ca2dc65))
* Configure golangci-lint in GitHub Actions ([#10](https://github.com/unfunco/anthropic-sdk-go/issues/10)) ([83b8766](https://github.com/unfunco/anthropic-sdk-go/commit/83b8766c74f0058da0446619279e3db0ec3cb41e))
* Fix a broken automated commit process ([#19](https://github.com/unfunco/anthropic-sdk-go/issues/19)) ([a7f0d4e](https://github.com/unfunco/anthropic-sdk-go/commit/a7f0d4ed2c142ec30751d0e118cbe2fe190ccf93))
* Make version updates less fiddly ([#25](https://github.com/unfunco/anthropic-sdk-go/issues/25)) ([d7b24a9](https://github.com/unfunco/anthropic-sdk-go/commit/d7b24a931e1330e5181cd1af8e87f7391fe527bf))
* Pull before committing version changes ([#20](https://github.com/unfunco/anthropic-sdk-go/issues/20)) ([c1bda69](https://github.com/unfunco/anthropic-sdk-go/commit/c1bda6992e6601f15452e3a2ae6d1dbda2238dbc))
* Reformat code in the README document ([#14](https://github.com/unfunco/anthropic-sdk-go/issues/14)) ([696f732](https://github.com/unfunco/anthropic-sdk-go/commit/696f732979b4d2903b2d45ce911d0393264511ee))
* Simplify user-agent and version updates ([#16](https://github.com/unfunco/anthropic-sdk-go/issues/16)) ([9ebd683](https://github.com/unfunco/anthropic-sdk-go/commit/9ebd683211c4f13800a69f4bf6f15c5b65cde5fc))
* Update release-please-action organisation ([#26](https://github.com/unfunco/anthropic-sdk-go/issues/26)) ([a01ee1f](https://github.com/unfunco/anthropic-sdk-go/commit/a01ee1f656720bb51ddd12db6f536c7f209ef524))

## 0.1.0 (2024-04-18)


### New features

* Add the initial Anthropic SDK implementation ([#1](https://github.com/unfunco/anthropic-sdk-go/issues/1)) ([b3bf332](https://github.com/unfunco/anthropic-sdk-go/commit/b3bf332e64d75e8485b26a6ecc609a497f183cbd))


### Miscellaneous

* Add an update-user-agent composite action ([#7](https://github.com/unfunco/anthropic-sdk-go/issues/7)) ([5cf68b7](https://github.com/unfunco/anthropic-sdk-go/commit/5cf68b7f966da669477f3d2412a10d189489de3f))
* Add some repository boilerplate ([bfb93b6](https://github.com/unfunco/anthropic-sdk-go/commit/bfb93b6023bb9369ef42fc1fc2d28172b9a35d2c))
* Apply the MIT license and add a README ([1f87886](https://github.com/unfunco/anthropic-sdk-go/commit/1f87886531a5c17a06c60a646ea0771872144686))
* Configure an initial release process ([#2](https://github.com/unfunco/anthropic-sdk-go/issues/2)) ([b49ea86](https://github.com/unfunco/anthropic-sdk-go/commit/b49ea8602d1ca2c033a761cf7ad2ca20bc26fe7d))
* Improve documentation in source files ([#4](https://github.com/unfunco/anthropic-sdk-go/issues/4)) ([8124555](https://github.com/unfunco/anthropic-sdk-go/commit/8124555f31e647f6de0d4a236b27dfb9324970a5))
* Improve installation and usage instructions ([#9](https://github.com/unfunco/anthropic-sdk-go/issues/9)) ([23c4ab4](https://github.com/unfunco/anthropic-sdk-go/commit/23c4ab4f32c3347f062944270f85c580813ea66c))
* Initialise the repository ([c2b5cc2](https://github.com/unfunco/anthropic-sdk-go/commit/c2b5cc271f4deec7c6f5671a467130be96754b8a))
* Push version changes to the release branch ([#6](https://github.com/unfunco/anthropic-sdk-go/issues/6)) ([2d5abaf](https://github.com/unfunco/anthropic-sdk-go/commit/2d5abaf1b589609c2325cc767f51015c0eb78fc5))
* Reformat the LICENSE.md document ([#8](https://github.com/unfunco/anthropic-sdk-go/issues/8)) ([725d14a](https://github.com/unfunco/anthropic-sdk-go/commit/725d14ac813b34225a5e282280d0f3d2736c6c8a))
* Update version when releases are published ([#5](https://github.com/unfunco/anthropic-sdk-go/issues/5)) ([82fe560](https://github.com/unfunco/anthropic-sdk-go/commit/82fe56045cb9fd2e09891a523d92dbcf5bc5d9a9))
