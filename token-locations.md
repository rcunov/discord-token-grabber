# Where to find Discord tokens
Discord keeps the session token in browser local storage, not with cookies. Even though the Discord app is an Electron app, which is based on Chromium, the Discord desktop app stores the token encrypted on disk in a folder separate from the rest of the browser's local storage. This separate folder is `$env:APPDATA\discord\` whereas the rest of Chromium browser local storage would be something like `$env:LOCALAPPDATA\BraveSoftware\Brave-Browser\User Data\`. With Firefox, it stores the session token unencrypted with Firefox local storage (probably - more on that in the Firefox section).

## Discord app / Chromium
These are the paths to the Discord desktop app local storage files I mention in the [README](README.md).
```
$env:APPDATA\discord\Local Storage\leveldb\*.ldb
$env:APPDATA\discord\Local Storage\leveldb\*.log
```

## Firefox
I haven't written the code to support Firefox yet, but I've dug into it and it's pretty silly. Here's the path where I found my Discord token in Firefox local storage:
```
$env:APPDATA\Mozilla\Firefox\Profiles\*.default-release\storage\default\https+++discord.com\ls\data.sqlite
```
Take note of the `*.default-release` part of the path. The local storage is segmented based on browser profile, so that part of the path needs to match whatever browser profile you want to grab the Discord tokens from. I assumed that you'd want to pull from the default profile or you could iterate through all of them too. 

The plaintext tokens are stored in that `data.sqlite` database within a table called `data`. You can get the token from that database with a simple SELECT query: `SELECT value FROM data WHERE key='token';`

The tokens in this Firefox SQLite DB are not encrypted on my machine. On my machine, I'm logged in to my main account with the desktop app and Brave, which is my primary browser. On Firefox, I'm logged in to an alt account. Since the Firefox tokens are not in the separate Discord folder that has my main account data, my thinking is that there are two possibilities: 
1. Tokens are stored unencrypted in Firefox local storage (not in the separate folder that Discord typically uses) when you log in with Firefox. This is what I'm leaning to since cookies are also stored unencrypted in Firefox.
2. Since I was first logged in with my main account, Discord saw there were already login tokens in the separate folder and since it was already in use, stored the session tokens in Firefox local storage as a fallback. 

I could test this in a VM, but I haven't done that yet. Todo, I guess.
