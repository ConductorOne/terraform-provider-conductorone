# Post-regen patches

Unified-diff patches applied by `make gen` after Speakeasy generates the SDK + provider code, in filename order. Each patch hand-fixes a Speakeasy regression that returns on every regen until upstream resolves it. The patch files themselves carry a commit-message header explaining the regression — read those, not a duplicate index here.

## Adding a patch

1. Apply the change manually after `make gen`, verify it fixes the regression and that any tripwire test in `internal/provider/speakeasy_regen_test.go` flips to passing.
2. `git diff <files...> > patches/NN-<short-description>.patch` (or use `git format-patch` from a commit so the message header is preserved).
3. Confirm a clean re-run: `make gen` should succeed end-to-end.

**Cut the diff against the regen baseline, not against your branch.** `git diff` on a
branch that already carries an earlier attempt produces hunks whose pre-image is your own
intermediate commit — text Speakeasy never emits. Those patches apply and reverse cleanly
within the branch and fail on the next real regen, which nothing in CI exercises: CI runs
`make generate` (tfplugindocs), never `make gen`. Patches 04 and 05 shipped this way and
had to be re-cut.

**A hand-written file in a generated directory belongs in `.genignore`, not in a patch.**
Do not write a patch that re-creates it. `.genignore` already preserves `provider.go`,
`token_source.go` and `extra_sdk_options.go` that way; `update_mask.go` joins them. A
new-file hunk is a bet that regen deletes the file, and it fails with "already exists" the
moment that bet is wrong.

## Removing a patch

When upstream Speakeasy resolves the regression and we bump the pinned CLI past the fix:

1. Delete the patch file.
2. Run `make gen`; the corresponding tripwire test should still pass without the patch (proves the regression is actually gone upstream, not just hidden).
3. Update KNOWN_ISSUES.md and remove the tripwire test if it's now redundant.
