# `fddp`: Facebook Downloaded Data Processor

Facebook allows you to download "a copy of what you've shared on Facebook".
This includes a lot of Facebook data, including Messenger (chat) messages.

fddp is an experiment about doing some data mining on these Facebook messages.

## Getting your Facebook archive
You can acquire your archive on [Facebook settings page](https://www.facebook.com/settings),
it is at the end under "Download a copy of your Facebook data", then follow the steps to save it.

Once you have this zip file, open it and place the `messages.htm` inside the `personal` folder  

> I have crafted a sample (under `samples/sample.html`) that is identical to the format that Facebook uses to distribute the `messages.htm`. This is an easier way to get started then download your FB archive but is way less populated.

## Setup
`fddp` uses [godep](https://github.com/tools/godep) for dependency management. It is not required to run.
You may install it using `make deps`

### Clone
```
git clone https://github.com/CraftThatBlock/fddp.git
cd fddp
```
### Convert
You must convert your Html message file to Json before using it.
```
go build
./fddp convert samples/sample.html samples/sample.json
```
This will turn samples.html to Json format and save it. You can use `-i` to indent (pretty print).

## Notes
Running with ~350k messages from my Facebook runs in about 6-7 seconds on an average-high end desktop computer (with SSD).