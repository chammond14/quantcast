# Most Actvie Cookie Finder

A command line application written in Go to search a cookie logfile for the most active cookie on a certain date.

The tool attempts to infer the format of the input data by looking for file extensions in the file path. At the moment, only CSV files are supported, and the tool will default to attempting to read a CSV file if an extension is not provided.

# Contents
1. [Prerequisites](#Prerequisites)
2. [Building and Running](#Building-and-Running)
3. [Command Line Arguments](#Command-line-arguments)
4. [Accepted Data Formats](#Accepted-Data-Formats)
5. [Output](#Output)


# Prerequisites

Go (1.22 or greater) must be installed.

[Installing Go](https://go.dev/doc/install)

# Building and Running
Navigate in the terminal to the directory where this repository has been cloned, and then run the following two commands:

```
go build -o most-active-cookie

./most-active-cookie -f <file> -d <date>
```
This builds a binary with the file name `most-active-cookie`, which is then run with command line arguments (see below for details)

## Testing

This tool contains automated tests. They can be run using the following command:

```
go test
```

# Command Line Arguments

Two command line arguments are required:


| Argument        | Description           
| ------------- |:-------------:|
| -f      | The path to file containing cookie log entries |
| -d      | A date in the format yyyy-mm-dd     |

An example cookie log file which can be used has been included in this repo, named `logfile.csv`

# Accepted Data Formats

### Cookie Log File

The path to the cookie log file is provided as the `-f` flag. If the file exists in the project root directory, just the name can be given:

```
./most-active-cookie -f ../../logfile.csv -d <date>
./most-active-ccokie -f logfile.csv -d <date>
```

#### CSV
Currently only CSV files are supported. Each line in the file must take the structure `cookie,timestamp`:

```
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
```


### Input Date

The date for which the most active cookie should be found. It must be provided in the form `yyyy-mm-dd`:

```
./most-active-cookie -f logfile.csv -d 2018-12-09
```

# Output

If the program runs without issue, results will be printed to the standard output. The application finds the most active cookie, or, in the event of multiple cookies having the same frequency, multiple cookies.

```
./most-active-cookie -f logfile.csv -d 2018-12-09
AtY0laUfhglK3lC7
SAZuXPGUrfbcn5UA
```

In the event of an error, the output will inform of a problem:

```
./most-active-cookie -f logfile.csv -d 2018-12-09
Failed to process, see logs for detail
```

Detailed logs are written to a file in the project directory called `log.log`.