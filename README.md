# Go PiGPIO Socket Interface

Implementation of the [PiGPIO](http://abyz.me.uk/rpi/pigpio/) socket to communicate with the PiGPIO daemon. Can be used
to communicate over network or localhost.

```go
import "github.com/BxNiom/go-pigpio"
```

For documentation have a look at [docs/pigpio-reference.md](./docs/pigpio-reference.md)

### Features

* Communication via sockets allows remote control pigpiod over network
* Almost complete implementation of [python pigpio](http://abyz.me.uk/rpi/pigpio/python.html) (see features to
  implement)
* No need of any C-Headers for compilation

---

### Progress

|                 | Title             | Progress                          | Description                                                      | Status       |
|-----------------|-------------------|-----------------------------------|------------------------------------------------------------------|--------------|
| :green_circle:  | connection socket | ![](https://progress-bar.dev/100) | create a socket and connect to daemon                            | __Finished__ |
| :green_circle:  | gpio mode         | ![](https://progress-bar.dev/100) | set gpio mode                                                    | __Finished__ |
| :green_circle:  | read/write        | ![](https://progress-bar.dev/100) | read/write gpio value                                            | __Finished__ |
| :green_circle:  | PWM               | ![](https://progress-bar.dev/100) | software PWM                                                     | __Finished__ |
| :green_circle:  | hardware          | ![](https://progress-bar.dev/100) | hardware PWM and clock                                           | __Finished__ |
| :green_circle:  | wave              | ![](https://progress-bar.dev/100) | add/create/modify waves                                          | __Finished__ |
| :green_circle:  | i2c               | ![](https://progress-bar.dev/100) | communication with I2C devices                                   | __Finished__ |
| :green_circle:  | spi               | ![](https://progress-bar.dev/100) | communication with SPI devices                                   | __Finished__ |
| :green_circle:  | filter            | ![](https://progress-bar.dev/100) | noise and glitch filters                                         | __Finished__ |
| :green_circle:  | serial            | ![](https://progress-bar.dev/100) | open/read/write to tty                                           | __Finished__ |
| :green_circle:  | callbacks         | ![](https://progress-bar.dev/100) |                                                                  | __Finished__ |
| :green_circle:  | scripts           | ![](https://progress-bar.dev/100) |                                                                  | __Finished__ |
| :green_circle:  | error codes       | ![](https://progress-bar.dev/100) | implement pigpiod error codes                                    | __Finished__ |
| :yellow_circle: | documentation     | ![](https://progress-bar.dev/50)  | currently using comments from pypigpio<br/>modify to fit go code |              |

:white_circle: Todo | :yellow_circle: Work in progress | :orange_circle: Features todo | :green_circle: Done
