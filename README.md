# fddp: Facebook Downloaded Data Processor

Facebook allows you to download "a copy of what you've shared on Facebook".
This includes a lot of Facebook data, including Messenger (chat) messages.

fddp is an experiment about doing some data mining on these Facebook messages.

## Getting your Facebook archive
You can acquire your archive on [Facebook settings page](https://www.facebook.com/settings),
it is at the end under "Download a copy of your Facebook data", then follow the steps to save it.

Once you have this zip file, open it and place the `messages.htm` inside the `personal` folder  

---

The coding quality is pretty terrible right now. The aim is getting it working first. I'm pretty bad at Go but I love it

I have crafted a sample (under `samples/sample.html`) that is identical to the format that Facebook uses to distribute the `messages.htm`. This is an easier way to get started then download your FB archive but is way less populated.