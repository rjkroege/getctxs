Recent versions of MacOS offer something called "Focus": [Set up a
Focus on Mac - Apple Support
(CA)](https://support.apple.com/en-ca/guide/mac-help/mchl613dc43f/mac).

This repository contains a simple tool for [Alfred - Productivity App
for macOS](https://www.alfredapp.com/) to read the set of defined
Focus modes and provide that as an autocomplete-capable list. I used
this to create an Alfred workflow to set my current focus.

# Implementation Notes
Apple doesn't seem to currently offer a pubic API to interrogate or
control the Focus mode. I believe that the internal component for this
is called `DoNotDisturb`. This has a private framework. Internet
searches suggest that the available Focus modes are listed in a JSON
file in `~Library/DoNotDisturb/DB/ModeConfigurations.json`. If there's
a nicer way to get the Focus mode than reading this file, feel free to
suggest it in an issue.

# Build

```
go install github.com/rjkroege/getctxs
```

should probably do the trick to stick it in your default `GOBIN`.
