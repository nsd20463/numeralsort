# numeralsort

Sort text strings so that text strings containing decimal numbers sort as
expected by humans.

For example, the values

  "file1.ext"
  "file10.ext"
  "file100.ext"
  "file2.ext"
  "file20.ext"
  "file3.ext"

sort to

  "file1.ext"
  "file2.ext"
  "file3.ext"
  "file10.ext"
  "file20.ext"
  "file100.ext"

This is useful when filenames contain embedded numbers without '0' prefixes which would make a regular text sort come out in the desired order.

