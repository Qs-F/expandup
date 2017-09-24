# expandup

A simple html generator along with global config file.

## About config file

Config files are put at `~/.expandup/`.

## Purpose

Riot against present complex frontend.  
What we want to do, is to generate one index file, and use already designed components, and serve it.  
Inspired by Go.

## How to use

What you have to do is only put your comment on `*.html`.

```
<!DOCTYPE html>
<html>
  <head>
    <!-- EXPANDUP RIOT -->
    <!-- EXPANDUP DEFAULT_STYLES -->
  </head>
  <body>
    <!-- EXPANDUP RIOT_IMPLEMENT -->
  </body>
</html>
```

After you wrtie like above, save as `index.html`, and run `expandup index.html`, you could have like this result as rewrited `index.html` as YOUR STDOUT.  
If you want to save it, please put `-w` option before file name. This is `gofmt` style.

```
<!DOCTYPE html>
<html>
  <head>
    <!-- EXPANDUP RIOT -->
    <script src="https://cdn.jsdelivr.net/riot/3.2/riot+compiler.min.js"></script>
    <!-- (END OF EXPANDUP) -->
    <!-- EXPANDUP DEFAULT_STYLES -->
    <style>
@font-face {
  font-family: 'GenShin';
  src: url('//de-liker.global.ssl.fastly.net/fonts/products/GenShinGothic-P-ExtraLight.woff') format('woff');
  font-weight: 100;
}

@font-face {
  font-family: 'GenShin';
  src: url('//de-liker.global.ssl.fastly.net/fonts/products/GenShinGothic-P-Regular.woff') format('woff');
  font-weight: 400;
}

@font-face {
  font-family: 'GenShin';
  src: url('//de-liker.global.ssl.fastly.net/fonts/products/GenShinGothic-P-Heavy.woff') format('woff');
  font-weight: 900;
}

:root {
  --dark: rgb(64, 64, 64);
  --light: rgb(248, 248, 248);
  --white: rgb(255, 255, 255);
  --black: rgb(0, 0, 0);
  --font: "Roboto", "GenShin", sans-serif;
  --font-title: "Roboto Condensed", "GenShin", sans-serif;
  --font-mono: "Roboto Mono", "GenShin", monospace;
}
    </style>
    <!-- (END OF EXPANDUP) -->
  </head>
  <body>
    <!-- EXPANDUP RIOT_IMPLEMENT -->
    <script>
      riot.mount('*')
    </script>
    <!-- (END EXPANDUP) -->
  </body>
</html>
```

## How to configure

Put your config file to your `~/.expandup` directory.  
Config file style is very simple.

**Definition**s

- `COMMAND_NAME` means `<!-- EXPANDUP COMMAND_NAME -->`
- contents means placeholder texts.

**Usage**

Write your contents, and save it as `COMMAND_NAME` in `~/.expandup/`.

And, if you put `.expandup.imports`, expandup can detect it, and include your another expandup config files. You must put directory name, not file name.  
By using this, git integration becomes so easy.

Here is a sample of `.expandup.imports`.

```
github.com/Qs-F/hoge
git.de-liler.com/Qs-F/hoge
```

Any happend errors are ignored. If you have some trouble and want to see log, use `-v` option.

## Auto update

head-up is easy. once your COMMAND file is updated, and re-run expandup on the same file, you can update it.
