# fddp: Facebook Downloaded Data Processor [![Build Status](https://travis-ci.org/CraftThatBlock/fddp.svg)](https://travis-ci.org/CraftThatBlock/fddp)

Facebook allows you to download "a copy of what you've shared on Facebook".
This includes a lot of Facebook data, including Messenger (chat) messages (what we are interested in)

fddp is an experiment about doing some data mining (or other things) on Facebook messages from your own account.

## Getting your Facebook archive
You can acquire your archive on [Facebook settings page](https://www.facebook.com/settings),
it is at the end under "Download a copy of your Facebook data", then follow the steps to save it.

Once you have this zip file, open it and place the `messages.htm` inside the `personal` folder  

### [Samples](https://github.com/CraftThatBlock/fddp/tree/master/samples)
I have crafted samples for you to use to test.
These are identical to how Facebook's archives distribute their `messages.htm`.

> This is an easier way to get started then download your FB archive but is way less populated (only a few messages).

## Setup
`fddp` uses [godep](https://github.com/tools/godep) for dependency management. It is not required to run.
You may install it using `make deps`

### Install
You must first get the project.
```
git clone https://github.com/CraftThatBlock/fddp.git && cd fddp
```

You can then build it and install it to `$GOPATH/bin` (must be in `$PATH`)
```
go build && go install
```

## Usage

### Convert
You must convert your HTML message file to JSON before doing anything with it. It will also clean it.
```
fddp convert personal/messages.htm personal/messages.json
```

This will turn `personal/messages.htm` to JSON format and save it (under `personal/messages.json`).

You can use `-i` (or `--indent`) to indent (pretty print). This is not recommended on big data set as it adds useless storage bulk
(see [Example File Size](https://github.com/CraftThatBlock/fddp#example-file-size) for increase).


### Count
You must input a JSON file (use convert command first). You may use many flags at the same time.
```
fddp count [flags] input.json
```

| Name     | Flag               |
|----------|--------------------|
| Threads  | `--thread`, `-t`   |
| Messages | `--messages`, `-m` |
| Words    | `-words`, `-w`     |


## Notes

#### Example File Size 
Sample size of this is from my Facebook, ~350k messages.
Running on an average-high end desktop CPU (4770K) and SSD.

| Command             | Time       |
|---------------------|------------|
| Convert             | ~7.5s      |
| Convert (with `-i`) | ~8s        |
| Count (all type)    | ~850-950ms |

| Files Size             | Size   |
|------------------------|-------:|
| `messages.htm`         | 70.3MB |
| `messages.json`        | 36.4MB |
| `messages-indent.json` | 47.1MB |