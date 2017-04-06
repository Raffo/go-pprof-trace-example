# Go trace example

This repo contains just an example on how to use the Go tracing tool. It aims at being short, simple, with clear commands and a short piece of code. It also contains some links to some more details articles. 

## Profiling 

Profiling based on CPU: 

```
go tool pprof http://127.0.0.1:8080/debug/pprof/profile
```

Profiling based on heap:

```
go tool pprof http://127.0.0.1:8080/debug/pprof/heap
```

NOTE: use the `-seconds` option to limit the time of tracing. Take into account that if you use read or write timeout for your application, this will also limit your capabilities of getting the profiling data. 

Interesting option for memory profiling: 

```
Sample value selection option (for heap profiles):
  -inuse_space      Display in-use memory size
  -inuse_objects    Display in-use object counts
  -alloc_space      Display allocated memory size
  -alloc_objects    Display allocated object counts
```

When launched an interactive mode, use `top` to see the top entries and `web` to get a nice graph in your browser. 
For more information, a more in detail article: 

http://artem.krylysov.com/blog/2017/03/13/profiling-and-optimizing-go-web-applications/

## Tracing 

Get the tracing data: 

```
curl http://localhost:8080/debug/pprof/trace?seconds=5 > trace.out
```

Start the tracing interface: 

```
go tool trace -http=':8081' trace http://localhost:8080/debug/pprof/trace
```

NOTE: the -http option specifies the address to use for the Web API for the tracing tool. For more information on the tool, read this interesting article: https://making.pusher.com/go-tool-trace/

