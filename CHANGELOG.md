# Changelog

## [2.0.0](https://github.com/CamiloValderruten/faultline/compare/v1.7.0...v2.0.0) (2026-08-27)


### ⚠ BREAKING CHANGES

* footer, which is too sharp a knife for a project at this stage. Add 'Ask the maintainer first.' to the table row and a paragraph spelling out the policy: contributors don't unilaterally mark breaking changes; the maintainer accumulates them and decides when to cut a major. Also documents the Release-As: footer for forcing a specific version when needed.

### Features

* add agent.wait_for_tools collaborator inject mode ([#23](https://github.com/CamiloValderruten/faultline/issues/23)) ([ffc8706](https://github.com/CamiloValderruten/faultline/commit/ffc87063cd9de0840a9466bd7932976cda15253a))
* add Discord collaborator channel with rich messaging ([7195d64](https://github.com/CamiloValderruten/faultline/commit/7195d640e40516642293bdcd7d736dec27b5b61f))
* add Discord voice channel config fields ([0ef5719](https://github.com/CamiloValderruten/faultline/commit/0ef5719be149d76b2c98eb7f39224222b156f965))
* add email_fetch tool ([f1b4ab6](https://github.com/CamiloValderruten/faultline/commit/f1b4ab6f6cde2349247841b29f6845cf33767b30))
* add priority inbox and authenticated webhook ([4a50941](https://github.com/CamiloValderruten/faultline/commit/4a509416057d84d8bdd412509c14e257759c0a23))
* add priority inbox and authenticated webhook ([a7d6dbf](https://github.com/CamiloValderruten/faultline/commit/a7d6dbfbea8750f6b1133462a098f1dafea46860))
* add priority inbox and authenticated webhook ([#27](https://github.com/CamiloValderruten/faultline/issues/27)) ([4a50941](https://github.com/CamiloValderruten/faultline/commit/4a509416057d84d8bdd412509c14e257759c0a23))
* add pull-only peer messaging between Faultline processes ([b96c14e](https://github.com/CamiloValderruten/faultline/commit/b96c14ecd41c35cf4b161a5495b4087db07a5eb1))
* add pull-only peer messaging between Faultline processes ([4962e03](https://github.com/CamiloValderruten/faultline/commit/4962e03dc8ab6b8e5582318379686d7147d15dc1))
* add rebuild_indexes tool and refactor reconcile/rebuild helpers ([27aea57](https://github.com/CamiloValderruten/faultline/commit/27aea575d51f7509af3697e878a81b75f40f03e9))
* add release pipeline (release-please + goreleaser) ([3075763](https://github.com/CamiloValderruten/faultline/commit/3075763c74284f2fe9f2447a2f2b5724003e8c89))
* add release pipeline (release-please + goreleaser) ([0a1b86d](https://github.com/CamiloValderruten/faultline/commit/0a1b86d82e82af0b822eadbb971cddd383ae3217))
* add sandbox_shell tool for arbitrary shell command execution ([40c0922](https://github.com/CamiloValderruten/faultline/commit/40c0922d017847e2fb8d11a23754c17d84430b15))
* add sandbox_shell tool for arbitrary shell command execution ([efe41cb](https://github.com/CamiloValderruten/faultline/commit/efe41cb2d7d4f58c13f762437e8b83df9943f964))
* add scheduled tasks, Docker deployment, and sandbox cron ([9521090](https://github.com/CamiloValderruten/faultline/commit/952109077747dca592f9101744c265b0cd5a0070))
* add self-update support ([cc7090e](https://github.com/CamiloValderruten/faultline/commit/cc7090eb6c9b6da4601bb3a5c5be132519fcc5d6))
* add self-update support ([1739d61](https://github.com/CamiloValderruten/faultline/commit/1739d61bab2a551858a099442343315c04e4437f))
* add semantic memory search with persisted vector index ([ff012dc](https://github.com/CamiloValderruten/faultline/commit/ff012dc0f08ffe7eb62afbf4b2e6bec5d1898fba))
* add semantic memory search with persisted vector index ([6265089](https://github.com/CamiloValderruten/faultline/commit/62650891ae24f9c555af3a3f129ea94bdbb21225))
* **admin:** in-place config editor with validate / save / restart ([#34](https://github.com/CamiloValderruten/faultline/issues/34)) ([ec1608e](https://github.com/CamiloValderruten/faultline/commit/ec1608ec02382ca2b61a7df515e7e3669c1c3a93))
* **admin:** live agent inspector + tool-call feed on dashboard ([dcb252c](https://github.com/CamiloValderruten/faultline/commit/dcb252c2a4739ab832b66cf5345ea273a7ba0cc0))
* **admin:** matrix-themed dashboard with sidebar nav and reflective config form ([#43](https://github.com/CamiloValderruten/faultline/issues/43)) ([faa6eed](https://github.com/CamiloValderruten/faultline/commit/faa6eed303620a015acfd80b8998e6ed0185703d))
* **admin:** per-skill enable/disable persisted to skills.toml ([#32](https://github.com/CamiloValderruten/faultline/issues/32)) ([9c842bc](https://github.com/CamiloValderruten/faultline/commit/9c842bc4fd9ce9f6f2cc820b287ea4921e5c93b7))
* **admin:** scaffold HTTP admin UI with login + session auth ([2bcba17](https://github.com/CamiloValderruten/faultline/commit/2bcba17b966eac0ec13cc5bcba1c67fadf8e533f))
* **admin:** self-update card with apply-now button ([#33](https://github.com/CamiloValderruten/faultline/issues/33)) ([f61e87a](https://github.com/CamiloValderruten/faultline/commit/f61e87a3f2e866e7e1f8379933e92b79c1c5d30c))
* Agent Skills support ([#24](https://github.com/CamiloValderruten/faultline/issues/24)) ([eaad5a0](https://github.com/CamiloValderruten/faultline/commit/eaad5a0eb8d43f80c709b97e474f10fe3def3fad))
* custom Arch-based multi-runtime sandbox image ([#23](https://github.com/CamiloValderruten/faultline/issues/23)) ([41bee48](https://github.com/CamiloValderruten/faultline/commit/41bee48f1bd97c14e76ea5acdc3a97bc3186667d))
* Deepgram SpeakOggOpus for Discord voice playback ([b64854d](https://github.com/CamiloValderruten/faultline/commit/b64854de47e10a5936a2b6afbb1c7599b31c6766))
* Discord live voice channel with silence-gated turns ([a3451c3](https://github.com/CamiloValderruten/faultline/commit/a3451c309becd73f4842287a65905906581bc199))
* Discord modals, send_file, voice bubbles, typing, entity selects ([c2aa010](https://github.com/CamiloValderruten/faultline/commit/c2aa010cff70e4e1efbf2dcb7cd1a568fa9db40b))
* Discord modals, send_file, voice bubbles, typing, entity selects ([e7b4f73](https://github.com/CamiloValderruten/faultline/commit/e7b4f737eab22377f9cfcb724a136731a31ecefb))
* Discord voice notes via Deepgram STT/TTS ([955cfb7](https://github.com/CamiloValderruten/faultline/commit/955cfb7f28b3987e54659acfdebaf62d4337c052))
* Docker-backed daemon_* tools with alerts and skills ([7688b03](https://github.com/CamiloValderruten/faultline/commit/7688b034fa11531159a68ea7720a042db05e579e))
* hybrid BM25+vector search for search_available_tools ([bbb7666](https://github.com/CamiloValderruten/faultline/commit/bbb766696419ea11f918bac22e7b52a89d2da674))
* hybrid BM25+vector search for search_available_tools ([e83b87c](https://github.com/CamiloValderruten/faultline/commit/e83b87caa6eb55ef12992047df0695a2bcb695f7))
* inject collaborator channel guide into cycle context ([22d6756](https://github.com/CamiloValderruten/faultline/commit/22d67569de678f72ccbea1139e3814822a2c8842))
* inject sandbox env and ship gh in sandbox image ([8e474b2](https://github.com/CamiloValderruten/faultline/commit/8e474b2e13ad5efb28e769ecfdb4038c8fa313ba))
* **mcp:** add OAuth setup for HTTP servers ([ca1ec0f](https://github.com/CamiloValderruten/faultline/commit/ca1ec0f8e8c39a89139c041a846c2ed1700a326a))
* **mcp:** require current config for updates ([e8ddaee](https://github.com/CamiloValderruten/faultline/commit/e8ddaee57a1ba87ef202a811f47311d3924793b3))
* **mcp:** support sandboxed and streamable servers ([#35](https://github.com/CamiloValderruten/faultline/issues/35)) ([5350f78](https://github.com/CamiloValderruten/faultline/commit/5350f7848306f45e5197e9b7eb325ff5a75d78b5))
* optional peer message inject delivery mode ([901554b](https://github.com/CamiloValderruten/faultline/commit/901554bbe4cd765c66e41ff0ee9bf99a764f28e4))
* paragraph-aligned semantic indexing with adaptive batching ([#19](https://github.com/CamiloValderruten/faultline/issues/19)) ([e5d7918](https://github.com/CamiloValderruten/faultline/commit/e5d7918539d43b80f499101538c171c237e1bcc5))
* prevent tool error loops with schema feedback and forced compaction ([#29](https://github.com/CamiloValderruten/faultline/issues/29)) ([93a18c1](https://github.com/CamiloValderruten/faultline/commit/93a18c102068a6540b3343ae1fb24a8b8c41ffc0))
* **prompts:** identity/operating split, changelog, softer compaction ([#41](https://github.com/CamiloValderruten/faultline/issues/41)) ([a09a0d5](https://github.com/CamiloValderruten/faultline/commit/a09a0d5a2d7ce9466e673ddee1f4816bf95af522))
* **publish:** HTML publishing harness with HTTP serving layer ([ed0ac3d](https://github.com/CamiloValderruten/faultline/commit/ed0ac3d272e2f93c0b7226c58637fa422505c95e))
* **publish:** HTTP serving layer for HTML publishing harness ([4a2ca75](https://github.com/CamiloValderruten/faultline/commit/4a2ca758eecdf556bcda2626c795f0e028aa8af3))
* **publish:** wire and harden HTML publishing harness ([f27390a](https://github.com/CamiloValderruten/faultline/commit/f27390a85b3e7200504d4b03d2c0e2d162014056))
* require delivered reply for collaborator messages ([c65712e](https://github.com/CamiloValderruten/faultline/commit/c65712e2df99a2fa40510975f65f0625414cf4b0))
* **skills:** add html-artifact publishing skill ([36749d8](https://github.com/CamiloValderruten/faultline/commit/36749d8fcab3922efcf79a6fe03414b88d33f1e7))
* **skills:** add html-artifact publishing skill ([22c76b7](https://github.com/CamiloValderruten/faultline/commit/22c76b7c6ad8382cfc9e3f9f4cfadb7b0f2a12ef))
* **skills:** security audit subagent + sandbox hardening ([#27](https://github.com/CamiloValderruten/faultline/issues/27)) ([93515d5](https://github.com/CamiloValderruten/faultline/commit/93515d58b3dc86f5666f8c9bea6033956704fd7a))
* subagent delegation ([#26](https://github.com/CamiloValderruten/faultline/issues/26)) ([c42b698](https://github.com/CamiloValderruten/faultline/commit/c42b6988b0e537d6e647467acdefef81d9e21b53))
* **telegram:** inbound photos, buttons, rich messages, and thinking control ([cfd8752](https://github.com/CamiloValderruten/faultline/commit/cfd875263d439bd833d59c3bfacae57111033e5f))
* **telegram:** inbound photos, buttons, rich messages, and thinking control ([d045492](https://github.com/CamiloValderruten/faultline/commit/d045492b53e44d1f352bec5865abd7bb98df443e))
* two-tier tool loading with search_available_tools ([e0335dc](https://github.com/CamiloValderruten/faultline/commit/e0335dc99677f44f8010ec24686ffd5a1191df83))
* two-tier tool loading with search_available_tools ([079d4c4](https://github.com/CamiloValderruten/faultline/commit/079d4c4abc0f0e76b1007eb5eff470de656d8cf5))
* untrusted-tool-output guard + prompt-migration delivery ([#36](https://github.com/CamiloValderruten/faultline/issues/36)) ([98819f1](https://github.com/CamiloValderruten/faultline/commit/98819f114d92d2ceb3174e9664f1bd5147c7553c))


### Bug Fixes

* backoff on LLM 429 instead of exiting ([0644c36](https://github.com/CamiloValderruten/faultline/commit/0644c36beebd5a5b5afd381673bef878250942e3))
* cap tool results and fail-safe empty compaction ([aba53f9](https://github.com/CamiloValderruten/faultline/commit/aba53f98b9a5c8b62fdfde21a1eee28558da9674))
* cap tool results and fail-safe empty compaction ([bcbf228](https://github.com/CamiloValderruten/faultline/commit/bcbf22800a6a7586d00fd41d5fe567a74ff5ca65))
* **ci:** fallback to GITHUB_TOKEN if RELEASE_PLEASE_TOKEN is unset ([ff5b23f](https://github.com/CamiloValderruten/faultline/commit/ff5b23f291661353273d8084368bff44bb5756d1))
* **ci:** set GH_REPO so gh release commands skip git detection ([7ebc3e5](https://github.com/CamiloValderruten/faultline/commit/7ebc3e5843ca4e87028addc74db9f1c723faea94))
* **ci:** set GH_REPO so release-notes step works without checkout ([#38](https://github.com/CamiloValderruten/faultline/issues/38)) ([7ebc3e5](https://github.com/CamiloValderruten/faultline/commit/7ebc3e5843ca4e87028addc74db9f1c723faea94))
* complete Discord DAVE MLS handshake ([24fe33d](https://github.com/CamiloValderruten/faultline/commit/24fe33d54e83f52fcb1449502e998afc5c43293b))
* correct broken backtick in sleep tool struct tag ([f005570](https://github.com/CamiloValderruten/faultline/commit/f005570499228d611694ea307283467185e41f1a))
* **discord:** disable message components after click ([59d1231](https://github.com/CamiloValderruten/faultline/commit/59d1231c339f6c080424a3000cc1eb7dc6d51e26))
* do not silently fall back to chat voice attachments in VC ([21ff340](https://github.com/CamiloValderruten/faultline/commit/21ff3409bda204823d4ba93a0da93655c1c6696e))
* log Discord voice close codes and fail DAVE wait fast ([96bdfd4](https://github.com/CamiloValderruten/faultline/commit/96bdfd4cc99a9829816b171771543aecb2d68273))
* **mcp:** restart stdio servers with fresh config ([0071563](https://github.com/CamiloValderruten/faultline/commit/00715630681ef2b703b528d0b82cf5bd57e4ae7a))
* properly integrate sleep tool into tools.go ([d7b7554](https://github.com/CamiloValderruten/faultline/commit/d7b7554b36f3a2bbe652260ac50a54336ecda085))
* properly integrate sleep tool into tools.go (clean version) ([9d6e0fd](https://github.com/CamiloValderruten/faultline/commit/9d6e0fd1ccfa7b246dcf353266c2d2848f3b19da))
* properly integrate sleep tool into tools.go with method ([3e92ab9](https://github.com/CamiloValderruten/faultline/commit/3e92ab9e4d91039bdb90a279dcb4ce5e904942d3))
* publish Arlo webhook on host loopback ([e7cfb5f](https://github.com/CamiloValderruten/faultline/commit/e7cfb5f261a05985e55736bc320d773296a26a3e))
* remove extra blank line (gofmt) ([e905454](https://github.com/CamiloValderruten/faultline/commit/e90545413d7872cf298aeb738ddeb57264e9b1dd))
* retry Discord gateway connect with backoff and retry transient LLM errors ([#28](https://github.com/CamiloValderruten/faultline/issues/28)) ([45e8f48](https://github.com/CamiloValderruten/faultline/commit/45e8f48b22436b521e2c67e1d07f5a927fd7e913))
* retry Discord voice joins and use host networking for Arlo ([e0e9cfc](https://github.com/CamiloValderruten/faultline/commit/e0e9cfc52bffac2412b04b18ae7bf730b5265d0c))
* **sandbox:** add html folder alias for publish root ([60ddd74](https://github.com/CamiloValderruten/faultline/commit/60ddd7472b4065cae67c59d87d494ea4b485bca1))
* **sandbox:** add html folder for HTML publishing ([e7934d1](https://github.com/CamiloValderruten/faultline/commit/e7934d1a90402bb0d7e152360e1d6a347e22ef52))
* scrub invalid tool-call JSON instead of crash-looping ([6fcbb78](https://github.com/CamiloValderruten/faultline/commit/6fcbb78e8bfe4bbe257d13fdff235455ae96519a))
* scrub invalid tool-call JSON instead of crash-looping ([a7ee5ee](https://github.com/CamiloValderruten/faultline/commit/a7ee5ee0feca6dde18da5bf32b052cf6782704bd))
* serialize Discord voice joins to stop instant disconnects ([6fc20b9](https://github.com/CamiloValderruten/faultline/commit/6fc20b9ad2731bb2f6e5911dd3f54dcf24fe8781))
* shutdown handling and reconcile observability ([#21](https://github.com/CamiloValderruten/faultline/issues/21)) ([81cc3b6](https://github.com/CamiloValderruten/faultline/commit/81cc3b6e3d147c122cd06f5072e310a561406ea7))
* **skills:** ban send_file for html-artifact canvases ([e4e1dea](https://github.com/CamiloValderruten/faultline/commit/e4e1dea54080bf8b5073dc96272ee5bc46f826e4))
* **skills:** ban send_file for html-artifact canvases ([13ced25](https://github.com/CamiloValderruten/faultline/commit/13ced25ba0b62fe072bd8005d125535f8c0f67c7))
* **skills:** require Discord link button for every canvas ([7f1948d](https://github.com/CamiloValderruten/faultline/commit/7f1948de6e9d21430909ab4eebde1cd821ca7849))
* **skills:** require Discord link button for every canvas ([617cac5](https://github.com/CamiloValderruten/faultline/commit/617cac505f2a737d175c1fa415e84a138bf921a1))
* stabilize Discord voice DAVE handshake ([2b50d9e](https://github.com/CamiloValderruten/faultline/commit/2b50d9e38ef4b17cd4f5b185eef8f65fe5a9a806))
* surface DAVE opcodes and avoid premature voice leave ([dbd0c04](https://github.com/CamiloValderruten/faultline/commit/dbd0c04a972a4eaa14be48ac2705490067c79f31))
* **update:** back off polling on GitHub rate-limit responses ([2f570e1](https://github.com/CamiloValderruten/faultline/commit/2f570e1ee7eb41a231edb408c770924860311a7d))
* use DAVE-capable discordgo fork for Discord voice ([c1fc1de](https://github.com/CamiloValderruten/faultline/commit/c1fc1de0dd94ca1d71fbaf6e2e7b31e06952de55))


### Refactors

* hexagonal architecture (ports & adapters) ([c1c3ea3](https://github.com/CamiloValderruten/faultline/commit/c1c3ea350e998251cfc3b8c71c67afd983862019))
* **tools:** extract HTML-to-markdown extraction to its own file ([4a634ab](https://github.com/CamiloValderruten/faultline/commit/4a634abf484a5e560a5a7d237ed773381050e0e5))


### Miscellaneous Chores

* tidy release-please config and document major-bump policy ([793f6ee](https://github.com/CamiloValderruten/faultline/commit/793f6ee50a7e21a78fe7c32db6537b82ebc5c276))

## [1.7.0](https://github.com/matjam/faultline/compare/v1.6.0...v1.7.0) (2026-05-02)


### Features

* **mcp:** support sandboxed and streamable servers ([#35](https://github.com/matjam/faultline/issues/35)) ([5350f78](https://github.com/matjam/faultline/commit/5350f7848306f45e5197e9b7eb325ff5a75d78b5))


### Bug Fixes

* **ci:** set GH_REPO so gh release commands skip git detection ([7ebc3e5](https://github.com/matjam/faultline/commit/7ebc3e5843ca4e87028addc74db9f1c723faea94))
* **ci:** set GH_REPO so release-notes step works without checkout ([#38](https://github.com/matjam/faultline/issues/38)) ([7ebc3e5](https://github.com/matjam/faultline/commit/7ebc3e5843ca4e87028addc74db9f1c723faea94))

## [1.6.0](https://github.com/matjam/faultline/compare/v1.5.0...v1.6.0) (2026-05-01)


### Features

* **admin:** in-place config editor with validate / save / restart ([#34](https://github.com/matjam/faultline/issues/34)) ([ec1608e](https://github.com/matjam/faultline/commit/ec1608ec02382ca2b61a7df515e7e3669c1c3a93))
* **admin:** live agent inspector + tool-call feed on dashboard ([dcb252c](https://github.com/matjam/faultline/commit/dcb252c2a4739ab832b66cf5345ea273a7ba0cc0))
* **admin:** per-skill enable/disable persisted to skills.toml ([#32](https://github.com/matjam/faultline/issues/32)) ([9c842bc](https://github.com/matjam/faultline/commit/9c842bc4fd9ce9f6f2cc820b287ea4921e5c93b7))
* **admin:** scaffold HTTP admin UI with login + session auth ([2bcba17](https://github.com/matjam/faultline/commit/2bcba17b966eac0ec13cc5bcba1c67fadf8e533f))
* **admin:** self-update card with apply-now button ([#33](https://github.com/matjam/faultline/issues/33)) ([f61e87a](https://github.com/matjam/faultline/commit/f61e87a3f2e866e7e1f8379933e92b79c1c5d30c))
* untrusted-tool-output guard + prompt-migration delivery ([#36](https://github.com/matjam/faultline/issues/36)) ([98819f1](https://github.com/matjam/faultline/commit/98819f114d92d2ceb3174e9664f1bd5147c7553c))


### Bug Fixes

* **update:** back off polling on GitHub rate-limit responses ([2f570e1](https://github.com/matjam/faultline/commit/2f570e1ee7eb41a231edb408c770924860311a7d))

## [1.5.0](https://github.com/matjam/faultline/compare/v1.4.0...v1.5.0) (2026-05-01)


### Features

* Agent Skills support ([#24](https://github.com/matjam/faultline/issues/24)) ([eaad5a0](https://github.com/matjam/faultline/commit/eaad5a0eb8d43f80c709b97e474f10fe3def3fad))
* **skills:** security audit subagent + sandbox hardening ([#27](https://github.com/matjam/faultline/issues/27)) ([93515d5](https://github.com/matjam/faultline/commit/93515d58b3dc86f5666f8c9bea6033956704fd7a))
* subagent delegation ([#26](https://github.com/matjam/faultline/issues/26)) ([c42b698](https://github.com/matjam/faultline/commit/c42b6988b0e537d6e647467acdefef81d9e21b53))

## [1.4.0](https://github.com/matjam/faultline/compare/v1.3.0...v1.4.0) (2026-05-01)


### Features

* custom Arch-based multi-runtime sandbox image ([#23](https://github.com/matjam/faultline/issues/23)) ([41bee48](https://github.com/matjam/faultline/commit/41bee48f1bd97c14e76ea5acdc3a97bc3186667d))


### Bug Fixes

* shutdown handling and reconcile observability ([#21](https://github.com/matjam/faultline/issues/21)) ([81cc3b6](https://github.com/matjam/faultline/commit/81cc3b6e3d147c122cd06f5072e310a561406ea7))

## [1.3.0](https://github.com/matjam/faultline/compare/v1.2.0...v1.3.0) (2026-05-01)


### Features

* paragraph-aligned semantic indexing with adaptive batching ([#19](https://github.com/matjam/faultline/issues/19)) ([e5d7918](https://github.com/matjam/faultline/commit/e5d7918539d43b80f499101538c171c237e1bcc5))

## [1.2.0](https://github.com/matjam/faultline/compare/v1.1.0...v1.2.0) (2026-05-01)


### Features

* add rebuild_indexes tool and refactor reconcile/rebuild helpers ([27aea57](https://github.com/matjam/faultline/commit/27aea575d51f7509af3697e878a81b75f40f03e9))
* add semantic memory search with persisted vector index ([ff012dc](https://github.com/matjam/faultline/commit/ff012dc0f08ffe7eb62afbf4b2e6bec5d1898fba))
* add semantic memory search with persisted vector index ([6265089](https://github.com/matjam/faultline/commit/62650891ae24f9c555af3a3f129ea94bdbb21225))

## [1.1.0](https://github.com/matjam/faultline/compare/v1.0.0...v1.1.0) (2026-05-01)


### Features

* add self-update support ([cc7090e](https://github.com/matjam/faultline/commit/cc7090eb6c9b6da4601bb3a5c5be132519fcc5d6))
* add self-update support ([1739d61](https://github.com/matjam/faultline/commit/1739d61bab2a551858a099442343315c04e4437f))

## 1.0.0 (2026-05-01)


### ⚠ BREAKING CHANGES

* footer, which is too sharp a knife for a project at this stage. Add 'Ask the maintainer first.' to the table row and a paragraph spelling out the policy: contributors don't unilaterally mark breaking changes; the maintainer accumulates them and decides when to cut a major. Also documents the Release-As: footer for forcing a specific version when needed.

### Features

* add email_fetch tool ([f1b4ab6](https://github.com/matjam/faultline/commit/f1b4ab6f6cde2349247841b29f6845cf33767b30))
* add release pipeline (release-please + goreleaser) ([3075763](https://github.com/matjam/faultline/commit/3075763c74284f2fe9f2447a2f2b5724003e8c89))
* add release pipeline (release-please + goreleaser) ([0a1b86d](https://github.com/matjam/faultline/commit/0a1b86d82e82af0b822eadbb971cddd383ae3217))
* add sandbox_shell tool for arbitrary shell command execution ([40c0922](https://github.com/matjam/faultline/commit/40c0922d017847e2fb8d11a23754c17d84430b15))
* add sandbox_shell tool for arbitrary shell command execution ([efe41cb](https://github.com/matjam/faultline/commit/efe41cb2d7d4f58c13f762437e8b83df9943f964))


### Miscellaneous Chores

* tidy release-please config and document major-bump policy ([793f6ee](https://github.com/matjam/faultline/commit/793f6ee50a7e21a78fe7c32db6537b82ebc5c276))
