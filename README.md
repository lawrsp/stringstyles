# stringstyles

## install

```bash
  go get -u "github/lawrsp/stringstyles"
```

## Usage

### CamelCase

```golang
import "github.com/lawrsp/stringstyles"

input := "the_quick_brown_fox_jumps_over_the_lazy_dog"
output := stringstyles.CamelCase(input)

// got:
// theQuickBrownFoxJumpsOverTheLazyDog
```

### SnakeCase

```golang
import "github.com/lawrsp/stringstyles"

input := "TheQuickBrownFoxJumpsOverTheLazyDog"
output := stringstyles.SnakeCase(input)

// got:
// the_quick_brown_fox_jumps_over_the_lazy_dog
```

### KebabCase

```golang
import "github.com/lawrsp/stringstyles"

input := "TheQuickBrownFoxJumpsOverTheLazyDog"
output := stringstyles.SnakeCase(input)

// got:
// the-quick-brown-fox-jumps-over-the-lazy-dog
```
