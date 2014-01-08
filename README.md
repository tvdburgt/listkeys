# listkeys

This program can be useful for recovering a key that is partially obfuscated. I
originally designed this to recover a Windows product key. This was done by
enumerating all possible keys and feeding them into a key checker (e.g., [The
Ultimate PID Checker](http://janek2012.eu/ultimate-pid-checker/)). The
applicability of `listkeys` is very limited and highly depends on the
readability of the key you want to recover and speed of veryifing each possible key.

## Installation

```bash
$ go get github.com/tvdburgt/listkeys
```

Or install the [http://bit.ly/listkeys](binary) directly, if you don't have Go installed.

## Usage 

`listkeys` works with **globbing**. The wildcard metacharacter `?` can be used
to substitute unidentified characters in the key:

```bash
$ listkeys HM7DF-G8XWM-J2VRG-4M3C4-GR27?
HM7DF-G8XWM-J2VRG-4M3C4-GR27A
HM7DF-G8XWM-J2VRG-4M3C4-GR27B
HM7DF-G8XWM-J2VRG-4M3C4-GR27C
...
HM7DF-G8XWM-J2VRG-4M3C4-GR278
HM7DF-G8XWM-J2VRG-4M3C4-GR279
HM7DF-G8XWM-J2VRG-4M3C4-GR270
```

By default, wildcards are expanded to `[A-Z0-9]`. This can be narrowed down by
using the global charset flag `-c`. For example:

```bash
$ listkeys -c ASDF HM7DF-G8XWM-J2VRG-4M3C4-GR27?
HM7DF-G8XWM-J2VRG-4M3C4-GR27A
HM7DF-G8XWM-J2VRG-4M3C4-GR27S
HM7DF-G8XWM-J2VRG-4M3C4-GR27D
HM7DF-G8XWM-J2VRG-4M3C4-GR27F
```

The charset space can also be narrowed down for individual characters. This is
done with Bash-like **brace expansion**:

```bash
$ listkeys HM7DF-G8XWM-J2VRG-{4HML}M3C4-GR27{POQ}
HM7DF-G8XWM-J2VRG-4M3C4-GR27P
HM7DF-G8XWM-J2VRG-4M3C4-GR27O
HM7DF-G8XWM-J2VRG-4M3C4-GR27Q
HM7DF-G8XWM-J2VRG-HM3C4-GR27P
HM7DF-G8XWM-J2VRG-HM3C4-GR27O
HM7DF-G8XWM-J2VRG-HM3C4-GR27Q
HM7DF-G8XWM-J2VRG-MM3C4-GR27P
HM7DF-G8XWM-J2VRG-MM3C4-GR27O
HM7DF-G8XWM-J2VRG-MM3C4-GR27Q
HM7DF-G8XWM-J2VRG-LM3C4-GR27P
HM7DF-G8XWM-J2VRG-LM3C4-GR27O
HM7DF-G8XWM-J2VRG-LM3C4-GR27Q
```

Obviously, this last method is preferred, as it greatly reduces the number of
possible keys, making it much more feasible for a reduced brute-force search.
