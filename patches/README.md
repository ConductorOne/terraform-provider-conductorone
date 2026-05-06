# Post-regen patches

Unified-diff patches applied by `make gen` after Speakeasy generates the SDK + provider code, in filename order. Each patch hand-fixes a Speakeasy regression that returns on every regen until upstream resolves it. The patch files themselves carry a commit-message header explaining the regression — read those, not a duplicate index here.

## Adding a patch

1. Apply the change manually after `make gen`, verify it fixes the regression and that any tripwire test in `internal/provider/speakeasy_regen_test.go` flips to passing.
2. `git diff <files...> > patches/NN-<short-description>.patch` (or use `git format-patch` from a commit so the message header is preserved).
3. Confirm a clean re-run: `make gen` should succeed end-to-end.

## Removing a patch

When upstream Speakeasy resolves the regression and we bump the pinned CLI past the fix:

1. Delete the patch file.
2. Run `make gen`; the corresponding tripwire test should still pass without the patch (proves the regression is actually gone upstream, not just hidden).
3. Update KNOWN_ISSUES.md and remove the tripwire test if it's now redundant.
