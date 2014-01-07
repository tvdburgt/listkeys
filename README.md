# permutekeys

This program can be useful for recovering a key that is partially obfuscated. I
originally designed this to recover a Windows product key. This was done by
enumerating all possible keys and feeding them into a key checker (e.g., [The
Ultimate PID Checker](http://janek2012.eu/ultimate-pid-checker/)).

## Installation

```bash
$ go get github.com/tvdburgt/permutekeys
```

## Usage 

If there's nothing to permute, the only possiblity is listed:

```bash
$ permutekeys HM7DF-G8XWM-J2VRG-4M3C4-GR27Z
HM7DF-G8XWM-J2VRG-4M3C4-GR27Z
```

It gets more interesting when a glob (using the wildcard character `?`) is
provided:

```bash
$ permutekeys HM7DF-G8XWM-J2VRG-4M3C4-GR27?
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
$ permutekeys -c ASDF HM7DF-G8XWM-J2VRG-4M3C4-GR27?
HM7DF-G8XWM-J2VRG-4M3C4-GR27A
HM7DF-G8XWM-J2VRG-4M3C4-GR27S
HM7DF-G8XWM-J2VRG-4M3C4-GR27D
HM7DF-G8XWM-J2VRG-4M3C4-GR27F
```

The charset space can also be narrowed down for individual characters. This is
done with Bash-like brace expansion:

```bash
$ permutekeys HM7DF-G8XWM-J2VRG-{4HML}M3C4-GR27{POQ}
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
permutations, making it much more feasible for a reduced brute-force search.
