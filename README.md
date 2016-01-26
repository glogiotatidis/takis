# Read me


Do you need to wait to MySQL to initialize before you run your tests
against it? Do you have `sleep` cmds with random wait values in your
scripts?

Takis waits for a port to open and then quits, giving control to you
again. Takis is a *completely* statically linked go program. Takis is
simple. Takis does the job. Be like Takis.


## Use it


Takis can be used directly

```
wget https://github.com/glogiotatidis/takis/raw/master/bin/takis
chmod +x takis
CHECK_PORT=80 CHECK_HOST=google.com ./takis
```

or from its docker image

```
docker run -e CHECK_PORT=80 -e CHECK_HOST=google.com giorgos/takis
```

Environment variables:

* **CHECK_PORT** required
* **CHECK_HOST** optional, defaults to "localhost"

## Footnotes

Thanks
[Adriaan](http://blog.adriaandejonge.eu/)
for this [great blog post](http://blog.adriaandejonge.eu/2014/07/05/xebia-blog-create-smallest-possible.html).
