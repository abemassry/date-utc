# date-utc
Convert times to UTC quickly on the command line

## Usage

If your system time is set to EDT
```
$> date-utc 1:15pm
Thu 03 Aug 2023 17:15:00 UTC
Thu 03 Aug 2023 01:15:00 PM EDT
Thu 03 Aug 2023 10:15:00 AM PDT

$> date-utc
Thu 03 Aug 2023 16:06:07 UTC

$> date-utc 10:30am
Thu 03 Aug 2023 14:30:00 UTC
Thu 03 Aug 2023 10:30:00 AM EDT
Thu 03 Aug 2023 07:30:00 AM PDT
```

## Install
1. clone this repo
2. `ln -s <location-cloned>/date-utc ~/bin/date-utc`
