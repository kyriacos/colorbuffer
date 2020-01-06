# ColorBuffer - Simple and fast colorbuffer for games and stuff

The motivation behind this package is that i found myself recreating it time and time again every time I needed just a simple colorbuffer so I could change some pixel values in a byte array and render stuff to the screen either with SDL or just in the terminal.

I tried using the image.RGBA package that's native to Go but I realized after running a few benchmarks it was actually quite slow. Even when using RGBAAt and SetRGBA directly. After a few experiments I just went back to maintaining a plain buffer myself and added a few convenience methods for it.

Check out my other project, a [Go Raycaster](https://github.com/kyriacos/go-raycaster) where I added a bit more detail on how this came about.

I have a few more ideas about extending it further. As always feedback and issues are always welcomed.

## Todo

- [ ] Add helper utils for parsing the hex colors
- [ ] Add go doc
- [ ] Add example
- [ ] Add Bounds method so it's closer to the image implementation?
- [ ] Add byte order support (rgba vs bgra) - benchmark
- [ ] Add CI
- [ ] Add quick byte check b[3]
- [ ] Parallel?
