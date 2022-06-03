# gbash
Bash wrapper shell to run shell inside systemd pam session

This a temporary hack to enable systemd user session in WSL when using [wsl-distrod](https://github.com/nullpo-head/wsl-distrod/).

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

`PATH` env passed from windows gets lost. As a workaround I placed them in `WINDOWS_PATH` env var. You can load them from `.bashrc`. I have this snippet at the start of `.bashrc`:

```bash
[ -z "$WINDOWS_PATH" ] || export PATH=$PATH:$WINDOWS_PATH
```