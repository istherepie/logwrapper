# Logwrapper

This is a simple log wrapper around the standard log library. 
(I use it as logging boilerplate!)

## Install

```
    go get github.com/istherepie/logwrapper
```

## Usage

```
    logwrapper.Trace("Tracing stuff...")
    logwrapper.Debug("Some logs for the ops guys...")
    logwrapper.Info("Informing everyone about things...")
    logwrapper.Warning("I am soooo warning you!!!!")
    logwrapper.Error("This really does not work!")

```

Please see the examples in the `examples` directory. 

```
    go run examples/app.go
```


## License

MIT © Steffen Park <dev@istherepie.com>