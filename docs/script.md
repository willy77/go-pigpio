# Script Compiler
Use pre defined or custom macros in scripts.

## Macro structure

Each macro can take up to 10 parameters and can return up to 10 values.
> Variables v130-v139 are reserved for macro parameters. <br>
> Return values are saved to v140-v149 <br>
> Macro sub routines starting at tag 1100 <br>

### JSON

```json
[
    {
        "name": "",
        "params": 0,
        "results": 0,
        "code": ""
    }
]
```

| Option  | Description                              | Type   |
|---------|------------------------------------------|--------|
| name    | macro name                               | string |
| params  | paramter count the macro requires (0-10) | int    |
| results | result count (0-10)                      | int    |
| code    | macro sub routine code                   | string |

#### Example

Macro
```json
[
    {
        "name": "per",
        "params": 2,  
        "results": 1,
        "code": "lda v130 mlt 100 div v131 sta v140"
    }
]
```
Code
```
per 200 50
ldr v0
```
creates following output
```
ld v130 200
ld v131 50
call 1100
ld v1 v140 
tag 1100 lda v130 mlt 100 div v131 sta v140 ret
```


## Default macros

Load with ```Compiler.LoadDefaultMacros()```

### ldr
Load macro result to variables

#### Syntax
```
ldr v0..129
```

#### Parameters
| Parameter | Description                                  |
|-----------|----------------------------------------------|
| v0..129   | variables to store previous macro results to | 


#### Example
```
ldr v0 v1 v2
```

### dec
Decode int32 to 4 bytes

#### Syntax
```
dec [v]
```

#### Parameters
| Parameter | Description          |
|-----------|----------------------|
| v         | variable or constant |

#### Example
```
dec v0
ldr v1 v2 v3 v4
```

### per
Calculate percentage x from max

#### Syntax
```
per [max] [x]
```

#### Parameters
| Parameter | Description       |
|-----------|-------------------|
| max       | max value (100%)  |
| x         | the current value |

#### Example
```
per 10 5
ldr v2          # 50%
```

### pera
Calculate percentage - result is saved to accumulator

#### Syntax
```
pera [max] [x]
```

#### Parameters
| Parameter | Description       |
|-----------|-------------------|
| max       | max value (100%)  |
| x         | the current value |

#### Example
```
pera 10 5
```

### lerp
Linear interpolation

#### Syntax
```
lerp [start] [target] [t]
```

#### Parameters
| Parameter | Description   |
|-----------|---------------|
| start     | start value   |
| target    | target value  |
| t         | timer (0-100) |

#### Example
```
lerp 0 10 50 
ldr v2
```

### lerpa
Linear interpolation - result is stored in accumulator

#### Syntax
```
lerpa [start] [target] [t]
```

#### Parameters
| Parameter | Description   |
|-----------|---------------|
| start     | start value   |
| target    | target value  |
| t         | timer (0-100) |

#### Example
```
lerpa 0 10 50 
```

### addx
Add x to y - Result is stored in accumulator

#### Syntax
```
addx [x] [y]
```

#### Parameters
| Parameter | Description                       |
|-----------|-----------------------------------|
| x         | constant or variable or parameter |
| y         | constant or variable or parameter |

#### Example
```
addx 0 v1 
```

### subx
Subtract y from x - Result is stored in accumulator

#### Syntax
```
subx [x] [y]
```

#### Parameters
| Parameter | Description                       |
|-----------|-----------------------------------|
| x         | constant or variable or parameter |
| y         | constant or variable or parameter |

#### Example
```
subx 0 v1 
```

### divx
Divide x by y - Result is stored in accumulator

#### Syntax
```
divx [x] [y]
```

#### Parameters
| Parameter | Description                       |
|-----------|-----------------------------------|
| x         | constant or variable or parameter |
| y         | constant or variable or parameter |

#### Example
```
divx 20 v1 
```

### mltx
Multiply x by y - Result is stored in accumulator

#### Syntax
```
mltx [x] [y]
```

#### Parameters
| Parameter | Description                       |
|-----------|-----------------------------------|
| x         | constant or variable or parameter |
| y         | constant or variable or parameter |

#### Example
```
mltx 20 v1 
```

