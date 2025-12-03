mapcli
===

`mapcli` is a sort of silly tool that allows you to perform word replacement for any command line tool.

For example, imagine that from time to time you make use of the following command:

```bash
aws s3 list-objects --bucket forgotten-objects --prefix /deep/in/the/depths
```

You might use `mapcli` to do something like this:

1. Create a yaml of words you want to map, like so, starting with your new word, mapping to a word
from the CLI.

```yaml
# /path/to/mapping.yaml

dark-overlords: aws
dragon-hoards: s3
plunder: list-objects
"--hoard": "--bucket"
"--summoning-charm": "--prefix"
```

2. Make sure that `mapcli` is on your path and then run it to create a new CLI mapping, passing in the path to your mapping config like so:

```
mapcli create dark-overlords /path/to/mapping.yaml
```

3. Ensure the right things are on your PATH.

You'll want to add `~/.mapcli/bin` by doing something like:

```bash
echo 'PATH=$PATH:~/.mapcli/bin' >> ~/.bashrc
```

4. Invoke your new command:

```
dark-overlords dragon-hoards plunder --hoard forgotten-objects --summoning-charm /deep/in/the/depths 
```

I did say it was sort of silly.

---

This CLI was primarily a way for me to learn my way around go. I had fun.
