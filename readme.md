# Binary Allocation Map Flags

This library will take a BAM as an integer and spit out the values of which the BAM is comprised.

## Usage

```golang
import "github.com/audibleblink/bamflags"

flags := int64(514)
bamflags.ParseInt(flags)
// => [ 512, 2 ]

bam := 514
bamflags.Contains(bam, int64(2))
// => true

bam := 520
bamflags.Contains(bam, int64(2))
// => false
```
