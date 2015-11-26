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

This is an easier way to get started then download your FB archive but is way less populated (only a few messages).

## Setup
`fddp` uses [godep](https://github.com/tools/godep) for dependency management. It is not required to run.
You may install it using `make deps`

### Install
You must first get the project.
```
git clone https://github.com/CraftThatBlock/fddp.git
cd fddp
```
### Convert
You must convert your Html message file to Json before doing any research.
```
go build
./fddp convert personal/messages.htm personal/messages.json
```

This will turn `personal/messages.htm` to Json format and save it (under `personal/messages.json`).

You can use `-i` to indent (pretty print). This is not recommended on big dataset as it adds useless storage bulk.

## Notes
Running with ~350k messages from my Facebook runs in about 6-7 seconds on an average-high end desktop computer (with SSD).

#### Example File Size
| Files                  | Size   |
|------------------------|-------:|
| `messages.htm`         | 70.3MB |
| `messages.json`        | 36.4MB |
| `messages-indent.json` | 47.1MB |