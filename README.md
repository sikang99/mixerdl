# mixerdl
note: this currently just downloads the raw mp4 file, speed improvements could be made by parallel downloading the m3u8 playlist and then concatenating with ffmpeg.
## Download

Go to the releases folder and look for the current build. Download `mixerdl.exe` if you are on Windows and `mixerdl` if want to use it on MacOS/Ubuntu.

## Requirements

No dependencies as long as you use the built binary. If you want to compile from source, you must run `go build mixerdl.go` or `go run mixerdl.go`

## Usage

You have to call mixerdl from the console.

Calling options:

- -url `-url="https://mixer.com/Kabby?vod=WVKDcVRHNEOFt3o7H0-l5g"` enter the url you want to download the vod from (e.g., https://mixer.com/Kabby?vod=WVKDcVRHNEOFt3o7H0-l5g)

You may have to mark `mixerdl` as an executable on your machine. If on MacOS or Linux, run `chmod +x mixerdl`.
