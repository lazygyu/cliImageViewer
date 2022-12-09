# ImageViewer in CLI

A command line image viewer

![Screen shot](https://i.imgur.com/yoTsRhx.png)

## Build

```bash
> go build .
```

## Usage

### Providing a file name to show
```bash
> gv [options] {image_file_path_name}
```

### Through a pipe

```bash
> cat {image_file_path_name} | gv [options]
```

### options

- `-i` : invert the image brightness (optional, default: false)
- `-w {number}` : limit the maximum width in pixels (optional, default: screen width)
- `-h {number}` : limit the maximum height in pixels (optional, default: screen height)

## So, Where should I use this?

### with `fzf`

You can use this to `fzf`'s preview

![fzf screenshot](https://imgur.com/nKbwuLD.gif)

```bash
> ls *.png | fzf --preview "gv -w 100 {}"
```

### insert an image into text contents

```text
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣤⣶⣶⣾⣿⢿⡿⣿⣷⣶⣶⣦⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣶⣿⣶⣾⡿⠟⣋⣡⣤⣖⡶⣞⣞⢷⣛⡾⢶⡳⠶⠦⣬⣉⠛⢿⡿⠿⠿⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢠⣿⠟⣠⠶⠻⠂⣴⠚⣡⣴⣶⣶⣶⣬⡑⢯⣟⡼⢃⣴⣾⣿⣿⣷⣦⡙⢷⣄⠙⠻⣳⡌⢻⣷⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢻⣿⠈⣿⡄⢠⣿⡁⡾⠉⠉⣿⣿⣿⣿⣿⢈⡿⠠⡏⠀⡘⣿⣿⣿⣿⡗⢸⣳⢧⠘⡷⠃⣼⡿⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣶⢀⣯⠿⣥⠘⣶⣶⣿⣿⣿⡿⢋⡼⠋⠧⠙⢷⣿⣿⣿⡿⠟⣡⡟⣶⢏⣇⢸⣿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⢸⣯⢞⣧⣟⣳⣤⣤⣴⢲⠎⣤⣤⣀⣤⣤⣍⠲⣶⡞⣯⡟⣽⣝⡮⣟⢾⠀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⢸⣟⣭⡟⣞⣳⡽⣞⡽⡽⣦⠉⣤⡄⣦⠌⣡⣼⣛⡾⡽⣭⣻⢽⣚⣯⢽⡂⣿⣟⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣏⢸⡳⢿⣹⢶⡻⣝⡾⣣⣟⢧⡷⣈⣤⣌⡴⢯⣛⣧⢿⣹⡞⣧⠷⣾⣹⢾⡁⢿⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣹⣿⢈⡿⣱⢿⣱⠾⣝⡾⡽⢶⣫⡟⣞⣧⠿⣭⠷⣞⣳⢯⣳⢟⣮⡽⣛⣧⢿⡁⢿⣟⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣀⣴⣾⣿⠀⣿⣚⣯⢶⣻⡼⣛⣮⠷⣽⡺⣵⢯⣳⡞⣧⡟⣾⣱⡟⣽⡞⣽⣫⣞⠷⡆⢽⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢸⣿⠁⢶⡿⠀⣯⢷⣫⣟⡼⣏⡾⣝⡾⣝⡾⣵⣻⢼⡻⣼⢻⣼⢻⡵⣛⣾⡹⢷⣹⢧⡇⢸⡗⠈⣿⡧⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠈⠿⣷⣶⣾⠐⣯⢷⣛⡿⡼⣶⢻⣝⡾⣹⠾⣽⢶⡻⣝⡶⣏⣷⢳⡾⣭⢿⣹⡳⣞⢯⡇⢸⣶⣿⠟⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡟⢸⣎⣟⣮⢟⣮⢟⣮⠷⣞⣧⢯⡗⡿⡼⡽⣞⢯⣞⠷⣽⡭⣷⣫⡽⢶⡳⣏⢸⣿⡃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⢸⣏⡿⣺⣝⡾⣣⢿⣹⡞⣯⡽⣝⡾⣎⡿⡵⢯⣾⣹⡳⢯⣻⢳⡞⣧⡟⡧⢸⣿⠄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣏⢸⡞⣯⡽⣞⢯⣞⡽⣳⣏⢷⣻⠼⣏⣾⢫⡾⣭⢟⣳⡞⣧⡟⣧⠿⣼⣫⠇⣸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⠐⣯⢷⣫⣟⢮⡷⣫⣞⣯⡽⣞⡽⢮⣗⢯⣞⢷⡳⣽⢞⣵⡻⣖⣯⢳⡻⢀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣧⠘⣿⣱⢯⣳⢯⣻⡼⣛⣷⢺⢯⡽⣞⢧⡿⣱⣞⢯⣳⣏⢿⣱⢯⡳⢁⣾⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡄⠙⢯⢷⡻⣼⢏⣾⢳⣏⡾⣏⣶⣛⡾⣝⣧⢯⡟⣵⣳⠭⢃⠰⣿⣯⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣟⠰⢩⡿⠓⣠⣌⣉⠛⠓⠯⠾⠽⠽⠾⠽⠞⠮⠓⢋⣉⣤⣦⡙⢿⡉⠄⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⣷⣶⣾⡿⠋⠉⠙⠛⠻⠿⠿⠿⠿⠿⠿⠿⠿⠛⠛⠉⠁⠙⠻⣿⣾⡿⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
```

Make a text image with this and paste it into text contents

```bash
# make an inverted image and copy it
> gv -i {image_file_path_name} | pbcopy
```

### etc
- including a cool logo in your cli command application
- just for fun

