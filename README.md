# gbash
Bash wrapper shell to run shell inside systemd pam session

This a temporary hack to enable systemd user session in WSL when using [wsl-distrod](https://github.com/nullpo-head/wsl-distrod/).

# Notice: Official microsoft support for systemd is available from WSL version 0.67.6
Here is the official [announcement article](https://devblogs.microsoft.com/commandline/systemd-support-is-now-available-in-wsl/). If you are running windows 11 and able to upgrade WSL to latest versions, I highly recommend using the official wsl systemd. I'd recommend using this only for **windows 10** users.

# Usage

`make config` will compile, install it to `/bin/gbash`, then configure shell using `chsh` and run `distrod enable`.
`make unconfig` to undo.

After running `make config` open a new terminal tab, you should be placed inside a systemd user session enabled shell. To verify

```bash
$ env | grep XDG
XDG_SESSION_TYPE=tty
XDG_SESSION_CLASS=user
XDG_SESSION_ID=c53
XDG_RUNTIME_DIR=/run/user/1000
```

You can see XDG env variables are set. You can also run `systemctl --user status`.



# Known Issue

1. ~~It fails when terminal is launched for the first time, as it can't find systemd.You can close the failed and open new tabs. It works from now on.~~
No longer an issue, workaround implemented.
