# CallCulator :see_no_evil:

CallCulator takes a list of calls and returns the total cost of these calls.

Note: The caller with the highest total call duration of the day will not be charged.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

The list of call should have the following format

```
time_of_start;time_of_finish;call_from;call_to
```

Example:
```
09:11:30;09:15:22;+351914374373;+351215355312
15:20:04;15:23:49;+351217538222;+351214434422
16:43:02;16:50:20;+351217235554;+351329932233
17:44:04;17:49:30;+351914374373;+351963433432
```

### Installing

First, clone the project:

```
git clone git@github.com:andre-mota/CallCulator.git
```

Then run:

```
go build
```

### Running

After cloning and building, just do:

```
./CallCulator input_file
```

Example:

```
./CallCulator data/testdata.csv
```
