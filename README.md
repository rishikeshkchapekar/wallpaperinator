# WALLPAPERINATOR

### A simple script for linux based systems that changes the desktop background based on the user's interests

## Usage

1. Currently uses Unsplash API. Theme has been hardcoded to Star Wars. Change the query parameter in the API call to fit your needs
2. Strictly meant for personal use only. 
3. Built for Gnome desktops, so to use for other types, you'll have to alter the code on this line: `cmd = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", arg)`
4. The Unsplash API key is stored as an environment variable named `UNSPLASH_API`. Make sure you have this environment variable set before running the code
5. No additional setups required. Simply run the `main.go` file, and it'll do its thing

## Contributions

1. Contributions to support other OS are welcome. When doing so, create a new folder and put all source files for that build in there