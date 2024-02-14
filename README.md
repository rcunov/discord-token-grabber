# discord-token-grabber
This is intended to be a clear, easy to understand example of a Discord token grabber.

## Explanation / me whining about stuff
I wanted to challenge myself by writing some malware and thought a Discord token grabber would be a decent (i.e., fairly easy) place to start. I looked online for some examples to get me pointed in the right direction and could not find a _simple, straightforward example_ of how to get Discord tokens from a standard installation of the Discord desktop app on the local machine. I quickly learned the overall process, which isn't too complicated: 
- Look through the `leveldb` folder on the machine (for Windows, that's `$env:APPDATA\discord\Local Storage\leveldb`)
- Do a regex search on all `.ldb` or `.log` files in that folder (not sure if the `.log` files ever have the token, but can't hurt to look anyways) that returns the encrypted token values
- Decrypt the token values

Sounds simple enough, right? Yeah, I wouldn't be having a tantrum in this README if that were true. A big problem I encountered was that the code for other Discord token grabber projects I used for inspiration was either not documented at all, intentionally obfuscated (it is malware so that's fair I guess), or both. So trying to parse out the logic was much more difficult than I hoped it would be.

### Regex search
Each example I found used some cryptic regex search pattern to look through the `.ldb` files and that regex wouldn't work for me (I tried opening the file in VS Code and searching with that pattern - no joy). I think this was because the examples I found were in Python, C#, Javascript, and C and each language treats regex just a tiny bit differently (or the projects were just outdated). I found one [project](https://github.com/moonD4rk/HackBrowserData) that stole browser data written in Go, but it was **far** too complicated for my use case and trying to rip code from it took way more effort than it was worth. So I had to make my own regex that seems to get the encrypted token values from those files. It's much more simple than all the examples I've seen, which concerns me, but it works with my own Discord tokens so if it ain't broke don't fix it.

### Decrypting the token values
The tokens are encrypted at rest, so we need to decrypt them. Here is the encryption process we need to reverse:
1. The tokens are AES-GCM encrypted 
2. The decryption key is then itself encrypted with [DPAPI](https://en.wikipedia.org/wiki/Data_Protection_API)
3. The encrypted decryption key is Base64 encoded
4. The encoded, encrypted decryption key is stored as JSON under `os_crypt.encrypted_key` in the Local State file (a file common to Chromium browsers. For the Discord desktop app, this Local State file is located at `$env:APPDATA\discord\Local State`).

You have to be super careful here with these encrypted token values. Messing with a single byte in the sequence will give an error with the decryption process. That means you have to get the Base64 encoded decryption key exactly, the encrypted token values exactly, and if anything goes wrong you'll have precisely zero idea what caused it.

### Detailed decryption process
This seems rude, but honestly just read the code if you want to better understand the decryption process in practice. This README explains the general process in pretty good detail, but there's a bit more complexity to the actual implementation - namely that you have to cut up some of the encrypted byte sequences in a specific way. I documented the code pretty well so if you've gotten this far reading the code shouldn't be hard.

## Usage
Currently this project only runs on Windows, but hoping to add more functionality over time.

### Normal way
Do `go build .` and `.\discord-token-grabber.exe` from a CMD or Powershell window to print out the Discord token for your installation of the Discord desktop app. 

### Sneaky way
`go build -ldflags -H=windowsgui -o silent.exe .` will build an .exe called `silent.exe` that won't spawn any windows when a user clicks on the .exe to run it. It tricks the machine into thinking it's a GUI app even if it doesn't have a GUI so it hides terminal windows. Useful if someone were to have changed the logic in the `main()` function to send the token back to some C2 infra and didn't want to spook a user with temporary terminal windows.