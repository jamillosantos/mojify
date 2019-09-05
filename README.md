# mojify

`mojify` receives input from the stdin or execute a command replacing all github
emojis by its unicode. Hence, the emoji is rendered.

![Log](https://user-images.githubusercontent.com/142870/64313563-68a73f00-cf82-11e9-9189-7a363e645263.png)

## How to Use

**"Piping"**

    $ git log --oneline | mojify

**As a command**

    $ mojify git log --oneline

## Limitations

* It does not support colors yet.
* Aliases are not understood.

## Installation

You can use [gobin](https://github.com/myitcv/gobin) to make the installation:

```sh
gobin github.com/jamillosantos/mojify
```

## Thanks to

I would like to thank https://github.com/kyokomi/emoji. I grab the emoji codebase from them.

Many thanks!