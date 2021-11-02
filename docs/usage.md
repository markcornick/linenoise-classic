# Usage

```
$ linenoise-classic -help
Usage of linenoise-classic:
  -digit
        Include digits (default true)
  -length int
        Length to generate (default 16)
  -lower
        Include lowercase letters (default true)
  -upper
        Include uppercase letters (default true)

$ linenoise-classic
h6ECtbDZPnRddHV7
$ linenoise-classic -length=8
XdWod8f8
$ linenoise-classic -length=64
QhESpeyPDidxV9kFNCrJqeMa4XUYbET4B3s5oGA8kYsV6XwDKHrCL7wojGZm9gj5
$ linenoise-classic -length=0
2017/10/31 16:04:47 Invalid length - must be an integer greater than 2
$ linenoise-classic -lower=false
387HNFDEUW4YGMZA
$ linenoise-classic -upper=false -length=8
hcsym6tj
$ linenoise-classic -lower=false -upper=false -length=32
92992759356835354563826487673794
$ linenoise-classic -lower=false -upper=false -digit=false
2021/10/02 21:38:36 must include at least one of -lower, -upper and/or -digit
exit status 1
```
