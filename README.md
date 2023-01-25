# Go experiment: Does `signal.Notify()` block default signal handling?

## Usage

```shell
make run
```

## Result

Yes, it disables the default signal handler for any signals you register with `Notify()`.
Watch as Ctrl+C no longer terminates the application:

```
$ make run
go build -o troll .
./troll
^C2023/01/25 17:37:45 Signal received: interrupt
^C2023/01/25 17:37:46 Signal received: interrupt
^C2023/01/25 17:37:48 Signal received: interrupt
^C2023/01/25 17:37:49 Signal received: interrupt
^C2023/01/25 17:37:49 Signal received: interrupt
^C2023/01/25 17:37:50 Signal received: interrupt
```

Pressing Ctrl+Z to put it in the background and then running `kill %1` will defeat the troll.
