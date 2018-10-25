# IsHeadless

This little Go package provides a function `isheadless.IsHeadless() bool`.
This function tells whether the Go program is running in a headless environment.

## Examples
On Windows, `isheadless.IsHeadless()` always returns `false`.
(Note that this is not correct as it ignores the edge case of running on Windows 10 IoT Core headless.)

On all other systems, `isHeadless.IsHeadless()` returns `false` if an environment variable named `DISPLAY` exists.

## Limitations

- Currently, this function always returns `false` on Windows.
  The fact that Windows 10 IoT Core can be headless is ignored and unsupported.
- Currently, this function does not support Android.
  On Android, it will always return `true`, no matter whether this is headless Android or "normal" Android.
- Currently, this function always inspects the `DISPLAY` variable on all other operating systems, assuming X11 or an equivalent.
  This is wrong on those operating systems which provide a UI through means independent of a `DISPLAY` environment variable.
