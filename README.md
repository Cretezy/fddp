# fddp: Facebook Downloaded Data Processor [![Build Status](https://travis-ci.org/Cretezy/fddp.svg)](https://travis-ci.org/Cretezy/fddp)

Facebook allows you to download "a copy of what you've shared on Facebook".
This includes a lot of Facebook data, including Messenger (chat) messages (what we are interested in)

fddp is an experiment about doing some data mining (or other things) on Facebook messages from your own account.

## Getting your Facebook archive
You can acquire your archive on [Facebook settings page](https://www.facebook.com/settings),
it is at the end under "Download a copy of your Facebook data", then follow the steps to save it.

Once you have this zip file, open it and place the `messages.htm` inside the `personal` folder  

### [Samples](https://github.com/Cretezy/fddp/tree/master/samples)
I have crafted samples for you to use to test.
These are identical to how Facebook's archives distribute their `messages.htm`.

> This is an easier way to get started then download your FB archive but is way less populated (only a few messages).

## Setup
`fddp` uses [godep](https://github.com/tools/godep) for dependency management. It is not required to run.

### Install
- Setup your `$GOPATH`
- `git clone https://github.com/Cretezy/fddp.git && cd fddp && go build`
- Enjoy! Check if everything works with `./fddp`. You must be in the `fddp` directory to run commands.

## Usage

### Web UI
- To start the Web UI, run `./fddp server`
- Visit `http://localhost:3000/`

You can switch the port using the `PORT` environment variable.

### Commandline

#### Convert
Converts a HTML message file (ex: Facebook's `messages.htm`) to JSON.

You must convert your HTML message file to JSON before doing anything with it. It will also clean it.
```
./fddp convert personal/messages.htm personal/messages.json
```

This will turn `personal/messages.htm` to JSON format and save it (under `personal/messages.json`).

You can use `-i` (or `--indent`) to indent (pretty print). This is not recommended on big data set as it adds useless storage bulk
(see [Example File Size](https://github.com/Cretezy/fddp#example-file-size) for increase).


#### Count
Counts threads/messages/words in a data set.

You must input a JSON file (use convert command first). You may use many flags at the same time.
```
./fddp count [flags] input.json
```

| Name     | Flag               |
|----------|--------------------|
| Threads  | `--thread`, `-t`   |
| Messages | `--messages`, `-m` |
| Words    | `-words`, `-w`     |


#### Compare
Shows the difference between 2 data sets (in count, not data).

You must input 2 JSON file (use convert command first).
```
./fddp compare samples/sample.json samples/sample-indent.json
```

#### List
List tops people you have messaged.

You must input a JSON file (use convert command first).
```
./fddp list samples/sample.json
```

Default shows top `50` but using `-c` (or `--count`) followed by a custom number you may change the number of threads displayed.

## Notes

### Purpose
The purpose of this project is to:
- Gain experience (Go, etc)
- Have a better interface to read past Facebook messages
- Analyse Facebook messages as "big data"

### Example File Size 
Sample size of this is from my Facebook, ~350k messages.
Running on an average-high end desktop CPU (4770K) and SSD.

Note: I think that reading the actual file (~36MB) from disk is the main bottleneck for these stats.
I estimate that reading from file (and parsing the JSON) actually takes around 800ms or so,
except the convert which takes a much longer time (parsing HTML/messages is a lot slower) and compare, which needs to open 2 files.

| Command                | Time       |
|------------------------|------------|
| Convert                | ~8s      |
| Convert (with `-i`)    | ~8.5s        |
| Count (all type)       | ~850-950ms |
| Compare (self vs self) | ~1.8s      |
| List                   | ~850-900ms |

| Files Size             | Size   |
|------------------------|-------:|
| `messages.htm`         | 70.3MB |
| `messages.json`        | 36.4MB |
| `messages-indent.json` | 47.1MB |
