# history-wallpaper

"History Wallpaper" - a project that uses search queries from your browser history to find and set images
as your wallpaper. Automatically update your desktop background with relevance to your interests.

## Dependencies

- Wayland
- Chromium\Firefox-based browser
- Utils swaybg\wbg
- Internet connection
- Go (for building)

## Install

```sh
make install
```

## Usage

```
hw
```

## Examples of usage

```
Usage of hw:
  -browser string
        browser to use. Available: [vivaldi chrome chromium brave opera firefox] (default "vivaldi")
  -follow string
        follow a time interval and update wallpaper. e.g. 1h, 1m, 30s
  -history-file string
        browser history file to use. Auto detect if empty (only for chromium based browsers)
        e.g ~/.mozilla/icecat/gxda4hpz.default-1672760493248/formhistory.sqlite
  -resolution string
        resolution to use. e.g. 1920x1080 (default "1920x1080")
  -save-image-dir string
        directory to save image to (default "/home/labile/Pictures")
  -search-phrase string
        search phrase to use
  -wp-api string
        wallpaper api to use. Available: [unsplash] (default "unsplash")
  -wp-tool string
        wallpaper tool to use. Available: [swaybg wbg] (default "swaybg")
```

Update your wallpaper every hour with wbg manager:

```
hw -wp-tool wbg -follow 1h
```

Update your wallpaper every 30 minute with swaybg manager:

```
hw -follow 30m
```

Use ghw as dynamic wallpaper:

```
hw -follow 1h -search-phrase space
```

Direct path to browser history file (firefox-based)
```
hw -browser firefox -history-file ~/.mozilla/icecat/gxda4hpz.default-1672760493248/formhistory.sqlite
```

Direct path to browser history file (chromium-based)
```
hw -browser vivaldi -history-file ~/.config/vivaldi/Default/History
```

Add sway autostart:

```
exec hw [options]
```

## TODO (maybe never)

- [x] Add swaybg tool
- [ ] Add logger
- [x] Add wbg tool
- [x] Unsplash image provider
- [x] Chromium based browser support
- [x] Firefox support
- [ ] Resolution auto detection
- [ ] Google image provider without API key
- [ ] Pinterest?