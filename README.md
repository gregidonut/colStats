# Col Stats

## Abstract 
Col Stats is a CLI tool that executes statistical operations
a CSV file.

it can receive optional input parameters each with a default
value including:
- `-col` The column on which to execute the operation (defaults to 1)
- `-op` The operation to execute on the selected column. The operations
    include but are not limited to:
  - `sum` Calculates the sum of all values,
  - `avg` Determines the average value of the column.

in addition to the optional flags, this tool accepts any number of
file names to process. If the user provides more than one file name, the
tool combines the results for the same column in all files.
