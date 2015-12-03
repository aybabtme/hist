# hist

Plots histograms for values read from stdin.

## usage

```bash
hist -re '(\d+\.\d+)'                      # regexp used to match
     -format (unix|plain|elapsed-seconds)  # func used to pretty-print matches
     -bins 30                              # how many lines for the histogram
     -width 20                             # how wide is the histogram
     preview    # preview what will be matched in the input
     plot       # plot the matches in the input
```

## installation

### linux

```bash
wget -qO- https://github.com/aybabtme/hist/releases/download/0.1/hist_linux.tar.gz | tar xvz
```

### darwin

```bash
wget -qO- https://github.com/aybabtme/hist/releases/download/0.1/hist_darwin.tar.gz | tar xvz
```

## example

```bash
$ cat file.log

[5572220.674985] cron[32806]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5572340.722262] cron[32807]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5572820.910143] cron[32808]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5573421.145055] cron[32809]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5574021.379941] cron[32810]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5574621.614755] cron[32811]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
[5575221.849840] cron[32812]: segfault at 19b10208 ip 00007f105ab49479 sp 00007ffd1ccd5690 error 4 in ld-2.19.so[7f105ab3a000+23000]
# ellided

$ hist < file.log -re '(\d+\.\d+)' plot
5174062.26s-5198083.57s  2.59%  ████████████████████████████████████████▏  49
5198083.57s-5222104.89s  2.54%  ███████████████████████████████████████▏   48
5222104.89s-5246126.21s  2.48%  ██████████████████████████████████████▍    47
5246126.21s-5270147.53s  2.48%  ██████████████████████████████████████▍    47
5270147.53s-5294168.85s  2.48%  ██████████████████████████████████████▍    47
5294168.85s-5318190.17s  2.54%  ███████████████████████████████████████▏   48
5318190.17s-5342211.48s  2.48%  ██████████████████████████████████████▍    47
5342211.48s-5366232.80s  2.48%  ██████████████████████████████████████▍    47
5366232.80s-5390254.12s  2.59%  ████████████████████████████████████████▏  49
5390254.12s-5414275.44s  2.48%  ██████████████████████████████████████▍    47
5414275.44s-5438296.76s  2.43%  █████████████████████████████████████▋     46
5438296.76s-5462318.08s  2.54%  ███████████████████████████████████████▏   48
5462318.08s-5486339.39s  2.54%  ███████████████████████████████████████▏   48
5486339.39s-5510360.71s  2.43%  █████████████████████████████████████▋     46
5510360.71s-5534382.03s  2.48%  ██████████████████████████████████████▍    47
5534382.03s-5558403.35s  2.54%  ███████████████████████████████████████▏   48
5558403.35s-5582424.67s  2.54%  ███████████████████████████████████████▏   48
```


## license

MIT
